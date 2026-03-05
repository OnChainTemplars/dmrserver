package handlers

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/pd0mz/go-dmr/homebrew"
)

func LinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /peer/link")

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	h := r.Context().Value("homebrew").(*homebrew.Homebrew)

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	port, err := strconv.Atoi(r.FormValue("port"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	addr := net.UDPAddr{
		IP:   net.ParseIP(r.FormValue("ip")),
		Port: port,
	}

	peer := &homebrew.Peer{
		ID:       uint32(id),
		Addr:     &addr,
		AuthKey:  []byte(r.FormValue("auth_key")),
		Incoming: true,
		Status:   homebrew.AuthNone,
	}

	h.Link(peer)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
