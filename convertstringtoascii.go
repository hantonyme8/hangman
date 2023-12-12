package hangman

import "fmt"

func StringToAsciiArt(path string, liste []string) {
	/*
		Entrée : string, liste de string
		Sortie : liste de string
		Objectif : affiche l'ascii art voulue
	*/
	ascii := OpenWords(path)
	for k := 0; k < 8; k++ {
		for i := 0; i < len(liste); i++ {
			code := []rune(liste[i])
			index := (code[0]-32)*8+code[0]-31
			for j := int(index)+k; j < int(index)+k+1; j++ {
				fmt.Print(ascii[j])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
