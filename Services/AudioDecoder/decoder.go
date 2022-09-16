package AudioDecoder

import (
	"bufio"
	"github.com/hajimehoshi/go-mp3"
	"github.com/viert/go-lame"
	"log"
	"os"
)

func run(enc *lame.Encoder, file string) {
	inf, err := os.Open(file)
	if err != nil {
		log.Println("2", err)
	}
	defer inf.Close()
	d, err := mp3.NewDecoder(inf)
	if err != nil {
		log.Println("4", err)
	}
	r := bufio.NewReader(d)
	r.WriteTo(enc)
}
func cycle(slice []string, enc *lame.Encoder, pauseCount int) {
	pause := "./Audio/1s.mp3"
	for i, _ := range slice {
		run(enc, "./Audio/"+slice[i])
		for i := 0; i < pauseCount; i++ {
			run(enc, pause)
		}
	}
}

func Decoder() {
	slice := []string{"butter.mp3", "salt.mp3"}
	of, err := os.Create("./Audio/output.mp3")
	if err != nil {
		log.Println("1", err)
	}
	defer of.Close()
	enc := lame.NewEncoder(of)
	defer enc.Close()
	cycle(slice, enc, 1)
}
