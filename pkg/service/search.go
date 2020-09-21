package service

import (
	"context"
	"os"

	"github.com/chuross/taisho/pkg/model/search"
	"golang.org/x/xerrors"
	"google.golang.org/api/customsearch/v1"
)

const (
	SearchTypeImage = "image"
)

type SearchType string

func Search(ctx context.Context, keyword string, searchType SearchType) (*search.Result, error) {

	client, err := customsearch.NewService(ctx)
	if err != nil {
		return nil, xerrors.Errorf("customsearch init failed: %w", err)
	}

	req := client.Cse.List().
		Q(keyword).
		Cx(os.Getenv("TAISHO_SEARCH_ID")).
		SearchType(string(searchType))

	call, err := req.Do()
	if err != nil {
		return nil, xerrors.Errorf("image search failed: %w", err)
	}

	links := make([]*string, 0)
	for _, item := range call.Items {
		links = append(links, &item.Link)
	}

	return &search.Result{
		Links: links,
	}, nil
}
