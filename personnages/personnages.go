package personnages

import (
	"fmt"
	"math/rand"
	"time"
)

// ----------------------
// Personnages
// ----------------------
type Personnage struct {
	Nom     string
	Vie     int
	Force   int
	Protégé bool
}

func (p *Personnage) EstVivant() bool {
	return p.Vie > 0
}

func (p *Personnage) SubirDegats(degats int) {
	p.Vie -= degats
	if p.Vie < 0 {
		p.Vie = 0
	}
}

func (p *Personnage) Attaquer(cible *Personnage) {
	if cible.Protégé {
		fmt.Printf("%s attaque %s mais l’attaque est bloquée par le bouclier ! 🛡️\n", p.Nom, cible.Nom)
		cible.Protégé = false
		return
	}

	degats := rand.Intn(51)
	fmt.Printf("%s attaque %s et inflige %d dégâts !\n", p.Nom, cible.Nom, degats)
	cible.SubirDegats(degats)
}

// ----------------------
// Inventaire & Objets
// ----------------------
func AfficherInventaire(name string, inventory map[string]int) {
	fmt.Printf("=== Inventaire de %s ===\n", name)
	if len(inventory) == 0 {
		fmt.Println("Inventaire vide.")
	} else {
		for item, qty := range inventory {
			fmt.Printf(" - %s x%d\n", item, qty)
		}
	}
	fmt.Println("----------------------")
}

func UtiliserObjet(inventory map[string]int, objet string, user *Personnage, cible *Personnage) {
	if qty, ok := inventory[objet]; ok && qty > 0 {
		switch objet {
		case "Bouclier":
			fmt.Printf("%s utilise un Bouclier 🛡️ et sera protégé de la prochaine attaque.\n", user.Nom)
			user.Protégé = true
			inventory[objet]--

		case "Potion de vie":
			soin := rand.Intn(21) + 5
			user.Vie += soin
			if user.Vie > 100 {
				user.Vie = 100
			}
			fmt.Printf("%s boit une potion de vie et récupère %d PV ❤️ (PV actuels : %d).\n", user.Nom, soin, user.Vie)
			inventory[objet]--

		case "Potion de poison":
			if cible != nil {
				degats := rand.Intn(16) + 5
				fmt.Printf("%s utilise une potion de poison 💀 sur %s et inflige %d dégâts.\n", user.Nom, cible.Nom, degats)
				cible.SubirDegats(degats)
				inventory[objet]--
			} else {
				fmt.Println("Pas de cible pour utiliser le poison !")
			}

		default:
			fmt.Println("Cet objet n’a pas encore d’effet défini.")
		}

		if inventory[objet] <= 0 {
			delete(inventory, objet)
		}
	} else {
		fmt.Printf("Vous n'avez pas de %s dans votre inventaire.\n", objet)
	}
}

// ----------------------
// Joueur & Ennemi
// ----------------------
type Joueur struct {
	Personnage
	Niveau int
}

type Ennemi struct {
	Personnage
	Type string
}

// ----------------------
// Combat
// ----------------------
func LancerCombat() {
	rand.Seed(time.Now().UnixNano())

	joueur := Joueur{
		Personnage: Personnage{Nom: "Héros", Vie: 100, Force: 20},
		Niveau:     1,
	}
	ennemi := Ennemi{
		Personnage: Personnage{Nom: "Gobelin", Vie: 50, Force: 10},
		Type:       "Monstre",
	}

	invJoueur := map[string]int{
		"Potion de vie":    2,
		"Potion de poison": 1,
		"Bouclier":         1,
	}
	invEnnemi := map[string]int{
		"Potion de vie":    1,
		"Potion de poison": 1,
		"Bouclier":         0,
	}

	for joueur.EstVivant() && ennemi.EstVivant() {
		// Tour joueur
		fmt.Println("\n--- Tour du joueur ---")
		fmt.Printf("PV %s : %d | PV %s : %d\n", joueur.Nom, joueur.Vie, ennemi.Nom, ennemi.Vie)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Print("Choisis une action : ")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			joueur.Attaquer(&ennemi.Personnage)
		} else if choix == 2 {
			AfficherInventaire(joueur.Nom, invJoueur)
			fmt.Print("Quel objet utiliser ? ")
			var objet string
			fmt.Scan(&objet)
			UtiliserObjet(invJoueur, objet, &joueur.Personnage, &ennemi.Personnage)
			goto EnnemiTour
		} else {
			fmt.Println("Choix invalide, tu perds ton tour !")
		}

		if !ennemi.EstVivant() {
			fmt.Println(ennemi.Nom, "est vaincu ! 🎉")
			break
		}

	// Tour ennemi
	EnnemiTour:
		fmt.Println("\n--- Tour de l’ennemi ---")
		if rand.Intn(2) == 0 || len(invEnnemi) == 0 {
			ennemi.Attaquer(&joueur.Personnage)
		} else {
			objets := []string{"Potion de vie", "Potion de poison", "Bouclier"}
			objet := objets[rand.Intn(len(objets))]
			UtiliserObjet(invEnnemi, objet, &ennemi.Personnage, &joueur.Personnage)
		}

		if !joueur.EstVivant() {
			fmt.Println(joueur.Nom, "est mort ! 💀")
			break
		}
	}

	fmt.Println("\nCombat terminé ⚔️")
}
