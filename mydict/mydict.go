package mydict

import "errors"

// Dictionary type (type은 method를 가질 수 있다)
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
