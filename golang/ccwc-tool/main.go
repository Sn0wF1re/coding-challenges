package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	argLen := len(os.Args)

	switch argLen {
	case 3:
		fileName := os.Args[2]
		switch os.Args[1] {
		case "-c":
			byteCount := countBytes(fileName)
			fmt.Println(byteCount, fileName)
		case "-l":
			lineCount, err := countLines(fileName)
			if err != nil {
				fmt.Println("Error counting lines: ", err)
			}
			fmt.Println(lineCount, fileName)
		case "-w":
			wordCount, err := countWords(fileName)
			if err != nil {
				fmt.Println("Error counting words: ", err)
			}
			fmt.Println(wordCount, fileName)
		case "-m":
			runeCount, err := countRunes(fileName)
			if err != nil {
				fmt.Println("Error counting runes: ", err)
			}
			fmt.Println(runeCount, fileName)
		}
	case 2:
		fileName := os.Args[1]
		lineCount, err := countLines(fileName)
		if err != nil {
			fmt.Println("Error counting lines: ", err)
		}
		wordCount, err := countWords(fileName)
		if err != nil {
			fmt.Println("Error counting words: ", err)
		}
		byteCount := countBytes(fileName)
		fmt.Println(lineCount, wordCount, byteCount, fileName)
	}
}

func countBytes(fileName string) int64 {
	info, err := os.Stat(fileName)

	if err != nil {
		fmt.Println("Error checking the file:", err)
		return 0
	}

	if info.IsDir() {
		fmt.Println("The specified path is a directory, not a file.")
		return 0
	}

	return info.Size()
}

func countLines(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	return lineCount, scanner.Err()
}

func countWords(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount, scanner.Err()
}

func countRunes(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	runeCount := 0
	for scanner.Scan() {
		runeCount++
	}
	return runeCount, scanner.Err()
}
