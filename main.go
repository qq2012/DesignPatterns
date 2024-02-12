package main

import (
	abstractfactory "DesignPatterns/abstract_factory"
	"DesignPatterns/builder"
	"DesignPatterns/singleton"
	"fmt"
)

func main() {
	fmt.Println("hey")
	fmt.Println(abstractfactory.Run())
	fmt.Println(builder.Run())
	fmt.Println(singleton.Run())
	fmt.Println("done")
}
