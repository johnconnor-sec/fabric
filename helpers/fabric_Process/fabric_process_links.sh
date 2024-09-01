#!/bin/bash
# Note: Use the 'MakeMoC' tool to make a MoC of the Playlist

# Function to display usage information
usage() {
    echo "Usage: $0 -i INPUT_FILE -o OUTPUT_DIR"
    echo "  -i INPUT_FILE   : Path to the input file containing YouTube links"
    echo "  -o OUTPUT_DIR   : Path to the directory where output files will be saved"
    echo "Example: $0 -i file.md -o /home/john/dir/dir/dir"
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

# Create the output directory if it does not exist
mkdir -p "$OUTPUT_DIR"

# Read the input file line by line
while IFS= read -r line; do
  # Check if the line contains a video link
  if [[ $line =~ \[\!\[\]\(https://.*youtube\.com.*\)\] ]]; then
    # Extract the video URL from the line
    VIDEO_URL=$(echo "$line" | grep -oP '(?<=\[\!\[\]\().*?(?=\)\])')

    echo "Processing URL: $VIDEO_URL"  # Debug statement

    # Use yt-dlp to get the video title
    VIDEO_TITLE=$(yt-dlp --get-title "$VIDEO_URL")

    echo "Video title: $VIDEO_TITLE"  # Debug statement

    # Replace spaces with underscores and remove special characters in the title
    CLEAN_TITLE=$(echo "$VIDEO_TITLE" | tr -cd '[:alnum:]_')

    echo "Clean title: $CLEAN_TITLE"  # Debug statement

    # Define the output file for this iteration
    OUTPUT_FILE="$OUTPUT_DIR/${CLEAN_TITLE}.md"

    echo "Output file: $OUTPUT_FILE"  # Debug statement

    # Run the fabric subprocess and save the output
    yt --transcript "$VIDEO_URL" | fabric > "$OUTPUT_FILE"

    echo "Processed $VIDEO_URL and saved output to $OUTPUT_FILE"
  fi
done < "$INPUT_FILE"

echo "All video links have been processed."
