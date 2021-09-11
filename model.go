package main

import (
	"encoding/json"
	"strings"
	"sync"
	"time"
)

type Response struct {
	Data     []Data  `json:"data"`
	Included []Field `json:"included"`
}

func (r Response) toRadio() []Radio {
	data := r.Data
	included := r.Included
	radios := make([]Radio, len(data))

	var wg sync.WaitGroup
	for i, datum := range data {
		wg.Add(1)

		go func(i int, datum Data) {
			defer wg.Done()
			attr := datum.Attributes
			category := datum.getCategory(included)
			media := datum.getMedia(included)
			audio := media.getAudio()
			radio := Radio{
				ID:          datum.ID,
				Title:       attr.Title,
				Description: datum.getDescription(),
				Thumb:       datum.getThumb(),
				PublishAt:   datum.getPublish(),
				Category:    category.Name,
				Audio:       audio,
				Length:      remoteContentLength(audio),
				Duration:    attr.Duration,
			}
			radios[i] = radio
		}(i, datum)
	}

	wg.Wait()
	return radios
}

type Data struct {
	ID         string `json:"id"`
	Attributes struct {
		Title       string `json:"title"`
		Description string `json:"desc"`
		Thumb       string `json:"thumb"`
		PublishedAt string `json:"published-at"`
		Duration    int    `json:"duration"`
		Content     string `json:"content"`
	} `json:"attributes"`
	Relationships struct {
		Category Relation `json:"category"`
		Media    Relation `json:"media"`
	} `json:"relationships"`
}

func (d Data) getThumb() string {
	thumb := d.Attributes.Thumb
	if !isUrlValid(thumb) {
		thumb = imageDomain + thumb
	}
	return thumb
}

func (d Data) getDescription() string {
	attr := d.Attributes

	var builder strings.Builder
	builder.WriteString(attr.Description)
	builder.WriteString("\\n")
	if len(attr.Content) > 0 {
		var block ContentBlock
		if err := json.Unmarshal([]byte(attr.Content), &block); nil == err {
			for _, b := range block.Blocks {
				if strings.Contains(b.Text, "时间轴") {
					continue
				}
				builder.WriteString(b.Text)
				builder.WriteString("\\n")
			}
		}
	}

	desc := builder.String()
	return cleanString(desc)
}

func (d Data) getPublish() time.Time {
	pub := d.Attributes.PublishedAt
	t, _ := time.Parse(time.RFC3339, pub)
	return t
}

func (d Data) getCategory(fields []Field) Category {
	category := d.Relationships.Category
	typ := category.Data.Type
	id := category.Data.ID
	for _, field := range fields {
		if field.Type == typ && field.ID == id {
			return field.Attributes.Category
		}
	}
	return Category{}
}

func (d Data) getMedia(fields []Field) Media {
	media := d.Relationships.Media
	typ := media.Data.Type
	id := media.Data.ID
	for _, field := range fields {
		if field.Type == typ && field.ID == id {
			return field.Attributes.Media
		}
	}
	return Media{}
}

type ContentBlock struct {
	Blocks []struct {
		Text string `json:"text"`
	} `json:"blocks"`
}

type Relation struct {
	Data struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}

type Field struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Category
		Media
	} `json:"attributes"`
}

type Category struct {
	Name string `json:"name"` // type: categories
}

type Media struct {
	Audio string `json:"audio"` // type: media
}

func (m Media) getAudio() string {
	audio := m.Audio
	if !isUrlValid(audio) {
		audio = audioDomain + audio
	}
	return audio
}

type Radio struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumb       string    `json:"thumb"`
	PublishAt   time.Time `json:"publish_at"`
	Category    string    `json:"category"`
	Audio       string    `json:"audio"`
	Length      int64     `json:"length"`
	Duration    int       `json:"duration"`
}

// ByPublishAt implements sort.Interface based on the Radio PublishAt field.
type ByPublishAt []*Radio

func (a ByPublishAt) Len() int           { return len(a) }
func (a ByPublishAt) Less(i, j int) bool { return a[i].PublishAt.After(a[j].PublishAt) }
func (a ByPublishAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
