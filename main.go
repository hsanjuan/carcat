package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ipld/go-car"
	"github.com/ipld/go-car/util"
)

func main() {
	out := bufio.NewWriterSize(os.Stdout, 1024*1024)
	defer out.Flush()

	if len(os.Args) < 2 {
		fmt.Println("Usage carcat file1.car file2.car... > merged.car")
		os.Exit(1)
	}
	lastF, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer lastF.Close()

	lastBuf := bufio.NewReader(lastF)
	lastHeader, err := car.ReadHeader(lastBuf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	finalHeader := &car.CarHeader{
		Roots:   lastHeader.Roots,
		Version: 1,
	}

	if err := car.WriteHeader(finalHeader, out); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, carFileName := range os.Args[1:] {
		f, err := os.Open(carFileName)
		if err != nil {
			log.Fatal(err)
			return
		}

		buf := bufio.NewReaderSize(f, 1024*1024)
		_, err = car.ReadHeader(buf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for {
			bs, err := util.LdRead(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = util.LdWrite(out, bs)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
	}
}
