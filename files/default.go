package files

import (
	"io/ioutil"
)

// ReadTextFile simple, non-buffered, file reader func
func ReadTextFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
