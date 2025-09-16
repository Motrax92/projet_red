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
