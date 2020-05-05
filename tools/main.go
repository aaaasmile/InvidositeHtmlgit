package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github/aaasmile/vidotools/util"
)

// semplice programma per convertire la directory out in unaltra iso-8890 per il deployment
// su kickers.fabbricadigitale.it
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type FileCreator struct {
	Debug bool
}

func (fc *FileCreator) copyFile(src, dst string) error {
	if fc.Debug {
		log.Println("Copy file ", src, dst)
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func main() {
	dirSrc := "D:\\Projects\\GItHub\\InvidositeHtmlgit\\out\\snapshot"
	//dirOut := "D:\\Projects\\GItHub\\InvidositeHtmlgit\\out\\conv-iso8890"

	pathItems := []string{"/"}
	extToCon := []string{".html", ".txt", ".xml"}
	extNoTouch := []string{".jpg", ".png"}
	fw := util.FileWalker{Debug: true}
	toConvs := make([]string, 0, 300)
	copyOnlys := make([]string, 0, 200)
	resAll := fw.GetAllFiles(pathItems, dirSrc, func(filename string) bool {
		ext := strings.ToLower(filepath.Ext(filename))
		if stringInSlice(ext, extToCon) {
			toConvs = append(toConvs, filename)
			return true
		} else if stringInSlice(ext, extNoTouch) {
			copyOnlys = append(copyOnlys, filename)
			return true
		}
		return false
	})
	for _, item := range toConvs {
		fmt.Println(item)
	}
	log.Println("Files to convert are conv/copy", len(toConvs), len(copyOnlys), len(resAll))
	log.Println("Scanned dirs ", fw.DirProc)
}
