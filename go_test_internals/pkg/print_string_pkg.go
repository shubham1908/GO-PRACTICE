package pkg

import (
	"errors"
	"log"
)

func FuncToTestFromPkg(isReturn bool) (string, error) {
	if isReturn {
		return "", errors.New("error on test")
	}
	log.Printf("printing something")
	return "test passed", nil
}
