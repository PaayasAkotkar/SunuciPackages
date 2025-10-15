package sunucistring

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// @RETURN: shift any given value at provided index
func StringShift(shift_v string, from string, at_index int) string {
	str := strings.Split(from, "")
	str = Shift(shift_v, str, at_index)
	from = strings.Join(str, "")
	return from
}

// @RETURN: shift any given value at provided index
// @NOTE: you can use it for any given array type
func Shift[T any](shift_v T, from []T, at_index int) []T {
	temp := []T{}

	/**normal swapping process*/
	temp = append(temp, from...)
	// prepends the value at given index
	from = append(from[:at_index], shift_v)
	// appends back the value at given index after the appended value
	from = append(from, temp[len(from)-1:]...)
	/**endof swapping procss*/
	return from
}

// @RETURN: [][]string of between the limits (opne, close]
func Pattern(str string, _open string, close_ string) [][]string {
	pattern := regexp.QuoteMeta(_open) + `(.*?)` + regexp.QuoteMeta(close_)
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	s := re.FindAllStringSubmatch(str, -1)
	return s
}

/**process of enum**/
type EIndex int

const (
	first_index EIndex = iota
	last_index
	mid_index
)

/** end of process of enum**/

// @RETURN: value at the first, middle, last in the file
// @NOTE: not all
func IO_See_First_Middle_Last(file_path string, index EIndex) string {
	element := ""
	fs, err := os.OpenFile(file_path, os.O_RDONLY, 0644)
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
	case first_index:
		{
			element = string(i1Char[:fChar])
			break
		}
	case mid_index:
		{
			element = string(iChar[fChar/2])
			break
		}
	case last_index:
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

// @RETURN: index at the first, middle, last in the file
// @NOTE: not all
func IO_Get_First_Middle_Last(file_path string, index EIndex) int {
	element_index := 1
	fs, err := os.OpenFile(file_path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	fls, err := fs.Stat()
	if err != nil {
		panic(err)
	}

	// get the max len of the file
	iChar := make([]byte, fls.Size())

	fChar2, err := fs.Read(iChar)
	if err != nil {
		panic(err)
	}

	switch index {
	case first_index:
		{

			element_index = 0
			break
		}
	case mid_index:
		{
			element_index = int(iChar[fChar2/2])

			break
		}
	case last_index:
		{
			element_index = int(iChar[fChar2-1])
			break
		}
	default:
		{
			element_index = int(iChar[fChar2/2])
		}

	}
	return element_index
}

// @RETURN: converts file data to string
func FileDataToString(filepath string) []string {
	// o.App, os.Create
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
		// with this we will get the data at each line
		if trimmedLine != "" {
			lines = append(lines, trimmedLine)
		}
	}
	return lines
}

// @RETURN: index of found value
func GetIndex(str []string, search string) int {
	found := -1
	for r := range str {
		if str[r] == search {
			found = r
		}
	}
	if found > 0 {
		return found
	} else {
		return found - 2
	}
}

// @RETURN: replaces the value at given string
func Replace(str []string, search string, replace string) {
	// if not found it will not do anything
	found := -1
	for r := range str {
		if str[r] == search {
			found = r
			if found > 0 {
				str[found] = replace
			}
		}
	}

}

// @RETURN: last index of the string[]
func LastIndex(str []string) int {
	lasti, err := -1, 2
	for r := range str {
		lasti = r
	}
	if lasti == -1 {
		panic(err)
	}
	return lasti
}

// @RETURN: 2nd last index
func SecondLastIndex(str []string, search string) int {
	inx := []int{}
	for r := range str {
		if str[r] == search {
			inx = append(inx, r)
		}
	}
	f := -1
	if len(inx) > 1 {
		f = inx[len(inx)-2]
	} else if len(inx) == 1 {
		f = inx[0]
	} else {
		f = 0
	}
	return f
}

// RETURN: number of times the element repeated
func ElementRepeated(str []string, search string) int {
	inx := []int{}
	repeated := 0
	for r := range str {
		if str[r] == search {
			inx = append(inx, r)
		}
	}
	fmt.Println("lenght: ", len(inx))
	if len(inx) == 2 {
		fmt.Println("lenght: ", len(inx))
		repeated = len(inx) - 2

	}
	if len(inx) > 2 {
		fmt.Println("lendaa: ", len(inx))
		repeated = len(inx) - 1
	}
	return repeated

}

// @RETURN: true if the value is in the string
func Includes(str []string, search string) bool {
	has := false
	for r := range str {
		if str[r] == search {
			has = true
		}
	}
	return has
}

// @RETURN: last index of repeated value in a string
func GetLastRepeationIndex(str []string, search string) int {
	found := []int{}
	foundinx := -1
	for r := range str {
		if str[r] == search {
			found = append(found, r)
		}
	}
	if len(found) != 0 {
		foundinx = found[len(found)-1]
	}
	return foundinx
}
