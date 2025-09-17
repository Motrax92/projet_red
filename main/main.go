package main

import (
	"fmt"
	"red/musique"
	"red/pagedegarde"
	"red/personnages"
	"red/prologue"
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

		var choix int
		fmt.Print("👉 Entrez un numéro de menu : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Println("🎮 Le jeu commence !")

			// Prologue
			prologue.Show()

			// Lancer la musique en arrière-plan
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				musique.Musique()
			}()

			// --- Lancer le combat tour par tour ---
			personnages.LancerCombat()

			wg.Wait()
			fmt.Println("🎵 La musique est terminée.")

		case 2:
			// Paramètres > Langue
			for {
				fmt.Println("------------------------------------------------")
				fmt.Println("            P A R A M È T R E S  >  L A N G U E ")
				fmt.Println("------------------------------------------------")
				fmt.Println("1. Français (actuel)")
				fmt.Println("2. Anglais")
				fmt.Println("3. Retour")
				fmt.Println("------------------------------------------------")
				var choixLang int
				fmt.Print("👉 Choisissez une option : ")
				fmt.Scanln(&choixLang)

				switch choixLang {
				case 1, 3:
					break
				case 2:
					fmt.Println("On parle français ici.")
				default:
					fmt.Println("❌ Choix invalide, réessayez.")
				}
				break
			}

		case 3:
			fmt.Println("👋 Au revoir !")
			return

		default:
			fmt.Println("❌ Choix invalide, réessayez.")
		}

		fmt.Println() // ligne vide pour lisibilité
	}
}
