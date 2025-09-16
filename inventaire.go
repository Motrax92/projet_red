package red

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name        string
	Class       string
	Level       int
	HPMax       int
	HP          int
	Inventory   map[string]int
	Skills      []string
	Gold        int
	Equipment   Equipment
	InventoryCap int
	InventoryUpgrades int
}

func InitCharacterDefault() *Character {
	c := &Character{
		Name:      "Aventurier",
		Class:     "Elfe",
		Level:     1,
		HPMax:     100,
		HP:        40,
		Inventory: map[string]int{"Potion de vie": 3},
		Skills:    []string{"Coup de poing"},
		Gold:      100,
		InventoryCap: 10,
		InventoryUpgrades: 0,
	}
	return c
}

func InitCharacterInteractive() (*Character, error) {
	name := AskString("Choisis ton nom (lettres seulement) : ")
	if name == "" {
		return nil, errors.New("nom invalide")
	}
	name = strings.Title(strings.ToLower(name))

	fmt.Println("Choisis une classe :")
	fmt.Println("1) Humain (HP max 100)")
	fmt.Println("2) Elfe   (HP max 80)")
	fmt.Println("3) Nain   (HP max 120)")
	choice := AskInt("> ")

	var hpMax int
	var class string
	switch choice {
	case 1:
		hpMax = 100
		class = "Humain"
	case 2:
		hpMax = 80
		class = "Elfe"
	case 3:
		hpMax = 120
		class = "Nain"
	default:
		fmt.Println("Choix invalide, classe Elfe par défaut.")
		hpMax = 80
		class = "Elfe"
	}

	startHP := hpMax / 2
	c := &Character{
		Name:      name,
		Class:     class,
		Level:     1,
		HPMax:     hpMax,
		HP:        startHP,
		Inventory: map[string]int{"Potion de vie": 1},
		Skills:    []string{"Coup de poing"},
		Gold:      100,
		InventoryCap: 10,
		InventoryUpgrades: 0,
	}
	return c, nil
}

func (c *Character) DisplayInfo() {
	fmt.Println("----- Infos du personnage -----")
	fmt.Printf("Nom  : %s\n", c.Name)
	fmt.Printf("Classe : %s\n", c.Class)
	fmt.Printf("Niveau : %d\n", c.Level)
	fmt.Printf("PV : %d / %d\n", c.HP, c.HPMax)
	fmt.Printf("Argent : %d pièces\n", c.Gold)
	fmt.Printf("Inventaire (%d/%d) :\n", c.InventorySize(), c.InventoryCap)
	for item, qty := range c.Inventory {
		fmt.Printf(" - %s x%d\n", item, qty)
	}
	fmt.Printf("Compétences : %v\n", c.Skills)
	fmt.Printf("Equipement : tête=%s torse=%s pieds=%s\n", c.Equipment.Head, c.Equipment.Torso, c.Equipment.Feet)
	fmt.Println("-------------------------------")
}

func (c *Character) InventorySize() int {
	sum := 0
	for _, q := range c.Inventory {
		sum += q
	}
	return sum
}

func (c *Character) AddInventory(item string, qty int) error {
	if c.InventorySize()+qty > c.InventoryCap {
		return errors.New("inventaire plein")
	}
	c.Inventory[item] += qty
	return nil
}

func (c *Character) RemoveInventory(item string, qty int) error {
	cur := c.Inventory[item]
	if cur < qty {
		return errors.New("quantité insuffisante")
	}
	if cur == qty {
		delete(c.Inventory, item)
	} else {
		c.Inventory[item] = cur - qty
	}
	return nil
}

func (c *Character) UsePotion() {
	if c.Inventory["Potion de vie"] <= 0 {
		fmt.Println("Tu n'as pas de Potion de vie.")
		return
	}
	_ = c.RemoveInventory("Potion de vie", 1)
	healed := 50
	c.HP += healed
	if c.HP > c.HPMax {
		c.HP = c.HPMax
	}
	fmt.Printf("Tu utilises une Potion de vie. PV : %d / %d\n", c.HP, c.HPMax)
}

func (c *Character) UsePoisonPot() {
	if c.Inventory["Potion de poison"] <= 0 {
		fmt.Println("Tu n'as pas de Potion de poison.")
		return
	}
	_ = c.RemoveInventory("Potion de poison", 1)
	fmt.Println("Tu es empoisonné ! -10 PV par seconde pendant 3s.")
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c.HP -= 10
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("PV : %d / %d\n", c.HP, c.HPMax)
		if c.IsDead() {
			fmt.Println("Tu es mort suite au poison.")
			c.Respawn()
			return
		}
	}
}

func (c *Character) IsDead() bool {
	return c.HP <= 0
}

func (c *Character) Respawn() {
	c.HP = c.HPMax / 2
	fmt.Printf("Tu ressuscites avec %d / %d PV.\n", c.HP, c.HPMax)
}

func (c *Character) LearnSpell(spell string) {
	for _, s := range c.Skills {
		if s == spell {
			fmt.Println("Tu connais déjà ce sort.")
			return
		}
	}
	c.Skills = append(c.Skills, spell)
	fmt.Printf("Nouveau sort appris : %s\n", spell)
}

func (c *Character) UpgradeInventorySlot() error {
	if c.InventoryUpgrades >= 3 {
		return errors.New("limite d'améliorations atteinte (3 max)")
	}
	if c.Gold < 30 {
		return errors.New("pas assez d'or pour acheter l'augmentation")
	}
	c.Gold -= 30
	c.InventoryCap += 10
	c.InventoryUpgrades += 1
	fmt.Printf("Capacité d'inventaire augmentée à %d (améliorations utilisées %d/3)\n", c.InventoryCap, c.InventoryUpgrades)
	return nil
}
