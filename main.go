package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	endpoint := "https://gg.protoku.io/rng/c29d7b9f-114f-475b-b8db-946d7d96e1dd"

	cmd := exec.Command("git", "diff", "--staged", "--", ":!package-lock.json", ":!yarn.lock")

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	data := map[string]string{"payload": string(output)}

	jsonData, err := json.Marshal(data)
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

	fmt.Println(response.Message)
}
