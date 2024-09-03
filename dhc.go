package main

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

// Helper function to compute the MD5 hash of a file
func md5Hash(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return ""
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Printf("Error hashing file %s: %v\n", filePath, err)
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}

// Collect file hashes from a directory
func collectFileHashes(directory string) map[string][]string {
	fileHashes := make(map[string][]string)
	var mu sync.Mutex
	var wg sync.WaitGroup

	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				hash := md5Hash(p)
				if hash != "" {
					mu.Lock()
					fileHashes[hash] = append(fileHashes[hash], p)
					mu.Unlock()
				}
			}(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory %s: %v\n", directory, err)
	}

	wg.Wait()
	return fileHashes
}

// Write file hashes to a CSV file
func writeCSV(fileHashes map[string][]string, outputFile string) {
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating CSV file %s: %v\n", outputFile, err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for hash, paths := range fileHashes {
		if len(paths) > 1 {
			record := append([]string{hash}, paths...)
			if err := writer.Write(record); err != nil {
				fmt.Printf("Error writing to CSV file %s: %v\n", outputFile, err)
			}
		}
	}
}

func main() {
	drive := flag.String("d", "", "Drive letter (e.g., 'D:\\') of the removable drive.")
	output := flag.String("o", "", "Path to the output CSV file.")
	flag.Parse()

	if *drive == "" || *output == "" {
		fmt.Println("Drive and output CSV file path are required.")
		return
	}

	if _, err := os.Stat(*drive); os.IsNotExist(err) {
		fmt.Printf("The drive %s does not exist.\n", *drive)
		return
	}

	fmt.Printf("Collecting file hashes from drive %s...\n", *drive)
	fileHashes := collectFileHashes(*drive)

	fmt.Printf("Writing results to %s...\n", *output)
	writeCSV(fileHashes, *output)

	fmt.Println("Processing complete.")
}
