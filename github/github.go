package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	fmt.Println(githubInfo(ctx, "weitecklee"))
}

// func demo() {
// 	resp, err := http.Get("https://api.github.com/users/weitecklee")
// 	if err != nil {
// 		log.Fatalf("error: %s", err)
// 		/*
// 			log.Printf("error: %s", err)
// 			os.Exit(1)
// 		*/
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalf("error: %s", resp.Status)
// 	}
// 	fmt.Printf("Content-Type : %s\n", resp.Header.Get("Content-Type"))
// 	/*
// 		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
// 			log.Fatalf("error: can't copy - %s", err)
// 		}
// 	*/
// 	var r Reply
// 	dec := json.NewDecoder(resp.Body)
// 	if err := dec.Decode(&r); err != nil {
// 		log.Fatalf("error: can't decode - %s", err)
// 	}
// 	// fmt.Println(r)
// 	fmt.Printf("%#v\n", r)
// }

func githubInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	// resp, err := http.Get(url)
	// NOTE: always use context when making network requests
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}
	defer resp.Body.Close()
	// var r Reply
	var r struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.NumRepos, nil
}

// type Reply struct {
// 	Name string
// 	// Public_Repos int
// 	NumRepos int `json:"public_repos"`
// }
