package main

import (
	"testing"
	"fmt"
)

func TestRepositories(t *testing.T) {
	repos, err := repositories("diogok")
	if err != nil{
		t.Error("Listing repos error")
		fmt.Printf("%+v\n", err)
	} else if repos == nil {
		t.Error("Repos is nil")
	} else if len(repos) != 50 {
		t.Error("Wrong number of repos")
	} else {
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
}

func TestCost(t *testing.T) {
	repos, _ := repositories("diogok")
	cost := calcCost(repos)
	if cost != 55 {
		t.Error("Wrong cost")
		fmt.Printf("%d",cost)
	}
}

func TestDevelopers(t *testing.T) {
	devs, err := developers("diogo",0)
	if err != nil{
		t.Error("Listing developers error")
	} else if devs == nil {
		t.Error("Devs is nil")
	} else if len(devs) != 10 {
		t.Error("Wrong number of devs")
	} else {
		devs, err = developers("diogok",0)
		if devs[0].Cost != 55 {
			t.Error("Problem calc cost in developers")
			fmt.Printf("%+v\n", devs[0])
		}
	}
}

