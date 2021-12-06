package main

import (
	"encoding/binary"
	"math"
	"os"
)

const (
	samplesPerSec = 44100
	tau           = 2 * math.Pi

	end = 1.0e-1

	C3  = 130.813
	D3  = 146.832
	E3  = 164.814
	F3  = 174.614
	Fs3 = 184.997
	G3  = 195.998
	A3  = 220.000
	B3  = 246.942

	C4  = 261.626
	Cs4 = 277.183
	D4  = 293.665
	E4  = 329.628
	F4  = 349.228
	Fs4 = 369.994
	G4  = 391.995
	A4  = 440.000
	B4  = 493.883

	C5  = 523.251
	Cs5 = 554.365
	D5  = 587.330
	Ds5 = 622.254
	E5  = 659.255
	F5  = 698.456
	Fs5 = 739.989
	G5  = 783.991
	A5  = 880.000
	B5  = 987.767
)

func main() {
	file := "out.bin"
	f, _ := os.Create(file)
	generate(A5, 2, f)
	generate(Fs5, 1, f)
	generate(G5, 1, f)

	generate(A5, 2, f)
	generate(Fs5, 1, f)
	generate(G5, 1, f)

	generate(A5, 1, f)
	generate(A4, 1, f)
	generate(B4, 1, f)
	generate(Cs5, 1, f)

	generate(D5, 1, f)
	generate(E5, 1, f)
	generate(Fs5, 1, f)
	generate(G5, 1, f)

	generate(Fs5, 2, f)
	generate(D5, 1, f)
	generate(E5, 1, f)

	generate(Fs5, 2, f)
	generate(Fs4, 1, f)
	generate(G4, 1, f)

	generate(A4, 1, f)
	generate(B4, 1, f)
	generate(A4, 1, f)
	generate(G4, 1, f)

	generate(A4, 1, f)
	generate(Fs4, 1, f)
	generate(G4, 1, f)
	generate(A4, 1, f)

	generate(G4, 2, f)
	generate(B4, 1, f)
	generate(A4, 1, f)

	generate(G4, 2, f)
	generate(Fs4, 1, f)
	generate(E4, 1, f)

	generate(Fs4, 1, f)
	generate(E4, 1, f)
	generate(D4, 1, f)
	generate(E4, 1, f)

	generate(Fs4, 1, f)
	generate(G4, 1, f)
	generate(A4, 1, f)
	generate(B4, 1, f)

	generate(G4, 2, f)
	generate(B4, 1, f)
	generate(A4, 1, f)

	generate(B4, 2, f)
	generate(Cs5, 1, f)
	generate(D5, 1, f)

	generate(A4, 1, f)
	generate(B4, 1, f)
	generate(Cs5, 1, f)
	generate(D5, 1, f)

	generate(E5, 1, f)
	generate(Fs5, 1, f)
	generate(G5, 1, f)
	generate(A5, 1, f)
}

func generate(frequency float64, soundLength int, file *os.File) {
	samples := (soundLength * samplesPerSec) / 4
	damping := math.Pow(end, 1.0/float64(samples))
	for i := 0; i < samples; i++ {
		sample := math.Sin((tau * frequency * float64(i)) / float64(samplesPerSec))
		sample = sample * math.Pow(damping, float64(i))
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
