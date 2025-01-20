# Duplicate File Finder

This tool helps to identify duplicate files on a removable drive (or any other directory) by computing the MD5 hash of each file and comparing them. The results are written to a CSV file, where each entry lists the MD5 hash and the paths of the duplicate files.

## Features
- Computes the MD5 hash of each file in a specified directory.
- Identifies duplicate files based on their MD5 hash.
- Outputs the results into a CSV file.

## Requirements
- Go 1.x or later
- A removable drive (or any directory) to scan for duplicate files.

## Usage

### Command-line Arguments

- `-d <drive>`: The path to the directory (e.g., drive letter) you want to scan for duplicate files (e.g., `D:\\`).
- `-o <output>`: The path to the CSV file where duplicate file information will be saved.

### Example

```bash
go run main.go -d "D:\" -o "duplicates.csv"
