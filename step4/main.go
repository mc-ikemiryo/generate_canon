package main

import (
	"encoding/binary"
	"math"
	"os"
)

const (
	soundLength   = 1
	samplesPerSec = 44100
	tau           = 2 * math.Pi

	end = 1.0e-2

	C4 = 261.626
	D4 = 293.665
	E4 = 329.628
	F4 = 349.228
	G4 = 391.995
	A4 = 440.000
	B4 = 493.883
	C5 = 523.251
)

func main() {
	file := "out.bin"
	f, _ := os.Create(file)
	generate(C4, f)
	generate(D4, f)
	generate(E4, f)
	generate(F4, f)
	generate(G4, f)
	generate(A4, f)
	generate(B4, f)
	generate(C5, f)
}

func generate(frequency float64, file *os.File) {
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
