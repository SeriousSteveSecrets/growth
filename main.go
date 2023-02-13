package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type Myfile struct {
	name   string
	path   string
	size   int64
	growth int64
}

func fileList(s string) []Myfile {
	//get list of files, save as slice with custom type
	var fileslice []Myfile
	files, err := os.ReadDir(s)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		info, _ := file.Info()
		fname := info.Name()
		fpath := s + info.Name()
		fsize := info.Size()
		var new Myfile = Myfile{fname, fpath, fsize, 0}
		fileslice = append(fileslice, new)
	}
	return fileslice
}

func fileGrowth(fileslice []Myfile){
	for i, file := range fileslice {
		//check file size, new file size minus old file size, apply to growth, order, print
		stat, _ := os.Stat(file.path)
		g := stat.Size() - file.size
		fileslice[i].growth = file.growth + g
	}
	sort.Slice(fileslice, func(i, j int) bool {
		return fileslice[i].growth > fileslice[j].growth
	})
	for _, item := range fileslice {
		fmt.Printf("%+v\n", item)

	}
	fmt.Printf("\n")
}

func main() {
	fl := fileList(os.Args[1])
	for {
		fileGrowth(fl)
		time.Sleep(10 * time.Second)
	}	
}
