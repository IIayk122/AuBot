package main

import (
	"fmt"

	au "github.com/IIayk122/AuBot/au"
)

func main() {
	a := au.Search("asus+k56cb")
	//fmt.Println(a)

	fmt.Println(a.Count)
	for _, row := range a.List {
		fmt.Println(row.ID, row.Name, row.Images[0].Origin, row.BlitzPrice, row.CurrentPrice)
	}

}
