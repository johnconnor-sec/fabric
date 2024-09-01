#!/bin/bash

# Load configuration
source="$HOME/tools/my_youtube_playlists/get_saved_videos/gsv-py/.config/config.sh"

# Lock file to prevent multiple instances
LOCKFILE="/tmp/run_script.lock"

# Function to clean up resources
cleanup() {
    rm -f "$LOCKFILE"
    deactivate 2>/dev/null  # Suppress error if deactivate fails
    exit 1
}

# Trap signals for cleanup
trap cleanup INT TERM EXIT

# Check if lock file exists
if [ -e "$LOCKFILE" ]; then
    echo "Script is already running."
    exit 1
fi

# Create lock file
touch "$LOCKFILE"

# Activate virtual environment
source "$HOME/tools/my_youtube_playlists/get_saved_videos/gsv-py/venv/bin/activate"

# Define log file
LOG_DIR="$HOME/tools/my_youtube_playlists/get_saved_videos/logs"
mkdir -p "$LOG_DIR"  # Create the directory if it doesn't exist
LOG_FILE="$LOG_DIR/script.log"

# Python script
PYTHON_SCRIPT="$HOME/tools/my_youtube_playlists/get_saved_videos/gsv-py/app.py"

# Run Python script and log output with timestamp
{
    echo "$(date '+%Y-%m-%d %H:%M:%S') - Starting script"
    /usr/bin/python3 "$PYTHON_SCRIPT"
    echo "$(date '+%Y-%m-%d %H:%M:%S') - Script finished"
} >> "$LOG_FILE" 2>&1

# Deactivate virtual environment
deactivate

# Remove lock file
rm -f "$LOCKFILE"

# Send notification
notify-send "$(date '+%Y-%m-%d %H:%M:%S') YouTube Playlist Update" "The playlist update script has finished running."

# Cleanup trap
trap - INT TERM EXIT

exit 0
