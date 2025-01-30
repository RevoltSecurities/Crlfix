package vulnerable

import (
	"net/http"
)

func Vulnerable(responsed *http.Response) (string, string, bool) {
	for header, values := range responsed.Header {
		for _, value := range values {
			if value == "whoami=revolt" {
				return header, value, true
			}
		}
	}
	return "", "", false
}
