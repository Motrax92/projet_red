package game

import (
	"fmt"
	"red/personnages"
 // ⚠️ adapte selon ton chemin exact
)

func Games() {
	fmt.Println("------------------------------------------------")
	fmt.Println("                     M E N U                    ")
	fmt.Println("------------------------------------------------")
	fmt.Println("1. Jouer")
	fmt.Println("2. Paramètres")
	fmt.Println("3. Aides")
	fmt.Println("------------------------------------------------")

	// Lire le choix de l'utilisateur
	var choix int
	fmt.Print("👉 Entrez un numéro de menu : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		fmt.Println("🎮 Jeu lancé !")
		personnages.LancerCombat() // 👈 démarre le combat
	case 2:
		fmt.Println("⚙️ Menu Paramètres")
	case 3:
		fmt.Println("🛠️ Menu Settings")
	default:
		fmt.Println("❌ Choix invalide.")
	}
}
