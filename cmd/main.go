package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"rgr/internal/defuzz"
	"rgr/internal/function"
	"rgr/internal/graph"
)

type Data struct {
	Data []function.Function `json:"data"`
}

func main() {
	jsonFile, err := os.Open("base.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var data Data

	json.Unmarshal(byteValue, &data)

	for i, v := range data.Data {
		fmt.Println(i, v)
	}

	fmt.Println("----------------------------------------")

	var digits []int
	for _, v := range data.Data {

		tmpF := v.FuzzNums()

		for i, v := range tmpF {
			fmt.Printf("%d: %.5f\n", i, v)
		}

		tmpD := int(defuzz.MiddleMaximum(tmpF))

		fmt.Printf("Mu = %d\n", tmpD)

		digits = append(digits, tmpD)

		fmt.Println("----------------------------------------")
	}
	fmt.Println(digits)

	// TODO: graph
	tree := graph.BinaryTree{}

	for _, v := range digits {
		tree.Insert(v)
	}

	graph.Print(os.Stdout, tree.Root, 7, 'M')
}
