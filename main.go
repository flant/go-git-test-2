package main

import (
	"fmt"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	repository, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}

	fromHash := plumbing.NewHash("c73fd0c6d5669affde76a78c60e4cc372d39abc1")
	fromCommitObj, err := repository.CommitObject(fromHash)
	if err != nil {
		panic(err)
	}

	toHash := plumbing.NewHash("1fa54caeab28ce28d5aa9d0eef9c72c734218cfe")
	toCommitObj, err := repository.CommitObject(toHash)
	if err != nil {
		panic(err)
	}

	originPatch, err := fromCommitObj.Patch(toCommitObj)
	if err != nil {
		panic(err)
	}

	// diff.DefaultContextLines = 0
	//_ = originPatch.String()
	fmt.Printf("%s", originPatch.String())
}
