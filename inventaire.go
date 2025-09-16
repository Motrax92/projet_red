package red

import (
	"strings"
	"fmt"	
)

type Equipement struct {
	head string
	Torso string 
	Fett string 
}

type Character struct {
	Name string 
	Class string 
	Level int
	HPMax int
	HP int
	Inventory map[string]int
	Skills []string
	Gold int
	Equipement Equipement
	InventoryCap int
	InventoryUpgrades int 
}

func InitCharacterDeflaut() *Character {
	c := &Character{
		Name: 
	}
}
