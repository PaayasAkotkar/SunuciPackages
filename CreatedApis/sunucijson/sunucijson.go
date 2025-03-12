package sunucijson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sunuciApi/sunucistring"
)

// @RETURN: writes the json array in the given jFile
// @NOTE: the file suppose to be in the directory
// @NOTE: right now the json data ends with a comma but in future this issue will be solved
// @NOTE: it is for the array json-format
func PushJData(jFile string, i interface{}) {

	g, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	fs, err := os.OpenFile(jFile, os.O_APPEND|os.O_RDWR|os.O_SYNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	fs.Write([]byte(g))
	fs.Seek(-1, io.SeekEnd)
	fs.Write([]byte(","))
	fs.Seek(1, io.SeekStart)
	fs.Write([]byte("\n"))
}

/**
* @NOTE: it is not test with object Json-Format
* object Json-format:[ "compilerOption":{"strict":"true"},""..]
* it is designed for any array Json-format
* [{},{},{}]
* runs the data into json file
 */
func ToJsonFormat(path string) {
	lines := sunucistring.FileDataToString(path)
	j := ""
	j = strings.Join(lines, "")
	// string to rune to work with characters
	ru := []rune(j)

	// find the eof in json file in-between
	comp := regexp.MustCompile("]{")
	// finding the index to easily remove
	inx := comp.FindIndex([]byte(j))
	comp2 := regexp.MustCompile(`\[`)
	inx2 := comp2.FindIndex([]byte(j))

	// important else the range out of index
	if len(inx) != 0 {
		ru[inx[0]] = '\n'
	}
	// append the opening tag to the json file
	if len(inx2) == 0 {
		// in-short push front or prepend
		ru = append([]rune{'['}, ru...)
	}
	// check for the last char having eof
	if ru[len(ru)-1] != ']' {
		ru = append(ru, ']') // append the closing tag
	}

	// convert the char to string
	j = string(ru)
	// write the file
	err := os.WriteFile(path, []byte(j), 0644)
	if err != nil {
		panic(err)
	}

}

// @NOTE what you wish to separate must send in raw string
// raw string: “
// @NOTE: if data {"username":"afafa"} than please send via "username
// @NOTE: it is a case-insensitive
// @NOTE: it is coded for the array json-format not for object json-format
// separates a single data
// meaning it cannot pass a single data of a client rather it will pass all the username name of client if wished
// @@TODO:make sure to remove whitespace by using regex and just replace All rather than using rune[]
func JsontoExcel(Json_file string, separate string, CSV_file string, bValuesOnly, bQuotationMark bool, bRemoveComma bool,
	_key bool, _keynvalue bool, _value bool, row bool) string {
	lines := sunucistring.FileDataToString(Json_file) // conv data from json to string
	str := strings.Join(lines, "")                    // []string to string conv

	/** base process
	* a base process is what the client has requested to separate
	* for example base process for getting username only
	 */
	// pattern for searching through regex
	// it basically says: (x,y) meaning involve everything b/w this range
	pat := regexp.QuoteMeta(separate) + (`.*?`) + (`,`)
	re, err := regexp.Compile(pat)
	if err != nil {
		panic(err)
	}
	// separtion process for the base
	us := re.FindAllStringSubmatch(str, -1)
	fulldata := []string{}
	for _, r := range us {
		fulldata = append(fulldata, r...)
	}
	// []string to string conv
	// if \n not given the file will be in the column format
	// menaing in the horiztall format in the csv
	keys_and_values := ""
	if row {
		keys_and_values = strings.Join(fulldata, "\n")
	} else {
		keys_and_values = strings.Join(fulldata, "")
	}
	/** base process end*/

	/**manipulation process
	* a mainpulation process: nothing but separting the key meaning the base
	* for example separate is username than it is the key
	* for example sepearate is username than it's value with key
	* for example separte is username than it can have value
	**/

	pat2, err := regexp.Compile(separate + "\": ")
	if err != nil {
	}
	usernames := pat2.Split(keys_and_values, -1)

	// values stored at key username
	values := strings.Join(usernames, "")

	pat3, err := regexp.Compile(separate)
	if err != nil {
	}

	keys := pat3.Find([]byte(keys_and_values))
	// key name at those values stored
	users := string(keys)

	/**end of manipulation processs*/
	ru := []rune(keys_and_values)
	ru2 := []rune(values)
	ru3 := []rune(users)

	// remove the quoataion mark
	for r := range ru3 {
		if ru3[r] == '"' {
			ru3[r] = ' '
		}
	}
	// for like removing commas and quotation mark
	for r := range ru {
		switch true {
		case bQuotationMark && bRemoveComma:
			{
				if ru[r] == '"' {
					ru[r] = ' '
				}
				if ru[r] == ',' {
					ru[r] = ' '
				}
				break
			}
		case bQuotationMark:
			{
				if ru[r] == '"' {
					ru[r] = ' '
				}
				break
			}
		case bRemoveComma:
			{
				if ru[r] == ',' {
					ru[r] = ' '
				}
				break
			}
		default:
			{

			}
		}

	}
	// for like removing commas and quotation mark
	for r := range ru2 {
		switch true {
		case bQuotationMark && bRemoveComma:
			{
				if ru2[r] == '"' {
					ru2[r] = ' '
				}
				if ru2[r] == ',' {
					ru2[r] = ' '
				}
				break
			}
		case bQuotationMark:
			{
				if ru2[r] == '"' {
					ru2[r] = ' '
				}
				break
			}
		case bRemoveComma:
			{
				if ru2[r] == ',' {
					ru2[r] = ' '
				}
				break
			}
		default:
			{

			}
		}

	}
	keys_and_values = string(ru)
	values = string(ru2)
	users = string(ru3)

	switch true {
	case _key:
		{
			return users
		}
	case _value:
		{
			return values

		}
	case _keynvalue:
		{
			return keys_and_values
		}
	default:
		{
			return keys_and_values
		}
	}

}

// incomplete
// @TODO: just do for loop and write in the file where even number goes to the row and odd number to the column
func WriteCSV(csv_file string) {
	fp := "data.json"

	// getting the single key in column format
	un := JsontoExcel(fp, `"Username`, "", true, true, true, true, false, false, true)
	pas := JsontoExcel(fp, `"Password`, "", true, true, true, true, false, false, true)
	ph := JsontoExcel(fp, `"PhoneNumber`, "", true, true, true, true, false, false, true)
	ln := JsontoExcel(fp, `"LastNumber`, "", true, true, true, true, false, false, true)
	fs, err := os.OpenFile(csv_file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
	}
	defer fs.Close()
	// this will write the key in the column format in the csv
	fs.Write([]byte(un))
	fs.Write([]byte(pas))
	fs.Write([]byte(ph))
	fs.Write([]byte(ln))

	// getting the key's data in the row format
	unv := JsontoExcel(fp, `"Username`, "", false, true, false, false, false, true, false)
	// pasv := JsontoExcel(fp, `"Password`, "", false, true, false, false, false, true, false)

	// replacing all the whitespace
	// @NOTE: this issue will be solved in the future
	unv = strings.ReplaceAll(unv, " ", "")
	// pasv = strings.ReplaceAll(pasv, " ", "")
	re, err := regexp.Compile(",")
	if err != nil {
	}
	// only getting the values rather than comma
	users := re.Split(unv, -1)

	// this loop suupose to merge str[0]str2[0] rather it merges str[len-1...][str0]
	lk := []string{}
	for i := 0; i < len(users); i++ {
		if i%2 == 0 {
			lk = sunucistring.Shift("A", users, i)
		} else if i%3 == 0 || i%5 == 0 {
			lk = sunucistring.Shift("A", users, i)
		}

	}
	str := strings.Join(lk, "\n")
	fmt.Println(str)
}
