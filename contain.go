package hangman

func Contain(words, letter string) []int {
	/* 
	Entrée : string
	Sortie : liste d'entier
	Objectif : renvoie l'indice s'il est contenue dans le mot
	*/
	liste_index := []int{}
	count := 0
	for i := 0; i < len(words); i++ {
		if i+len(letter) <= len(words) {
			a := words[i:len(letter)+i]
			for j := 0; j < len(letter); j++ {
				if a[j] == letter[j] {
					count++
				} else {
					count = 0
				}
				if count == len(letter) {
					for k := 0; k < count; k++ {
						liste_index = append(liste_index, i+k)
					}
				}
			}
		}
	}
	return liste_index
}

func ContainLetter(words, letter string) []int {
	/* 
	Entrée : string
	Sortie : liste d'entier
	Objectif : renvoie l'indice de la lettre qui est contenue dans le mot
	*/
	listew := []rune(words)
	liste := []int{}
	for i := 0; i < len(listew); i++ {
		if string(listew[i]) == letter {
			liste = append(liste, i)
		}
	}
	return liste
}