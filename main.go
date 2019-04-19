package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	s := server{
		apiKey:    os.Getenv("TOKEN"),
		userAgent: os.Getenv("USER_AGENT"),
	}
	fmt.Println(s.apiKey)
	fmt.Println(s.userAgent)

	http.HandleFunc("/hello-world", s.helloWorld)
	http.HandleFunc("/", s.chrisHandler)
	http.HandleFunc("/CORK", s.ryanHandler)
	http.ListenAndServe(":8080", nil)
}

type server struct {
	apiKey    string
	userAgent string
}

func (s *server) helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusAccepted)

	if _, err := w.Write([]byte("hello max")); err != nil {
		println(err)
	}

}

func (s *server) chrisHandler(w http.ResponseWriter, r *http.Request) {
	leagueStandings := getFantasyChris()
	mlb, err := getMLBAPI(s.apiKey, s.userAgent)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	populateScores(leagueStandings, mlb)
	leagueStandings.Rank()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	render(leagueStandings, w)
}

func (s *server) ryanHandler(w http.ResponseWriter, r *http.Request) {
	leagueStandings := getFantasyRyan()
	mlb, err := getMLBAPI(s.apiKey, s.userAgent)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	populateScores(leagueStandings, mlb)
	leagueStandings.Rank()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	render(leagueStandings, w)
}
