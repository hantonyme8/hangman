package hangman

import (
	"bufio"
	"os"
	"os/exec"
)

func Input() []string {
	/* 
	Entrée : input terminal
	Sortie : liste de string
	Objectif : renvoie ce qu'ecrit le joueur
	*/
	scanner := bufio.NewScanner(os.Stdin)
	var text []string
	scanner.Scan()         // Permet d'écouter l'entrée standard
	text = StringToSlice(scanner.Text()) // Permet de récuperer le texte entrée par l'utilisateur
	return ToUpper(text)
}
func Clear() {
	/*
		Entrée : rien
		Sortie : clear
		Objectif : clear le terminal
	*/
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
