package utils

import "log"

func TestErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
