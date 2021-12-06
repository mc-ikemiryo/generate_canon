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

	C3  = 130.813
	D3  = 146.832
	E3  = 164.814
	F3  = 174.614
	Fs3 = 184.997
	G3  = 195.998
	A3  = 220.000
	As3 = 233.082
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
	generate(A3, D4, Fs4, f)
	generate(A3, D4, Fs4, f)

	generate(A3, Cs4, E4, f)
	generate(A3, Cs4, E4, f)

	generate(B3, D4, Fs4, f)
	generate(B3, D4, Fs4, f)

	generate(Fs3, A3, Cs4, f)
	generate(Fs3, A3, Cs4, f)

	generate(G3, B3, D4, f)
	generate(G3, B3, D4, f)

	generate(Fs4, A3, D4, f)
	generate(Fs4, A3, D4, f)

	generate(G3, B3, D4, f)
	generate(G3, B3, D4, f)

	generate(A3, Cs4, E4, f)
	generate(A3, Cs4, E4, f)
}

func generate(frequency1, frequency2, frequency3 float64, file *os.File) {
	samples := samplesPerSec * soundLength
	damping := math.Pow(end, 1.0/float64(samples))
	for i := 0; i < samples; i++ {
		sample := 0.333*math.Sin((tau*frequency1*float64(i))/samplesPerSec) +
			0.333*math.Sin((tau*frequency2*float64(i))/samplesPerSec) +
			0.333*math.Sin((tau*frequency3*float64(i))/samplesPerSec)

		sample = sample * math.Pow(damping, float64(i))
		buf := make([]byte, 4)

		// バイト順序=LittleEndian
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
