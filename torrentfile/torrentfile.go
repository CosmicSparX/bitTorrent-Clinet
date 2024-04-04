package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"github.com/CosmicSparX/bencode-parser"
)

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func Open(path string) (TorrentFile, error) {
	bto, err := bencodeParser.OpenTorrent(path)
	if err != nil {
		return TorrentFile{}, err
	}
	return toTorrentFile(bto)
}

func hash(i *bencodeParser.BencodeInfo) ([20]byte, error) {
	var buf bytes.Buffer
	err := bencodeParser.Marshal(&buf, *i)
	if err != nil {
		return [20]byte{}, err
	}
	h := sha1.Sum(buf.Bytes())
	return h, nil
}

func splitPieceHashes(i *bencodeParser.BencodeInfo) ([][20]byte, error) {
	hashLen := 20 // Length of SHA-1 hash
	buf := []byte(i.Pieces)
	if len(buf)%hashLen != 0 {
		err := fmt.Errorf("received malformed pieces of length %d", len(buf))
		return nil, err
	}
	numHashes := len(buf) / hashLen
	hashes := make([][20]byte, numHashes)

	for i := 0; i < numHashes; i++ {
		copy(hashes[i][:], buf[i*hashLen:(i+1)*hashLen])
	}
	return hashes, nil
}

func toTorrentFile(bto bencodeParser.BencodeTorrent) (TorrentFile, error) {
	infoHash, err := hash(&bto.Info)
	if err != nil {
		return TorrentFile{}, err
	}
	pieceHashes, err := splitPieceHashes(&bto.Info)
	if err != nil {
		return TorrentFile{}, err
	}
	t := TorrentFile{
		Announce:    bto.Announce,
		InfoHash:    infoHash,
		PieceHashes: pieceHashes,
		PieceLength: bto.Info.PieceLength,
		Length:      bto.Info.Length,
		Name:        bto.Info.Name,
	}
	return t, nil
}

func main() {
	bto, err := Open("D:\\Programming stuff\\Projects\\Go\\bitTorrent Client\\torrentfile\\archlinux-2019.12.01-x86_64.iso.torrent")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bto.Announce, "hi")
}
