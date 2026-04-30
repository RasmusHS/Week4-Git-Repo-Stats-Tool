package output

import (
	"fmt"
	"os"
	"text/tabwriter"

	"git-stats/stats"
)

func PrintAuthorStats(authors []stats.AuthorStats) {
	fmt.Println("\n=== Commits Per Author ===")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Author\tEmail\tCommits\tAdded\tRemoved")
	fmt.Fprintln(w, "------\t-----\t-------\t-----\t-------")
	for _, a := range authors {
		fmt.Fprintf(w, "%s\t%s\t%d\t%d\t%d\n", a.Name, a.Email, a.Commits, a.Added, a.Removed)
	}
	w.Flush()
}

func PrintActivity(entries []stats.ActivityEntry) {
	fmt.Println("\n=== Activity Over Time ===")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Date\tCommits")
	fmt.Fprintln(w, "----\t-------")
	for _, e := range entries {
		fmt.Fprintf(w, "%s\t%d\n", e.Date.Format("2006-01-02"), e.Commits)
	}
	w.Flush()
}

func PrintFileChurn(files []stats.FileChurn) {
	fmt.Println("\n=== Most Edited Files ===")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Path\tEdits\tAdded\tRemoved")
	fmt.Fprintln(w, "----\t-----\t-----\t-------")
	for _, f := range files {
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\n", f.Path, f.Edits, f.TotalAdded, f.TotalRemoved)
	}
	w.Flush()
}

func PrintCommitSize(s stats.CommitSizeStats) {
	fmt.Println("\n=== Average Commit Size ===")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "Total Commits:\t%d\n", s.TotalCommits)
	fmt.Fprintf(w, "Avg Lines Added:\t%.1f\n", s.AvgAdded)
	fmt.Fprintf(w, "Avg Lines Removed:\t%.1f\n", s.AvgRemoved)
	fmt.Fprintf(w, "Avg Files Changed:\t%.1f\n", s.AvgFiles)
	w.Flush()
}
