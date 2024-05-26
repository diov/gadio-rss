package main

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
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
	radios = filterChannel(radios)
	rss.Channel = generateChannel(radios)
	return rss
}

func filterChannel(radios []*Radio) []*Radio {
	keywords := []string{"试听", "录音笔"}
	filtered := make([]*Radio, 0)
	for i := range radios {
		radio := radios[i]
		title := radio.Title

		skip := false
		for j := range keywords {
			if strings.Contains(title, keywords[j]) {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		filtered = append(filtered, radio)
	}
	return filtered
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
