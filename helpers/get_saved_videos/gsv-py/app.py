import os
import time
import httplib2
import logging
from googleapiclient.discovery import build
from googleapiclient.http import HttpRequest

# Setup logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

# Suppress warnings from the googleapiclient module
logging.getLogger('googleapiclient.discovery_cache').setLevel(logging.ERROR)
logging.getLogger('google.auth.transport.requests').setLevel(logging.ERROR)

# Disable the file cache
def build_request(http, *args, **kwargs):
    return HttpRequest(http, *args)

# YouTube Data API setup
API_KEY = "AIzaSyA0E0ifJ3WtiVDsnyj-VPang2pHn6JXoKQ" # Add your API key here as a string
youtube = build('youtube', 'v3', developerKey=API_KEY, requestBuilder=build_request)

# Obsidian setup
OBSIDIAN_VAULT_PATH = r'/home/john/Documents/Obsidian Vault/'  # Ensure the path is correct

# List of playlist IDs you want to monitor mapped to their respective file paths
PLAYLIST_IDS = {
    'PL-qEoXFQG4wuaN4huGKJTeORyGOrN26Yx': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Obsidian Videos.md'),
    'PL-qEoXFQG4wtk7qBocZPIMB0j3xgpit7F': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/AI Videos.md'),
    'PL-qEoXFQG4wuj-t1BwimNtTMDBg8aWLJt': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Networking Videos.md'),
    'PL-qEoXFQG4wsgKqexYOF9kSAOLXURRBDV': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Docker Videos.md'),
    'PL-qEoXFQG4wucRrp1r3xx7udQHmRI7xhJ': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Web Videos.md'),
    'PL-qEoXFQG4wtKE8fRsSkR7tWD5Qk2KkyV': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Homelab Videos.md'),
    'PL-qEoXFQG4wvkq990fFa-YonPTGhmDAB5': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Python Videos.md'),
    'PL-qEoXFQG4wv0Jy1Jq8aWjb2-GE0fgqrc': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Hacker Videos.md'),
    'PL-qEoXFQG4wsGPvWNSQ_ulzISK3JDso7S': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Database Videos.md'),
    'PL-qEoXFQG4wtPiHtUsPMElKQibaL6PbKT': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Business Videos.md'),
    'PL-qEoXFQG4wuQoGkmFOzSD0yfJZCY7QRd': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Agents Videos.md'),
    'PL-qEoXFQG4wuRbP19ToXsCp9lgt2U1SxA': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Ollama Videos.md'),
    'PL-qEoXFQG4wsBAxcqKVttqDPhmb66_2NB': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Prompting Videos.md'),
    'PL-qEoXFQG4wuYOlwspyHfXwVFPHGXe2W3': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Automation Videos.md'),
    'PL-qEoXFQG4wsgQ-9lqmHZ1LL1poMCtQhj': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/RAG Videos.md'),
    'PL-qEoXFQG4wtDVPpK2FfdsP7c9r4cIfPH': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Langchain Videos.md'),
    'PL-qEoXFQG4wu8fWgLK4AyS0Ub3nx_meVJ': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/No-Code Videos.md'),
    'PL-qEoXFQG4wsLlwL4xqP9V2yOEhIziz1n': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Fine Tuning Videos.md'),
    'PL-qEoXFQG4wu41nga8EyWj0M5r1nwIK7Z': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Red AI Videos.md'),
    'PL-qEoXFQG4wtNGbvU5ksWhxwhZZbXphFB': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/Web Scraping Videos.md'),
    'PL-qEoXFQG4wsQDX05QMTePY6amXvdGQ-m': os.path.join(OBSIDIAN_VAULT_PATH, '4_Reference Notes/CrewAI Videos.md'),
}

def get_latest_videos():
    latest_videos = []
    for playlist_id in PLAYLIST_IDS.keys():
        try:
            request = youtube.playlistItems().list(
                part='snippet',
                maxResults=50,
                playlistId=playlist_id
            )
            response = request.execute()
            while 'nextPageToken' in response:
                nextPageToken = response['nextPageToken']
                nextPage = youtube.playlistItems().list(
                    part="snippet",
                    playlistId=playlist_id,
                    maxResults=50,
                    pageToken=nextPageToken
                ).execute()
                response['items'] += nextPage['items']
                if 'nextPageToken' not in nextPage:
                    break
                response['nextPageToken'] = nextPage['nextPageToken']

            for item in response['items']:
                video_id = item['snippet']['resourceId']['videoId']
                video_url = f"https://www.youtube.com/watch?v={video_id}"
                video_title = item['snippet']['title']
                latest_videos.append((playlist_id, f"![{video_title}]({video_url})", None))
        except Exception as e:
            logging.error(f"An error occurred while fetching videos for playlist {playlist_id}: {e}")
    return latest_videos

def update_obsidian_notes(videos):
    # Group videos by playlist ID
    videos_by_playlist = {}
    for playlist_id, video_md, _ in videos:
        if playlist_id not in videos_by_playlist:
            videos_by_playlist[playlist_id] = []
        videos_by_playlist[playlist_id].append(video_md)

    for playlist_id, video_list in videos_by_playlist.items():
        file_path = PLAYLIST_IDS[playlist_id]
        cached_links = set()

        if os.path.exists(file_path):
            with open(file_path, 'r') as file:
                content = file.readlines()
                cached_links = {line.strip() for line in content}

        new_videos = [video_md for video_md in video_list if video_md not in cached_links]

        if new_videos:
            with open(file_path, 'a') as file:
                for video_md in new_videos:
                    file.write(f"{video_md}\n\n")
            logging.info(f"Added {len(new_videos)} new videos to {file_path}")
        else:
            logging.info(f"No new videos to add for playlist {playlist_id}")

def main():
    logging.info("Starting to fetch latest videos.")
    latest_videos = get_latest_videos()
    if not latest_videos:
        logging.info("No new videos found. Exiting.")
        return
    logging.info(f"Fetched {len(latest_videos)} videos.")
    update_obsidian_notes(latest_videos)
    logging.info("Finished updating Obsidian notes.")

if __name__ == '__main__':
    main()
