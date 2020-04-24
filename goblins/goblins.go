package goblins

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"gopkg.in/yaml.v2"
)

//Generator for gobos
func Generator(goblinCount, mountCount int, md bool) bool {
	if goblinCount < 0 && mountCount < 0 {
		return false
	}

	gg := GoblinGen{}

	goblinBits, err := ioutil.ReadFile("goblins/goblins.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(goblinBits), &gg)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("loaded: %v\n\n", gg)
	randy := rand.New(rand.NewSource(time.Now().UnixNano()))
	gg.Randy = randy

	if goblinCount > 0 {
		goblins := gg.GenGoblins(goblinCount)
		goblins.Print(md)
	}

	if mountCount > 0 {
		mounts := gg.GenGoblinMounts(mountCount)
		mounts.Print(md)
	}
	return true
}

//GoblinGen yaml file
type GoblinGen struct {
	Goblin struct {
		Name       []string
		Appearance []string
		Equipment  []string
		Behavior   []string
	}
	Mount struct {
		Feature []string
	}
	Randy *rand.Rand
}

//Print em
func (goblins Goblins) Print(md bool) {
	if md {
		fmt.Printf("## %d Goblins!\n\n", len(goblins))
		fmt.Println("|   No | Name | Appear | Equip | Behavior |")
		fmt.Println("|-----:|:-----|:-------|:------|:---------|")
	}
	for i, gob := range goblins {
		if md {
			fmt.Printf("| %d | %s | %s | %s | %s |\n", 1+i, gob.name, gob.appearance, gob.equipment, gob.behavior)
			if i%6 == 5 {
				fmt.Printf("| %[1]s | %[1]s | %[1]s | %[1]s | %[1]s |\n", " - ")
			}
			continue
		}
		fmt.Printf("      name: %s\n", gob.name)
		fmt.Printf("appearance: %s\n", gob.appearance)
		fmt.Printf(" equipment: %s\n", gob.equipment)
		fmt.Printf("  behavior: %s\n", gob.behavior)
		fmt.Println()
	}
	fmt.Println()
}

//Print print em
func (mounts Mounts) Print(md bool) {
	if md {
		fmt.Printf("## %d Goblin Mounts!\n\n", len(mounts))
		fmt.Println("|   No | Description |")
		fmt.Println("|-----:|:-----|")
	}
	for i, mount := range mounts {
		if md {
			fmt.Printf("| %d | %s |\n", 1+i, mount)
			if i%6 == 5 {
				fmt.Printf("| %[1]s | %[1]s |\n", " - ")
			}
			continue
		}
		fmt.Printf("description: %s\n", mount)
		fmt.Println()
	}
	fmt.Println()
}

func (gg *GoblinGen) randomString(s []string) string {
	//log.Printf("rolling for len: %d - %v\n", len(s), s)
	i := gg.Randy.Intn(len(s))
	//log.Printf("rolled: %d, for len: %d - %v\n", i, len(s), s)
	return s[i]
}

//Goblin short, ugly, mean
type Goblin struct {
	name       string
	appearance string
	equipment  string
	behavior   string
}

//Goblins gobos
type Goblins []Goblin

//Mounts gobo mounts
type Mounts []string

//GenGoblins make some gobos
func (gg *GoblinGen) GenGoblins(count int) Goblins {
	gobos := make([]Goblin, 0, count)
	for i := 0; i < count; i++ {
		gobos = append(gobos,
			Goblin{
				name:       gg.randomString(gg.Goblin.Name),
				appearance: gg.randomString(gg.Goblin.Appearance),
				equipment:  gg.randomString(gg.Goblin.Equipment),
				behavior:   gg.randomString(gg.Goblin.Behavior),
			})
	}
	return gobos
}

//GenGoblinMounts mounts
func (gg *GoblinGen) GenGoblinMounts(count int) Mounts {
	descr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		descr = append(descr, gg.randomString(gg.Mount.Feature))
	}
	return descr
}
