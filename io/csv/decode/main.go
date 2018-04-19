package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	_ "os"
	"strings"

	"github.com/BobbyShaftoe/gotour/io/csv/types"
	_ "github.com/BobbyShaftoe/gotour/io/csv/types"
	log "github.com/sirupsen/logrus"
)

const (
	csvDb = "./data/resource_catalogue.csv"
)

func main() {
	log.Info("CSV decoding")

	f, err := os.Open(csvDb)
	if nil != err {
		log.Fatal(err)
	}
	defer f.Close()

	// New reader and read header line
	r := csv.NewReader(f)
	header, err := r.Read()
	fmt.Printf("\nCSV Fields: %v\n\n", strings.Join(header, ", "))

	for {
		csvRecord, err := r.Read()
		if nil == err {
			process(csvRecord)
		} else if io.EOF == err {
			break
		} else {
			log.Fatal(err)
		}
	}

}

func process(ss []string) {
	u := &types.Resources{}
	u.FromCSV(ss)
	fmt.Println(u.Path, u.Resource)
}
