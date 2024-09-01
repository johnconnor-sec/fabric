#!/bin/bash

# Check if yt-dlp and jq are installed
command -v yt-dlp >/dev/null 2>&1 || { echo >&2 "yt-dlp is required but not installed. Aborting."; exit 1; }
command -v jq >/dev/null 2>&1 || { echo >&2 "jq is required but not installed. Aborting."; exit 1; }

# Function to display usage information
usage() {
    echo "Usage: $0 <playlist_url> <output_file>"
    echo "Example: $0 https://youtube.com/playlist?list=PLk6vOUIjcauWAzYx5zn5JTnDL9R-Osk_H ~/output.md"
    exit 1
}

# Check if correct number of arguments is provided
if [ "$#" -ne 2 ]; then
    usage
fi

# Assign arguments to variables
PLAYLIST_URL="$1"
OUTPUT_FILE="$2"

# Create the directory for the output file if it doesn't exist
OUTPUT_DIR=$(dirname "$OUTPUT_FILE")
mkdir -p "$OUTPUT_DIR" || { echo >&2 "Failed to create directory $OUTPUT_DIR. Aborting."; exit 1; }

# Activate virtual environment if it exists
if [ -f "/home/john/tools/venv/bin/activate" ]; then
    source /home/john/tools/venv/bin/activate
fi

# Clear the output file if it already exists, or create it if it doesn't
> "$OUTPUT_FILE" || { echo >&2 "Failed to create or clear $OUTPUT_FILE. Aborting."; exit 1; }

# Use yt-dlp to fetch the video URLs from the playlist
yt-dlp -j --flat-playlist "$PLAYLIST_URL" | jq -r '.id' | while read -r VIDEO_ID; do
    # Construct the full URL for each video
    VIDEO_URL="https://www.youtube.com/watch?v=$VIDEO_ID"
    # Append the formatted link to the output file
    echo "![]($VIDEO_URL)" >> "$OUTPUT_FILE"
    # Add a blank line after each link
    echo "" >> "$OUTPUT_FILE"
done

# Check if any links were actually written
if [ ! -s "$OUTPUT_FILE" ]; then
    echo "Warning: No links were written to $OUTPUT_FILE. The playlist might be empty or inaccessible."
fi

# Deactivate virtual environment if it was activated
if [ -n "$VIRTUAL_ENV" ]; then
    deactivate
fi

echo "All video links have been saved to $OUTPUT_FILE"
