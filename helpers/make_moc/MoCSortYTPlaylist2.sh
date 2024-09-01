#!/bin/bash

# Set stricter error handling
set -euo pipefail

# Define the input file containing video links
INPUT_FILE="/home/john/tools/my_youtube_playlists/Black_Hat_Python.md"

# Define the directory containing the files to be sorted and linked
FILE_DIR="/home/john/tools/my_youtube_playlists/Black_Hat_Python"

# Create a temporary file for the updated video links
TEMP_FILE=$(mktemp)

# Ensure temporary file is deleted on script exit
trap 'rm -f "$TEMP_FILE"' EXIT

# Get the list of files in the directory, sort them by the numerical value in the file name
SORTED_FILES=$(ls -v "$FILE_DIR")

# Initialize a counter
COUNTER=0

# Process the input file
while IFS= read -r line; do
  if [[ "$line" =~ \[\!\[\]\(https:// ]]; then
    echo "$line" >> "$TEMP_FILE"
    
    FILE=$(echo "$SORTED_FILES" | sed -n "$((COUNTER+1))p")
    FILE_LINK="[[$FILE]]"
    
    if ! grep -Fxq "$FILE_LINK" "$INPUT_FILE"; then
      if [ -n "$FILE" ]; then
        echo "$FILE_LINK" >> "$TEMP_FILE"
        echo "" >> "$TEMP_FILE"
        ((COUNTER++))
      fi
    fi
  else
    echo "$line" >> "$TEMP_FILE"
  fi
done < "$INPUT_FILE"

# Replace the original file with the updated one
mv "$TEMP_FILE" "$INPUT_FILE"

echo "Video links have been updated with corresponding file links."
