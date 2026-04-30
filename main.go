package main

import (
	"flag"
	"fmt"
	"os"

	"git-stats/git"
	"git-stats/output"
	"git-stats/parser"
	"git-stats/stats"
)

func main() {
	// Define command-line flags for the repository path and churn limit,
	// with default values and descriptions to guide the user when running the tool from the command line
	repoPath := flag.String("repo", ".", "Path to your local git repository")
	churnLimit := flag.Int("churn-limit", 10, "number of top files to show in churn report")
	flag.Parse()

	// Execute the git log command to retrieve the raw commit data from the specified repository path,
	// and handle any errors that occur during this process by printing an error message and exiting with a non-zero status code to indicate failure to the user
	raw, err := git.Log(*repoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Parse the raw git log output into a structured list of Commit objects,
	// and handle any errors that occur during parsing by printing an error message and exiting with a non-zero status code to indicate failure to the user
	commits, err := parser.ParseLog(raw)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing log: %v\n", err)
		os.Exit(1)
	}

	// If no commits are found in the repository, print a message to the user and exit gracefully with a zero status code to indicate successful execution without any data to analyze
	if len(commits) == 0 {
		fmt.Println("No commits found.")
		return
	}

	// Print the total number of commits analyzed to provide feedback to the user about the scope of the analysis being performed
	fmt.Printf("Analyzed %d commits\n", len(commits))

	// Generate and print the various statistics based on the parsed commits,
	// including commits per author, activity over time, file churn, and average commit size, using the functions defined in the stats package to compute the statistics and the output package to format and display the results in a readable format for the user
	output.PrintAuthorStats(stats.CommitsPerAuthor(commits))
	output.PrintActivity(stats.ActivityOverTime(commits))
	output.PrintFileChurn(stats.TopFileChurn(commits, *churnLimit))
	output.PrintCommitSize(stats.AverageCommitSize(commits))
}
