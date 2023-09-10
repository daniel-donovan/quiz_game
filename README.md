# A simple quiz game implemented in Go, from 'Gophercises' (https://courses.calhoun.io/lessons/les_goph_01)

This project implements exercise 1 from https://courses.calhoun.io/lessons/les_goph_01, where you create a simple quiz game.  

The quiz itself is simple in it's default state, consisting of small addition problems. The focus of the project was not to implement a challenging quiz, but to create a program using Go. This solution was created on my own, without referencing the answers presented on the website, only reading the problem.  

Lessons/concepts I didn't know before beginning, but discovered throughout implementation:
- 'flag' package and how it can easily parse options for you and generates it's own help
- how to easily read a .csv file by importing the encoding instead of implementing my own parser
- input channels, goroutines, and concurrency to get the timer working. I was originally experimenting with context, but I believe that's more useful for web responses than waiting on input

Compile using Go:  
- go build quiz.go

Command line options:  
- \-h: print help  
- \-csv: path\to\a\custom\.csv  
- \-limit: a custom time limit per question
