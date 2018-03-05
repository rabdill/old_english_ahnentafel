package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

// Person - one family member
type Person struct {
	ID         int    `json:"id"`
	Father     int    `json:"father"`
	Mother     int    `json:"mother"`
	Ahnentafel int    `json:"ahnentafel"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Birth      string `json:"birth"`
	Death      string `json:"death"`
	Bio        string `json:"bio"`
	Notes      string `json:"notes"`
}

// Corpus - all the family members to be numbered
type Corpus struct {
	People        []*Person
	DoubleEntries map[int]int
}

func (corpus *Corpus) findByID(id int) *Person {
	for i := 0; i < len(corpus.People); i++ {
		if corpus.People[i].ID == id {
			return corpus.People[i]
		}
	}
	return nil
}

func (person *Person) evaluateLine(family *Corpus, current int) {
	if person.Ahnentafel == 0 {
		person.Ahnentafel = current
		if person.Father != 0 {
			family.findByID(person.Father).evaluateLine(family, (person.Ahnentafel * 2))
		}
		if person.Mother != 0 {
			family.findByID(person.Mother).evaluateLine(family, ((person.Ahnentafel * 2) + 1))
		}
	} else {
		// If a person appears multiple times in the tree,
		// re-assign their ahnentafel number to the lowest
		// value it can be, and re-evaluate their ancestors
		// to match the new, lower numbers
		if current > person.Ahnentafel {
			family.DoubleEntries[current] = person.Ahnentafel
		} else {
			family.DoubleEntries[person.Ahnentafel] = current
			person.Ahnentafel = current
			if person.Father != 0 {
				family.findByID(person.Father).evaluateLine(family, (person.Ahnentafel * 2))
			}
			if person.Mother != 0 {
				family.findByID(person.Mother).evaluateLine(family, ((person.Ahnentafel * 2) + 1))
			}
		}
	}
}

func (corpus Corpus) printList() {
	// associate everyone with their ahnentafel number
	numbers := map[int]string{}
	for _, person := range corpus.People {
		numbers[person.Ahnentafel] = person.Name
	}
	for entry, ref := range corpus.DoubleEntries {
		numbers[entry] = strconv.Itoa(ref)
	}
	// get a slice of all the sorted numbers
	var keys []int
	for number := range numbers {
		keys = append(keys, number)
	}
	sort.Ints(keys)
	// print them in order
	for _, k := range keys {
		fmt.Printf("%d: %s\n", k, numbers[k])
	}

	// TODO: Remove chains of double entries. If there's a
	// double entry at 500 pointing at entry 20, we don't
	// also need a double entry at 1000 pointing at 40.
}

func main() {
	rawIn, err := ioutil.ReadFile("family.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fam Corpus
	json.Unmarshal(rawIn, &fam)
	if len(fam.People) == 0 {
		fmt.Println("Couldn't find anyone in the file.")
		os.Exit(1)
	}
	fam.DoubleEntries = make(map[int]int)
	fam.People[0].evaluateLine(&fam, 1)

	fam.printList()
}
