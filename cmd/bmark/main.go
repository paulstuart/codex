package main

import (
	"encoding/binary"
	"log"
	"os"
	"time"
)

type Dummy struct {
	Idx int64
	A   int64
	B   int64
	C   int64
	D   int64
}

func main() {
	const totalSize = 256 * 1024 * 1024 * 1024
	const dummySize = 40
	const totalRecs = totalSize / dummySize
	const filename = "/mnt/g/bigfiletest.dat"

	now := time.Now()
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	var dummy Dummy
	for i := int64(0); i < totalRecs; i++ {
		dummy.Idx = i
		if err = binary.Write(f, binary.LittleEndian, dummy); err != nil {
			log.Fatalln(err)
		}
	}
	log.Printf("wrote %d records in %s", totalRecs, time.Since(now))
	if err = f.Close(); err != nil {
		log.Fatalln(err)
	}
}
