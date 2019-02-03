package main

import (
	"fmt"
)

func main() {
	qt, _ := BuildQuadTree("./samples/jungle.jpg")
	fmt.Printf("%+v\n", qt)
}
