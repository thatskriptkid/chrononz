package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	gore "github.com/goretk/gore"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: chrononz.exe PATH_TO_FILE")
		os.Exit(0)
	}

	myFile := os.Args[1]
	
	
	fp, err := filepath.Abs(myFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse the filepath: %s.\n", err)
		os.Exit(1)
	}

	f, err := gore.Open(fp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when opening the file: %s.\n", err)
		os.Exit(1)
	}
	defer f.Close()

	pkgs, err := f.GetVendors()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when parsing packages: %s.\n", err)
		os.Exit(1)
	}

	vsInfo, err := GetVendorsInfo(pkgs)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get vendor info: %s.\n", err)
		os.Exit(1)
	}

	max := time.Time{}

	for _, vInfo := range vsInfo {
		fmt.Printf("%s %s\n", vInfo.PkgName, vInfo.Date.String())

		if vInfo.Date.After(max) {
			max = vInfo.Date
		}
	}

	fmt.Printf("Approximate (minimum) Timestamp equals = %s\n", max.String())
	
}