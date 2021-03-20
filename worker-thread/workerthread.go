package worker-thread

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var workerPoolSize = 3
var pages = make(chan int, workerPoolSize)
var teamAppearances = make(map[string]int)

const (
	competition = "UEFA%20Champions%20League"
	year        = 2011
)

type ResponseData struct {
	Page            int             `json:"page"`
	PerPage         int             `json:"per_page"`
	Total           int             `json:"total"`
	TotalPages      int             `json:"total_pages"`
	FootballMatches []FootballMatch `json:"data"`
}

type FootballMatch struct {
	Competition string `json:"competition"`
	Year        int    `json:"year"`
	Round       string `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Team1Goals  string `json:"team1goals"`
	Team2Goals  string `json:"team2goals"`
}

func AddTeamToMap(matches []FootballMatch) {
	for _, match := range matches {
		_, isTeam1 := teamAppearances[match.Team1]
		if isTeam1 {
			val := teamAppearances[match.Team1]
			teamAppearances[match.Team1] = val + 1
		} else {
			teamAppearances[match.Team1] = 1
		}

		_, isTeam2 := teamAppearances[match.Team2]
		if isTeam2 {
			val := teamAppearances[match.Team2]
			teamAppearances[match.Team2] = val + 1
		} else {
			teamAppearances[match.Team2] = 1
		}
	}
}

//GetFootballData get football data
func GetFootballData(competition string, year int, page int) ResponseData {
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?competition=%s&year=%d&page=%d", competition, year, page)
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseData ResponseData
	json.Unmarshal(body, &responseData)

	return responseData
}

//AllocateJobs - create jobs to be done - sender
func AllocateJobs(totalPages int) {
	for i := 2; i <= totalPages; i++ {
		pages <- i
	}

	close(pages)
}

//create workers 
func CreateWorkerThread(noOfWorkers int) map[string]int {
	var wg sync.WaitGroup
	for i := 1; i <= noOfWorkers; i++ {
		//means add the number of worker semaphore
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()

	return teamAppearances
}

//receive jobs
func worker(wg *sync.WaitGroup) {
	for page := range pages {
		resp := GetFootballData(competition, year, page)
		AddTeamToMap(resp.FootballMatches)
	}
	wg.Done()
}
