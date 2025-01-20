Duplicate File Finder

This Go program scans a specified directory or drive for duplicate files based on their MD5 hash. The program identifies duplicate files by comparing the hash values of the files and outputs the results into a CSV file.
Features

    Scans directories and subdirectories for files.
    Computes MD5 hash for each file.
    Identifies duplicate files by comparing MD5 hashes.
    Outputs duplicate file information into a CSV file, with the hash value and the paths of the duplicate files.

Prerequisites

    Go 1.18 or later installed.
    A compatible operating system (Linux, macOS, Windows).
    A removable drive or any directory to scan for duplicate files.

Installation

    Clone this repository:

git clone https://github.com/yourusername/duplicate-file-finder.git

Navigate to the project directory:

cd duplicate-file-finder

Build the program:

    go build -o duplicate-file-finder main.go

    Run the program.

Usage

To run the program, use the following command format:

./duplicate-file-finder -d <drive or directory path> -o <output CSV file path>

Flags

    -d (required): The path to the drive or directory to scan (e.g., D:\\ for Windows, /path/to/directory for Linux/macOS).
    -o (required): The path to the output CSV file where the results will be saved.

Example

./duplicate-file-finder -d D:\ -o duplicates.csv

This command will scan the D:\ drive for duplicate files and save the results in the duplicates.csv file.
Output Format

The output CSV file will contain the following columns:

    MD5 hash: The MD5 hash of the file.
    File paths: A list of paths to files that have the same MD5 hash.

Each row represents a group of duplicate files.

Example:

<MD5_HASH>, <Path_1>, <Path_2>
<MD5_HASH>, <Path_3>, <Path_4>

How It Works

    The program accepts a drive or directory path as input.
    It recursively traverses the directory and computes the MD5 hash for each file.
    It stores files with identical hashes and outputs them in the CSV format.
    The program runs concurrently, processing multiple files in parallel for faster execution.

Contributing

Contributions are welcome! If you have suggestions, bug fixes, or improvements, feel free to open an issue or submit a pull request.
