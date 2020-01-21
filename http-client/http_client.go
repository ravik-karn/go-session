package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name    string
	Company string
}

var user User

func HTTPClient(url string) (*User, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %s", err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read request erroe: %s", err)
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		return nil, fmt.Errorf("umnarshal error: %s", err)
	}

	return &user, nil
}
