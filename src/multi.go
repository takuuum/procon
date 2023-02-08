package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	header := bytes.NewBufferString("---HEADER---\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("---FOOTER---\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)

	var buffer bytes.Buffer
	reader = bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)

	ioutil.ReadAll(teeReader)
	io.Copy(os.Stdout, teeReader)
	fmt.Println(buffer.String())
}
