package patcher

import "os"

type Patcher struct{}

func (_ Patcher) PatchFile(path string, url string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	urls, err := FindURLs(data)
	if err != nil {
		return err
	}
	for _, index := range urls {
		Patch(data, index, url)
	}
	err = os.WriteFile(path+".patched", data, 0644)
	return err
}

func Patch(data []byte, index int, url string) {
	urlBytes := []byte(url + "\000")
	for i := 0; i < len(urlBytes); i++ {
		data[index+i] = urlBytes[i]
	}
}
