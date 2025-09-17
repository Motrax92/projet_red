package prologue

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const prologueText = `Les néons clignotent dans la nuit.
Un battement sourd résonne, comme le cœur d’une machine géante.
BOUM... BOUM... BOUM...

Soudain, le décor s’illumine :
un héros apparaît, silhouette scintillante, pixels vibrants.
Face à lui, un gobelin glitché, moitié créature, moitié bug du système.

Les circuits s’illuminent, les couleurs explosent.
Le temps ralentit.
Un seul mot traverse l’écran :

Éclairs. Étincelles. Énergie pure.
Le héros serre son arme digitale.
Le gobelin ricane, distordu par le bruit électronique.

Ici, pas de fuite.
Seulement un duel où chaque note, chaque coup,
résonne comme une pulsation de l’univers.

La musique s’emballe.
Le combat commence.`

// typeWriter affiche le texte caractère par caractère pour un effet "terminal/arcade"
func typeWriter(s string, delay time.Duration) {
	for _, r := range s {
		fmt.Printf("%c", r)
		time.Sleep(delay)
	}
}

// Show affiche le prologue au lancement, puis attend que l'utilisateur appuie sur Entrée
func Show() {
	// Effet: petite "intro" visuelle
	fmt.Print("\033[2J\033[H") // clear + home (ANSI)
	typeWriter(prologueText+"\n\n", 12*time.Millisecond)

	fmt.Print("\n▶ Appuyez sur Entrée pour continuer... ")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
