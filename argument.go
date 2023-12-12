package hangman

import (
	"log"
	"os"
)

func Args() []string {
	/* 
	Entrée : rien 
	Sortie : liste de string
	Objectif : renvoie l'argument
	*/
	arg := os.Args[1:]
	return arg
}

func FlagLetterFile() string {
	/* 
	Entrée : rien 
	Sortie : string
	Objectif : renvoie le flag du caractère ascii spécifié
	*/
	arg := Args()
	var text string
	for i := 0; i < len(arg); i++ {
		if arg[i] == "-letterFile" || arg[i] == "--letterFile" {
			text = arg[i+1]
		}
	}
	return text
}

func FlagStartWith() string {
	/* 
	Entrée : rien 
	Sortie : string
	Objectif : renvoie le flag de la sauvegarde spécifié
	*/
	arg := Args()
	var text string
	for i := 0; i < len(arg); i++ {
		if arg[i] == "-startWith" || arg[i] == "--startWith" {
			text = arg[i+1]
		}
	}
	return text
}

func SaveInfo() []string {
	/* 
	Entrée : rien 
	Sortie : liste de string
	Objectif : renvoie les sauvegardes enregistrer
	*/
	liste_name := []string{}
	files, err :=os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if len(Contain(file.Name(), "_SAVE.txt")) != 0 {
			liste_name = append(liste_name, file.Name())
		}
	}
	return liste_name
}

func IsFile(path string) bool {
	/* 
	Entrée : string 
	Sortie : bool
	Objectif : renvoie vrai si c'est un dossier
	*/
	listeFile, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range listeFile {
		if file.Name() == path {
			return true
		}
	}
	return false
}
