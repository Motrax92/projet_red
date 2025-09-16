package red

import (
	"fmt"
)

// ----------------------
// Personnages
// ----------------------
type Personnage struct {
	Nom   string
	Vie   int
	Force int
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
	fmt.Printf("%s attaque %s et inflige %d dégâts !\n", p.Nom, cible.Nom, p.Force)
	cible.SubirDegats(p.Force)
}

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
	// Création du joueur
	joueur := Joueur{
		Personnage: Personnage{
			Nom:   "Héros",
			Vie:   100,
			Force: 20,
		},

		Niveau :1,
	}

	// Création d’un ennemi
	ennemi := Ennemi{
		Personnage: Personnage{
			Nom:   "Gobelin",
			Vie:   50,
			Force: 10,
		},
		Type: "Monstre",
	}

	// Déroulement du combat
	for joueur.EstVivant() && ennemi.EstVivant() {
		// Tour du joueur
		fmt.Println("\n--- Tour du joueur ---")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Print("Choisis une action : ")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			joueur.Attaquer(&ennemi.Personnage)
		} else if choix == 2 {
			continue // on saute le reste du tour (l'ennemi ne joue pas encore)
		} else {
			fmt.Println("Choix invalide, tu perds ton tour !")
		}

		// Vérifie si l’ennemi est mort
		if !ennemi.EstVivant() {
			fmt.Println(ennemi.Nom, "est vaincu !")
			break
		}

		// Tour de l’ennemi
		fmt.Println("\n--- Tour de l’ennemi ---")
		ennemi.Attaquer(&joueur.Personnage)
		if !joueur.EstVivant() {
			fmt.Println(joueur.Nom, "est mort !")
			break
		}
	}

	fmt.Println("\nCombat terminé ⚔️")
}
