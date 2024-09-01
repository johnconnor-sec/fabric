package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "os/exec"
    "regexp"
)

func main() {
    inputFile := flag.String("i", "", "Path to the input file containing YouTube links")
    outputDir := flag.String("o", "", "Path to the directory where output files will be saved")
    command := flag.String("cmd", "fabric", "Command to process the video link")

    flag.Parse()

    if *inputFile == "" || *outputDir == "" {
        fmt.Println("Error: Both INPUT_FILE and OUTPUT_DIR must be specified.")
        flag.Usage()
        os.Exit(1)
    }

    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Error: Input file '%s' does not exist.\n", *inputFile)
        os.Exit(1)
    }
    defer file.Close()

    os.MkdirAll(*outputDir, os.ModePerm)

    scanner := bufio.NewScanner(file)
    videoLinkPattern := regexp.MustCompile(`\[\!\[\]\(https://.*youtube\.com.*\)\]`)

    for scanner.Scan() {
        line := scanner.Text()
        if videoLinkPattern.MatchString(line) {
            videoURL := extractVideoURL(line)

            fmt.Printf("Processing URL: %s\n", videoURL)

            cmd := exec.Command("yt-dlp", "--get-title", videoURL)
            videoTitleBytes, _ := cmd.Output()
            videoTitle := string(videoTitleBytes)

            cleanTitle := sanitizeTitle(videoTitle)
            outputFile := fmt.Sprintf("%s/%s.md", *outputDir, cleanTitle)

            fmt.Printf("Output file: %s\n", outputFile)

            cmd = exec.Command("yt", "--transcript", videoURL)
            pipe, _ := cmd.StdoutPipe()
            cmd.Start()

            fabricCmd := exec.Command(*command)
            fabricCmd.Stdin = pipe
            fabricCmd.Stdout, _ = os.Create(outputFile)
            fabricCmd.Run()

            fmt.Printf("Processed %s and saved output to %s\n", videoURL, outputFile)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
    }

    fmt.Println("All video links have been processed.")
}

func extractVideoURL(line string) string {
    re := regexp.MustCompile(`(?<=\[\!\[\]\().*?(?=\)\])`)
    return re.FindString(line)
}

func sanitizeTitle(title string) string {
    re := regexp.MustCompile(`[^[:alnum:]_]+`)
    return re.ReplaceAllString(title, "")
}

