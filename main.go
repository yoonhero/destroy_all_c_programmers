package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/yoonhero/destroy_all_c_programmers/utils"
)

func main() {
	dirname := "."
	fmt.Scanf("%s", &dirname)
	splited_dirname := strings.Split(dirname, "")
	if splited_dirname[len(splited_dirname)-1] != "/" {
		dirname = dirname + "/"
	}
	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	utils.HandleErr(err)

	for _, file := range files {
		selectedFile := strings.Split(file.Name(), ".")
		if selectedFile[len(selectedFile)-1] == "c" || selectedFile[len(selectedFile)-1] == "h" {

			dat, err := ioutil.ReadFile(dirname + file.Name())
			utils.HandleErr(err)
			data := string(dat)
			data = strings.Replace(data, ";", "z", -1)
			err = ioutil.WriteFile(dirname+file.Name(), []byte(data), 0644)
			utils.HandleErr(err)
		}
	}
}
