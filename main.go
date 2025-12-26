package main

import (
	"log"
	
	"bes-chromie/src/serve"
	"bes-chromie/src/capture"
)

func main() {

	_, err := capture.New("http://149.28.13.238:51480/canvas/textbasis/yycjunk?")
	if err!=nil {
		log.Fatal(err)
	}

	serve.Run()
	
}
