package main

import (
	"fmt"
	"os"
)

func main() {
	//files -> collection of bytes
	//read -> read from file
	//write -> write to file

	//open a file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name())        // file name
	// fmt.Println("file or folder:", fileInfo.IsDir())  // true if folder
	// fmt.Println("file size:", fileInfo.Size())        // in bytes
	// fmt.Println("file permissions:", fileInfo.Mode()) //file mode
	// fmt.Println("last modified:", fileInfo.ModTime()) // last modified

	//read from file - method1
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// fileInfo, err := f.Stat()
	// buf := make([]byte, fileInfo.Size())
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := range buf {
	// 	println("data", d, string(buf[i]))
	// }

	// println("data", d, buf)

	//read from file - method2 - for small files
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// println("data :", string(data))

	//read folders

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1) // -1 for all files, 0 for only directories

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())
	// }

	//create a file

	// f, err := os.Create("newFile.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// //append to file
	// // f.WriteString("hello world")
	// // f.WriteString("how are you!!!")

	// //f.Seek(0, 0) // replace content

	// bytes := []byte("hello Golang")
	// f.Write(bytes)

	// read from  a file and write to another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destinationFile, err := os.Create("newFile.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destinationFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destinationFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()

	// fmt.Println("written to new file successfully")

	//delete a file
	err := os.Remove("newFile.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("file deleted successfully")

}
