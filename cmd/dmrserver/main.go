package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/OnChainTemplars/dmrserver/internal/config"
	"github.com/OnChainTemplars/dmrserver/internal/handlers"
	"github.com/pd0mz/go-dmr/homebrew"
)

func contextMiddleware(next http.Handler, h *homebrew.Homebrew) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "homebrew", h)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func httpServer(h *homebrew.Homebrew, c *config.Config) {
	mux := http.NewServeMux()

	mux.HandleFunc("/peers", handlers.PeersHandler)
	mux.HandleFunc("/peer/link", handlers.LinkHandler)
	mux.HandleFunc("/", handlers.IndexHandler)

	addr := fmt.Sprintf("%s:%d", c.IP, c.HTTP.Port)

	log.Printf("HTTP Server is running at %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, contextMiddleware(mux, h)))
}

func main() {
	config := config.Parse("./config.toml")

	repeaterConfig := homebrew.RepeaterConfiguration{
		Callsign: config.Repeater.Callsign,
	}

	addr := net.UDPAddr{
		IP:   net.ParseIP(config.IP),
		Port: config.Repeater.Port,
	}

	h, err := homebrew.New(&repeaterConfig, &addr)
	if err != nil {
		log.Fatalln(err)
	}

	go httpServer(h, config)

	log.Printf("Homebrew repeater is running at %s\n", addr.String())

	log.Fatal(h.ListenAndServe())
}
