#!/bin/bash

# Function to display usage information
usage() {
    echo "Usage: $0 -i INPUT_FILE -o OUTPUT_DIR"
    echo "  -i INPUT_FILE   : Path to the input file containing YouTube links"
    echo "  -o OUTPUT_DIR   : Path to the directory containing files to be linked"
    echo "Example: $0 -i /path/to/links.md -o /path/to/transcripts/"
    exit 1
}

# Parse command line arguments
while getopts "i:o:" opt; do
    case "$opt" in
        i) INPUT_FILE=$OPTARG ;;
        o) OUTPUT_DIR=$OPTARG ;;
        *) usage ;;
    esac
done

# Check if required arguments are provided
if [ -z "$INPUT_FILE" ] || [ -z "$OUTPUT_DIR" ]; then
    echo "Error: Both INPUT_FILE and OUTPUT_DIR must be specified."
    usage
fi

# Check if input file exists
if [ ! -f "$INPUT_FILE" ]; then
    echo "Error: Input file '$INPUT_FILE' does not exist."
    exit 1
fi

# Create output directory if it doesn't exist
if [ ! -d "$OUTPUT_DIR" ]; then
    echo "Output directory '$OUTPUT_DIR' does not exist. Creating it now."
    mkdir -p "$OUTPUT_DIR"
    if [ $? -ne 0 ]; then
        echo "Error: Failed to create output directory '$OUTPUT_DIR'."
        exit 1
    fi
    echo "Output directory '$OUTPUT_DIR' created successfully."
fi

# Set stricter error handling
set -euo pipefail

# Create a temporary file for the updated video links
TEMP_FILE=$(mktemp)

# Ensure temporary file is deleted on script exit
trap 'rm -f "$TEMP_FILE"' EXIT

# Get the list of files in the directory, sort them by the numerical value in the file name
SORTED_FILES=$(ls -v "$OUTPUT_DIR")

# Initialize a counter and a flag for changes
COUNTER=0
CHANGES_MADE=false

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
                CHANGES_MADE=true
            fi
        else
            echo "$FILE_LINK" >> "$TEMP_FILE"
            echo "" >> "$TEMP_FILE"
            ((COUNTER++))
        fi
    else
        echo "$line" >> "$TEMP_FILE"
    fi
done < "$INPUT_FILE"

# Check if any changes were made
if [ "$CHANGES_MADE" = true ]; then
    # Replace the original file with the updated one
    mv "$TEMP_FILE" "$INPUT_FILE"
    echo "Video links have been updated with corresponding file links in '$INPUT_FILE'."
else
    echo "No changes were necessary. All video links already have corresponding file links."
    # Remove the temporary file as we didn't use it
    rm "$TEMP_FILE"
fi

# Report on any unlinked files
TOTAL_FILES=$(echo "$SORTED_FILES" | wc -l)
if [ $COUNTER -lt $TOTAL_FILES ]; then
    echo "Warning: There are $(($TOTAL_FILES - $COUNTER)) files in '$OUTPUT_DIR' that were not linked."
    echo "This may be because there are more files than video links, or because some links already had file associations."
fi
