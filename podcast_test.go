package main

import (
	"testing"
	"time"
)

var item = Item{
	Title:           "青春的狂躁与成长的迷茫，都在《FLCL》里了",
	Description:     CData{"《FLCL》就是这样一部动画，它把青春期的幻想展现出来，在这之上又加入些许对于未来的迷茫，同时配上一堆日本另类摇滚音乐，让看的人深陷其中。本期节目，我们想把这部21年前的动画推荐给你。"},
	PubDate:         &PubDate{time.Now()},
	EnclosureUrl:    "http://alioss.gcores.com/uploads/audio/c476eec0-44f4-44ca-969b-77269c92e49a.mp3",
	EnclosureLength: "56965329",
	EnclosureType:   "audio/mpeg",
	Guid:            "http://alioss.gcores.com/uploads/audio/c476eec0-44f4-44ca-969b-77269c92e49a.mp3",
}

var channel = Channel{
	Title:       "机核网 GADIO 游戏广播",
	Description: "机核网 www.gcores.com",
	Link:        "http://feed.tangsuanradio.com/gadio.xml",
	Language:    "zh-cn",
	Copyright:   "www.gcores.com",
	ItunesChannel: ItunesChannel{
		Image:    ItunesImage{"https://image.gcores.com/07bf7f05-a9df-457c-b450-397e3942e75c.jpg"},
		Category: ItunesCategory{"Sports"},
		Explicit: false,
	},
	Item: make([]*Item, 0),
}

func TestGenChannel(t *testing.T) {
	channel.Item = append(channel.Item, &item)
	channel.Item = append(channel.Item, &item)
	rss := defaultRss()
	rss.Channel = &channel
	xml, err := rss.Xml()
	if nil != err {
		t.Error(err)
		return
	}
	t.Log(xml)
}
