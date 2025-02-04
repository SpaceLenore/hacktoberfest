package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bmizerany/pat"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Submission struct {
	ID        int       `json:"pull_request_number"`
	Title     string    `json:"title"`
	Login     string    `json:"user"`
	State     string    `json:"state"`
	RepoName  string    `json:"repo_name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type Submissions struct {
	Total        int          `json:"count"`
	PullRequests []Submission `json:"pull_requests"`
}

var ghClient *github.Client

func commonHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-Content-Type-Options", "no-sniff")
		h.ServeHTTP(w, r)
	}
}

func logHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		h.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(map[string]string{"status": "ok"})
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func PullRequestHandler(w http.ResponseWriter, r *http.Request) {
	var s Submissions
	ghUser := r.URL.Query().Get(":ghUser")
	_, _, err := ghClient.Users.Get(oauth2.NoContext, ghUser)
	if err != nil {
		resp, _ := json.Marshal(map[string]string{"error": "User not found"})
		w.WriteHeader(http.StatusNotFound)
		w.Write(resp)
		return
	}
	opt := &github.ListOptions{PerPage: 100}
	events, _, err := ghClient.Activity.ListEventsPerformedByUser(oauth2.NoContext, ghUser, true, opt)
	if err != nil {
		fmt.Println(err)
	}
	for _, event := range events {
		repo_name := *event.Repo.Name
		if *event.Type == "PullRequestEvent" && inTimeSpan(*event.CreatedAt) {
			pull_request := event.Payload().(*github.PullRequestEvent).PullRequest
			submission := Submission{ID: *pull_request.Number, Title: *pull_request.Title, Login: *pull_request.User.Login, URL: *pull_request.HTMLURL, State: *pull_request.State, RepoName: repo_name, CreatedAt: *pull_request.CreatedAt}
			if submission.State == "open" {
				s.PullRequests = append(s.PullRequests, submission)
			}
		}
	}
	s.Total = len(s.PullRequests)
	resp, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func inTimeSpan(check time.Time) bool {
	return check.Month() == time.October
}

func portNumber() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return port
}

func init() {
	ghToken := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_API_TOKEN")})
	gh_oauthClient := oauth2.NewClient(oauth2.NoContext, ghToken)
	ghClient = github.NewClient(gh_oauthClient)
}

func main() {
	m := pat.New()
	m.Get("/api/:ghUser", http.HandlerFunc(PullRequestHandler))
	m.Get("/api/", http.HandlerFunc(StatusHandler))

	middleware := commonHeaders(logHandler(m))
	http.Handle("/api/", middleware)

	fs := http.FileServer(http.Dir("static/dist"))
	http.Handle("/", fs)

	http.ListenAndServe(":"+portNumber(), nil)

}
