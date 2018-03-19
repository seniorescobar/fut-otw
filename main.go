package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// https://www.easports.com/fifa/ultimate-team/api/fut/item?jsonParamObject={"page":1,"quality":"marquee","position":"LF,CF,RF,ST,LW,LM,CAM,CDM,CM,RM,RW,LWB,LB,CB,RB,RWB"}

type jsonParamObject struct {
	Page     int    `json:"page"`
	Quality  string `json:"quality"`
	Position string `json:"position"`
}

type futPlayerItemList struct {
	Count        int
	Items        []futItem
	Page         int
	TotalPages   int
	TotalResults int
	Type         string
}

type futItem map[string]interface{}

func main() {
	items, err := getPTGItems()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(items))

}

func getPTGItems() ([]futItem, error) {
	page := 1

	var items []futItem

	for {
		tmpURL, err := getURL(page)
		if err != nil {
			return nil, err
		}

		resp, err := http.Get(tmpURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var futObj futPlayerItemList
		json.Unmarshal(body, &futObj)

		items = append(items, futObj.Items...)

		if futObj.Page == futObj.TotalPages {
			break
		}
		page++
	}

	return items, nil
}

func getURL(page int) (string, error) {
	tmpURL, err := url.Parse("https://www.easports.com/fifa/ultimate-team/api/fut/item")
	if err != nil {
		return "", err
	}

	jsonParams := jsonParamObject{
		Page:     page,
		Quality:  "marquee",
		Position: strings.Join([]string{"LF", "CF", "RF", "ST", "LW", "LM", "CAM", "CDM", "CM", "RM", "RW", "LWB", "LB", "CB", "RB", "RWB", "GK"}, ","),
	}
	jsonParamsStr, err := json.Marshal(jsonParams)
	if err != nil {
		return "", err
	}

	v := url.Values{}
	v.Set("jsonParamObject", string(jsonParamsStr))

	tmpURL.RawQuery = v.Encode()

	return tmpURL.String(), nil
}
