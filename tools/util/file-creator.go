package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

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

func (fc *FileCreator) PrepareDirTree(dirs []string, root string) {
	for _, dd := range dirs {
		dout := path.Join(root, dd)
		if info, err := os.Stat(dout); err == nil && info.IsDir() {
			if fc.Debug {
				log.Println("Dir already exist ", dd)
			}
		} else {
			err := os.Mkdir(dout, 0755)
			if err != nil {
				log.Fatal("Error unable to create dir ", dout, err)
			}
			log.Println("Dir created", dout)
		}
	}
	log.Println("Directory tree ok on ", root, len(dirs))
}

func (fc *FileCreator) CopyAllFiles(list []string, dirOut, dirSrc string) {
	for _, item := range list {
		src := path.Join(dirSrc, item)
		dst := path.Join(dirOut, item)
		if err := fc.copyFile(src, dst); err != nil {
			log.Fatal("Copy error ", err)
		}
	}
	log.Println("Copy file success for ", len(list))
}

func (fc *FileCreator) ConvertToIsoAllFiles(list []string, dirOut, dirSrc string) {
	for _, item := range list {
		src := path.Join(dirSrc, item)
		dst := path.Join(dirOut, item)
		if err := copyWithTranslationIso(src, dst); err != nil {
			log.Fatal("Unable to create destination ", dst, err)
		}
		if fc.Debug {
			log.Println("File copied with iso 8859", item)
		}
	}
	log.Println("Copy file success for ", len(list))
}

func copyWithTranslationIso(src, dst string) error {
	fs, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fs.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	//r := charmap.ISO8859_1.NewEncoder()
	iso := ISOLatin1(InvalidSkip)
	r := iso.NewEncoder()
	_, err = io.Copy(r.Writer(destination), fs)
	if err != nil {
		log.Printf("*** Trouble in %s with error %v", dst, err)
		return nil
	}
	return err
}
