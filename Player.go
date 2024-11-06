package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Players = make(map[string]*Player)
)

type Player struct {
	Name                    string //PrimaryID // not nullable
	Username                string
	Years                   int
	Health                  int
	PrimaryAbilityRessource int
}

func (p Player) display() {
	fmt.Printf("Name: %v\n Username: %v\n Years: %v\n Health: %v\n PrimaryAbilityRessource: %v\n ", p.Name, p.Username, p.Years, p.Health, p.PrimaryAbilityRessource)
}

func (p Player) save() {
	NameofFile := p.Name
	NameofFileWithExt := fmt.Sprintf("%s.yml", NameofFile)
	d, err := yaml.Marshal(&p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	file, err := os.Create(NameofFileWithExt)
	if err != nil {
		log.Fatalf("Error when creating file : %v", err)
	}
	os.WriteFile(NameofFileWithExt, d, 0600)
	defer file.Close()
}

func (p Player) del() {
	delete(Players, p.Name)
	NameofFile := p.Name
	NameofFileWithExt := fmt.Sprintf("%s.yml", NameofFile)

	err := os.Remove(NameofFileWithExt)
	if err != nil {
		log.Fatalf("Error when delete file: %v", err)
	}

	fmt.Printf("file succesfuly destroyed: %s\n", NameofFileWithExt)
}

func (p Player) loadfromfile() {
	NameofFile := p.Name
	NameofFileWithExt := fmt.Sprintf("%s.yml", NameofFile)

	file, err := os.Open(NameofFileWithExt)
	if os.IsNotExist(err) {
		fmt.Printf("File %s not found, creating and saving player.\n", NameofFileWithExt)
		p.save()
		return
	} else if err != nil {
		log.Fatalf("Error when opening file: %v", err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error when reading in file : %v", err)
	}

	err = yaml.Unmarshal([]byte(content), &p)
	if err != nil {
		log.Fatalf("error when unmarshal: %v", err)
	}

	Players[p.Name] = &p

}

func AddPlayer() Player {
	fmt.Printf("\n what's your name ? : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)
	scan.Scan()
	ChoosenName := scan.Text()
	log.Printf("your real name : %v", ChoosenName)

	fmt.Printf("\n what's your username ? : ")
	scan.Scan()
	ChoosenUsername := scan.Text()
	log.Printf("your username : %v", ChoosenUsername)

	p := Player{
		Name:     ChoosenName,
		Username: ChoosenUsername,
	}

	Players[ChoosenName] = &p

	Players[ChoosenName].loadfromfile()

	return *Players[ChoosenName]
}

/// function out of class

func playerLoad(name string) Player {
	//this is only to check if exist on map
	var PlayertoReturn Player
	if Players[name] != nil {
		Players[name].display()
		PlayertoReturn = *Players[name]
	} else {
		PlayertoReturn = AddPlayer()
	}
	return PlayertoReturn
}

/// seed

func seeding() {
	p1 := Player{
		Name:                    "Bob",
		Username:                "Bobby",
		Years:                   24,
		Health:                  250,
		PrimaryAbilityRessource: 25,
	}
	p2 := Player{
		Name:                    "Clara",
		Username:                "Raclette",
		Years:                   20,
		Health:                  150,
		PrimaryAbilityRessource: 50,
	}
	p3 := Player{
		Name:                    "Julien",
		Username:                "PerceBallon",
		Years:                   22,
		Health:                  450,
		PrimaryAbilityRessource: 10,
	}
	//we add only p1 / p2 to the map
	Players[p1.Name] = &p1
	Players[p2.Name] = &p2
	//we create p1 / p3 yaml
	p1.save()
	p3.save()
}

/*
from the github
	yaml to struct
	err := yaml.Unmarshal([]byte(data), &p)
	if err != nil {
	log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	struct to yaml
	d, err := yaml.Marshal(&p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
	m := make(map[interface{}]interface{})

	yaml to struct
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
	struct to yaml
	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
*/
