package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func writeData(data []string) (string, error) {
	filename := "./data" + strconv.Itoa(rand.Int()) + ".txt"
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("writing file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error on file \"%v\" close", file.Name())
		}
	}(file)
	err = doWrite(file, data)
	if err != nil {
		return "", fmt.Errorf("writing file: %v\n", err)
	}
	return filename, nil
}

func doWrite(file *os.File, data []string) error {
	var builder strings.Builder
	for i := 0; i < len(data); i++ {
		builder.WriteString(data[i])
		if i < len(data)-1 {
			builder.WriteString("\n")
		}
	}
	code, err := file.WriteString(builder.String())
	fmt.Println(code)
	return err
}

func main() {
	todo := []string{"TODO", "1. Keep calm.", "2. Code some Go."}
	filename, err := writeData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("All data stored in  \"%v\" file\n", filename)
}
