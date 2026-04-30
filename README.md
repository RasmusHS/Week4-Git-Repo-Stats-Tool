# Week4-Git-Repo-Stats-Tool
A simple CLI tool that analyzes a local Git repository and produces statistics: commits per author, activity over time, file churn (most-edited files), and average commit size.
 
Built as a learning project focused on process execution in Go (`os/exec`), text parsing, data aggregation, and tabular output.

## Usage
 
```bash
# Analyze the current directory
go run .
 
# Analyze a specific repo
go run . --repo /path/to/repo
 
# Limit the number of files shown in churn report (default: 10)
go run . --repo /path/to/repo --churn-limit 20
```
 
## Output
 
The tool produces four reports:
 
- **Commits Per Author** — commit count, lines added/removed per contributor
- **Activity Over Time** — daily commit counts sorted chronologically
- **Most Edited Files** — top files by number of edits with total lines added/removed
- **Average Commit Size** — average lines added, removed, and files changed per commit
## How It Works
 
No Git libraries are used. The tool shells out to `git log` with `--pretty=format` and `--numstat`, then parses the output. The format uses null byte delimiters (`%x00`) to safely separate fields without colliding with commit message content.
 
## Project Structure
 
```
git-stats/
├── main.go          // CLI flags, orchestration
├── git/
│   └── git.go       // executes git commands
├── parser/
│   └── parser.go    // parses git log output into structs
├── models/
│   └── models.go    // Commit, FileChange
├── stats/
│   └── stats.go     // aggregation logic
└── output/
    └── table.go     // tabwriter formatted terminal output
```
 
## Requirements
 
- Go 1.22+
- Git installed and available on PATH
