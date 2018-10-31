package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"bytes"
)

func main() {

	// 参考  https://www.cnblogs.com/majianguo/p/8016426.html
	//readContent()
	//tempFile()
	//ReadFile()

	listFolder("../")

}

func readContent() {
	r := strings.NewReader("test string")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)

	fileName := "test.txt"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fb, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", fb)

	fb1, err := ioutil.ReadFile(fileName)
	fmt.Printf("%s\n", fb1)
}

func tempFile() {

	//TempDir在指定的dir目录新创建以prefix为前缀的临时目录，返回新临时目录的路径。如果dir为空字符串，
	// TempDir使用默认的存放临时文件的目录（os.TempDir，操作系统临时目录）。多个程序同时调用TempDir不会返回相同目录。
	// 调用者负责删除不再需要的临时目录。
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")

	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")

	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {

		log.Fatal(err)
	}

	//TempFile在dir目录中新创建以prefix为前缀的临时文件，打开文件用于读写，返回os.File指针。
	// 如果dir为空字符串，TempFile使用默认的存放临时文件的目录（os.TempDir，操作系统临时目录）。
	// 多个程序同时调用TempFile不会选择相同的文件。调用者负责删除不再需要的临时文件。
	tempFile, err := ioutil.TempFile("", "example1")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(content); err != nil {
		log.Fatalln(err)
	}

	if err := tempFile.Close(); err != nil {
		log.Fatalln(err)
	}
}

func ReadFile() {
	dir := "."
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func listFolder(dir string) (err error) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		fmt.Println(fi.Name())
		if fi.IsDir() {
			err = listFolder(filepath.Join(dir, fi.Name()))
			if err != nil {
				return err
			}
		}
	}
	return
}

func noCloser() {
	bf := bytes.NewBufferString("test")
	r := ioutil.NopCloser(bf)
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
}
