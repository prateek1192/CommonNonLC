package main

import (
	"fmt"
	"keyValueRippling/keyValueNestedTransaction"
)

func main() {
	tMap := keyValueNestedTransaction.NewKVStore()

	tMap.Set("a", "1")
	tMap.Set("b", "2")
	getA, _ := tMap.Get("a")
	getB, _ := tMap.Get("b")
	fmt.Println("After setting a and b:", getA, getB)

	tMap.Begin()
	tMap.Set("a", "3")
	tMap.Delete("b")
	fmt.Println("During transaction (a=3, b deleted):", getA, getB)

	tMap.Commit()
	fmt.Println("After commit:", getA, getB)

	tMap.Begin()
	tMap.Set("a", "4")
	tMap.Rollback()
	fmt.Println("After rollback:", getA)
}
