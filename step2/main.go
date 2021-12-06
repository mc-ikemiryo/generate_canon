package main

import (
	"encoding/binary"
	"math"
	"os"
)

const (
	soundLength   = 5
	samplesPerSec = 44100
	frequency     = 440
	tau           = 2 * math.Pi
)

func main() {
	file, _ := os.Create("out.bin")

	samples := samplesPerSec * soundLength
	for i := 0; i < samples; i++ {
		sample := math.Sin((tau * frequency * float64(i)) / samplesPerSec)
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
