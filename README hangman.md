
# Hangman
This is the Read-me of the hangman project


## Rules

Rules of the hangman game:

- You have to find the hidden words with only the few letters available
- if you enter a letter who is in the word,it will be revealed as many time as it is in the search word
- you can also enter the full word if you want
- every time you failed a stickman hanged will start to appear, becoming more and more completed
- if the hanged man appear totally, you have lose ,the hanged man died


## Installation

- go version go1.21.1
- git latest version

## Deployment

-first do this in your control panel(Debian etc..)

```bash
  git clone https://ytrack.learn.ynov.com/git/daethan/hangman.git
```
Then do 
```bash
    cd Hangman/
```

```bash
    go mod init hangman
```
to run the programm of the project

```bash
go run main/main.go
```




## Authors

- [@daethan](https://ytrack.learn.ynov.com/git/daethan)

