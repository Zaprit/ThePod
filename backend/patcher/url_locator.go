package patcher

import (
	"bytes"
	"strings"
)

func FindURLs(data []byte) (map[string]int, error) {

	occurrences := findAllOccurrences(data, []string{"http://", "https://"})

	urls := make(map[string]int)

	for _, index := range occurrences {
		if index == 0 {
			continue
		}
		url := ""
		for i := index; data[i] != 0; i++ {
			url += string(data[i])
		}
		urls[url] = index
	}

	for url := range urls {
		if strings.Contains(url, "%") {
			delete(urls, url)
		}
		if len(url) <= 8 {
			delete(urls, url)
		}
		if !strings.HasPrefix(url, "http") {
			delete(urls, url)
		}
	}
	return urls, nil
}

func findAllOccurrences(data []byte, searches []string) []int {
	results := make([]int, 1)
	for _, search := range searches {
		searchData := data
		term := []byte(search)
		for x, d := bytes.Index(searchData, term), 0; x > -1; x, d = bytes.Index(searchData, term), d+x+1 {
			results = append(results, x+d)
			searchData = searchData[x+1:]
		}
	}
	return results
}
