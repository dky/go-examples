package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs int
}

type Artist struct {
	Name, Genre string
}


func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}


func main() {
	me := &Artist{Name: "Don", Genre: "Ukulele", Songs: 43}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
