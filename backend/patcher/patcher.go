package patcher

import (
	"bytes"
	"os"
)

const bufSize = 4096

func FindURLs(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	findAllOccurrences(data, []string{""})
	return nil
}

func findAllOccurrences(data []byte, searches []string) map[string][]int {
	results := make(map[string][]int)
	for _, search := range searches {
		searchData := data
		term := []byte(search)
		for x, d := bytes.Index(searchData, term), 0; x > -1; x, d = bytes.Index(searchData, term), d+x+1 {
			results[search] = append(results[search], x+d)
			searchData = searchData[x+1 : len(searchData)]
		}
	}
	return results
}
