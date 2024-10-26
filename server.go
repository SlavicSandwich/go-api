package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodGet:
		p.ProcessScore(w, r, player)
	case http.MethodPost:
		p.ProcessWin(w, r, player)
	}
}

func (p *PlayerServer) ProcessScore(w http.ResponseWriter, r *http.Request, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) ProcessWin(w http.ResponseWriter, r *http.Request, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
