package news

import (
	"context"
	"github.com/cloudfoundry-attic/jibber_jabber"
	"github.com/mmcdole/gofeed"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"strings"
)

var parser = gofeed.NewParser()

var articleCache = make(map[string][]Item)

func GetArticles(ctx context.Context, server string) []Item {
	if items, exists := articleCache[server]; exists {
		return items
	}

	feed, err := parser.ParseURL(server)
	if err != nil {
		runtime.LogErrorf(ctx, "failed to fetch feed: %s", err.Error())
		return nil
	}

	var items []Item

	lang, err := jibber_jabber.DetectLanguage()
	if err != nil {
		runtime.LogErrorf(ctx, "failed to detect language: %s", err.Error())
		lang = "en"
	}

	for _, item := range feed.Items {
		if feed.FeedType == "rss" {
			if lang != "en" {
				if strings.HasSuffix(item.Link, lang) || !strings.Contains(item.Link, "?lang=") {
					continue
				}
			} else {
				if strings.Contains(item.Link, "?lang=") {
					continue // if the article is not english then continue
				}
			}
		} else {
			if feed.Language != lang {
				continue
			}
		}

		items = append(items, Item{
			Title:           item.Title,
			Description:     item.Description,
			Content:         item.Content,
			Link:            item.Link,
			Links:           item.Links,
			Updated:         item.Updated,
			UpdatedParsed:   item.UpdatedParsed,
			Published:       item.Published,
			PublishedParsed: item.PublishedParsed,
			GUID:            item.GUID,
			Image:           item.Image,
			Categories:      item.Categories,
			Enclosures:      item.Enclosures,
			Custom:          item.Custom,
		})
	}

	articleCache[server] = items

	return items
}
