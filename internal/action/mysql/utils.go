package mysql

import (
	"errors"
	"regexp"
)

//root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local
func getDbFromEndpoint(endpoint string) (string, error) {
	// compile := regexp.MustCompile(`([a-zA-Z0-9~!@#$%^&*_]+):([a-zA-Z0-9!@#$%^&*_]+)@tcp\(\d.\d.\d.\d:\d\)/([a-zA-Z0-9_]+)`)
	compile := regexp.MustCompile(`([a-zA-Z0-9~!@#$%^&*_]+):([a-zA-Z0-9!@#=$%^&*_]+)@tcp\(\d+.\d+.\d+.\d+:\d+\)/([a-zA-Z0-9_]+)`)
	subMatch := compile.FindStringSubmatch(endpoint)
	if len(subMatch) < 4 {
		return "", errors.New("db not found")
	}
	return subMatch[3], nil
}
