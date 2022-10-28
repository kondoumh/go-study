package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("must pass in a git repo to build")
		os.Exit(1)
	}
	repo := os.Args[1]
	if err := build(repo); err != nil {
		fmt.Println(err)
	}
}

func build(repoUrl string) error {
	fmt.Printf("Building %s\n", repoUrl)

	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	repo := client.Git(repoUrl)
	src, err := repo.Branch("main").Tree().ID(ctx)
	if err != nil {
		return err
	}

	workdir := client.Host().Workdir()

	golang := client.Container().From("golang:latest")
	golang = golang.WithMountedDirectory("src", src).WithWorkdir("src")


	oses := []string{"linux", "darwin"}
	arches := []string{"amd64", "arm64"}

	for _, goos := range oses {
		for _, goarch := range arches {
			path := fmt.Sprintf("build/%s/%s/", goos, goarch)
			outpath := filepath.Join(".", path)
			err = os.MkdirAll(outpath, os.ModePerm)
			if err != nil {
				return err
			}
			build := golang.WithEnvVariable("GOOS", goos)
			build = build.WithEnvVariable("GOARCH", goarch)
			build = build.Exec(dagger.ContainerExecOpts{
				Args: []string{"go", "build", "-o", path},
			})
			output, err := build.Directory(path).ID(ctx)
			if err != nil {
				return err
			}
		
			_, err = workdir.Write(ctx, output, dagger.HostDirectoryWriteOpts{Path: path})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
