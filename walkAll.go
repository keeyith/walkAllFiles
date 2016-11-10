package main

import (
	  "path/filepath"
	  "os"
	  "flag"
	  "fmt"
	  "strings"
)

func visit(path string, f os.FileInfo, err error) error {
	//BaseName
	aBaseName := f.Name()
	aName := strings.Split(aBaseName,".")
	aBaseName = strings.Replace(aName[0],"&","&amp;",-1)
	
	//Extension
	anExtension := ""
	if len(anExtension) > 0 {
		anExtension = aName[1]
}
	//Directory
	aDirectory := strings.TrimSuffix(path,aBaseName)
	
	//Write Time
	aTimeString :=  fmt.Sprintf("%v", f.ModTime())
	timeStrings := strings.Split(aTimeString," ")
	dateStrings := strings.Split(timeStrings[0],"-")
	aWriteTime := dateStrings[1] + "/" + dateStrings[2] + "/" + dateStrings[0] + " " + timeStrings[1]
	
	//Size
	aLength := fmt.Sprintf("%v", f.Size())

	//Mode
	aMode := fmt.Sprintf("%v", f.Mode())

	fmt.Printf("<File>\n")
	if len(aBaseName) > 0 {
		fmt.Printf("<BaseName>%s</BaseName>\n", aBaseName)
	}
	if len(anExtension) > 0 {
		fmt.Printf("<Extension>.%s</Extension>\n", anExtension)
	}
	if len(aDirectory) > 0 {
		fmt.Printf("<Directory>%s</Directory>\n", aDirectory)
	}
	if len(aLength) > 0 {
		fmt.Printf("<Length>%v</Length>\n", aLength)
	}
	if len(aMode) > 0 {
		fmt.Printf("<Mode>%s</Mode>\n", aMode)
	}
	if len(aWriteTime) > 0 {
		fmt.Printf("<lastWriteTime>%s</lastWriteTime>\n", aWriteTime)
	}
	fmt.Printf("</File>\n")

	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fmt.Printf("<FileSystem>\n")
	err := filepath.Walk(root, visit)
	fmt.Printf("</FileSystem>\n")
	anErrMsg := fmt.Sprintf("%v", err)

	if len(anErrMsg) > 0 {
		fmt.Printf("<File>\n")
		fmt.Printf("<Error> %v", err)
		fmt.Printf("</Error>\n")
		fmt.Printf("<File>\n")
	
	}
}
