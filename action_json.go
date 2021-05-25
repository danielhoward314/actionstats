package main

import (
	"encoding/json"
	"log"
)

type ActionMsg struct {
	Action string `json:"action"`
	Time   int    `json:"time"`
}

func FromJSON(s string) (a ActionMsg, e error) {
	b := []byte(s)
	var aMsg ActionMsg
	err := json.Unmarshal(b, &aMsg)
	return aMsg, err
}

type ActionAvg struct {
	Action string `json:"action"`
	Avg    int    `json:"avg"`
}

type StatsMsg []*ActionAvg

func ToJSON(avgs []map[string]int) string {
	var jsonData []byte
	jsonData, err := json.Marshal(avgs)
	if err != nil {
		log.Println(err)
	}
	return string(jsonData)
}
