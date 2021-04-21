package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func writeData(data []string, lineLen int) (string, error) {
	if lineLen <= 0 {
		panic("incorrect line length")
	}
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
	err = doWrite(file, data, lineLen)
	if err != nil {
		return "", fmt.Errorf("writing file: %v\n", err)
	}
	return filename, nil
}

func doWrite(file *os.File, data []string, lineLen int) error {
	var builder strings.Builder
	i := 0
	for builder.Len() < lineLen {
		str := data[i]
		i += 1
		if builder.Len()+len(str) <= lineLen {
			builder.WriteString(str)
		} else {
			l := lineLen - builder.Len()
			builder.WriteString(str[:l])
		}
	}
	code, err := file.WriteString(builder.String())
	fmt.Printf("Writen %d symbols\n", code)
	return err
}

func process(data []string, length int) {
	defer func() {
		val := recover()
		if val != nil {
			log.Printf("panicing: %v\n", val)
		}
	}()
	filename, err := writeData(data, length)
	fmt.Println("this 'll not execute on panic")
	if len(filename) > 0 {
		fmt.Printf("file \"%v\" was created;\n", filename)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	todo := []string{"TODO", "1. Keep calm.", "2. Code some Go."}
	process(todo, 10)
	process([]string{}, -1)
	fmt.Println("normal execution continues")
}
