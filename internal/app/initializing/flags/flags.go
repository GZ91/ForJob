package flags

import "flag"

func ReadFlags() (string, string, string, string, string) {
	addressServer := flag.String("a", "localhost:8080", "Run Address server")
	addressServerURL := flag.String("b", "http://localhost:8080/", "Address server for URL")
	lvlLogs := flag.String("l", "info", "log level")
	connectionStringDB := flag.String("d", "", "connection string for postgresql")
	SecretKey := flag.String("k", "SecretKey", "connection string for postgresql")

	flag.Parse()
	return *addressServer, *addressServerURL, *lvlLogs, *connectionStringDB, *SecretKey
}
