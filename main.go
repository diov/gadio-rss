package main

import "log"

func main() {
	err := setupManager()
	if nil != err {
		log.Fatalln(err)
	}
	fetch(0, true)
}
