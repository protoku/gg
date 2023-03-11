package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	Payload string `json:"payload"`
}

type Response struct {
	Message string `json:"message"`
}

func GetCommitMessage(diff string) string {
	endpoint := "https://gg.protoku.io/rng/c29d7b9f-114f-475b-b8db-946d7d96e1dd"

	request := Request{Payload: diff}

	jsonData, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return strings.ReplaceAll(response.Message, `"`, "")
}
