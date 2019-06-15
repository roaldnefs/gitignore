// Copyright © 2019 Roald Nefs <info@roaldnefs.com>

package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

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

		// Get the preferred languange from the argument
		search := strings.ToLower(args[0]) + ".gitignore"
		filePath := wd + "/.gitignore"

		// Fetch gitignore templates from the github.com/github/gitignore repository
		client := github.NewClient(nil)
		_, repositoryContent, _, err := client.Repositories.GetContents(context.Background(), "github", "gitignore", "", nil)
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

		// Create the local .gitignore file
		out, err := os.Create(filePath)
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
