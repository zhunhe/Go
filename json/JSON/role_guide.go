package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Roles struct {
	Roles []Role `json:"roles"`
}

type Role struct {
	RoleName  string   `json:"roleName"`
	RoleGuide []string `json:"roleGuide"`
	Max       int      `json:"max"`
	Faction   string   `json:"faction"`
}

func main() {
	// Open JSON file
	jsonFile, err := os.Open("role_guide.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	// JSON -> struct
	var roles Roles
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &roles)
	// Print JSON data
	for i := 0; i < len(roles.Roles); i++ {
		fmt.Println("index: " + strconv.Itoa(i))
		fmt.Println("RoleName: " + roles.Roles[i].RoleName)
		fmt.Println("Max: " + strconv.Itoa(roles.Roles[i].Max))
		for j := 0; j < len(roles.Roles[i].RoleGuide); j++ {
			fmt.Println("RoleGuide(" + strconv.Itoa(j) + "): " + roles.Roles[i].RoleGuide[j])
		}
		fmt.Println("Faction: " + roles.Roles[i].Faction)
		fmt.Println()
	}
}
