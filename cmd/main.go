package main

import (
	"crypto/sha1"
	"os"
)

const PIECE_LENGTH = 256 * 1024 // 256 kB

type FileMetadata struct {
	Filename string
	FileSize int64
}

type TorrentMetadata struct {
}

func main() {

	args := os.Args
	parseFile(args[1])

}

func parseFile(filename string) ([]byte, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fStats, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fSize := fStats.Size()

	pieceNum := fSize/PIECE_LENGTH + fSize&PIECE_LENGTH

	pieces := make([]byte, pieceNum*20)
	buffer := make([]byte, PIECE_LENGTH)
	for idx := range pieceNum {
		_, err := f.Read(buffer)
		if err != nil {
			return nil, err
		}
		hash := sha1.Sum(buffer)
		copy(pieces[idx*int64(20):idx*int64(20)+20], hash[:])
	}

	return pieces, nil
}
