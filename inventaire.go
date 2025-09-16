package red

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)


func AskString(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func AskInt(prompt string) int {
	for {
		s := AskString(prompt)
		if s == "" {
			return 0
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Entrée invalide, réessaie.")
			continue
		}
		return n
	}
}


type Character struct {
	Name      string
	Class     string
	Level     int
	HPMax     int
	HP        int
	Inventory map[string]int
	Gold      int
	Skills    []string
}

type Monster struct {
	Name   string
	HPMax  int
	HP     int
	Attack int
}


func InitCharacterInteractive() *Character {
	name := AskString("Choisis ton nom : ")
	name = strings.Title(strings.ToLower(name))

	fmt.Println("Choisis une classe :")
	fmt.Println("1) Humain (100 PV max)")
	fmt.Println("2) Elfe   (80 PV max)")
	fmt.Println("3) Nain   (120 PV max)")
	choice := AskInt("> ")

	var class string
	var hpMax int
	switch choice {
	case 1:
		class, hpMax = "Humain", 100
	case 2:
		class, hpMax = "Elfe", 80
	case 3:
		class, hpMax = "Nain", 120
	default:
		fmt.Println("Choix invalide, Elfe par défaut.")
		class, hpMax = "Elfe", 80
	}

	return &Character{
		Name:      name,
		Class:     class,
		Level:     1,
		HPMax:     hpMax,
		HP:        hpMax / 2,
		Inventory: map[string]int{"Potion de vie": 1},
		Gold:      10,
		Skills:    []string{"Coup de poing"},
	}
}

func (c *Character) DisplayInfo() {
	fmt.Println("----- Personnage -----")
	fmt.Printf("Nom    : %s\n", c.Name)
	fmt.Printf("Classe : %s\n", c.Class)
	fmt.Printf("Niveau : %d\n", c.Level)
	fmt.Printf("PV     : %d / %d\n", c.HP, c.HPMax)
	fmt.Printf("Or     : %d pièces\n", c.Gold)
	fmt.Println("Inventaire :")
	for item, qty := range c.Inventory {
		fmt.Printf(" - %s x%d\n", item, qty)
	}
	fmt.Println("----------------------")
}

func (c *Character) UsePotion() {
	if c.Inventory["Potion de vie"] > 0 {
		c.Inventory["Potion de vie"]--
		c.HP += 50
		if c.HP > c.HPMax {
			c.HP = c.HPMax
		}
		fmt.Printf("%s utilise une Potion de vie (+50 PV). PV : %d/%d\n", c.Name, c.HP, c.HPMax)
	} else {
		fmt.Println("Tu n’as pas de Potion de vie !")
	}
}


func MerchantMenu(c *Character) {
	for {
		fmt.Println("----- Marchand -----")
		fmt.Println("Or :", c.Gold)
		fmt.Println("1) Potion de vie (3 pièces)")
		fmt.Println("0) Retour")
		choice := AskInt("> ")
		switch choice {
		case 1:
			if c.Gold >= 3 {
				c.Gold -= 3
				c.Inventory["Potion de vie"]++
				fmt.Println("Achat effectué : Potion de vie")
			} else {
				fmt.Println("Pas assez d’or !")
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}


func InitGoblin() *Monster {
	return &Monster{
		Name:   "Gobelin d'entraînement",
		HPMax:  40,
		HP:     40,
		Attack: 5,
	}
}

func TrainingFight(c *Character) {
	g := InitGoblin()
	turn := 1
	for {
		fmt.Printf("\n=== Tour %d ===\n", turn)
		fmt.Println("1) Attaquer")
		fmt.Println("2) Utiliser potion")
		fmt.Println("0) Fuir")
		choice := AskInt("> ")

		switch choice {
		case 1:
			fmt.Println(c.Name, "attaque avec Coup de poing ! (-5 PV)")
			g.HP -= 5
			if g.HP < 0 {
				g.HP = 0
			}
			fmt.Printf("%s : %d/%d PV\n", g.Name, g.HP, g.HPMax)
		case 2:
			c.UsePotion()
		case 0:
			fmt.Println("Tu fuis le combat...")
			return
		default:
			fmt.Println("Choix invalide")
		}

		if g.HP <= 0 {
			fmt.Println("Le gobelin est vaincu !")
			return
		}


		fmt.Printf("%s attaque ! (-%d PV)\n", g.Name, g.Attack)
		c.HP -= g.Attack
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("%s : %d/%d PV\n", c.Name, c.HP, c.HPMax)

		if c.HP <= 0 {
			fmt.Println("Tu es mort... Game Over")
			return
		}

		turn++
	}
}

//
// ==== MENU PRINCIPAL ====
//
func GameLoop() {
	var player *Character
	for {
		fmt.Println("\n--- Menu principal ---")
		fmt.Println("1) Créer personnage")
		fmt.Println("2) Infos personnage")
		fmt.Println("3) Marchand")
		fmt.Println("4) Combat d’entraînement")
		fmt.Println("0) Quitter")
		choice := AskInt("> ")

		switch choice {
		case 1:
			player = InitCharacterInteractive()
			fmt.Println("Personnage créé !")
		case 2:
			if player == nil {
				fmt.Println("Aucun personnage.")
			} else {
				player.DisplayInfo()
			}
		case 3:
			if player == nil {
				fmt.Println("Aucun personnage.")
			} else {
				MerchantMenu(player)
			}
		case 4:
			if player == nil {
				fmt.Println("Aucun personnage.")
			} else {
				TrainingFight(player)
			}
		case 0:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

