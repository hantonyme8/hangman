package hangman

import (
	"strconv"
)

func SliceToString(liste []string) string {
	/*
		Entrée : liste de string
		Sortie : string
		Objectif : renvoie la string d'une liste de string
	*/
	str := ""
	for i := 0; i < len(liste); i++ {
		str += liste[i]
	}
	return str
}

func SliceToStringSpace(liste []string) string {
	/*
		Entrée : liste de string
		Sortie : string
		Objectif : renvoie la string d'une liste de string avec un espace
	*/
	str := ""
	for i := 0; i < len(liste); i++ {
		str += liste[i]
		if i != len(liste)-1 {
			str += " "
		}
	}
	return str
}

func StringToSlice(str string) []string {
	/*
		Entrée : string
		Sortie : liste de string
		Objectif : renvoie la liste de string d'une string
	*/
	liste := []string{}
	rmot := []rune(str)
	for i := 0; i < len(rmot); i++ {
		liste = append(liste, string(rmot[i]))
	}
	return liste
}

func SliceIntToSliceString(liste []int) []string {
	/*
		Entrée : liste d'entier
		Sortie : liste de string
		Objectif : renvoie la liste de string d'une liste d'entier
	*/
	listeStr := []string{}
	for i := 0; i < len(liste); i++ {
		listeStr = append(listeStr, strconv.Itoa(liste[i]))
	}
	return listeStr
}

func CompareString(str1, str2 []string) bool {
	/*
		Entrée : listes de string
		Sortie : bool
		Objectif : renvoie vrai si les listes de string sont égales
	*/
	if len(str1) == 0 || len(str2) == 0 {
		return false
	}
	if len(str1)*2-1 != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i*2] {
			return false
		}
	}
	return true
}

func ToUpper(str []string) []string {
	/*
		Entrée : liste de string
		Sortie : liste de string
		Objectif : renvoie toute les lettres en majuscule
	*/
	listerune := []rune(SliceToString(str))
	for i := 0; i < len(listerune); i++ {
		if listerune[i] >= 'a' && listerune[i] <= 'z' {
			listerune[i] -= 32
		}
	}
	motstr := string(listerune)
	motstrf := StringToSlice(motstr)
	return motstrf
}

func IsAlphabet(s []string) bool {
	/*
		Entrée : liste de string
		Sortie : bool
		Objectif : renvoie vrai si c'est en alphabet latin
	*/
	numeric := []rune(SliceToString(s))
	for i := 0; i < len(numeric); i++ {
		if numeric[i] >= 'a' && numeric[i] <= 'z' || numeric[i] >= 'A' && numeric[i] <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
