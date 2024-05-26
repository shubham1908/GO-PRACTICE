package main

import (
	"errors"
	"log"
)

func funcToTest(isReturn bool) (string, error) {
	if isReturn {
		return "", errors.New("error on test")
	}
	log.Printf("printing something")
	return "test passed", nil
}
