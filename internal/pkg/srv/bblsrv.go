package srv

import (
	"fmt"
	"log"
	"net/http"
)

type Chatroom struct {
	Port   string
	MaxUsr int
	MaxMsg int
	MaxCh  int
}

func (cr *Chatroom) String() string {
	return fmt.Sprintf("port: %20s\nmax users: %13d\nmax messages: %11d\nmax characters: %10d\n", cr.Port, cr.MaxUsr, cr.MaxMsg, cr.MaxCh)
}

func (cr *Chatroom) history(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "history")
}

func (cr *Chatroom) send(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "send")
}

func (cr *Chatroom) subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "subscribe")
}

func (cr *Chatroom) Run() {
	http.HandleFunc("/history", cr.history)
	http.HandleFunc("/send", cr.send)
	http.HandleFunc("/subscribe", cr.subscribe)
	fmt.Printf("Chatroom up\n%s", cr)
	log.Fatal(http.ListenAndServe(":"+cr.Port, nil))
}
