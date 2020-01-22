package main

import (
	"context"
	"fmt"
)

type Release struct{}
type Option struct{}

type GitHub interface {
	CreateRelease(ctx context.Context, opt *Option) (string, error)
	GetRelease(ctx context.Context, tag string) (string, error)
	DeleteRelease(ctx context.Context, relaeseID int) error
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

func (ghr *GhRelease) GetRelease(ctx context.Context, tag string) (string, error) {
	version, err := ghr.c.GetRelease(ctx, tag)
	if err != nil {
		return "", fmt.Errorf("failed to get release: %v", err)
	}
	return version, nil
}

func (ghr *GhRelease) DeleteRelease(ctx context.Context, releaseID int) error {
	if err := ghr.c.DeleteRelease(ctx, releaseID); err != nil {
		return fmt.Errorf("failed to delete release: %v", err)
	}
	return nil
}