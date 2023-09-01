package service

import "strings"

var exceptionList = []string{"services", "token", "shortlink", "links"}

func CheckURL(URL string) bool {
	for _, val := range exceptionList {
		if strings.HasPrefix(URL, val) || strings.HasPrefix(URL, "/"+val) {
			return false
		}
	}
	return true
}
