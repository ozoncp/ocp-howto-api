package main

import (
	"fmt"
	"os"

	"github.com/ozoncp/ocp-howto-api/internal/howto"
	"github.com/ozoncp/ocp-howto-api/internal/utils"
)

func openFile(name string, times int) error {
	open := func(name string) error {
		file, err := os.Open(name)
		if err != nil {
			return err
		}
		defer file.Close()
		fmt.Println(file.Name())
		return nil
	}
	for i := 0; i < times; i++ {
		err := open(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	fmt.Println("Howto API. Author: Ivan Levin")

	if openFile("test.txt", 5) == nil {
		fmt.Println("File opened successfully.")
	} else {
		fmt.Println("File open failed.")
	}

	howtos := []howto.Howto{
		*howto.New(1, "question0", "answer0"),
		*howto.New(1, "question1", "answer1"),
		*howto.New(1, "question2", "answer2"),
	}
	for _, v := range howtos {
		fmt.Println(v)
	}

	fmt.Println(utils.SplitToBulks(howtos, 2))

	m, err := utils.ConvertToMap(howtos)
	if err != nil {
		fmt.Printf("Convert to map failed: %v", err)
		fmt.Println()
	}
	for key, value := range m {
		fmt.Printf("%v:\r\n%v", key, value)
		fmt.Println()
	}
}
