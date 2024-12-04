package pdf

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

type ImageOptions struct {
	ImageType             string
	ReadDpi               bool
	AllowNegativePosition bool
}

func Createpdf(Name string, Prime string, ImgPath string, OutputName string) {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pageWidth, pageHeight := pdf.GetPageSize()
	options := gofpdf.ImageOptions{
		ImageType:             "",
		ReadDpi:               false,
		AllowNegativePosition: false,
	}
	pdf.Image(ImgPath, 20, 55, 170, 0, false, "", 0, "")
	imagePath := "./image/wantedVierge.png"
	pdf.ImageOptions(imagePath, 0, 0, pageWidth, pageHeight, false, options, 0, "")

	pdf.SetFont("Arial", "B", 45)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetY(260)
	pdf.WriteAligned(0, 0, Prime, "C")

	pdf.SetFont("Arial", "B", 48)
	pdf.SetTextColor(122, 29, 29)
	pdf.SetY(228)
	pdf.WriteAligned(0, 0, Name, "C")

	err := pdf.OutputFileAndClose(OutputName)

	if err != nil {
		fmt.Println("Erreur lors de la création du PDF:", err)
	} else {
		fmt.Println("PDF créé avec succès.")
	}
}
