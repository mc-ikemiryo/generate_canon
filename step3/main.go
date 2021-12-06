package main

import (
	"encoding/binary"
	"math"
	"os"
)

const (
	soundLength   = 5
	samplesPerSec = 44100
	tau           = 2 * math.Pi
	frequency     = 440

	end = 1.0e-5
)

func main() {
	file, _ := os.Create("out.bin")

	samples := samplesPerSec * soundLength
	damping := math.Pow(end, 1.0/float64(samples))
	for i := 0; i < samples; i++ {
		sample := math.Sin((tau * frequency * float64(i)) / samplesPerSec)
		sample = sample * math.Pow(damping, float64(i))
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
