package utils

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func Set(urls []string) []string {
	sets := make(map[string]bool)
	var results []string

	for _, url := range urls {
		if _, ok := sets[url]; !ok {
			sets[url] = true
			results = append(results, url)
		}
	}
	return results
}

func Pathadder(baseurl, path string) (string, error) {
	Baseurl, err := url.Parse(baseurl)
	if err != nil {
		return "", err
	}
	if strings.HasSuffix(Baseurl.Path, "/") {
		Baseurl.Path += path
	} else {
		Baseurl.Path += "/" + path
	}
	return Baseurl.String(), nil
}

func IsStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func IsPermission(filename string) (bool, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		if os.IsPermission(err) {
			return false, fmt.Errorf("You don't have any write permission in this %s file", filename)
		}
		return false, err
	}
	defer file.Close()
	return true, nil
}
