// llama.go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type LlamaRequest struct {
    Prompt string `json:"prompt"`
}

type LlamaResponse struct {
    Move string `json:"move"`
}

func getLlamaMove(boardState string) (string, error) {
    url := "http://localhost:8000/getMove" // URL of the LLaMA API
    prompt := "Suggest a chess move based on this board state: " + boardState
    
    requestBody, err := json.Marshal(LlamaRequest{Prompt: prompt})
    if err != nil {
        return "", err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var llamaResponse LlamaResponse
    if err := json.NewDecoder(resp.Body).Decode(&llamaResponse); err != nil {
        return "", err
    }

    return llamaResponse.Move, nil
}

