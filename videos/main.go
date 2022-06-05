package main

import "fmt"

func main() {
	videos := getVideos()

	for i := range videos {
		videos[i].Description = videos[i].Description + "\n go f yourself"
	}

	fmt.Println(videos)

	saveVideos(videos)
}
