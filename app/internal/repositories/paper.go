package repositories

import (
	"app/internal/utils"

	"github.com/mmcdole/gofeed"
)

type PaperRepository struct {
	FeedParser *gofeed.Parser
	Logger     *utils.Logger
}

func NewPaperRepository(parser *gofeed.Parser, logger *utils.Logger) *PaperRepository {
	return &PaperRepository{
		FeedParser: parser,
		Logger:     logger,
	}
}

// fetchRSS関数
func (r *PaperRepository) FetchRSS(url string) ([]*gofeed.Item, error) {
	feed, err := r.FeedParser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}
