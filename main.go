package main

import (
	"flag"
	"log"
	"os"
)

var flags struct {
	ForceRefresh bool
	Token        string
	Output       string
}

func main() {
	flag.BoolVar(&flags.ForceRefresh, "R", false, "Force refresh record")
	flag.StringVar(&flags.Token, "T", "", "Github token")
	flag.StringVar(&flags.Output, "O", "gcores.xml", "Output feed file path")
	flag.Parse()

	log.SetOutput(os.Stdout)
	if err := setupDbManager(); nil != err {
		log.Fatalln(err)
	}
	if "" != flags.Token {
		setupGitManager(flags.Token)
		if err := gitMgr.getPreviousArtifact(); nil != err {
			log.Printf("Fetch previous db failed: %v\n", err)
		}
	}
	if flags.ForceRefresh {
		_ = dbMgr.Drop()
	}

	fetch(0)

	all, err := dbMgr.All()
	if nil != err {
		log.Fatalln(err)
	}

	radios := generateRadios(all)
	rss := generateRss(radios)
	xml, err := rss.Xml()
	if nil != err {
		log.Fatalln(err)
	}
	// Bypass if content didn't modified
	if size := fileSize(flags.Output); int64(len(xml)) == size {
		log.Println("Content didn't modified, ignore this round")
		return
	}

	if err = writeFile(flags.Output, xml); nil != err {
		log.Fatalln(err)
	}

	if err := gitMgr.pushFeedFile(flags.Output); nil != err {
		log.Fatalln(err)
	} else {
		log.Println("New feed has pushed to wiki")
	}
}
