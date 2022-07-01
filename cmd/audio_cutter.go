package main

import (
	"github.com/sandovalrr/mediacutter/cutter"
	"github.com/sandovalrr/mediacutter/models"
)

func main() {
	audioCutter := cutter.NewAudioCutter(models.CutterOption{
		Path:      "/home/rahma/Downloads/SS149_Kassywedding.mp3",
		Samples:   5,
		ChunkPath: "/home/rahma/speech_files",
	})

	audioCutter.Split()
}
