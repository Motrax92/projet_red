package red

import (
	"fmt"
)

type Character struct {
	Name      string
	HP        int
	HPMax     int
	Inventory map[string]int
}

func NewCharacter(name string) *Character {
	return &Character{
		Name:      name,
		HPMax:     100,
		HP:        50,
		Inventory: map[string]int{"Potion de vie": 2, "Potion de poison": 1},
	}
}

// Afficher inventaire
func (c *Character) ShowInventory() {
	fmt.Println("----- Inventaire -----")
	for item, qty := range c.Inventory {
		fmt.Printf(" - %s x%d\n", item, qty)
	}
	fmt.Println("----------------------")
}

// Ajouter un objet
func (c *Character) AddItem(item string, qty int) {
	c.Inventory[item] += qty
	fmt.Printf("%s a obtenu %d %s\n", c.Name, qty, item)
}

// Retirer un objet
func (c *Character) RemoveItem(item string, qty int) {
	if c.Inventory[item] >= qty {
		c.Inventory[item] -= qty
		if c.Inventory[item] == 0 {
			delete(c.Inventory, item)
		}
		fmt.Printf("%s a perdu %d %s\n", c.Name, qty, item)
	} else {
		fmt.Println("Pas assez d'objets à retirer !")
	}
}

// Utiliser une potion
func (c *Character) UsePotion() {
	if c.Inventory["Potion de vie"] > 0 {
		c.Inventory["Potion de vie"]--
		c.HP += 50
		if c.HP > c.HPMax {
			c.HP = c.HPMax
		}
		fmt.Printf("%s utilise une Potion de vie (+50 PV). PV : %d/%d\n", c.Name, c.HP, c.HPMax)
	} else {
		fmt.Println("Pas de Potion de vie disponible.")
	}
}

func main() {
	player := NewCharacter("Héros")

	player.ShowInventory()
	player.UsePotion()
	player.ShowInventory()

	player.AddItem("Potion de vie", 1)
	player.RemoveItem("Potion de poison", 1)
	player.ShowInventory()
}
