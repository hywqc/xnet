package udp

import (
	"fmt"
	"log"

	"golang.org/x/net/html"
)

func Send() {
	fmt.Printf("udp send 5")
	doc, err := html.Parse(nil)
	if err != nil {
		log.Fatal(err)
	}
	_ = doc

}
