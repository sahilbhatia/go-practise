package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Define ANSI color codes as constants
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

// Mapping from color name (lowercase) to ANSI codes
var colorMap = map[string]string{
	"red":    ColorRed,
	"green":  ColorGreen,
	"yellow": ColorYellow,
	"blue":   ColorBlue,
	"purple": ColorPurple,
	"cyan":   ColorCyan,
	"white":  ColorWhite,
}

func main() {
	dir, listMode, allMode, colorName := parseFlags()

	colorCode, ok := getColorCode(colorName)
	if !ok {
		log.Printf("Warning: invalid color '%s' provided. Using default color 'green'.\n", colorName)
		colorCode = ColorGreen
	}

	if err := listDirectory(dir, listMode, allMode, colorCode); err != nil {
		log.Fatal(err)
	}
}

func parseFlags() (dir string, listMode bool, allMode bool, colourName string) {
	dirPtr := flag.String("d", ".", "path to dir whose contents needs to be listed")
	listPtr := flag.Bool("l", false, "show one file per line")
	allPtr := flag.Bool("a", false, "include hidden files in the list")
	colorPtr := flag.String("c", "green", "color to use for filenames (red, green, yellow, blue, purple, cyan, white)")

	flag.Parse()

	return *dirPtr, *listPtr, *allPtr, strings.ToLower(*colorPtr)
}

// getColorCode returns the ANSI color code for a name and whether it's valid
func getColorCode(colorName string) (string, bool) {
	code, ok := colorMap[colorName]

	return code, ok
}

func listDirectory(dir string, listMode, allMode bool, colorCode string) error {
	if dir == "." {
		fmt.Printf("Listing contents of current directory\n\n")
	} else {
		fmt.Printf("Listing contents of '%s' directory\n\n", dir)
	}

	files, err := readVisibleFiles(dir, allMode)
	if err != nil {
		return err
	}

	printFiles(dir, files, listMode, colorCode)

	return nil
}

func readVisibleFiles(dir string, allMode bool) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []os.DirEntry
	for _, entry := range entries {
		name := entry.Name()
		if !allMode && strings.HasPrefix(name, ".") {
			continue
		}
		files = append(files, entry)
	}

	return files, nil
}

// colorize returns the text wrapped in ANSI color escape sequences
func colorize(text, color string) string {
	return color + text + ColorReset
}

func printFiles(dirPath string, files []os.DirEntry, listMode bool, colorCode string) {
	if listMode {
		for _, file := range files {
			filename := file.Name()

			info, err := os.Stat(filepath.Join(dirPath, filename))
			if err != nil {
				log.Fatal(err)
			}

			mode := info.Mode().String()
			size := info.Size()
			dateTime := info.ModTime().Format("02 Jan 15:04")
			coloredFileName := "\t" + colorize(filename, colorCode)

			fmt.Printf("%s\t%10d bytes\t%10v\t%s\n", mode, size, dateTime, coloredFileName)
		}

		return
	}

	for i, file := range files {
		// fmt.Print(colorize(file.Name(), colorCode))
		fmt.Print(file.Name())

		if i != len(files)-1 {
			fmt.Print("\t")
		}
	}

	fmt.Println()
}
