package csvparser

import (
	"MyGeneratorPrime/pdf"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ParseCSV(filepath string) string {
	data, err := os.Open(filepath)
	if err != nil {
		fmt.Print("Erreur dans la lecture du csv ", err, "\n")
	}
	reader := csv.NewReader(data)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Print("Erreur dans la lecture du csv ", err, "\n")
	}
	var choice string
	fmt.Println("Do you want a pdf for all pirates ? (Y/n)")
	fmt.Scanln(&choice)
	if strings.ToUpper(choice) == "Y" {
		for _, line := range records {
			result := []string{line[0], line[1], line[2]}
			OutputName := string(result[0]) + "Prime.pdf"
			pdf.Createpdf(result[0], result[1], result[2], OutputName)
		}
	}
	var choice2 string
	if strings.ToLower(choice) == "n" {
		fmt.Println("What pirate do you want ?")
		fmt.Scanln(&choice2)
		for _, line := range records {
			if strings.Contains(line[0], strings.ToUpper(choice2)) {
				result := []string{line[0], line[1], line[2]}
				OutputName := string(result[0]) + "Prime.pdf"
				pdf.Createpdf(result[0], result[1], result[2], OutputName)

			}

		}

	}
	var r string
	return r

}
