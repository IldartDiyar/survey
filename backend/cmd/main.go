package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"survey/config"
	"survey/internal/infra"
	"survey/internal/repo"
	"survey/internal/service"
)

func main() {
	wordPtr := flag.String("config", "./config/config.yml", "-config=./config/config.yml")
	flag.Parse()

	config, err := config.LoadConfig(wordPtr)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	infra, err := infra.New(config)
	if err != nil {
		log.Fatalf("Error initializing infrastructure: %v", err)
	}

	repo := repo.New(infra)

	s := service.New(repo)

	h := http.NewServeMux()

	h.HandleFunc("/survey", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			resp, err := s.SurveyService.GetAnswers()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			bb, err := json.Marshal(resp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write(bb)
		case http.MethodPost:
			bb, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			var req map[string]string

			if err := json.Unmarshal(bb, req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			if err := s.SurveyService.ProcessSurveyResponse(req); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}

			w.WriteHeader(200)
		default:
		}

	})

	log.Fatal(http.ListenAndServe(":8089", h))
}
