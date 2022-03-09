package main

import (
	"fmt"
	lrucache "go-playground/LRUCache"
	array "go-playground/arrays"
	article "go-playground/articles"
	hashmap "go-playground/hashmaps"
	linkedlist "go-playground/linkedlists"
)

func main() {

	// Run by using `go run .`

	fmt.Println("Go Playground!")

	// Variable and functions declared in same package is available to other files in same package.
	fmt.Println("\nImported Name: ", Name)

	// Variable and functions declared in different package needs to be imported.
	// import "go mod/folder name"
	// reference the variable by package_name.var_name
	fmt.Println("Imported Article", article.Article)

	hashmap.DesignHashMapTest()

	array.TwoSumTest()

	linkedlist.TestConstructLinkedList()

	lrucache.TestLRUCache()

	linkedlist.TestReverse()

	linkedlist.TestMergeLinkedList()
}
