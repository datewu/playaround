package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	createFile()
	deleteFile()
	checkExistence()
	renameFile()
	copyFile()
	writeToFile()
	writeToFileWithIOUtil()
	writeToFileWithBufferedWriter()
	readFile()
	readFileAgain()
	readWithBufferedReader()
	readWithScanner()
	createDir()
	createDirs()
	deleteDir()
	dirTraversal()
}

func createFile() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	fmt.Println("Created:", f.Name())
}

func deleteFile() {
	fn := "file.txt"
	f, err := os.Create(fn)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	err = os.Remove(fn)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Deleted:", fn)
}

func checkExistence() {
	fi, err := os.Stat("file.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Does not exist")
		}
		log.Fatalln(err)
	}
	fmt.Printf("Exists, last modified %v\n", fi.ModTime())
}

func renameFile() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	err = os.Rename(f.Name(), "rename.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func copyFile() {
	originF, err := os.Open("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer originF.Close()

	newF, err := os.Create("copy.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer newF.Close()

	byteWritten, err := io.Copy(newF, originF)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Bytes written:", byteWritten)

	if err := newF.Sync(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Copy finished")
}

func writeToFile() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if _, err := f.Write([]byte("Errors are values.\n")); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Write to file finished")
}

func writeToFileWithIOUtil() {
	b := []byte("Clear is better than clever.\n")
	if err := ioutil.WriteFile("myfile.txt", b, 0666); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Write to file with ioutil finished")
}

func writeToFileWithBufferedWriter() {
	f, err := os.Create("panic.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bufferdW := bufio.NewWriter(f)
	if _, err := bufferdW.WriteString("Donit panic.\n"); err != nil {
		log.Fatalln(err)
	}
	log.Println("Buffered:", bufferdW.Buffered())
	log.Println("Avilable:", bufferdW.Available())

	if err := bufferdW.Flush(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Write to file with bufferedWriter finished")
}

func readFile() {
	f, err := os.Open("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}

func readFileAgain() {
	bs, err := ioutil.ReadFile("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

func readWithBufferedReader() {
	f, err := os.Open("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	bs, err := br.ReadBytes('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))

	bs, err = br.ReadBytes('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

func readWithScanner() {
	f, err := os.Open("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	ln := 0
	for s.Scan() {
		ln++
		fmt.Printf("%d - %s", ln, s.Text())
	}

	if s.Err() != nil {
		log.Fatalln(err)
	}
	fmt.Println("Read with scanner finished")
}

func createDir() {
	if err := os.Mkdir("mydir", 0744); err != nil {
		log.Fatalln(err)
	}
	fi, err := os.Stat("mydir")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("is 'dir' a directory?; %t", fi.IsDir())
}

func createDirs() {
	dirs := []string{"dir", "sub", "sub3", "sub998"}
	path := filepath.Join(dirs...)
	if err := os.MkdirAll(path, 0744); err != nil {
		log.Fatalln(err)
	}
}

func deleteDir() {
	if err := os.Remove("mydir"); err != nil {
		log.Fatalln(err)
	}

	if err := os.RemoveAll("dir"); err != nil {
		log.Fatalln(err)
	}
}

func dirTraversal() {
	file, err := os.Create("combined.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	bw := bufio.NewWriter(file)

	f := func(p string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		log.Println(p)
		bs, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}

		if _, err := bw.Write(bs); err != nil {
			return err
		}
		bw.Flush()
		return nil
	}

	if err := filepath.Walk("proverbs", f); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Traverse dirs finished")
}
