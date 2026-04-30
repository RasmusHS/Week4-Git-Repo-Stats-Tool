package models

import "time"

// Commit represents a single git commit with its metadata and associated file changes
type Commit struct {
	Hash        string
	Author      string
	Email       string
	Date        time.Time
	Subject     string
	FileChanges []FileChange
}

// FileChange represents the changes made to a single file in a commit
type FileChange struct {
	Added   int
	Removed int
	Path    string
}
