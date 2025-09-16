package red

import "fmt"

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

// Fonction publique qui lance un combat
func LancerCombat() {
	// Création du joueur
	joueur := Joueur{
		Personnage: Personnage{
			Nom:   "Héros",
			Vie:   100,
			Force: 20,
		},
		Niveau: 1,
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
		joueur.Attaquer(&ennemi.Personnage)
		if !ennemi.EstVivant() {
			fmt.Println(ennemi.Nom, "est vaincu !")
			break
		}

		ennemi.Attaquer(&joueur.Personnage)
		if !joueur.EstVivant() {
			fmt.Println(joueur.Nom, "est mort !")
			break
		}
	}

	fmt.Println("Combat terminé ⚔️")
}
