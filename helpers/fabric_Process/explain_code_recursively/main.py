#!/usr/bin/python3
import os
import subprocess
import sys
import argparse

def process_file(input_filepath, output_dir, ignored_paths, prepend_string):
    if any(ignored_path in input_filepath for ignored_path in ignored_paths):
        print(f"Ignored file: {input_filepath}")
        return

    try:
        # Use 'cat' to read the file content
        cat_process = subprocess.Popen(['cat', input_filepath], stdout=subprocess.PIPE)
        cat_output, _ = cat_process.communicate()

        # Combine prepend_string and file content
        combined_input = prepend_string + '\n\n' + cat_output.decode()

        # Use 'fabric' to process the combined input
        fabric_process = subprocess.Popen(
            ['fabric', '-m', 'claude-3-5-sonnet-20240620', '-p', 'explain_code'],
            stdin=subprocess.PIPE, stdout=subprocess.PIPE
        )
        output, _ = fabric_process.communicate(input=combined_input.encode())

        # Create output file path
        filename = os.path.basename(input_filepath)
        output_filepath = os.path.join(output_dir, filename + '.md')

        # Write the output to the new .md file
        with open(output_filepath, 'wb') as f:
            f.write(output)

        print(f"Processed file: {input_filepath} -> {output_filepath}")

    except Exception as e:
        print(f"Error processing file {input_filepath}: {e}")

def process_directory(input_dir, output_dir, ignored_paths, prepend_string):
    for root, dirs, files in os.walk(input_dir):
        for file in files:
            input_filepath = os.path.join(root, file)
            process_file(input_filepath, output_dir, ignored_paths, prepend_string)

def main():
    parser = argparse.ArgumentParser(description="Process a directory of files and save the output to a new directory.")
    parser.add_argument('input_dir', default='', help='The input directory to process')
    parser.add_argument('output_dir', default='', help='The directory to save the output .md files')
    parser.add_argument('-prompt', default='', help='String to prepend to the output')
    parser.add_argument('-ignore', default='', help='Comma-separated list of files and directories to ignore')

    args = parser.parse_args()

    input_dir = args.input_dir
    output_dir = args.output_dir
    prepend_string = args.prompt
    ignored_paths = [path.strip() for path in args.ignore.split(',') if path.strip()]

    if not ignored_paths:
        print("No paths to ignore were provided.")
    else:
        print(f"The following paths will be ignored: {', '.join(ignored_paths)}")

    if not os.path.isdir(input_dir):
        print(f"Input directory '{input_dir}' does not exist.")
        sys.exit(1)

    if not os.path.isdir(output_dir):
        os.makedirs(output_dir)

    process_directory(input_dir, output_dir, ignored_paths, prepend_string)

if __name__ == "__main__":
    main()
