package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var agents_par_classe = map[string][]string{
	"Duelliste":  {"Jett", "Phoenix", "Neon", "Raze", "Reyna", "Yoru"},
	"Controleur": {"Astra", "Brimstone", "Omen", "Viper"},
	"Initiateur": {"Breach", "KAY/O", "Skye", "Sova", "Fade"},
	"Sentinelle": {"Chamber", "Cypher", "Killjoy", "Sage"},
}

func randomAgent() {
	classe_aleatoire := rand.Intn(len(agents_par_classe))
	index := 0
	for _, classe := range agents_par_classe {
		if index == classe_aleatoire {
			fmt.Println(classe[rand.Intn(len(classe))])
		}

		index += 1
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	if len(os.Args) > 1 {
		classe_demande, found := agents_par_classe[os.Args[1]]

		if !(found) {
			log.Printf("Classe %s ne fait pas partie de la liste des classes qui existent, renvoie d'un agent aléatoire à la place.", os.Args[1])
			randomAgent()
		} else {
			fmt.Println(classe_demande[rand.Intn(len(classe_demande))])
		}

	} else {
		randomAgent()
	}
}
