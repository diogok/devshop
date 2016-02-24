package main

import (
	"testing"
)

func TestRepositories(t *testing.T) {
	repos, err := repositories("diogok")
	if err != nil{
		t.Error("Listing repos error")
	}
	if repos == nil {
		t.Error("Repos is nil")
	}
	if len(repos) != 50 {
		t.Error("Wrong number of repos")
	}
	found := false
	for _, repo := range repos {
		if *repo.Name == "restserver" {
			found = true
			if *repo.Stars != 31 {
				t.Error("Wrong star count")
			}
			if *repo.Forks != 9 {
				t.Error("Wrong fork count")
			}
		}
	}

	if found == false {
		t.Error("Repo not found")
	}
}

func TestDevelopers(t *testing.T) {
	devs, err := developers("diogo",0)
	if err != nil{
		t.Error("Listing developers error")
	}
	if devs == nil {
		t.Error("Devs is nil")
	}
	if len(devs) != 10 {
		t.Error("Wrong number of devs")
	}
}

