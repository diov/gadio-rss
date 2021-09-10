package main

type ItunesChannel struct {
	Image    ItunesImage    `xml:"itunes:image,omitempty"`
	Category ItunesCategory `xml:"itunes:category,omitempty"`
	Explicit bool           `xml:"itunes:explicit,omitempty"`
}

type ItunesImage struct {
	Href string `xml:"href,attr"`
}

type ItunesCategory struct {
	Text string `xml:"text,attr"`
}

type ItunesItem struct {
	Duration int         `xml:"itunes:duration,omitempty"`
	Image    ItunesImage `xml:"itunes:image,omitempty"`
}
