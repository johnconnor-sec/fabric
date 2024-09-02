package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	apiKey             = "YOUR_YOUTUBE_API_KEY_HERE"
	obsidianVaultPath  = "/home/john/Documents/Obsidian_Vault/"
)

var playlistIDs = map[string]string{
	"PL-qEoXFQG4wuaN4huGKJTeORyGOrN26Yx": filepath.Join(obsidianVaultPath, "4_Reference Notes/Obsidian Videos.md"),
}

type Video struct {
	PlaylistID string
	Markdown   string
}

func getLatestVideos(service *youtube.Service) ([]Video, error) {
	var latestVideos []Video

	for playlistID := range playlistIDs {
		nextPageToken := ""
		for {
			call := service.PlaylistItems.List([]string{"snippet"}).
				MaxResults(50).
				PlaylistId(playlistID)

			if nextPageToken != "" {
				call = call.PageToken(nextPageToken)
			}

			response, err := call.Do()
			if err != nil {
				log.Printf("Error fetching videos for playlist %s: %v", playlistID, err)
				break
			}

			for _, item := range response.Items {
				videoID := item.Snippet.ResourceId.VideoId
				videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
				videoTitle := item.Snippet.Title
				markdown := fmt.Sprintf("![%s](%s)", videoTitle, videoURL)
				latestVideos = append(latestVideos, Video{PlaylistID: playlistID, Markdown: markdown})
			}

			nextPageToken = response.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
	}

	return latestVideos, nil
}

func updateObsidianNotes(videos []Video) error {
	videosByPlaylist := make(map[string][]string)
	for _, video := range videos {
		videosByPlaylist[video.PlaylistID] = append(videosByPlaylist[video.PlaylistID], video.Markdown)
	}

	for playlistID, videoList := range videosByPlaylist {
		filePath := playlistIDs[playlistID]
		cachedLinks := make(map[string]bool)

		content, err := os.ReadFile(filePath)
		if err == nil {
			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				cachedLinks[strings.TrimSpace(line)] = true
			}
		}

		var newVideos []string
		for _, videoMD := range videoList {
			if !cachedLinks[videoMD] {
				newVideos = append(newVideos, videoMD)
			}
		}

		if len(newVideos) > 0 {
			f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				return fmt.Errorf("error opening file %s: %v", filePath, err)
			}
			defer f.Close()

			for _, videoMD := range newVideos {
				if _, err = f.WriteString(videoMD + "\n\n"); err != nil {
					return fmt.Errorf("error writing to file %s: %v", filePath, err)
				}
			}
			log.Printf("Added %d new videos to %s.", len(newVideos), filePath)
		} else {
			log.Printf("No new videos to add for playlist %s", playlistID)
		}
	}

	return nil
}

func main() {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	log.Println("Starting to fetch latest videos.")
	latestVideos, err := getLatestVideos(service)
	if err != nil {
		log.Fatalf("Error fetching latest videos: %v", err)
	}

	if len(latestVideos) == 0 {
		log.Println("No new videos found. Exiting.")
		return
	}

	log.Printf("Fetched %d videos.", len(latestVideos))

	if err := updateObsidianNotes(latestVideos); err != nil {
		log.Fatalf("Error updating Obsidian notes: %v", err)
	}

	log.Println("Finished updating Obsidian notes.")
}

