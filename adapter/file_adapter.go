package adapter

import (
	"bufio"
	"os"
)

type FileAdapter struct {
	filename string
	Lines    []string
}

func NewFileAdapter() *FileAdapter {
	return &FileAdapter{filename: "./dic.csv"}
}


func (this *FileAdapter) OpenFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
}

func (this *FileAdapter) fileScan(file *os.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		this.Lines = append(this.Lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (this *FileAdapter) ReadByLine() error {
	file, err := os.Open(this.filename)
	defer file.Close()
	if err != nil {
		return err
	}
	if err = this.fileScan(file); err != nil {
		return err
	}
	return nil
}

func (this *FileAdapter) writelines(file *os.File) error {
	for _, line := range this.Lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *FileAdapter) WriteByLine() error {
	file, err := os.Create(this.filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err = this.writelines(file); err != nil {
		return err
	}
	return nil
}
