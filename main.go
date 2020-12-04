package main

import (
	"fmt"
	"os"

	"github.com/antchfx/xmlquery"
)

func processErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	filePath := os.Args[1]
	xPathExpression := os.Args[2]

	f, err := os.Open(filePath)
	processErr(err)

	doc, err := xmlquery.Parse(f)
	processErr(err)

	nodes, err := xmlquery.QueryAll(doc, xPathExpression)
	processErr(err)

	for _, node := range nodes {
		fmt.Println(node.Data)
	}
}
