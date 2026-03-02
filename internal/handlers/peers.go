package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pd0mz/go-dmr/homebrew"
)

type Peer struct {
	ID     uint32 `json:"id"`
	Addr   string `json:"addr"`
	Status string `json:"status"`
}

func PeersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /peers")

	h := r.Context().Value("homebrew").(*homebrew.Homebrew)

	var peers = make([]*Peer, 0)
	for _, p := range h.Peer {
		peer := &Peer{
			ID:     p.ID,
			Addr:   p.Addr.String(),
			Status: p.Status.String(),
		}

		peers = append(peers, peer)
	}

	data := struct {
		Peers []*Peer `json:"peers"`
	}{
		Peers: peers,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
