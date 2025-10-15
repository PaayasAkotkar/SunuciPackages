// Package dataset handful of algorithms to work with string and databse
package dataset

import (
	"encoding/json"
	"fmt"
	"maps"
	"sort"
	"strings"
)

//	SearchSystem_ handful of hand to store the data
//
// this is the parent
// key=[]string
type SearchSystem_ struct {

	// events, main-categories, sub-categories, fields, items, sub-items
	// note: items and list-of-items can be called as mapped items
	set map[string]map[string]map[string]map[string]map[string][]string
	// example:
	// category: sports
	// field: international
	// book: cricket
	// item: india
	// sub-items: kohli, sharma, bumrah
	defaults map[string]bool
}

type Parcel struct {
	Event       string              `json:"Event"`
	Category    string              `json:"Category"`
	Field       string              `json:"Field"`
	Book        string              `json:"Book"`
	MappedItems map[string][]string `json:"MappedItems"`
}

// Docs returns the documentation of the keywords
func (ss *SearchSystem_) Docs() string {
	doc := `
	the interface is build upon 7 main category:
	#event: any-key that is related to the database
	#category
	#field
	#book
	#item
	#sub-item
	understanding: 
    suppose to include cricket team india players:
	+++key: 2011-world-cup // new event
    +++event: sports // new category
    +++main-category: cricket // new book
    +++sub-category:international // new field
    +++item: india 
    +++sub-items: kohli, dhoni, tendulkar

	rules:
	# key must be the named after the events
	example: to store icc world cup 2011 events
	key= icc world cup 2011 ✅
	key= 2011 icc world cup ❎
	# use a single map containing the named of the specific to put the items in-it
	example: to store icc world cup 2011 players
	✅
	map[india]=[]string{....}
	map[austrilia]=[]string{....}
	❎
	map[ind]=[]string{...}
	map[australia-icc-world-cup-2011]=[]string{...}
	`
	return doc
}

// Constructor default Constructor
func (ss *SearchSystem_) Constructor() *SearchSystem_ {
	events := []string{"2025-international-cricket-team"}
	category := []string{"sports", "entertainment"}
	books := []string{"cricket", "movies"} // for sports its cricket
	fields := []string{"international", "national", "domestic"}
	// for cricket international players
	MapItems := map[string][]string{
		"india": {"rohit sharma", "virat kohli", "jasprit bumrah", "ravindra jadeja", "mohammad shami",
			"mohammed siraj", "kannur lokesh rahul",
			"shubman gill", "hardik pandya", "suryakumar yadav", "rishabh pant", "kuldeep yadav", "axar patel", "yasIncludesvi jaiswal",
			"rinku singh", "tilak verma", "ruturaj gaikwad", "shardul thakur", "shivam dube", "ravi bishnoi", "jitesh sharma",
			"washington sundar", "mukesh kumar", "sanju samson", "arshdeep singh", "kona srikar bharat", "prasidh kirshna", "avesh khan",
			"rajat patidar", "sarfaraz khan", "dhruv jurel"},

		"austrillia": {"sean abbot", "xavier bartlett", "scott boland", "alex carey",
			"pat cummins", "nathan ellis", "cameron green", "aaron hardie",
			"josh hazlewood", "travis head", "josh inglis",
			"usman khawaja", "marnus labuschagne", "nathan lyon", "mitchell marsh",
			"glenn maxwell", "lance morris", "todd murphy", "jhye richardson",
			"matt short", "steve smith", "mitchell starc", "adam zampa"},

		"england": {
			"Gus Atkinson", "Harry Brook", "Jos Buttler", "Joe Root", "Jamie Smith", "Ben Stokes", "Mark Wood",
			"Rehan Ahmed", "Jofra Archer", "Jonny Bairstow", "Shoaib Bashir", "Brydon Carse", "Zak Crawley",
			"Sam Curran", "Ben Duckett", "Will Jacks", "Jack Leach", "Liam Livingstone", "Ollie Pope", "Matthew Potts",
			"Adil Rashid", "Phil Salt", "Olly Stone", "Josh Tongue", "Reece Topley",
			"Chris Woakes", "Jacob Bethell", "Josh Hull", "John Turner"},

		// to do: make sure the input for the names like new zeland south africa
		// to be converted into newzeland southafrica
		// note: player having three names or four to be converted into two names making it first name and last name only
		"newzeland": {
			"Tom Blundell", "Michael Bracewell", "Mark Chapman", "Josh Clarkson", "Jacob Duffy",
			"Matt Henry", "Kyle Jamieson", "Tom Latham", "Daryl Mitchell", "Henry Nicholls", "Will O’Rourke",
			"Ajaz Patel", "Glenn Phillips", "Rachin Ravindra", "Mitchell Santner",
			"Ben Sears", "Nathan Smith", "Ish Sodhi", "Tim Southee", "Will Young",
		},

		"southafrica": {
			"Temba Bavuma", "David Bedingham", "Nandre Burger", "Gerald Coetzee",
			"Tony de Zorzi", "Reeza Hendricks", "Marco Jansen", "Keshav Maharaj",
			"Kwena Maphaka", "Aiden Markram", "Wiaan Mulder", "Senuran Muthusamy", "Lungi Ngidi",
			"Kagiso Rabada", "Ryan Rickelton", "Tristan Stubbs", "Kyle Verreynne",
			"Lizaad Williams", "David Miller", "Rassie van der Dussen",
		},
		"westindies": {
			"Alick Athanaze", "Kraigg Brathwaite", "Keacy Carty", "Tagenarine Chanderpaul", "Joshua Da Silva", "Jason Holder",
			"Shai Hope", "Akeal Hosein", "Alzarri Joseph", "Brandon King", "Kyle Mayers",
			"Gudakesh Motie", "Nicholas Pooran", "Rovman Powell", "Kemar Roach", "Jayden Seales", "Romario Shepherd",
		},
		// note this not the acutal 2025 squad tho
		"srilanka": {
			"Charith Asalanka", "Pathum Nissanka", "Avishka Fernando",
			"Kusal Mendis", "Kamindu Mendis", "Janith Liyanage", "Nishan Madushka",
			"Nuwanidu Fernando", "Wanindu Includesaranga", "Maheesh Theekshana", "Dunith Wellalage", "Jeffrey Vandersay", "Asitha Fernando", "Lahiru Kumara", "Mohamed Shiraz", "Eshan Malinga",
		},
	}
	d := map[string]bool{}
	d["2025-international-cricket-team"] = true
	set := map[string]map[string]map[string]map[string]map[string][]string{}
	set[events[0]] = map[string]map[string]map[string]map[string][]string{
		category[0]: {
			fields[0]: {
				books[0]: MapItems,
			},
		},
	}
	ss.set = set
	ss.defaults = d
	return ss
}

// See the default value
func (ss *SearchSystem_) See() map[string]map[string]map[string]map[string]map[string][]string {
	_copy := make(map[string]map[string]map[string]map[string]map[string][]string)
	maps.Copy(_copy, ss.set)
	return _copy
}

// AEvents all the events in the dataset
func (ss *SearchSystem_) AEvents() []string {
	doms := make([]string, 0, len(ss.set))
	for r := range ss.set {
		doms = append(doms, r)
	}
	return doms
}

// ACategory all the events in the dataset
func (ss *SearchSystem_) ACategory() []string {
	doms := make([]string, 0, len(ss.set))
	for events := range ss.set {
		for categoires := range ss.set[events] {
			doms = append(doms, categoires)
		}
	}
	return doms
}

// AFields all the sub categories in the dataset
func (ss *SearchSystem_) AFields(_sort bool) []string {
	SC := []string{}
	set := ss.set
	for events := range set {
		for categoires := range set[events] {
			for fields := range set[events][categoires] {
				SC = append(SC, fields)
			}
		}
	}
	if _sort {
		sort.Strings(SC)
	}
	return SC
}

// ABooks all the items in the dataset
func (ss *SearchSystem_) ABooks(_sort bool) []string {
	B := []string{}
	set := ss.set
	for events := range set {
		for categoires := range set[events] {
			for fields := range set[events][categoires] {
				for books := range set[events][categoires][fields] {
					B = append(B, books)
				}
			}
		}
	}
	if _sort {
		sort.Strings(B)
	}
	return B
}

// AFileds all the fields in the dataset
// func (ss *SearchSystem_) AFileds(_sort bool) []string {
// 	F := []string{}
// 	set := ss.set
// 	for r := range set {
// 		for r2 := range set[r] {
// 			for r3 := range set[r][r2] {
// 				for r4 := range set[r][r2][r3] {
// 					for r5 := range set[r][r2][r3][r4] {
// 						F = append(F, r5)
// 					}
// 				}

// 			}
// 		}
// 	}
// 	if _sort {
// 		sort.Strings(F)
// 	}
// 	return F
// }

// AItems all the items in the dataset
func (ss *SearchSystem_) AItems(_sort bool) []string {
	I := []string{}
	set := ss.set
	for events := range set {
		for categoires := range set[events] {
			for fields := range set[events][categoires] {
				for books := range set[events][categoires][fields] {
					for items := range set[events][categoires][fields][books] {
						I = append(I, items)
					}
				}
			}
		}
	}
	if _sort {
		sort.Strings(I)
	}
	return I
}

// ASubItems all the sub items in the dataset
func (ss *SearchSystem_) ASubItems(_sort bool) []string {
	SI := []string{}
	set := ss.set
	for events := range set {
		for categoires := range set[events] {
			for fields := range set[events][categoires] {
				for books := range set[events][categoires][fields] {
					for items := range set[events][categoires][fields][books] {
						SI = append(SI, set[events][categoires][fields][books][items]...)
					}
				}
			}
		}
	}
	if _sort {
		sort.Strings(SI)
	}
	return SI
}

// IncludesField true if the value found in the dataset
// TIP you can even pass the forename, surname, or any other character that is separated via whitespace(for example: apple pie)
// NOTE it searches through the list all fields
func (ss *SearchSystem_) IncludesField(item string) bool {
	item = strings.ToLower(item)
	items := ss.AFields(false)
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// IncludesItem true if the value found in the dataset
// NOTE it searches through the list all items
func (ss *SearchSystem_) IncludesItem(item string) bool {
	item = strings.ToLower(item)
	items := ss.AItems(false)
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// IncludesBook true if the value found in the dataset
// NOTE it searches through the list all sub category
func (ss *SearchSystem_) IncludesBook(item string) bool {
	item = strings.ToLower(item)
	items := ss.AFields(false)
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// IncludesCategory true if the value found in the dataset
// NOTE it searches through the list all main category
func (ss *SearchSystem_) IncludesCategory(item string) bool {
	item = strings.ToLower(item)
	items := ss.ABooks(false)
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// IncludesEvent true if the value found in the dataset
// NOTE it searches through the list all events
func (ss *SearchSystem_) IncludesEvent(item string) bool {
	item = strings.ToLower(item)
	items := ss.AEvents()
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// PosItem the position or index of the items from whole collection
// if not found returns -1
func (ss *SearchSystem_) PosItem(item string, _in []string) int {
	items := _in
	found := 0
	for !Includes(items, item) {
		found += 1
		if found > len(items) {
			found = -1
			break
		}
	}
	return found
}

func (ss *SearchSystem_) mapEtoLower(m map[string][]string) map[string][]string {
	for _, v := range m {
		for r, r2 := range v {
			v[r] = strings.ToLower(r2)
		}
	}
	return m
}

// Include  adds the new key and respective value in it
func (ss *SearchSystem_) Include(event string, category string, book string, MappedItems map[string][]string, international bool, domestic bool, national bool) {
	for r := range ss.set {
		if event == r {
			panic(event + "cannot override the key")
		}
	}
	// all the elements to lower cases
	newCategory := strings.ToLower(category)
	newBook := strings.ToLower(book)
	newMappedItems := ss.mapEtoLower(MappedItems)

	set := map[string]map[string]map[string]map[string]map[string][]string{}

	if international && national && domestic || !international && !national && !domestic {
		panic("accepts only a single true")
	} else if international && national {
		panic("accepts only a single true")

	} else if national && domestic {
		panic("accepts only a single true")
	}

	switch true {
	case international:
		set[event] = map[string]map[string]map[string]map[string][]string{
			newCategory: {
				"international": {
					newBook: newMappedItems,
				},
			},
		}
	case national:
		set[event] = map[string]map[string]map[string]map[string][]string{
			newCategory: {
				"national": {
					newBook: newMappedItems,
				},
			},
		}
	case domestic:
		set[event] = map[string]map[string]map[string]map[string][]string{
			newCategory: {
				"domestic": {
					newBook: newMappedItems,
				},
			},
		}
	}
	ss.set = set
}

// RemoveEvent removes the key and associated data with it
func (ss *SearchSystem_) RemoveEvent(event string) {
	s := SearchSystem_{}
	for r := range s.defaults {
		fmt.Println(r, event)
		if r == event {
			panic("cannot remove default event")
		}
	}
	if !ss.IncludesEvent(event) {
		panic("there's no domin named " + event + " exists")
	}
	delete(ss.set, event)
}

// ReplaceItemInEvent replace the item in the event in the item category
func (ss *SearchSystem_) ReplaceItemInEvent(event string, search string, replace string) {
	s := ss.set
	for category := range s[event] {
		for field := range s[event][category] {
			for book := range s[event][category][field] {
				for item := range s[event][category][field][book] {
					for _, items := range s[event][category][field][book][item] {
						if search == items {
							Replace(s[event][category][field][book][item], search, replace)
						}

					}
				}
			}
		}
	}
	ss.set[event] = s[event]
}

// EventsCount total number of keys
func (ss *SearchSystem_) EventsCount() int {
	count := len(ss.AEvents())
	return count
}

// MainCategoriesCount total number of categories
func (ss *SearchSystem_) MainCategoriesCount() int {
	count := len(ss.ABooks(false))
	return count
}

// FieldsCount returns total number of items
func (ss *SearchSystem_) FieldsCount() int {
	count := len(ss.AFields(false))
	return count
}

// ItemCount total number of items
func (ss *SearchSystem_) ItemCount() int {
	count := len(ss.AItems(false))
	return count
}

// BookCount returns total number of items
func (ss *SearchSystem_) BookCount() int {
	count := len(ss.ASubItems(false))
	return count
}

// Pack current data to string
func (ss *SearchSystem_) Pack(event string) string {

	_, category, field, book, mappedItems := ss.Manager(event)

	token := Parcel{
		Event: event, Category: category, Field: field, Book: book, MappedItems: mappedItems,
	}

	packed, _ := json.Marshal(&token)
	parcel := string(packed)
	return parcel
}

// // Shift shifts the value in the key or sorts the collection if true
// func (ss *SearchSystem_) Shift(event string, inItem string, appenditem string, _sort bool) {
// 	s := ss.set
// 	if ss.defaults[event] {
// 		panic("cannot shift value in default event")
// 	}
// 	for r := range s[event] {
// 		for r2 := range s[event][r] {
// 			for r3 := range s[event][r][r2] {
// 				for r4 := range s[event][r][r2][r3] {
// 					for r5 := range s[event][r][r2][r3][r4] {
// 						fmt.Println(r5, inItem)
// 						if r5 == inItem {
// 							s[event][r][r2][r3][r4][r5] = append(s[event][r][r2][r3][r4][r5], appenditem)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	ss.set[event] = s[event]
// }

// Match true if the element that matches in the given key
func (ss *SearchSystem_) Match(inItem string, search string) bool {
	// equalize the string
	s := ss.set
	search = strings.ToLower(search)
	found := false
	event := ""
	for _event := range ss.set {
		event = _event
	}
	// spearate the names in first and last
	_names := []string{}
	for category := range s[event] {
		for field := range s[event][category] {
			for book := range s[event][category][field] {
				for item := range s[event][category][field][book] {
					for _, items := range s[event][category][field][book][item] {
						if items == inItem {
							fmt.Println(true)
							_names = append(_names, s[event][category][field][book][item]...)
						}
					}
				}
			}
		}
	}
	_names = ParseWords(_names)

	for _, r := range _names {
		if r == search {
			found = true
		}
	}
	return found
}

// Manager separates the event, category, field, book, mapped item lists
func (ss *SearchSystem_) Manager(event string) (string, string, string, string, map[string][]string) {
	_category, _field, _book, _items := "", "", "", make(map[string][]string)

	if _, ok := ss.set[event]; ok {
		for category := range ss.set[event] {
			_category = category
			for field := range ss.set[event][category] {
				_field = field
				for book := range ss.set[event][category][field] {
					_book = book
					maps.Copy(_items, ss.set[event][category][field][book])

				}
			}
		}
	} else {
		event = "event not found"
	}
	return event, _category, _field, _book, _items
}

func (ss *SearchSystem_) LazyView() (string, string, string, string, map[string][]string) {
	event, _category, _field, _book, _items := "", "", "", "", make(map[string][]string)

	for _event := range ss.set {
		event = _event
	}

	if _, ok := ss.set[event]; ok {
		for category := range ss.set[event] {
			_category = category
			for field := range ss.set[event][category] {
				_field = field
				for book := range ss.set[event][category][field] {
					_book = book
					maps.Copy(_items, ss.set[event][category][field][book])

				}
			}
		}
	} else {
		event = "event not found"
	}
	return event, _category, _field, _book, _items
}
