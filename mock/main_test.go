package main

import (
	"context"
	"fmt"
	"testing"
)

type fakeGitHub struct {
	GitHub
	FakeCreateRelease func(ctx context.Context, opt *Option) (string, error)
	FakeGetRelease    func(ctx context.Context, tag string) (string, error)
}

func (c *fakeGitHub) CreateRelease(ctx context.Context, opt *Option) (string, error) {
	return c.FakeCreateRelease(ctx, opt)
}

func (c *fakeGitHub) GetRelease(ctx context.Context, tag string) (string, error) {
	return c.FakeGetRelease(ctx, tag)
}

func TestGhRelease_CreateNewRelease(t *testing.T) {
	fakeclient := &fakeGitHub{
		FakeCreateRelease: func(ctx context.Context, opt *Option) (string, error) {
			return "v1.0", nil
		},
		FakeGetRelease: func(ctx context.Context, tag string) (string, error) {
			return "", fmt.Errorf("faild to get %v release!", tag)
		},
	}

	ghr := &GhRelease{c: fakeclient}

	release, err := ghr.CreateNewRelease(context.Background())
	if err != nil {
		t.Error(err)
	}
	_ = release
}
