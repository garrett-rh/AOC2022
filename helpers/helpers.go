package helpers

import "log"

func Check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
