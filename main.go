package main

import (
	"flag"
	"fmt"
	"os"

	"labrynth/goblins"
	"labrynth/stonewalls"
)

func main() {
	mdPtr := flag.Bool("md", false, "genrate markdown")
	goblinCountRef := flag.Int("goblins", 0, "how many to generate")
	goblinMountCountRef := flag.Int("goblin-mounts", 0, "how many to generate")
	stonewallsCountRef := flag.Int("stonewalls", 0, "how many between scenes generate for stone walls")
	flag.Parse()

	//log.Printf("FLAGS: md:%v goblins:%d mounts:%d stonewalls:%d\n", *mdPtr, *goblinCountRef, *goblinMountCountRef, *stonewallsCountRef)

	if *mdPtr {
		fmt.Printf("# Random Labyrinth Stuff!\n\n")
	}
	handled := stonewalls.Generator(*stonewallsCountRef, *mdPtr)
	handled = goblins.Generator(*goblinCountRef, *goblinMountCountRef, *mdPtr) || handled

	if !handled {
		fmt.Printf("nothign to do. try executing %s -help\n", os.Args[0])
	}
}
