package main

import (
	"encoding/binary"
	"io"
	"math"
	"os"
)

func main() {
	melodyFile, _ := os.Open("./melody.bin")
	defer melodyFile.Close()
	melodyF := []float32{}
	for {
		b := make([]byte, 4)
		_, err := melodyFile.Read(b)
		if err == io.EOF {
			break
		}
		u := binary.LittleEndian.Uint32(b)
		f := math.Float32frombits(u)
		melodyF = append(melodyF, f)
	}

	codeFile, _ := os.Open("./code.bin")
	defer codeFile.Close()
	codeF := []float32{}
	for {
		b := make([]byte, 4)
		_, err := codeFile.Read(b)
		if err == io.EOF {
			break
		}
		u := binary.LittleEndian.Uint32(b)
		f := math.Float32frombits(u)
		codeF = append(codeF, f)
	}

	file := "out.bin"
	f, _ := os.Create(file)
	samples := 705600
	for i := 0; i < samples; i++ {
		sample := 0.5*melodyF[i] + 0.5*codeF[i]
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		f.Write(buf)
	}
}
