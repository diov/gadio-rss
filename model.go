package main

import (
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

	for i, datum := range data {
		attr := datum.Attributes
		category := datum.getCategory(included)
		media := datum.getMedia(included)
		radio := Radio{
			ID:          datum.ID,
			Title:       attr.Title,
			Description: attr.Description,
			Thumb:       datum.getThumb(),
			PublishAt:   datum.getPublish(),
			Category:    category.Name,
			Audio:       media.getAudio(),
			Duration:    media.Duration,
		}
		radios[i] = radio
	}

	return radios
}

type Data struct {
	ID         string `json:"id"`
	Attributes struct {
		Title       string `json:"title"`
		Description string `json:"desc"`
		Thumb       string `json:"thumb"`
		PublishedAt string `json:"published-at"`
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
	Audio    string `json:"audio"`    // type: media
	Duration uint   `json:"duration"` // type: media
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
	Duration    uint      `json:"duration"`
}
