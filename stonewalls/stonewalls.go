package stonewalls

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"gopkg.in/yaml.v2"
)

//Generator for between the stone walls
func Generator(stonewallsCount int, md bool) bool {
	if stonewallsCount < 1 {
		return false
	}
	bsg := BetweenScenesGen{}

	bits, err := ioutil.ReadFile("stonewalls/stonewalls.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(bits), &bsg)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("loaded: %+v\n\n", bsg)
	randy := rand.New(rand.NewSource(time.Now().UnixNano()))
	bsg.Randy = randy

	if stonewallsCount > 0 {
		betweens := bsg.GenBetweens(stonewallsCount)
		betweens.Print(md)
	}
	return true
}

//Between in between the stone walls
type Between struct {
	Dressing, Encounter, Weater string
}

//Betweens several of them
type Betweens []Between

//BetweenScenesGen yaml file
type BetweenScenesGen struct {
	Dressing   []string
	Encounters []string
	Weather    []string
	Randy      *rand.Rand
}

//GenBetweens make some
func (bsg *BetweenScenesGen) GenBetweens(count int) Betweens {
	btw := make(Betweens, 0, count)
	for i := 0; i < count; i++ {
		btw = append(btw, Between{
			Dressing:  bsg.randomString(bsg.Dressing),
			Encounter: bsg.randomString(bsg.Encounters),
			Weater:    bsg.randomString(bsg.Weather),
		})
	}
	return btw
}

//randomString get one
func (bsg *BetweenScenesGen) randomString(s []string) string {
	i := bsg.Randy.Intn(len(s))
	return s[i]
}

//Print em
func (btw Betweens) Print(md bool) {
	if md {
		fmt.Printf("## %d Between Scenes in the Stone Walls (Area 1)!\n\n", len(btw))
		fmt.Println("|   No | Dressing | Encounter | Weather |")
		fmt.Println("|-----:|:---------|:----------|:--------|")
	}
	for i, swEnc := range btw {
		if md {
			fmt.Printf("| %d | %s | %s | %s |\n", 1+i, swEnc.Dressing, swEnc.Encounter, swEnc.Weater)
			if i%6 == 5 {
				fmt.Printf("| %[1]s | %[1]s | %[1]s | %[1]s |\n", " - ")
			}
			continue
		}
		fmt.Printf(" dressing: %s\n", swEnc.Dressing)
		fmt.Printf("encounter: %s\n", swEnc.Encounter)
		fmt.Printf("  weather: %s\n", swEnc.Weater)
		fmt.Println()
	}
	fmt.Println()
}
