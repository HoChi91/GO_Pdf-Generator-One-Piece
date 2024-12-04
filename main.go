package main

import (
	csvparser "MyGeneratorPrime/CsvParser"
	generateprime "MyGeneratorPrime/GeneratePrime"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Not enough arguments")
	} else if args[0] == "-cli" {
		generateprime.Generate()
	} else if args[0] == "-file" {
		var csvpath string
		fmt.Println("Enter the csv file path :")
		fmt.Scanln(&csvpath)
		csvparser.ParseCSV(csvpath)
	} else {
		log.Fatal("Argument does not exist")
	}
}
