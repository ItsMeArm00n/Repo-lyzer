package analyzer

import (
	"github.com/agnivo988/Repo-lyzer/internal/github"
)

func CommitsPerDay(commits []github.Commit) map[string]int {
	result := make(map[string]int)

	for _, c := range commits {
		day := c.Commit.Author.Date.Format("2006-01-02")
		result[day]++
	}

	return result
}
