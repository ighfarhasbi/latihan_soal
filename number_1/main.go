package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	teamName1 := "Barcelona"
	result, err := sumData(teamName1)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Team Name: %s, Total Goals: %d", teamName1, result)

}

type MatchData struct {
	Competition string `json:"competition"`
	Year        int    `json:"year"`
	Round       string `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Team1Goals  string `json:"team1goals"`
	Team2Goals  string `json:"team2goals"`
}

type Response struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
	Data       []MatchData `json:"data"`
}

func sumData(teamName1 string) (int, error) {
	encodeTeamName := url.QueryEscape(teamName1) // untuk encode spasi ke %20
	url1 := "https://jsonmock.hackerrank.com/api/football_matches?year=2011&team1=" + encodeTeamName + "&page=%d"
	url2 := "https://jsonmock.hackerrank.com/api/football_matches?year=2011&team2=" + encodeTeamName + "&page=%d"

	firstUrl := fmt.Sprintf(url1, 1)
	firstUrl2 := fmt.Sprintf(url2, 1)
	result, err := getData(firstUrl)
	if err != nil {
		return 0, err
	}
	result2, err := getData(firstUrl2)
	if err != nil {
		return 0, err
	}

	totalPage := result.TotalPages
	totalPage2 := result2.TotalPages
	sum := 0

	for page := 1; page <= totalPage; page++ {
		url := fmt.Sprintf(url1, page)
		result, err := getData(url)
		if err != nil {
			return 0, err
		}
		for _, match := range result.Data {
			goalInt, err := strconv.Atoi(match.Team1Goals)
			if err != nil {
				return 0, err
			}
			sum += goalInt
			// log.Printf("Team1Goals: %s, Team2Goals: %s",
			// 	match.Team1Goals, match.Team2Goals)
		}
	}

	for page := 1; page <= totalPage2; page++ {
		url := fmt.Sprintf(url2, page)
		result, err := getData(url)
		if err != nil {
			return 0, err
		}
		for _, match := range result.Data {
			goalInt, err := strconv.Atoi(match.Team1Goals)
			if err != nil {
				return 0, err
			}
			sum += goalInt
		}
	}

	return sum, nil
}

func getData(url string) (Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Response{}, err
	}

	return result, nil
}
