package app

import (
	"errors"
	"strings"
)

func QueryToArray(queryParam string) ([]string, error) {
	/**
	* Разделяет параметр на массив
	 */
	var err error
	if len(queryParam) < 1 {
		err = errors.New("Query is empty!")
		return []string{}, err
	}
	params := strings.Split(queryParam, ",")
	return params, err
}
