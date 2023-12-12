package hangman

import (
	"math/rand"
	"time"
)

func RandWord(path string) string {
	/* 
	Entrée : rien
	Sortie : string
	Objectif : renvoie un mot aléatoire à partir du fichier texte "words.txt"
	*/
	rand.Seed(time.Now().UnixNano())
	liste_words := OpenWords(path)
	ind_word := rand.Intn(len(liste_words))
	return liste_words[ind_word]
}
