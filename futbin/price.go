package futbin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// GetPrice returns player's price listed on futbin.
func GetPrice(playerId string) (int, error) {
	raw, err := fetch(playerId)
	if err != nil {
		return 0, err
	}

	return extractPrice(raw, playerId)
}

func fetch(playerId string) ([]byte, error) {
	url := fmt.Sprintf("https://www.futbin.com/19/playerPrices?player=%s", playerId)

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

func extractPrice(raw []byte, playerId string) (int, error) {
	// FIXME: find a way to replace interface{} with structs
	// so far, the first level are player ids, which cannot be hardcoded to golang structs
	var j map[string]interface{}
	if err := json.Unmarshal(raw, &j); err != nil {
		return 0, err
	}

	// TODO: handle type assertion errors
	players := j[playerId].(map[string]interface{})
	prices := players["prices"].(map[string]interface{})
	ps := prices["ps"].(map[string]interface{})

	lastPsPrice := ps["LCPrice"].(string)
	lastPsPrice = strings.Replace(lastPsPrice, ",", "", -1)

	return strconv.Atoi(lastPsPrice)
}
