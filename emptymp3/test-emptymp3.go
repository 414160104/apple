package main

import (
	"encoding/binary"
	"os"
	"time"
)

func main() {
	outputFilename := "empty.mp3"
	file, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writeWaveHeader(file, 44100)

	duration := time.Millisecond
	sampleRate := 44100
	numSamples := int(float64(sampleRate) * duration.Seconds())
	samples := make([]int16, numSamples)

	err = binary.Write(file, binary.LittleEndian, samples)
	if err != nil {
		panic(err)
	}

	println("Wave file created successfully: ", outputFilename)
}

func writeWaveHeader(file *os.File, sampleRate int) {
	header := []byte{
		'R', 'I', 'F', 'F', // Chunk ID
		0, 0, 0, 0, // Chunk Size (to be filled in later)
		'W', 'A', 'V', 'E', // Format
		'f', 'm', 't', ' ', // Subchunk1 ID
		16, 0, 0, 0, // Subchunk1 Size
		1, 0, // Audio Format (PCM)
		1, 0, // Num Channels
		0, 0, 0, 0, // Sample Rate (to be filled in later)
		0, 0, 0, 0, // Byte Rate (to be filled in later)
		2, 0, // Block Align
		16, 0, // Bits Per Sample
		'd', 'a', 't', 'a', // Subchunk2 ID
		0, 0, 0, 0, // Subchunk2 Size (to be filled in later)
	}

	binary.LittleEndian.PutUint32(header[4:8], uint32(len(header)-8))               // Chunk Size
	binary.LittleEndian.PutUint32(header[24:28], uint32(sampleRate))               // Sample Rate
	binary.LittleEndian.PutUint32(header[28:32], uint32(sampleRate*2*1))           // Byte Rate
	binary.LittleEndian.PutUint32(header[40:44], uint32(len(header)-44))           // Subchunk2 Size
	file.Write(header)
}
