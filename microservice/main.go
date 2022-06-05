package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetVideos)
	http.ListenAndServe(":8080", nil)
}

func GetVideos(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		videos := getVideos()

		videoBytes, err := json.Marshal(videos)

		if err != nil {
			panic(err)
		}

		w.Write(videoBytes)
	} else if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		var videos []video
		err = json.Unmarshal(body, &videos)

		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Bad request")
		}

		saveVideos(videos)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Internal Server Error")
		}
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not Supported")
	}
}
