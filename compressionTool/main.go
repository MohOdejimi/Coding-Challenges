package main

import (
	"fmt"
	"os"
	"bufio"
	"container/heap"

	"compressionTool/huffman"
	"compressionTool/minheap"
	"compressionTool/traversal"
)


func main() {
	args := os.Args[1:]
	filepath := args[0]

	file, err := os.Open(filepath)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not found")
		} else if os.IsPermission(err) {
			fmt.Println("Permisiion Denied")
		} else {
			fmt.Println(err.Error())
		}
	}
	defer file.Close() 

	scanner := bufio.NewScanner(file)	

	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	characterCount := make(map[rune]int) 

	for _, s := range input {
		for _, ch := range s {
			characterCount[ch]++
		}
	}

	mh := minheap.MinHeap{}

	for ch, freq := range characterCount {
		node := &minheap.TreeNode{
			Freq: freq,
			Char: ch,
		}
		heap.Push(&mh, node)
	}

	root := huffman.HuffManImplementation(mh)

	rootStackDeets := &traversal.Stack{
		Node: root,
		CurrentCode: "",
	}

	var codes = make(map[rune]string)

	traversal.StckSlice = append(traversal.StckSlice, rootStackDeets)
	traversal.PreOrderTraversal(traversal.StckSlice, codes)

	for ch, code := range codes {
		fmt.Printf("%c: %s\n", ch, code)
	}
}

