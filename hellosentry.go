package main

import (
	"fmt"
	"log"
	"os"

	"github.com/getsentry/raven-go"
)

func main() {
	f, err := os.Open("filename.ext")
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}
	fmt.Print(f)
}
