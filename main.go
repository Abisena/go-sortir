package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// Ambil path direktori sekarang
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Gagal mendapatkan direktori sekarang:", err)
		return
	}

	// Baca semua isi folder
	files, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Println("Gagal membaca folder:", err)
		return
	}

	// Loop tiap File
	for _, file := range files {
		if file.IsDir() || file.Name() == "main.go" {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		var destFolder string

		// Menentukan Folder Tujuan Berdasarkan Ekstensi
		switch ext {
		case ".jpg", ".jpeg", ".png", ".gif":
			destFolder = "images"
		case ".pdf", ".doc", ".docx", ".txt":
			destFolder = "docs"
		case ".mp4", ".avi", ".mkv":
			destFolder = "videos"
		default:
			destFolder = "others"
		}

		if _, err := os.Stat(destFolder); os.IsNotExist(err) {
			err := os.Mkdir(destFolder, os.ModePerm)
			if err != nil {
				fmt.Println("Gagal membuat folder:", destFolder, ":", err)
				continue
			}
		}

		srcPath := filepath.Join(currentDir, file.Name())
		dstPath := filepath.Join(currentDir, destFolder, file.Name())

		err := os.Rename(srcPath, dstPath)
		if err != nil {
			fmt.Println("Gagal Memindahkan File", err)
		} else {
			fmt.Println("Berhasil Memindahkan", file.Name(), "Ke", destFolder)
		}
	}
}
