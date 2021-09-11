package main

import "log"

func main() {
	err := setupManager()
	if nil != err {
		log.Fatalln(err)
	}
	fetch(0, false)

	all, err := mgr.All()
	if nil != err {
		log.Fatalln(err)
	}

	radios := generateRadios(all)
	rss := generateRss(radios)
	xml, err := rss.Xml()
	if nil != err {
		log.Fatalln(err)
	}
	err = writeFile("gcores.xml", xml)
	if nil != err {
		log.Fatalln(err)
	}
}
