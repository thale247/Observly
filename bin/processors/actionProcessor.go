package processors

import (
	"log"
	"time"
)

func ProcessAction() {
	log.Println("Processing Event")
	time.Sleep(10 * time.Second)
	log.Println("Processed Event")

}
