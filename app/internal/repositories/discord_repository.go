package repositories

import (
	"github.com/mmcdole/gofeed"
)

type DiscordRepository struct {
	FeedParser *gofeed.Parser
}

func NewDiscordRepository(parser *gofeed.Parser) *DiscordRepository {
	return &DiscordRepository{FeedParser: parser}
}

// fetchRSS method goes here
// fetchRSS関数
func (r *DiscordRepository) FetchRSS(url string) ([]*gofeed.Item, error) {
	feed, err := r.FeedParser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}
