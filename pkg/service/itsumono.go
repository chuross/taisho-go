package service

import (
	"context"
	"math/rand"

	"golang.org/x/xerrors"
)

func Itsumono(ctx context.Context) (*string, error) {
	res, err := Search(ctx, "寿司", SearchTypeImage)
	if err != nil {
		return nil, xerrors.Errorf("itsumono failed: %w", err)
	}
	if len(res.Links) > 0 {
		return res.Links[rand.Intn(len(res.Links))], nil
	}
	return nil, nil
}
