package main

import (
	"flag"
	"log"
	"os"
)

// flags
var (
	flagDry      = flag.Bool("dry", false, "dry run will only print changes without making them")
	flagNoRemove = flag.Bool("no-remove", false, "do not remove setup.go file after running")
)

var makeDirs = []string{
	"pkg", "internal", "cmd",
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	flag.Parse()

	if *flagDry {
		log.Println("running in dry mode")
	}
}

func main() {
	for _, dir := range makeDirs {
		log.Printf("creating dir '%s'\n", dir)
		if !*flagDry {
			checkErr(os.Mkdir(dir, 0755))
		}
	}

	if !*flagNoRemove {
		log.Printf("removing setup.go\n")
		if !*flagDry {
			checkErr(os.Remove("setup.go"))
		}
	}
}
