package main

import (
	"bytes"
	"encoding/xml"
	"time"
)

const (
	contentNs  = "http://purl.org/rss/1.0/modules/content/"
	wfwNs      = "http://wellformedweb.org/CommentAPI/"
	dcNs       = "http://purl.org/dc/elements/1.1/"
	atomNs     = "http://www.w3.org/2005/Atom"
	syNs       = "http://purl.org/rss/1.0/modules/syndication/"
	slashNs    = "http://purl.org/rss/1.0/modules/slash/"
	itunesNs   = "http://www.itunes.com/dtds/podcast-1.0.dtd"
	rawvoiceNs = "http://www.rawvoice.com/rawvoiceRssModule/"
	version    = "2.0"
)

// Rss wraps the given RSS channel.
type Rss struct {
	XMLName       xml.Name `xml:"rss"`
	ContentXmlns  string   `xml:"xmlns:content,attr,omitempty"`
	WfwXmlns      string   `xml:"xmlns:wfw,attr,omitempty"`
	DcXmlns       string   `xml:"xmlns:dc,attr,omitempty"`
	AtomXmlns     string   `xml:"xmlns:atom,attr,omitempty"`
	SyXmlns       string   `xml:"xmlns:sy,attr,omitempty"`
	SlashXmlns    string   `xml:"xmlns:slash,attr,omitempty"`
	ItunesXmlns   string   `xml:"xmlns:itunes,attr,omitempty"`
	RawVoiceXmlns string   `xml:"xmlns:rawvoice,attr,omitempty"`
	Version       string   `xml:"version,attr"`
	Channel       *Channel `xml:"channel"`
}

func defaultRss() *Rss {
	return &Rss{
		ContentXmlns:  contentNs,
		WfwXmlns:      wfwNs,
		DcXmlns:       dcNs,
		AtomXmlns:     atomNs,
		SyXmlns:       syNs,
		SlashXmlns:    slashNs,
		ItunesXmlns:   itunesNs,
		RawVoiceXmlns: rawvoiceNs,
		Version:       version,
	}
}

func (r *Rss) Xml() ([]byte, error) {
	var buf bytes.Buffer
	if _, err := buf.Write([]byte(xml.Header)); err != nil {
		return nil, err
	}
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	if err := enc.Encode(r); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type Channel struct {
	Title         string `xml:"title"`
	Description   string `xml:"description"`
	Link          string `xml:"link"`
	Language      string `xml:"language"`
	Copyright     string `xml:"copyright"`
	Image         *Image `xml:"image"`
	ItunesChannel `xml:",omitempty"`
	Item          []*Item `xml:"item"`
}

func (c *Channel) Itunes(category string, explicit bool) {
	c.ItunesChannel = ItunesChannel{
		Image:    ItunesImage{c.Image.Url},
		Category: ItunesCategory{category},
		Explicit: explicit,
	}
}

type Image struct {
	Url   string `xml:"url"`
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type Item struct {
	Title       string    `xml:"title"`
	Description CData     `xml:"description"`
	PubDate     *PubDate  `xml:"pubDate"`
	Enclosure   Enclosure `xml:"enclosure"`
	Guid        string    `xml:"guid"`
	ItunesItem
}

func (i *Item) Itunes(duration int, image string) {
	i.ItunesItem = ItunesItem{
		Duration: duration,
		Image:    ItunesImage{image},
	}
}

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type PubDate struct {
	time.Time
}

func (p PubDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(p.Format(time.RFC1123Z))); err != nil {
		return err
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type CData struct {
	Value string `xml:",cdata"`
}
