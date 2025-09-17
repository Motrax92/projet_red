package main

import (
	"fmt"
	"red/inventaire"
	"red/personnages"
	"red/musique"
	"red/pagedegarde"
	"sync"
)

func main() {
	for {
		// --- MENU PRINCIPAL ---
		fmt.Println(pagedegarde.PageDeGarde())
		fmt.Println("------------------------------------------------")
		fmt.Println("                 M E N U   P R I N C I P A L    ")
		fmt.Println("------------------------------------------------")
		fmt.Println("1. Jouer")
		fmt.Println("2. Paramètres")
		fmt.Println("3. Quitter")
		fmt.Println("------------------------------------------------")

		// Lecture du choix utilisateur
		var choix int
		fmt.Print("👉 Entrez un numéro de menu : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Println("🎮 Le jeu commence !")

			// Lancer la musique en arrière-plan avec WaitGroup
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				musique.Musique()
			}()

			// Exemple de personnage et inventaire
			name := "Héros"
			inventory := map[string]int{
				"Potion de vie":    2,
				"Potion de poison": 1,
				"Bouclier":         1,
			}

			// Afficher l'inventaire
			inventaire.Inventaire(name, inventory)

			// --- Sous-menu Inventaire ---
			for {
				fmt.Println("------------------------------------------------")
				fmt.Println("1. Utiliser un objet")
				fmt.Println("2. Continuer vers le combat")
				fmt.Println("------------------------------------------------")

				var choixInv int
				fmt.Print("👉 Entrez un numéro : ")
				fmt.Scanln(&choixInv)

				if choixInv == 1 {
					var objet string
					fmt.Print("Quel objet voulez-vous utiliser ? ")
					fmt.Scanln(&objet)

					inventaire.UtiliserObjet(inventory, objet)
					inventaire.Inventaire(name, inventory) // réaffiche après utilisation

				} else if choixInv == 2 {
					break // sortir de la boucle → aller au combat
				} else {
					fmt.Println("❌ Choix invalide, réessayez.")
				}
			}

			// Lancer le combat tour par tour
			personnages.LancerCombat()

			// Attendre la fin de la musique avant de revenir au menu
			wg.Wait()
			fmt.Println("🎵 La musique est terminée.")

		case 2:
			fmt.Println("⚙️ Menu Paramètres (en cours de développement...)")

		case 3:
			fmt.Println("👋 Au revoir !")
			return // quitte le programme

		default:
			fmt.Println("❌ Choix invalide, réessayez.")
		}

		fmt.Println() // saute une ligne pour la lisibilité
	}
}
