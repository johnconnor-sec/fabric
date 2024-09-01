#!/bin/bash

# Check if the user has provided both directories as arguments
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <input_directory> <output_directory>"
    exit 1
fi

INPUT_DIR="$1"
OUTPUT_DIR="$2"

# Create the output directory if it does not exist
mkdir -p "$OUTPUT_DIR"

# Iterate over each file in the input directory
# Use a while loop with read to handle spaces in filenames
find "$INPUT_DIR" -maxdepth 1 -type f | while IFS= read -r FILE; do
    # Extract the filename from the path
    FILENAME=$(basename "$FILE")

    # Run the subprocess and capture the output
    OUTPUT=$(cat "$FILE" | fabric < "$FILE")

    # Save the output to a new file in the output directory
    echo "$OUTPUT" > "$OUTPUT_DIR/$FILENAME"
done

echo "Processing complete."
