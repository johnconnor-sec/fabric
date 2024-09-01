import os
import subprocess
import json
import re

def download_video_info(video_url):
    result = subprocess.run(
        ['yt-dlp', '--skip-download', '--write-info-json', '--', video_url],
        capture_output=True, text=True
    )
    if result.returncode != 0:
        raise Exception("Failed to download video info")
    info_file = next(file for file in os.listdir() if file.endswith('.info.json'))
    with open(info_file, 'r', encoding='utf-8') as f:
        info = json.load(f)
    os.remove(info_file)
    return info

def download_full_transcript(video_url, temp_dir):
    output_template = os.path.join(temp_dir, '%(title)s.%(ext)s')
    
    cmd = [
        'yt-dlp', '--skip-download', '--write-auto-sub', '--sub-lang', 'en',
        '--sub-format', 'vtt', '--convert-subs', 'vtt',
        '--output', output_template, '--', video_url
    ]
    
    result = subprocess.run(cmd, capture_output=True, text=True)
    if result.returncode != 0:
        print(f"yt-dlp error output: {result.stderr}")
        raise Exception("Failed to download full transcript")
    
    vtt_files = [f for f in os.listdir(temp_dir) if f.endswith('.vtt')]
    if not vtt_files:
        raise FileNotFoundError("No .vtt file found after download")
    
    return os.path.join(temp_dir, vtt_files[0])

def extract_chapter_transcript(full_transcript_file, start_time, end_time):
    with open(full_transcript_file, 'r', encoding='utf-8') as f:
        lines = f.readlines()
    
    chapter_lines = []
    in_chapter = False
    
    for line in lines:
        if re.match(r'^\d{2}:\d{2}:\d{2}\.\d{3} --> \d{2}:\d{2}:\d{2}\.\d{3}', line):
            timestamp = line.split(' --> ')[0]
            time_seconds = sum(float(x) * 60 ** i for i, x in enumerate(reversed(timestamp.split(':'))))
            
            if start_time <= time_seconds < end_time:
                in_chapter = True
            elif time_seconds >= end_time:
                break
        
        if in_chapter:
            chapter_lines.append(line)
    
    return chapter_lines

def clean_vtt_text(vtt_text):
    clean_lines = []
    for line in vtt_text:
        clean_line = re.sub(r'<[^>]+>', '', line).strip()
        if clean_line and not re.match(r'^\d{2}:\d{2}:\d{2}\.\d{3} --> \d{2}:\d{2}:\d{2}\.\d{3}', clean_line):
            clean_lines.append(clean_line)
    return clean_lines

def remove_duplicates(text_list):
    seen = set()
    result = []
    for text in text_list:
        if text not in seen:
            seen.add(text)
            result.append(text)
    return result

def convert_vtt_to_md(vtt_lines, md_file):
    clean_lines = clean_vtt_text(vtt_lines)
    unique_lines = remove_duplicates(clean_lines)
    paragraph = ' '.join(unique_lines)

    with open(md_file, 'w', encoding='utf-8') as md:
        md.write(paragraph)

def main(video_url):
    temp_dir = 'temp_transcripts'
    os.makedirs(temp_dir, exist_ok=True)

    try:
        video_info = download_video_info(video_url)
        chapters = video_info.get('chapters', [])
        
        if not chapters:
            print("No chapters found in the video. The entire video will be treated as a single chapter.")
            chapters = [{'title': 'Full Video', 'start_time': 0}]
        
        full_transcript_file = download_full_transcript(video_url, temp_dir)
        print(f"Full transcript downloaded to: {full_transcript_file}")

        for i, chapter in enumerate(chapters):
            start_time = chapter['start_time']
            end_time = chapters[i + 1]['start_time'] if i + 1 < len(chapters) else float('inf')
            
            chapter_lines = extract_chapter_transcript(full_transcript_file, start_time, end_time)
            
            # Create numbered filename without slugifying the title
            chapter_number = i + 1
            # Replace hyphens with spaces in the title for natural formatting
            chapter_title = chapter['title'].replace('-', ' ')
            md_file = f"{chapter_number:02d}. {chapter_title}.md"
            
            convert_vtt_to_md(chapter_lines, md_file)
            print(f"Chapter '{chapter_title}' transcript saved to {md_file}")

    except Exception as e:
        print(f"An error occurred: {str(e)}")
    
    finally:
        # Clean up temporary files
        for file in os.listdir(temp_dir):
            os.remove(os.path.join(temp_dir, file))
        os.rmdir(temp_dir)

if __name__ == '__main__':
    video_url = input("Enter the YouTube video URL: ")
    main(video_url)
