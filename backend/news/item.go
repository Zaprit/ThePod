package news

import (
	"github.com/mmcdole/gofeed"
	"time"
)

type Item struct {
	Title           string              `json:"title,omitempty"`
	Description     string              `json:"description,omitempty"`
	Content         string              `json:"content,omitempty"`
	Link            string              `json:"link,omitempty"`
	Links           []string            `json:"links,omitempty"`
	Updated         string              `json:"updated,omitempty"`
	UpdatedParsed   *time.Time          `json:"updated_parsed,omitempty"`
	Published       string              `json:"published,omitempty"`
	PublishedParsed *time.Time          `json:"published_parsed,omitempty"`
	GUID            string              `json:"guid,omitempty"`
	Image           *gofeed.Image       `json:"image,omitempty"`
	Categories      []string            `json:"categories,omitempty"`
	Enclosures      []*gofeed.Enclosure `json:"enclosures,omitempty"`
	StupidEnclosure *gofeed.Enclosure   `json:"stupid_enclosure,omitempty"`
	Custom          map[string]string   `json:"custom,omitempty"`
}
