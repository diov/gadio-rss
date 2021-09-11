package main

import (
	"encoding/json"
	"sort"
	"strconv"
)

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
		Link:        "https://wiki.dio.wtf/gcores.xml",
		Language:    "zh-cn",
		Copyright:   "www.gcores.com",
		Image: &Image{
			Url:   "http://media.fmit.cn/feed/gadionewlogos.png",
			Title: "GCores Archive",
			Link:  "https://wiki.dio.wtf/gcores.xml",
		},
	}
	items := make([]*Item, len(radios))
	for i, radio := range radios {
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
		items[i] = item
	}
	channel.Item = items
	channel.Itunes("Games & Hobbies", false)
	return &channel
}
