package main             

import (
	"red"
	"fmt"
)

func main() {
	fmt.Println(red.PageDeGarde())
	red.Musique()
	//Création du personnage
	player := red.NewCharacter("Héros")

	// Affichage de l'inventaire
	player.ShowInventory()

	// Utilisation d'une potion de vie
	player.UsePotion()

	// Affichage de l'inventaire après utilisation
	player.ShowInventory()

	// Ajout et retrait d'objets
	player.AddItem("Potion de vie", 1)
	player.RemoveItem("Potion de poison", 1)

	// Affichage final de l'inventaire
	player.ShowInventory()
}

