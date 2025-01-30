package validate

import "net/url"

func Validurl(u string) bool {
	U, err := url.Parse(u)
	if err != nil || U.Scheme == "" || U.Host == "" {
		return false
	}
	return true
}
