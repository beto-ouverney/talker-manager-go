package test

import (
	"os"
	"testing"
)

func seedTalkers(t *testing.T) {
	jsonFile, err := os.ReadFile("./seed.json")
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./talkers.json")
	if err == nil {
		_, err = f.Write(jsonFile)
	}
	defer f.Close()
}

type Talk struct {
	WatchedAt string `json:"watchedAt,omitempty"`
	Rate      int    `json:"rate,omitempty"`
}

type Talker struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	ID   int    `json:"id,omitempty"`
	Talk Talk   `json:"talk,omitempty"`
}

type TestError struct {
	Message string `json:"message,omitempty"`
}
