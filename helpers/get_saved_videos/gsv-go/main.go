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
	"PL-qEoXFQG4wtk7qBocZPIMB0j3xgpit7F": filepath.Join(obsidianVaultPath, "4_Reference Notes/AI Videos.md"),
	"PL-qEoXFQG4wuj-t1BwimNtTMDBg8aWLJt": filepath.Join(obsidianVaultPath, "4_Reference Notes/Networking Videos.md"),
	"PL-qEoXFQG4wsgKqexYOF9kSAOLXURRBDV": filepath.Join(obsidianVaultPath, "4_Reference Notes/Docker Videos.md"),
	"PL-qEoXFQG4wucRrp1r3xx7udQHmRI7xhJ": filepath.Join(obsidianVaultPath, "4_Reference Notes/Web Videos.md"),
	"PL-qEoXFQG4wtKE8fRsSkR7tWD5Qk2KkyV": filepath.Join(obsidianVaultPath, "4_Reference Notes/Homelab Videos.md"),
	"PL-qEoXFQG4wvkq990fFa-YonPTGhmDAB5": filepath.Join(obsidianVaultPath, "4_Reference Notes/Python Videos.md"),
	"PL-qEoXFQG4wv0Jy1Jq8aWjb2-GE0fgqrc": filepath.Join(obsidianVaultPath, "4_Reference Notes/Hacker Videos.md"),
	"PL-qEoXFQG4wsGPvWNSQ_ulzISK3JDso7S": filepath.Join(obsidianVaultPath, "4_Reference Notes/Database Videos.md"),
	"PL-qEoXFQG4wtPiHtUsPMElKQibaL6PbKT": filepath.Join(obsidianVaultPath, "4_Reference Notes/Business Videos.md"),
	"PL-qEoXFQG4wuQoGkmFOzSD0yfJZCY7QRd": filepath.Join(obsidianVaultPath, "4_Reference Notes/Agents Videos.md"),
	"PL-qEoXFQG4wuRbP19ToXsCp9lgt2U1SxA": filepath.Join(obsidianVaultPath, "4_Reference Notes/Ollama Videos.md"),
	"PL-qEoXFQG4wsBAxcqKVttqDPhmb66_2NB": filepath.Join(obsidianVaultPath, "4_Reference Notes/Prompting Videos.md"),
	"PL-qEoXFQG4wuYOlwspyHfXwVFPHGXe2W3": filepath.Join(obsidianVaultPath, "4_Reference Notes/Automation Videos.md"),
	"PL-qEoXFQG4wsgQ-9lqmHZ1LL1poMCtQhj": filepath.Join(obsidianVaultPath, "4_Reference Notes/RAG Videos.md"),
	"PL-qEoXFQG4wtDVPpK2FfdsP7c9r4cIfPH": filepath.Join(obsidianVaultPath, "4_Reference Notes/Langchain Videos.md"),
	"PL-qEoXFQG4wu8fWgLK4AyS0Ub3nx_meVJ": filepath.Join(obsidianVaultPath, "4_Reference Notes/No-Code Videos.md"),
	"PL-qEoXFQG4wsLlwL4xqP9V2yOEhIziz1n": filepath.Join(obsidianVaultPath, "4_Reference Notes/Fine Tuning Videos.md"),
	"PL-qEoXFQG4wu41nga8EyWj0M5r1nwIK7Z": filepath.Join(obsidianVaultPath, "4_Reference Notes/Red AI Videos.md"),
	"PL-qEoXFQG4wtNGbvU5ksWhxwhZZbXphFB": filepath.Join(obsidianVaultPath, "4_Reference Notes/Web Scraping Videos.md"),
	"PL-qEoXFQG4wsQDX05QMTePY6amXvdGQ-m": filepath.Join(obsidianVaultPath, "4_Reference Notes/CrewAI Videos.md"),
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

