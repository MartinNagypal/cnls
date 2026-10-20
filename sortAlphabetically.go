package main

import (
	"os"
	"strings"
)

func sort(entries []os.DirEntry) []os.DirEntry {
	sortedEntries := make([]os.DirEntry, len(entries))
	copy(sortedEntries, entries)

	sorted := false
	for !sorted {
		sorted = true

		for i := 1; i < len(sortedEntries); i++ {
			if strings.ToLower(sortedEntries[i].Name()) < strings.ToLower(sortedEntries[i-1].Name()) {
				sortedEntries[i], sortedEntries[i-1] = sortedEntries[i-1], sortedEntries[i]
				sorted = false
			}
		}
	}

	return sortedEntries
}
