package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"log"
	"os/exec"
	"strconv"
)

func main() {

	//Generate PDF from HTMl
	GenerateFromHtml("examples/sources/html/document.html", "examples/result/pdf/_document.generated.pdf")
	OptimizePdf("examples/result/pdf/_document.generated.pdf", "examples/result/pdf/_document.generated.optimized.pdf")

	// Create single page PDFs for in.pdf in outDir using the default configuration.
	SplitPdf("examples/result/pdf/_document.generated.optimized.pdf", "examples/result/split/generated")

	OptimizePdf("examples/sources/pdf/example.pdf", "examples/result/pdf/example.optimized.pdf")
	SplitPdf("examples/result/pdf/example.optimized.pdf", "examples/result/split/example")

	inFiles := []string{
		"examples/result/split/example/example.optimized_1.pdf",
		"examples/result/split/example/example.optimized_2.pdf",
		"examples/result/split/example/example.optimized_3.pdf",
		"examples/result/split/example/example.optimized_4.pdf",
		"examples/result/split/example/example.optimized_5.pdf",
	}

	MergePdf(inFiles, "examples/result/merged/example.pdf")

	PdfToJpg("examples/result/split/example/example.optimized_1.pdf", "examples/result/jpg/example.jpg", 400, 600)

	ClearExamples()
}

func GenerateFromHtml(inFile string, outFile string) string {

	arguments := []string{
		inFile,
		outFile,
	}

	cmd := exec.Command("wkhtmltopdf", arguments...)
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		return outFile
	}

	return outFile

}

func SplitPdf(inFile string, outFolder string) string {

	//Apply Split
	err := api.SplitFile(inFile, outFolder, 1, nil)
	if err != nil {
		fmt.Println(err.Error())
		return outFolder
	}

	return outFolder
}

func OptimizePdf(inFile string, outFile string) string {

	arguments := []string{
		inFile,
		outFile,
	}

	//Fix any PDF errors and Optimize
	cmd := exec.Command("ps2pdf", arguments...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return outFile
	}

	return outFile
}

func MergePdf(inFiles []string, outFile string) string {

	parameters := []string{
		"-sDEVICE=pdfwrite",
		"-dNOPAUSE",
		"-dBATCH",
		"-dSAFER",
		"-sOutputFile=" + outFile,
	}

	parameters = append(parameters, inFiles...)

	cmd := exec.Command("gs", parameters...)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return outFile
}

func PdfToJpg(inFile string, outFile string, sizeX int, sizeY int) string {

	parameters := []string{
		"-sDEVICE=jpeg",
		"-dNOPAUSE",
		"-dBATCH",
		"-dSAFER",
		"-r" + strconv.Itoa(sizeX) + "x" + strconv.Itoa(sizeY),
		"-sOutputFile=" + outFile,
		inFile,
	}

	cmd := exec.Command("gs", parameters...)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return outFile
}

func ClearExamples() {

	cmd := exec.Command("rm", []string{
		"-rf",
		"examples/result/*",
	}...)
	err := cmd.Run()

	cmd = exec.Command("mkdir", []string{
		"examples/result/pdf",
		"examples/result/split",
		"examples/result/split/example",
		"examples/result/split/generated",
		"examples/result/merged",
	}...)

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
