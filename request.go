package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	pageSize    = 10
	userAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36 Edg/93.0.961.44"
	domain      = "www.gcores.com"
	imageDomain = "https://image.gcores.com/"
	audioDomain = "https://alioss.gcores.com/uploads/audio/"
)

func generateEndPoint(offset int) string {
	u := url.URL{
		Scheme: "https",
		Host:   domain,
		Path:   "/gapi/v1/radios",
	}
	params := map[string]string{
		"page[limit]":      strconv.Itoa(pageSize),
		"page[offset]":     strconv.Itoa(offset),
		"sort":             "-published-at",
		"include":          "media,category",
		"filter[list-all]": "0",
		"fields[radios]":   "title,desc,thumb,published-at,media,duration,category,content",
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func fetch(page int) {
	log.Printf("start fetch %d page radios", page+1)
	endPoint := generateEndPoint(page * pageSize)
	resp, err := http.Get(endPoint)
	if nil != err {
		log.Println("get radios failed", err)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if nil != err {
		log.Println("parse json failed", err)
		return
	}

	radios := response.toRadio()
	stop := shouldStop(radios)
	log.Printf("%d page fetch %d radios", page+1, len(radios))
	for _, radio := range radios {
		marshal, _ := json.Marshal(radio)
		if err := mgr.Insert([]byte(radio.ID), marshal); nil != err {
			log.Println(err)
		}
	}
	if stop {
		return
	}

	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fetch(page + 1)
}
