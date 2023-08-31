package service

var exceptionList = []string{"services", "token", "shortlink", "links"}

func CheckURL(URL string) bool {
	for _, val := range exceptionList {
		if val == URL || "/"+val == URL {
			return false
		}
	}
	return true
}
