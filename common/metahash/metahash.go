package metahash

import (
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/types/infohash"
)

// GetDefaultMetaHashes default torrents
func GetDefaultMetaHashes() []metainfo.Hash {
	return []metainfo.Hash{
		infohash.FromHexString("3f9aac158c7de8dfcab171ea58a17aabdf7fbc93"),
		infohash.FromHexString("41e6cd50ccec55cd5704c5e3d176e7b59317a3fb"),
		infohash.FromHexString("95c6c298c84fee2eee10c044d673537da158f0f8"),
		infohash.FromHexString("2e8e44068b254814ea1a7d4969a9af1d78e0f51f"),
		infohash.FromHexString("7e9d8be76259a0e0485721b426feff78761a70bb"),
		infohash.FromHexString("741e7c5fcfa9b2495f2d7dc3a430325987cd18bd"),
		infohash.FromHexString("e2b8ac8e7d5500850cfe550c97708f216fdf4791"),
		infohash.FromHexString("2fc6fc24f7d56d1def32ae0334d2df0cd3b855ea"),
		infohash.FromHexString("b84e74c1dbcc88a02c5b24a6f84383f353a2e1dd"),
	}
}

func NeedDropTorrents(old []metainfo.Hash, new []metainfo.Hash) []metainfo.Hash {
	if len(new) == 0 {
		return nil
	}

	newMap := make(map[metainfo.Hash]bool)
	for i := range new {
		newMap[new[i]] = true
	}

	var dropTorrents []metainfo.Hash
	for i := range old {
		if _, ok := newMap[old[i]]; !ok {
			dropTorrents = append(dropTorrents, old[i])
		}
	}
	return dropTorrents
}
