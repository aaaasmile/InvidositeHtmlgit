package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// semplice programma per convertire la directory out in unaltra iso-8890 per il deployment
// su kickers.fabbricadigitale.it
type FilterFnc func(path string) bool

type FileWalker struct {
	Debug bool
}

func (fw *FileWalker) GetAllFiles(rootItems []string, rootPath string, filter FilterFnc) []string {
	onlyFiles := []string{}
	for _, pathItem := range rootItems {
		log.Printf("Process item %s", pathItem)
		itemAbs := path.Join(rootPath, pathItem)
		if info, err := os.Stat(itemAbs); err == nil && info.IsDir() {
			arr := []string{}
			arr = fw.getFilesinDir(itemAbs, pathItem, arr, filter)
			if fw.Debug {
				fmt.Println("Dir process result: ", arr)
			}
			for _, ele := range arr {
				onlyFiles = append(onlyFiles, ele)
			}
		} else {
			if filter(pathItem) {
				onlyFiles = append(onlyFiles, pathItem)
			}
		}
	}
	return onlyFiles
}

func (fw *FileWalker) getFilesinDir(dirAbs string, dirRel string, ini []string, filter FilterFnc) []string {
	r := ini
	//log.Println("Scan dir ", dirAbs)
	files, err := ioutil.ReadDir(dirAbs)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		itemAbs := path.Join(dirAbs, f.Name())
		if info, err := os.Stat(itemAbs); err == nil && info.IsDir() {
			if fw.Debug {
				fmt.Println("** Sub dir found ", f.Name())
			}
			r = fw.getFilesinDir(itemAbs, path.Join(dirRel, f.Name()), r, filter)
		} else {
			if fw.Debug {
				fmt.Println("** file is ", f.Name())
			}
			fname := path.Join(dirRel, f.Name())
			if filter(fname) {
				r = append(r, fname)
			}
		}
	}
	return r
}

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
	fw := FileWalker{Debug: true}
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
}
