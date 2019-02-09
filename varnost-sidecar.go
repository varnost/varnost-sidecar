package main

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Error   *log.Logger
)

func logInit(
	infoHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	logInit(os.Stdout, os.Stderr)

	Info.Println("Starting Varnost Sidecar")

}