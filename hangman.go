package hangman

import (
	"fmt"
)

func Hangman() {
	var input []string
	alreadyuse := false
	affichage := []string{}
	path := Args()
	arret := false
	ascii_letter := FlagLetterFile()
	save := FlagStartWith()
	var mot []string
	var hidden []string
	var lettre_utiliser []string
	var errorhangman int
	var index int
	count := -1
	var contenu []int
	var dico HangManData
	if len(Contain(SliceToStringSpace(SaveInfo()), save)) == 0 {
		if len(path) == 0 {
			fmt.Println("No words file")
			return
		}
		dico = CreateNewGame(path[0])
	} else {
		dico = InitSaveGame(save)
	}
	mot = dico.ToFind
	hidden = dico.Word
	lettre_utiliser = dico.UsedLetter
	index = dico.Attempts
	errorhangman = dico.Hangman
	Clear()
	GameRule()
	Clear()
	for i := index; i < 9; i++ { // boucle principale
		if i == -1 {
			fmt.Println("Good Luck, you have 10 attempts.")
		}
		count++
		alreadyuse = false
		if SliceToString(input) == "-STOP" {
			fmt.Print("save name : ")
			input = Input()
			Stop(SliceToString(input), mot, hidden, lettre_utiliser, i, errorhangman)
			arret = true
			break
		}
		if CompareString(input, StringToSlice(SliceToStringSpace(mot))) || CompareString(mot, affichage) { // condition d'arret
			fmt.Print(color[3])
			fmt.Printf("Congrates ! You find the word %s, in %d attemps, with %d errors \n", SliceToString(mot), count, errorhangman+1)
			fmt.Print(color[0])
			break
		}
		if len(Contain(" "+SliceToString(lettre_utiliser)+" ", SliceToString(input))) == 0 {
			lettre_utiliser = append(lettre_utiliser, SliceToString(input), " ")
		} else {
			for j := 0; j < len(lettre_utiliser); j++ {
				if SliceToString(input) == lettre_utiliser[j] {
					fmt.Println(color[1], "Already choose", color[0])
					alreadyuse = true
					i--
				}
			}
			if len(input) == 1 && !alreadyuse {
				lettre_utiliser = append(lettre_utiliser, SliceToString(input), " ")
			}
		}
		if len(input) > 1 && !CompareString(input, mot) && IsAlphabet(input) && !alreadyuse {
			errorhangman++
			contenu = []int{}
		}
		if !IsAlphabet(input) {
			fmt.Println(color[1], "Choose only Latin alphabet.", color[0])
			i--
		}
		if len(contenu) > 0 || len(input) == 0 {
			i--
		} else if len(input) != 0 && len(contenu) == 0 && !alreadyuse && IsAlphabet(input) {
			fmt.Println(color[1])
			fmt.Printf("The letter %s is not present in the word \n", SliceToString(input))
			fmt.Println(color[0])
			errorhangman += 1
		}
		if errorhangman != -1 && errorhangman < 9 {
			PrintHangman(errorhangman)
		}
		affichage := PrintLetter(contenu, mot, hidden)
		if CompareString(mot, affichage) { // condition d'arret
			fmt.Print(color[3])
			fmt.Printf("Congrates ! You find the word %s, in %d attemps, with %d errors \n", SliceToString(mot), count, errorhangman+1)
			fmt.Print(color[0])
			break
		}
		if len(ascii_letter) == 0 {
			fmt.Println(SliceToString(affichage))
		} else {
			StringToAsciiArt(ascii_letter, affichage)
		}
		fmt.Println("used letter : ", SliceToString(lettre_utiliser))
		if 9-errorhangman <= 0 {
			fmt.Println("the word was :", SliceToString(mot))
			break
		} else {
			if 9-errorhangman <= 3 {
				fmt.Print(color[1])
				fmt.Printf("%d tries left\n", 9-errorhangman)
				fmt.Print(color[0])
			} else {
				fmt.Printf("%d tries left\n", 9-errorhangman)
			}
			fmt.Print("Choose : ")
			input = Input()
		}
		if len(input) <= 1 {
			contenu = ContainLetter(SliceToString(mot), SliceToString(input))
		}
		Clear()
	}
	if arret {
		fmt.Printf("Game Saved in %s_SAVE.txt.\n", SliceToString(input))
	}
	if save != "" {
		RemoveFile(save)
	}
}
