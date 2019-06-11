// Copyright © 2019 Roald Nefs <info@roaldnefs.com>

package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

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
		if len(args) < 1 {
			fmt.Println("gitignore needs at least one argument")
			os.Exit(1)
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// TODO validate the argument
		language := args[0]

		// TODO download the gitignore from:
		// https://raw.githubusercontent.com/github/gitignore/master/<language>.gitignore
		filepath := wd + "/.gitignore"

		err = downloadGitignore(filepath, language)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf(".gitignore created at %s\n", wd)
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

// downloadGitignore will download a raw .gitignore file from GitHub to a local file.
func downloadGitignore(filepath string, language string) error {
	// Build the URL
	url := "https://raw.githubusercontent.com/github/gitignore/master/" + language + ".gitignore"

	// Get the data
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	// Create the local .gitignore file
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	// Write the body to the .gitignore file
	_, err = io.Copy(out, response.Body)
	return err
}
