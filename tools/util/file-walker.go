package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type FilterFnc func(path string) bool

type FileWalker struct {
	Debug   bool
	DirProc []string
}

func (fw *FileWalker) GetAllFiles(rootItems []string, rootPath string, filter FilterFnc) []string {
	onlyFiles := []string{}
	fw.DirProc = make([]string, 0, 20)
	for _, pathItem := range rootItems {
		log.Printf("Process root %s", pathItem)
		itemAbs := path.Join(rootPath, pathItem)
		if info, err := os.Stat(itemAbs); err == nil && info.IsDir() {
			arr := []string{}
			arr = fw.getFilesinDir(itemAbs, pathItem, arr, "", filter)
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

func (fw *FileWalker) getFilesinDir(dirAbs string, dirRel string, ini []string, subpath string, filter FilterFnc) []string {
	r := ini
	//log.Println("Scan dir ", dirAbs)
	files, err := ioutil.ReadDir(dirAbs)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		itemAbs := path.Join(dirAbs, f.Name())
		if info, err := os.Stat(itemAbs); err == nil && info.IsDir() {
			nn := f.Name()
			if subpath != "" {
				nn = path.Join(subpath, f.Name())
			}
			fw.DirProc = append(fw.DirProc, nn)
			if fw.Debug {
				fmt.Println("** Sub dir found ", f.Name())
			}
			r = fw.getFilesinDir(itemAbs, path.Join(dirRel, f.Name()), r, nn, filter)
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
