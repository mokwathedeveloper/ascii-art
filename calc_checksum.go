package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

var expectedChecksum = map[string]string{
	"standard.txt":   "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

func ValidateFileChecksum(file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist. Downloading...\n", file)
		err := DownloadFile(file)
		if err != nil {
			return fmt.Errorf("error downloading file: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("error checking file existence: %w", err)
	} else {
		// File exists, calculate its checksum
		checksum, err := calculateChecksum(file)
		if err != nil {
			return fmt.Errorf("error calculating checksum: %w", err)
		}

		expected, ok := expectedChecksum[file]
		if !ok {
			return fmt.Errorf("no expected checksum defined for file: %s", file)
		}

		if checksum != expected {
			fmt.Printf("Checksum verification failed for file: %s. Downloading...\n", file)
			err := DownloadFile(file)
			if err != nil {
				return fmt.Errorf("error downloading file: %w", err)
			}
		}

		// fmt.Printf("Checksum verified for file %s\n", file)
	}
	return nil
}

func calculateChecksum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
