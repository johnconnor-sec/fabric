import os
import sys
import subprocess
import json

def check_command(command):
    """Check if a command is available."""
    return subprocess.call(["which", command], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL) == 0

def usage():
    """Display usage information."""
    print(f"Usage: {sys.argv[0]} <playlist_url> <output_file>")
    print(f"Example: {sys.argv[0]} https://youtube.com/playlist?list=PLk6vOUIjcauWAzYx5zn5JTnDL9R-Osk_H ~/output.md")
    sys.exit(1)

def main():
    # Check if yt-dlp and jq are installed
    if not check_command("yt-dlp"):
        print("yt-dlp is required but not installed. Aborting.", file=sys.stderr)
        sys.exit(1)
    if not check_command("jq"):
        print("jq is required but not installed. Aborting.", file=sys.stderr)
        sys.exit(1)

    # Check if correct number of arguments is provided
    if len(sys.argv) != 3:
        usage()

    # Assign arguments to variables
    playlist_url = sys.argv[1]
    output_file = sys.argv[2]

    # Create the directory for the output file if it doesn't exist
    output_dir = os.path.dirname(output_file)
    os.makedirs(output_dir, exist_ok=True)

    # Activate virtual environment if it exists
    venv_path = "/home/john/tools/venv/bin/activate"
    if os.path.exists(venv_path):
        exec(open(venv_path).read(), {'__file__': venv_path})

    # Clear the output file if it already exists, or create it if it doesn't
    open(output_file, 'w').close()

    # Use yt-dlp to fetch the video URLs from the playlist
    yt_dlp_command = f"yt-dlp -j --flat-playlist {playlist_url}"
    process = subprocess.Popen(yt_dlp_command, stdout=subprocess.PIPE, shell=True)

    with open(output_file, 'w') as f:
        for line in process.stdout:
            video_data = json.loads(line)
            video_id = video_data['id']
            video_url = f"https://www.youtube.com/watch?v={video_id}"
            f.write(f"![]({video_url})\n\n")

    # Check if any links were actually written
    if os.path.getsize(output_file) == 0:
        print(f"Warning: No links were written to {output_file}. The playlist might be empty or inaccessible.")

    # Deactivate virtual environment if it was activated
    if 'VIRTUAL_ENV' in os.environ:
        del os.environ['VIRTUAL_ENV']
        os.environ['PATH'] = os.environ['PATH'].split(os.pathsep)[1:]

    print(f"All video links have been saved to {output_file}")

if __name__ == "__main__":
    main()
