package sofascore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type (
	JsonObj struct {
		Tournaments []Tournament `json:"tournaments"`
	}
	Tournament struct {
		HasEventPlayerStatistics bool    `json:"hasEventPlayerStatistics"`
		Events                   []Event `json:"events"`
	}
	Event struct {
		PlayerMatchInfo PlayerMatchInfo `json:"playerMatchInfo"`
		StartTimestamp  int64           `json:"startTimestamp"`
	}
	PlayerMatchInfo struct {
		Rating string `json:"rating"`
	}
)

func GetRatings(playerId string) ([]float64, error) {
	raw, err := fetch(playerId)
	if err != nil {
		return nil, err
	}

	jsonObj, err := toJson(raw)
	if err != nil {
		return nil, err
	}

	return filterRatings(jsonObj)
}

func fetch(playerId string) ([]byte, error) {
	url := fmt.Sprintf("https://www.sofascore.com/player/%s/events/json", playerId)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func toJson(raw []byte) (JsonObj, error) {
	var jsonObj JsonObj
	if err := json.Unmarshal(raw, &jsonObj); err != nil {
		return JsonObj{}, err
	}

	return jsonObj, nil
}

func filterRatings(jsonObj JsonObj) ([]float64, error) {
	weekAgo := time.Now().Add(-7 * 24 * time.Hour)
	ratings := []float64{}

	for _, t := range jsonObj.Tournaments {
		if !t.HasEventPlayerStatistics {
			continue
		}

		for _, e := range t.Events {
			ts := time.Unix(e.StartTimestamp, 0)
			if ts.Before(weekAgo) {
				continue
			}

			r, err := strconv.ParseFloat(e.PlayerMatchInfo.Rating, 64)
			if err != nil {
				return nil, err
			}

			ratings = append(ratings, r)
		}
	}

	return ratings, nil
}
