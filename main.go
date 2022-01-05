package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/IzumiSy/go-fdkaac"
	uuid "github.com/google/uuid"
	"github.com/nareix/joy4/format/rtmp"
)

func main() {
	server := &rtmp.Server{}
	server.HandlePublish = func(conn *rtmp.Conn) {
		defer func() {
			conn.Close()
			log.Println("Connection closed")
		}()

		log.Println("HandlePublish")

		id := uuid.New()
		log.Printf("ID: %s", id)

		file, err := os.Create(fmt.Sprintf("result-%s.pcm", id))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// AAC LC/48000Hz/monoralなASCの設定でfdk-aacのデコーダを初期化
		decoder := fdkaac.NewAacDecoder()
		if err := decoder.InitRaw([]byte{0x11, 0x88}); err != nil {
			panic(err)
		}
		defer decoder.Close()

		// Read packets
		pcmBuffer := new(bytes.Buffer)
		pcmBuffer.Grow(256)
		for {
			if pkt, err := conn.ReadPacket(); err != nil {
				if err == io.EOF {
					break
				}
				return
			} else {
				pcmBuffer.Reset()
				if err := decoder.Decode(pkt.Data, pcmBuffer); err != nil {
					log.Fatal(err)
				}

				written, err := file.Write(pcmBuffer.Bytes())
				if err != nil {
					log.Fatal(err)
				}

				log.Printf("Chunk: %d --> PCM %d", len(pkt.Data), written)
			}
		}

		log.Println("Done")
	}

	log.Println("Server running...")
	log.Fatal(server.ListenAndServe())
}
