package main

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

const auditionKeyWord = "试听"

func generateRadios(data [][]byte) []*Radio {
	radios := make([]*Radio, len(data))
	for i, datum := range data {
		var radio Radio
		_ = json.Unmarshal(datum, &radio)
		radios[i] = &radio
	}
	sort.Sort(ByPublishAt(radios))
	return radios
}

func generateRss(radios []*Radio) *Rss {
	rss := defaultRss()
	rss.Channel = generateChannel(radios)
	return rss
}

func generateChannel(radios []*Radio) *Channel {
	channel := Channel{
		Title:       "GCores Archive",
		Description: "机核网 www.gcores.com (Archive)",
		Link:        "https://assets.dio.wtf/misc/gcores.xml",
		Language:    "zh-cn",
		Copyright:   "www.gcores.com",
		Image: &Image{
			Url:   "http://media.fmit.cn/feed/gadionewlogos.png",
			Title: "GCores Archive",
			Link:  "https://assets.dio.wtf/misc/gcores.xml",
		},
	}
	items := make([]*Item, 0)
	for i := range radios {
		radio := radios[i]
		title := radio.Title
		if strings.Contains(title, auditionKeyWord) {
			continue
		}
		item := &Item{
			Title:       radio.Title,
			Description: CData{radio.Description},
			PubDate:     &PubDate{radio.PublishAt},
			Enclosure: Enclosure{
				Url:    radio.Audio,
				Type:   "audio/mpeg",
				Length: strconv.FormatInt(radio.Length, 10),
			},
			Guid: radio.Audio,
		}
		item.Itunes(radio.Duration, radio.Thumb)
		items = append(items, item)
	}
	channel.Item = items
	channel.Itunes("Games & Hobbies", false)
	return &channel
}
