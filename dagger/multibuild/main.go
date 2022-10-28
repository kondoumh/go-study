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

	// 1. Get a context
	ctx := context.Background()
	// 2. Initialize dagger client
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	// 3. Clone the repo using Dagger
	repo := client.Git(repoUrl)
	src, err := repo.Branch("main").Tree().ID(ctx)
	if err != nil {
		return err
	}
	// 4. Load the golang image
	workdir := client.Host().Workdir()
	golang := client.Container().From("golang:latest")
	// 5. Mount the cloned repo to the golang image
	golang = golang.WithMountedDirectory("src", src).WithWorkdir("src")
	// 6. Create the output path on the host for the build
	path := "build/"
	outpath := filepath.Join(".", path)
	err = os.MkdirAll(outpath, os.ModePerm)
	if err != nil {
		return err
	}
	// 7. Do th go build
	golang = golang.Exec(dagger.ContainerExecOpts{
		Args: []string{"go", "build", "-o", "build/"},
	})

	// 8. Get build output from builder
	output, err := golang.Directory(path).ID(ctx)
	if err != nil {
		return err
	}

	// 9. Write the build output to the host
	_, err = workdir.Write(ctx, output, dagger.HostDirectoryWriteOpts{Path: path})
	if err != nil {
		return err
	}
	return nil
}
