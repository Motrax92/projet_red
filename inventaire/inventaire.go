package inventaire

import "fmt"

// Affiche l'inventaire
func Inventaire(name string, inventory map[string]int) {
	fmt.Printf("=== Inventaire de %s ===\n", name)
	for item, qty := range inventory {
		fmt.Printf(" - %s x%d\n", item, qty)
	}
	fmt.Println("----------------------")
}

// Utiliser un objet
func UtiliserObjet(inventory map[string]int, objet string) {
	if qty, ok := inventory[objet]; ok {
		if qty > 1 {
			inventory[objet] = qty - 1
			fmt.Printf("Vous avez utilisé un(e) %s. Il en reste %d.\n", objet, inventory[objet])
		} else {
			delete(inventory, objet)
			fmt.Printf("Vous avez utilisé votre dernier(e) %s. Plus aucun en stock.\n", objet)
		}
	} else {
		fmt.Printf("Vous n'avez pas de %s dans votre inventaire.\n", objet)
	}
}
