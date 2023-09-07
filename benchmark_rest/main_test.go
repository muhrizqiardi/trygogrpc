package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

type Input struct {
	Input int `json:"input"`
}

type Output struct {
	Result int `json:"result"`
}

func BenchmarkFactorialHandler(b *testing.B) {
	requestBody := Input{
		Input: 5,
	}
	requestBytes, _ := json.Marshal(requestBody)

	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest("POST", "http://localhost:9000/factorial", strings.NewReader(string(requestBytes)))
		if err != nil {
			fmt.Println("err:", err)
			b.Error("exp nil, got:", err)
		}
		resp, err := http.DefaultClient.Do(req)

		var result Output
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			b.Error("exp nil, got:", err)
		}

		fmt.Println("output:", result.Result)
	}
}
