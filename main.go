package main

import (
	"fmt"
	article "go-playground/articles"
)

func main() {
	fmt.Println("Go Playground!")

	// Variable and functions declared in same package is available to other files in same package.
	fmt.Println("Imported Name: ", Name)

	// Variable and functions declared in different package needs to be imported.
	// import "go mod/folder name"
	// reference the variable by package_name.var_name
	fmt.Println("Imported Article", article.Article)
}
