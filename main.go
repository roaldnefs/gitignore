package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

// global indicates to search globally useful gitignores
var global bool

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "gitignore [language name]",
	Short: "A tool for downloading .gitignore templates",
	Long: `Gitignore will create a new .gitignore file in the current working
directory.

Example: gitignore Python -> resulting in a new .gitignore file for Python.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure an argument is given
		if len(args) < 1 {
			fmt.Println("gitignore needs at least one argument")
			os.Exit(1)
		}

		// Get the working directory
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Get the preferred language from the argument
		search := strings.ToLower(args[0]) + ".gitignore"
		filePath := path.Join(wd, ".gitignore")

		// Build path for gitignore templates based on flags
		repositoryPath := ""
		if global {
			repositoryPath = "Global/"
		}

		// Fetch gitignore templates from the github.com/github/gitignore repository
		client := github.NewClient(nil)
		_, repositoryContent, _, err := client.Repositories.GetContents(context.Background(), "github", "gitignore", repositoryPath, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var gitignore *github.RepositoryContent

		// Compare the argument to the gitignore files available in the repository
		for _, file := range repositoryContent {
			if strings.Compare(strings.ToLower(*file.Name), search) == 0 {
				gitignore = file
				break
			}
		}

		// Throw an error if none matching gitignore file is found
		if gitignore == nil {
			fmt.Println("no matching gitignore found")
			os.Exit(1)
		}

		// Check for existing .gitignore file
		append := false
		if exists, err := exists(filePath); exists {
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// Ask the user to append the existing .gitignore file otherwise is will be overwritten
			if askForConfirmation("Do you want to append to the existing .gitignore?") {
				append = true
			}
		}

		// Check if the user is in the root of a git repository
		gitDirectory := path.Join(wd, ".git")
		if exists, err := exists(gitDirectory); !exists {
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// Exit if the user doesn't want to create a .gitignore file in the current directory
			if !askForConfirmation("This doesn't seem like a git repository. Do you want to continue?") {
				os.Exit(0)
			}
		}

		// Get the download URL for the gitignore template
		downloadURL := *gitignore.DownloadURL
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Get the gitignore template
		response, err := http.Get(downloadURL)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer response.Body.Close()

		// Set flags used for writing to the file
		fileFlag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
		if append {
			fileFlag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
		}

		// Create the local .gitignore file
		out, err := os.OpenFile(filePath, fileFlag, 0600)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer out.Close()

		// Write to contents of the gitignore template to the local .gitignore file
		_, err = io.Copy(out, response.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf(".gitignore created at %s\n", filePath)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&global, "global", "g", false, "Search globally useful gitignores")
}

// askForConfigrmation asks the user for confirmation
func askForConfirmation(s string) bool {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		answer, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		answer = strings.ToLower(strings.TrimSpace(answer))

		yesResponses := []string{"y", "yes"}
		noResponses := []string{"n", "no"}

		if stringInSlice(answer, yesResponses) {
			return true
		} else if stringInSlice(answer, noResponses) {
			return false
		}
	}
}

// stringInSlice checks is the string is in the list
func stringInSlice(s string, list []string) bool {
	for _, entry := range list {
		if entry == s {
			return true
		}
	}
	return false
}

// exists checks if the given path exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
