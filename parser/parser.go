package parser

import (
	"bufio"
	"git-stats/models"
	"strconv"
	"strings"
	"time"
)

func ParseLog(raw string) ([]models.Commit, error) {
	var commits []models.Commit
	var current *models.Commit

	scanner := bufio.NewScanner(strings.NewReader(raw))
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		// Null byte delimiter means this is a commit header line
		if strings.Contains(line, "\x00") {
			if current != nil {
				commits = append(commits, *current)
			}

			parts := strings.Split(line, "\x00") // Expecting 5 parts: hash, author, email, date, subject
			if len(parts) != 5 {
				continue
			}

			date, err := time.Parse(time.RFC3339, parts[3]) // Git's %aI format is ISO 8601 which is compatible with time.RFC3339
			if err != nil {
				continue
			}

			current = &models.Commit{ // Start a new commit record with the parsed header info
				Hash:    parts[0],
				Author:  parts[1],
				Email:   parts[2],
				Date:    date,
				Subject: parts[4],
			}
			continue
		}

		// Otherwise it's a numstat line: <added>\t<removed>\t<path>
		if current == nil {
			continue
		}

		fields := strings.Split(line, "\t") // Expecting 3 fields: added, removed, path
		if len(fields) != 3 {
			continue
		}

		// Binary files show "-" for added/removed
		added, _ := strconv.Atoi(fields[0])
		removed, _ := strconv.Atoi(fields[1])

		current.FileChanges = append(current.FileChanges, models.FileChange{ // Add this file change to the current commit
			Added:   added,
			Removed: removed,
			Path:    fields[2],
		})
	}

	// Don't forget the last commit
	if current != nil {
		commits = append(commits, *current)
	}

	// Return the list of commits and any scanning error
	return commits, scanner.Err()
}
