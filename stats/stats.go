package stats

import (
	"sort"
	"time"

	"git-stats/models"
)

// AuthorStats represents aggregated statistics for a single author across all their commits
type AuthorStats struct {
	Name    string
	Email   string
	Commits int
	Added   int
	Removed int
}

// ActivityEntry represents the number of commits made on a specific date, used for tracking activity over time
type ActivityEntry struct {
	Date    time.Time
	Commits int
}

// FileChurn represents the total number of edits and lines added/removed for a specific file across all commits, used to identify files with high churn
type FileChurn struct {
	Path         string
	Edits        int
	TotalAdded   int
	TotalRemoved int
}

// CommitSizeStats represents aggregated statistics about the size of commits,
// including total commits, total lines added/removed, and average lines added/removed per commit as well as average number of files changed per commit
type CommitSizeStats struct {
	TotalCommits int
	TotalAdded   int
	TotalRemoved int
	AvgAdded     float64
	AvgRemoved   float64
	AvgFiles     float64
}

// CommitsPerAuthor takes a list of commits and returns a sorted list of AuthorStats,
// showing the total number of commits and lines added/removed for each author identified by their email address and name
func CommitsPerAuthor(commits []models.Commit) []AuthorStats {
	m := make(map[string]*AuthorStats)

	for _, c := range commits {
		s, ok := m[c.Email]
		if !ok {
			s = &AuthorStats{Name: c.Author, Email: c.Email}
			m[c.Email] = s
		}
		s.Commits++
		for _, f := range c.FileChanges {
			s.Added += f.Added
			s.Removed += f.Removed
		}
	}

	result := make([]AuthorStats, 0, len(m))
	for _, s := range m {
		result = append(result, *s)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Commits > result[j].Commits
	})

	return result
}

// ActivityOverTime takes a list of commits and returns a sorted list of ActivityEntry,
// showing the number of commits made on each date, which can be used to track activity trends over time and identify periods of high or low activity in the repository
func ActivityOverTime(commits []models.Commit) []ActivityEntry {
	m := make(map[string]int)

	for _, c := range commits {
		day := c.Date.Format("2006-01-02")
		m[day]++
	}

	result := make([]ActivityEntry, 0, len(m))
	for day, count := range m {
		date, _ := time.Parse("2006-01-02", day)
		result = append(result, ActivityEntry{Date: date, Commits: count})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Date.Before(result[j].Date)
	})

	return result
}

// TopFileChurn takes a list of commits and returns a sorted list of FileChurn,
// showing the total number of edits and lines added/removed for each file across all commits, which can be used to identify files with high churn that may require refactoring or additional attention
func TopFileChurn(commits []models.Commit, limit int) []FileChurn {
	m := make(map[string]*FileChurn)

	for _, c := range commits {
		for _, f := range c.FileChanges {
			fc, ok := m[f.Path]
			if !ok {
				fc = &FileChurn{Path: f.Path}
				m[f.Path] = fc
			}
			fc.Edits++
			fc.TotalAdded += f.Added
			fc.TotalRemoved += f.Removed
		}
	}

	result := make([]FileChurn, 0, len(m))
	for _, fc := range m {
		result = append(result, *fc)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Edits > result[j].Edits
	})

	if limit > 0 && limit < len(result) {
		result = result[:limit]
	}

	return result
}

// AverageCommitSize takes a list of commits and returns a CommitSizeStats struct,
// showing the total number of commits, total lines added/removed, and average lines added/removed per commit as well as average number of files changed per commit, which can be used to understand the typical size of commits in the repository and identify trends in commit sizes over time
func AverageCommitSize(commits []models.Commit) CommitSizeStats {
	s := CommitSizeStats{TotalCommits: len(commits)}

	if s.TotalCommits == 0 {
		return s
	}

	totalFiles := 0
	for _, c := range commits {
		totalFiles += len(c.FileChanges)
		for _, f := range c.FileChanges {
			s.TotalAdded += f.Added
			s.TotalRemoved += f.Removed
		}
	}

	n := float64(s.TotalCommits)
	s.AvgAdded = float64(s.TotalAdded) / n
	s.AvgRemoved = float64(s.TotalRemoved) / n
	s.AvgFiles = float64(totalFiles) / n

	return s
}
