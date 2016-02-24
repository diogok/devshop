package main

import (
  "github.com/google/go-github/github"
	"math"
)

type Developer struct {
	Name *string `json:"name,omitempty"`
	ID *int `json:"id"`
	Login *string `json:"login"`
	Cost int `json:"cost"`
}

type Developers []Developer

type Repository struct {
	Name *string
	Stars *int
	Forks *int
	Watchers *int
}

type Repositories []Repository

func repositories(user string)  (Repositories,error) {
	client := github.NewClient(nil)

	opts := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 99, Page: 0},
	}

	result, _,err := client.Repositories.List(user,opts)

	if err != nil {
		return nil, err
	} else {
		repos := Repositories{}
		for _, repo := range result {
			repo := Repository{
				Name: repo.Name,
				Forks: repo.ForksCount,
				Stars: repo.StargazersCount,
				Watchers: repo.WatchersCount,
			}
			repos = append(repos,repo)
		}
		return repos, nil
	}
}

func calcCost(repos Repositories) int {
	cost := 0
	for _, repo := range repos {
		cost = cost + (1 + ( *repo.Stars * 2) + ( *repo.Watchers * 3 ) + ( *repo.Forks * 4 ))
	}
	return int( math.Floor( float64( cost ) * 0.1 ) )
}

func developers(query string, page int) (Developers, error) {
	client := github.NewClient(nil)

	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 10, Page: page},
	}

	rdevs, _,err := client.Search.Users(query,opts)

	if err != nil {
		return nil, err
	} else {
		devs := Developers{}
		for _, user := range rdevs.Users {
			repos,_ := repositories(*user.Login)
			cost := calcCost(repos)
			dev := Developer{
				Login: user.Login,
				Name: user.Name,
				ID: user.ID,
				Cost: cost,
			}
			devs = append(devs,dev)
		}
		return devs, nil
	}
}

