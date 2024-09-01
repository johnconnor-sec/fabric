#!/bin/bash

# Define the input file containing video links
INPUT_FILE="/home/john/tools/my_youtube_playlists/fabric_Process_Links/fabric_outputs/Black Hat Python.md"

# Define the directory containing the files to be sorted and linked
FILE_DIR="/home/john/tools/my_youtube_playlists/fabric_Process_Links/fabric_outputs/Black_Hat_Python/"

# Create a temporary file for the updated video links
TEMP_FILE="temp_video_links.txt"

touch $TEMP_FILE

# Get the list of files in the directory, sort them by the numerical value in the file name
SORTED_FILES=$(ls "$FILE_DIR" | sort -V)

# Initialize a counter
COUNTER=0

# Read the input file line by line
while IFS= read -r line; do
  # If the line contains a video link, process it
  if [[ "$line" =~ \[\!\[\]\(https:// )]]; then
    # Write the video link to the temporary file
    echo "$line" >> "$TEMP_FILE"
    
    # Get the corresponding sorted file
    FILE=$(echo "$SORTED_FILES" | sed -n "$((COUNTER+1))p")
    
    # Check if the file link already exists in the input file
    FILE_LINK="[[$FILE]]"
    if grep -Fxq "$FILE_LINK" "$INPUT_FILE"; then
      echo "Link $FILE_LINK already exists, skipping..."
    else
      # If a corresponding file exists and is not already linked, add it as a link with a newline
      if [ -n "$FILE" ]; then
        echo "$FILE_LINK" >> "$TEMP_FILE"
        echo "" >> "$TEMP_FILE" # Add a newline
        ((COUNTER++))
      fi
    fi
  else
    # Write any non-link lines directly to the temporary file
    echo "$line" >> "$TEMP_FILE"
  fi
done < "$INPUT_FILE"

# Replace the original video links file with the updated one
mv "$TEMP_FILE" "$INPUT_FILE"

echo "Video links have been updated with corresponding file links."
