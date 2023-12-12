package hangman

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintHangman(index int) string {
	/*
		Entrée : errorhangman (entier)
		Sortie : une slice de chaînes de caractères
		Objectif : renvoie la représentation graphique du pendu en fonction du nombre d'erreurs du joueur
	*/
	// 	content, err := os.Open("hangman.txt")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer content.Close()

	// 	scanner := bufio.NewScanner(content)
	// 	var positions []string

	// 	for scanner.Scan() {
	// 		line := scanner.Text()
	// 		// Ajoutez la ligne à la slice positions
	// 		positions = append(positions, line)
	// 	}

	//	if errorhangman >= 0 && errorhangman < len(positions) {
	//		// Retournez la position correspondant au nombre d'erreurs
	//		return positions[:errorhangman+1]
	//	} else {
	//
	//		// Gérer le cas où errorhangman est en dehors de la plage valide
	//		return []string{"Position non trouvée"}
	//	}
	var hangman string
	affiche := OpenWords("hangman.txt")
	index += index * 7
	for i := index; i < index+7; i++ {
		hangman += affiche[i] + "\n"
	}
	return hangman
}

func PrintHidden(word []string) []string {
	/*
		Entrée : liste de string
		Sortie : liste de string
		Objectif : renvoie des tirets à la place des lettres du mot à trouver
	*/
	liste := []string{}
	for i := 0; i < len(word); i++ {
		liste = append(liste, "_")
		if i != len(word)-1 {
			liste = append(liste, " ")
		}
	}
	return liste
}

func PrintLetter(liste []int, r_word, hidden []string) []string {
	/*
		Entrée : listes de string, liste d'entier
		Sortie : liste de string
		Objectif : renvoie les lettres trouver du mot
	*/
	for i := 0; i < len(liste); i++ {
		hidden[liste[i]*2] = r_word[liste[i]]
	}
	return hidden
}

func SelectN(word []string) []int {
	/*
		Entrée : string
		Sortie : liste d'entier
		Objectif : renvoie l'indice des lettres aléatoires
	*/
	rand.Seed(time.Now().UnixNano())
	liste := []int{}
	for i := 0; i < len(word)/2-1; i++ {
		r := rand.Intn(len(word))
		listeStr := SliceIntToSliceString(liste)
		rStr := SliceIntToSliceString([]int{r})
		contenu := Contain(SliceToString(listeStr), SliceToString(rStr))
		if len(contenu) > 0 {
			i--
		} else {
			liste = append(liste, r)
		}
	}
	return liste
}

var color = []string{
	string("\033[0m"),  // reset color
	string("\033[31m"), // red
	string("\033[32m"), // green
	string("\033[33m"), // yellow
	string("\033[34m"), // blue
	string("\033[35m"), // purple
	string("\033[36m"), // cyan
	string("\033[37m")} // white

func GameRule() {
	/*
		Entrée : rien
		Sortie : affiche
		Objectif : affiche les règles du jeu et ses fonctionnalités
	*/
	fmt.Println("You have to find the hidden word")
	fmt.Println("You can use ascii letters with the", color[3], "-letterFile", color[0], "flag when starting the game")
	fmt.Println("To save you just have to mark", color[3], "-stop", color[0])
	fmt.Println("To launch a saved game, simply fill it in with the", color[3], "-startWith", color[0], "flag followed by the name of the save game")
	fmt.Println("You can open the save folder by typing", color[3], "-saveInfo", color[0], "here")
	fmt.Print(color[1], "Press enter to start... ", color[0])
	input := Input()
	Clear()
	if SliceToString(input) == "-SAVEINFO" {
		if len(SaveInfo()) == 0 {
			fmt.Println("No file save")
		} else {
			fmt.Println("\nSave File : ")
			for i := 0; i < len(SaveInfo()); i++ {
				fmt.Println(SaveInfo()[i])
			}
		}
		fmt.Print(color[1], "\nPress enter to continu...", color[0])
		Input()
	}
}
