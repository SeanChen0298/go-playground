package main

import (
	"fmt"
	article "go-playground/articles"
	hashmap "go-playground/hashmap"
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

	hashmap.DesignHashMap()
	obj := hashmap.Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	obj.Put(2, 3)
	fmt.Println(obj.Get(2))
	obj.Remove(2)
	fmt.Println(obj.Get(2))

}
