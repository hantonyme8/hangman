package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func OpenWords(fichier string) []string {
    /* 
	Entrée : string
	Sortie : liste de string
	Objectif : renvoie ce qui est contenue dans un fichier texte
	*/
    listeW := []string{}
    file, err := os.Open(fichier)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        listeW = append(listeW, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return listeW
}

func WriteInFile(fichier, text string) {
    /*
	Entrée : string
	Sortie : affiche
	Objectif : écrit dans un fichier
	*/
    file, err := os.Create(fichier)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    file.WriteString(text)
    if err != nil {
        fmt.Println(err)
        file.Close()
        return
    }
}

func RemoveFile(path string) {
    /* 
	Entrée : string
	Sortie : supprime
	Objectif : enleve un fichier
	*/
    file := os.Remove(path)
    if file != nil {
        log.Fatal(file)
    }
}