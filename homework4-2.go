package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	//Firstly, compare firstName n lastName strings
	//then if they are equal for both Persons
	//choose a youngest Person

	if strings.Compare(p[i].firstName, p[j].firstName) == 0 {
		if strings.Compare(p[i].lastName, p[j].lastName) == 0 {
			return p[j].birthDay.Before(p[i].birthDay)
		} else {
			return strings.Compare(p[i].lastName, p[j].lastName) < 0
		}
	} else {
		return strings.Compare(p[i].firstName, p[j].firstName) < 0
	}

}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	ivanIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	ivanIvanovDate1, err := time.Parse("2006-Jan-02", "2010-Aug-10")
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate}, {"Ivan", "Ivanov", ivanIvanovDate1}}
	fmt.Println(sort.Sort(p))
	//fmt.Println(p)

}
