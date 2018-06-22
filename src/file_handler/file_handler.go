package file_handler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func CheckFile(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("The File %s does not exist", filename))
		return false
	} else {
		return true
	}
}

func ReadLines(path string) (lines []string, err error) {
	var (
		file *os.File
		part []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}
