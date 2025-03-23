package common

import "fmt"

func CheckError(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
	}
}
