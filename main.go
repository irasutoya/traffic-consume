package main

import (
	aslog "github.com/anacrolix/log"
	"github.com/anacrolix/torrent"

	"github.com/0x0129/traffic-consume/app/client"
	"github.com/0x0129/traffic-consume/common/metahash"
	"github.com/0x0129/traffic-consume/storage"
)

func main() {
	getVersion()

	cfg := torrent.NewDefaultClientConfig()
	cfg.DefaultStorage = new(storage.Client)
	cfg.ExtendedHandshakeClientVersion = string([]byte{0xde, 0xad, 0xbe, 0xef})
	cfg.Logger = aslog.Default.WithFilterLevel(aslog.Error)

	cli, err := client.New(cfg)
	if err != nil {

	}
	defer cli.Close()

	// Monitor stats
	go cli.Monitor()

	// add fake file
	cli.AddFakeTorrent()

	// // default torrents
	cli.AddTorrents(metahash.GetDefaultMetaHashes())

	// block wait completed
	cli.WaitAll()
}
