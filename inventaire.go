package red 
import "fmt"

// Fonction pour afficher l'inventaire d'un personnage
func Inventaire(name string, inventory map[string]int) {
	fmt.Printf("=== Inventaire de %s ===\n", name)
	for item, qty := range inventory {
		fmt.Printf(" - %s x%d\n", item, qty)
	}
	fmt.Println("----------------------")
}
