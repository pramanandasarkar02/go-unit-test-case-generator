package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// take flags as cmd line


	inputFilePath := flag.String("i", "", "-i <input_file_name>")
	flag.Parse()

	if *inputFilePath == "" {
		panic("input file required!")
	}
	inputFileName := filepath.Base(*inputFilePath)
	inputFileParts := strings.Split(inputFileName, ".")
	if len(inputFileParts) < 2 {
		log.Fatal("invalid file name format")
	}

	outputFileName := filepath.Dir(*inputFilePath) + "/" + inputFileParts[0] + "_test." + inputFileParts[1]

	fmt.Println("Output file name:", outputFileName)

	outputFile , err := os.Create(outputFileName)
	if err != nil{
		panic(err)
	}

	outputFile.WriteString("demo text is written \n in that \t\t\tfile")

	



	fmt.Println("cmd line flags loading...")

	// find functions and structs and their coresponding value

	fmt.Println("func and type structing...")


	// test function writting 
	
	fmt.Println("test function generating...")
	
	// writting file

	fmt.Println("writting file...")



}