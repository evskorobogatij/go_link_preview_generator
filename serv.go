package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GenServer struct{}

func NewGetServer() *GenServer {
	tg := &GenServer{}
	return tg
}

func (tg *GenServer) generateHandler(w http.ResponseWriter, req *http.Request) {
	type GenerateRequest struct {
		Url string `json:"url"`
	}

	if req.Method == http.MethodPost {
		//
		dec := json.NewDecoder(req.Body)
		dec.DisallowUnknownFields()
		var data GenerateRequest
		if err := dec.Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("%s %s \n", time.Now().Format(time.RFC3339), data.Url)
		// preview := PreviewInfo{Title: "site", Description: "info for site", PreviewUrl: "http://image"}
		var preview PreviewInfo
		err := GeneratePreview(data.Url, &preview)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		js, err := json.Marshal(preview)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

		// json.Marshal()
	} else {
		http.Error(w, fmt.Sprintf("expect method POST at /generate_preview, got %v", req.Method), http.StatusMethodNotAllowed)
		return
	}
}
