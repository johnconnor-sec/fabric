#!/bin/bash

# Check if a directory is provided as an argument
if [ $# -eq 0 ]; then
    echo "Please provide a directory path as an argument."
    echo "Usage: $0 /path/to/directory"
    exit 1
fi

# Store the provided directory path
directory="$1"

# Check if the provided path is a directory
if [ ! -d "$directory" ]; then
    echo "Error: '$directory' is not a valid directory."
    exit 1
fi

# Iterate through all files in the directory
for file in "$directory"/*; do
    # Check if it's a regular file
    if [ -f "$file" ]; then
        # Get the file name without the path
        filename=$(basename "$file")
        
        # Remove the file extension
        title="${filename%.*}"
        
        # Create the title line
        title="# $title"
        
        # Create a temporary file
        temp_file=$(mktemp)
        
        # Add the title to the temporary file, followed by the original content
        echo "$title" > "$temp_file"
        cat "$file" >> "$temp_file"
        
        # Replace the original file with the temporary file
        mv "$temp_file" "$file"
        
        echo "Added title to: $filename"
    fi
done

echo "Process completed."
