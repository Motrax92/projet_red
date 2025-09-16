package main

import (
	"fmt"
	"red"
	"sync"
)

func main() {
	// Affichage de la page de garde
	fmt.Println(red.PageDeGarde())

	// Préparer le WaitGroup pour la musique
	var wg sync.WaitGroup
	wg.Add(1) // On ajoute la goroutine musique

	// Lancer la musique en arrière-plan
	go func() {
		defer wg.Done()
		red.Musique() // boucle infinie tant que tu n'arrêtes pas le programme
	}()

	// Lancer les mini-jeux
	red.Games()

	// Exemple de personnage
	name := "Héros"
	inventory := map[string]int{
		"Potion de vie":    3,
		"Potion de poison": 4,
		"Bouclier":         6,
		"Boule de feu":		2,
	}

	// Affichage de l'inventaire
	red.Inventaire(name, inventory)

	// Lancer un combat
	fmt.Println("\n=== Début du combat ===")
	red.LancerCombat()

	// Attendre que la musique se termine avant de fermer le programme
	wg.Wait()
	fmt.Println("Programme terminé.")
}
