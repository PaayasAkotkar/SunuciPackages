package dataset

import (
	"bufio"
	"os"
	"strings"
)

// EIndex process of enum
type EIndex int

const (
	firstIndex EIndex = iota
	lastIndex
	midIndex
)

// IOSeeFirstMiddleLast return value at the first, middle, last in the file
// NOTE: not all
func IOSeeFirstMiddleLast(filePath string, index EIndex) string {
	element := ""
	fs, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	fls, err := fs.Stat()
	if err != nil {
		panic(err)
	}

	// get the index of the first char
	// i1Char := make([]byte, fls.Size()-(fls.Size()-1))
	i1Char := make([]byte, fls.Size()-(fls.Size()-1))

	// get the max len of the file
	iChar := make([]byte, fls.Size())

	// get the first char
	fChar, err := fs.Read(i1Char)
	if err != nil {
		panic(err)
	}

	fChar2, err := fs.Read(iChar)
	if err != nil {
		panic(err)
	}

	switch index {
	case firstIndex:
		{
			element = string(i1Char[:fChar])
			break
		}
	case midIndex:
		{
			element = string(iChar[fChar2])
			break
		}
	case lastIndex:
		{
			element = string(iChar[fChar2-1])
			break
		}
	default:
		{
			element = "not found"
		}

	}
	return element
}

// IOGetFirstMiddleLast return index at the first, middle, last in the file
// NOTE: not all
func IOGetFirstMiddleLast(filePath string, index EIndex) int {
	elementIndex := 1
	fs, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	fls, err := fs.Stat()
	if err != nil {
		panic(err)
	}

	//// get the max len of the file
	iChar := make([]byte, fls.Size())

	fChar2, err := fs.Read(iChar)
	if err != nil {
		panic(err)
	}

	switch index {
	case firstIndex:
		{

			elementIndex = 0
			break
		}
	case midIndex:
		{
			elementIndex = int(iChar[fChar2-2])

			break
		}
	case lastIndex:
		{
			elementIndex = int(iChar[fChar2-1])
			break
		}
	default:
		{
			elementIndex = int(iChar[fChar2-2])
		}

	}
	return elementIndex
}

// FileDataToString return converts file data to string
func FileDataToString(filepath string) []string {
	//// o.App, os.Create
	fs, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fs.Close()
	lines := []string{}
	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		//// with this we will get the data at each line
		if trimmedLine != "" {
			lines = append(lines, trimmedLine)
		}
	}
	return lines
}
