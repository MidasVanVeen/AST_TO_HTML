package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Attribute struct {
	Key string
	Value string
}

type Node struct {
	NodeType string
	Tag string
	Attrib []Attribute
	Children []Node
	Value string
}

func AST_to_HTML(AST Node) (string, error) {

	if AST.NodeType == "element" {
		attributes := ""
		for _,a := range AST.Attrib {
			attributes += a.Key + "=\"" + a.Value + "\""
		}
		html := "<" + AST.Tag + " " + attributes + ">"
		for _,c := range AST.Children {
			innerHtml, _ := AST_to_HTML(c)
			html += innerHtml
		}
		html += "</" + AST.Tag + ">"
		return html, nil
	}

	return AST.Value, nil
}

func main() {
	ASTFilename := os.Args[1]
	if ASTFilename == "" {
		panic("Enter a filename")
	}
	f, err := os.Open(ASTFilename)
	if err != nil {
		panic("Error reading file")
	}
	decoder := json.NewDecoder(f)
	var AST Node
	if err = decoder.Decode(&AST); err != nil {
		panic(err)
	}
	html,_ := AST_to_HTML(AST)
	fmt.Println(html)
}
