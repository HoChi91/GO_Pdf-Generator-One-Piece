package generateprime

import (
	"MyGeneratorPrime/pdf"
	"MyGeneratorPrime/pirates"
	"bufio"
	"fmt"
	"os"
)

func Generate() {
	var Img string
	var NamePdf string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a name for your pirate (in MAJ): ")
	Name, _ := reader.ReadString('\n')

	fmt.Println("Enter amount for the prime of your pirate : ")
	Prime, _ := reader.ReadString('\n')

	fmt.Println("Enter the path of the image :")
	fmt.Scanln(&Img)

	fmt.Println("Enter the name of the pdf file:")
	fmt.Scanln(&NamePdf)
	NamePdf = NamePdf + ".pdf"

	Pirate, err := pirates.New(Name, Prime, Img)
	if err != nil {
		fmt.Println(err)
	}
	pdf.Createpdf(Pirate.Name, Pirate.Prime, Pirate.Img, NamePdf)
}
