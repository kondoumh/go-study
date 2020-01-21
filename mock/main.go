package main

import (
	"context"
	"fmt"
)

type GitHub interface {
	CreateRelease(ctx context.Context, opt *Option) (string, error)
	GetRelease(ctx context.Context, tag string) (string, error)
	DeleteRelease(ctx context.Context, relaeaseID int) error
}

type GhRelease struct {
	c GitHub
}

func (ghr *GhRelease) CreateNewRelease(ctx context.Context) (*Release, error) {
	tag, err := ghr.c.CreateRelease(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create release: %v", err)
	}

	if _, err := ghr.c.GetRelease(ctx, tag); err != nil {
		return nil, fmt.Errorf("failed to get created release: %v", err)
	}

	return &Release{}, nil
}

type Option struct{}
type Release struct{}
