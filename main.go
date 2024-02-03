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
	S3Account    string
	S3AccessKey  string
	S3SecretKey  string
}

func main() {
	flag.BoolVar(&flags.ForceRefresh, "R", false, "Force refresh record")
	flag.StringVar(&flags.Token, "T", "", "Github token")
	flag.StringVar(&flags.Output, "O", "gcores.xml", "Output feed file path")
	flag.StringVar(&flags.S3Account, "A", "", "S3 account id")
	flag.StringVar(&flags.S3AccessKey, "K", "", "S3 access key")
	flag.StringVar(&flags.S3SecretKey, "S", "", "S3 secret key")
	flag.Parse()

	log.SetOutput(os.Stdout)
	if flags.Token != "" {
		setupGitManager(flags.Token)
		if err := gitMgr.getPreviousArtifact(); nil != err {
			log.Printf("Fetch previous db failed: %v\n", err)
		}
	}
	if err := setupDbManager(flags.ForceRefresh); nil != err {
		log.Fatalln(err)
	}
	if err := setupR2Manager(flags.S3Account, flags.S3AccessKey, flags.S3SecretKey); nil != err {
		log.Fatalln(err)
	}

	if err := fetch(0); nil != err {
		log.Fatalln(err)
	}

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

	if err := r2Mgr.uploadFeedFile(flags.Output); nil != err {
		log.Fatalln(err)
	} else {
		log.Println("New feed has uploaded to S3")
	}
}
