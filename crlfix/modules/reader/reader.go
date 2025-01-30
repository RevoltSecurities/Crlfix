package reader

import (
	"bufio"
	"os"
	"strings"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/utils"
)

func Reader(filename string) ([]string, error) {
	var urls []string
	filecontent, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer filecontent.Close()
	scanner := bufio.NewScanner(filecontent)
	for scanner.Scan() {
		line := scanner.Text()
		tline := strings.TrimSpace(line)
		urls = append(urls, tline)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return utils.Set(urls), nil
}
