package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// videos get subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs of get command
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "Youtube video ID")

	// videos add subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// inputs of add command
	addID := addCmd.String("id", "", "video id")
	addTitle := addCmd.String("title", "", "video title")
	addUrl := addCmd.String("url", "", "video url")
	addImageUrl := addCmd.String("imageurl", "", "video imageurl")
	addDesc := addCmd.String("description", "", "video description")

	if len(os.Args) < 2 {
		fmt.Println("insufficient number of arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getID)
	case "add":
		HandleAdd(addCmd, addID, addTitle, addUrl, addImageUrl, addDesc)
	default:
		fmt.Println("not a valid argument")
		os.Exit(1)
	}
}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		videos := getVideos()
		fmt.Printf("ID \t Title \t URL \t Image \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}
		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t Image \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
			}
		}
		return
	}

}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, description *string) {
	if *id == "" || *title == "" || *url == "" || *imageurl == "" || *description == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, description *string) {
	addCmd.Parse(os.Args[2:])
	ValidateVideo(addCmd, id, title, url, imageurl, description)

	video := video{
		Id:          *id,
		Title:       *title,
		Description: *description,
		Imageurl:    *imageurl,
		Url:         *url,
	}

	videos := getVideos()
	videos = append(videos, video)
	saveVideos(videos)
}
