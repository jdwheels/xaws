package internal

import "log"

func PrintStringArray(arr []*string) {
	for _, value := range arr {
		log.Println(*value)
	}
}
