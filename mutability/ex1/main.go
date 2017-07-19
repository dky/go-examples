package main

import (
	"fmt"
)

type Artist struct {
	Name, Genre string
	Songs int
}

//New release expects an object with the type Artist
func newRelease(a Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	//declare an object/type me and assign it values
	me := Artist{Name: "Don", Genre: "Ukulele", Songs: 51}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
