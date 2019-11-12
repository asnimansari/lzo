package main

import (
	"fmt"
	"github.com/cyberdelia/lzo"
	"io"
	"os"
)

func compress( path string) error {
	level := lzo.BestCompression
	input, err := os.Open(path)
	if err != nil {
		return err
	}
	output, err := os.Create(path + ".lzo")
	if err != nil {
		return err
	}
	compressor, err := lzo.NewWriterLevel(output, level)
	defer compressor.Close()
	compressor.Name = input.Name()
	if err != nil {
		return err
	}
	_, err = io.Copy(compressor, input)
	if err != nil {
		return err
	}
	return nil
}


func decompress(path string) error {
	input, err := os.Open(path)
	if err != nil {
		return err
	}
	decompressor, err := lzo.NewReader(input)
	if err != nil {
		return err
	}
	output, err := os.Create(decompressor.Name)

	fmt.Println(output.Name())
	if err != nil {
		return err
	}
	_, err = io.Copy(output, decompressor)
	if err != nil {
		return err
	}
	return nil
}



func main(){
	//compress("testdata/pg135.txt")
	decompress("testdata/pg1351234.txt.lzo")
}
