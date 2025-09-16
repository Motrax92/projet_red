package red

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func Musique() {
	// Ouvrir le fichier MP3
	f, err := os.Open("musique.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Décoder le MP3
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	// ⚠️ ne pas defer streamer.Close() ici, sinon ça coupe la boucle
	// on ferme manuellement plus tard si besoin

	// Initialiser le speaker
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	// Créer une boucle infinie
	loop := beep.Loop(-1, streamer)

	// Jouer la musique en boucle
	speaker.Play(loop)

	// Bloquer pour que le programme tourne en continu
	select {} // boucle infinie
}
