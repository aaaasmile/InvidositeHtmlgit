package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github/aaasmile/vidotools/util"
)

// semplice programma per convertire la directory out in unaltra iso-8859 per il deployment
// su kickers.fabbricadigitale.it
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	var debug = flag.Bool("debug", false, "Debug")
	var ddSrc = flag.String("dirsrc", "", "Directory source in UTF8")
	var ddDst = flag.String("dirdts", "", "Directory destination in ISO8859")
	flag.Parse()

	debugFW := *debug
	debugFC := *debug
	dirSrc := "D:\\Projects\\GItHub\\InvidositeHtmlgit\\out\\snapshot"
	if *ddDst != "" {
		dirSrc = *ddSrc
	}

	dirOut := "D:\\Projects\\GItHub\\InvidositeHtmlgit\\out\\conv-iso8859"
	if *ddDst != "" {
		dirOut = *ddDst
	}

	pathItems := []string{"/"}
	extToCon := []string{".html", ".txt", ".xml"}
	extNoTouch := []string{".jpg", ".png"}
	fw := util.FileWalker{Debug: debugFW}
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
	if *debug {
		for _, item := range toConvs {
			fmt.Println(item)
		}

	}

	log.Println("Files to convert are conv/copy", len(toConvs), len(copyOnlys), len(resAll))
	log.Println("Scanned dirs ", fw.DirProc)

	fc := util.FileCreator{Debug: debugFC}
	fc.PrepareDirTree(fw.DirProc, dirOut)
	fc.CopyAllFiles(copyOnlys, dirOut, dirSrc)
	fc.Debug = true
	fc.ConvertToIsoAllFiles(toConvs, dirOut, dirSrc)
}
