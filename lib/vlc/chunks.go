package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChuncks []BinaryChunck

type BinaryChunck string

type encodingTable map[rune]string

const chunksSize = 8

func NewBinChunks(data []byte) BinaryChuncks {

	res := make(BinaryChuncks, 0, len(data))

	for _, part := range data {
		res = append(res, NewBinChunk(part))
	}

	return res
}

func NewBinChunk(code byte) BinaryChunck {
	return BinaryChunck(fmt.Sprintf("%08b"))
}

func spliteByChuncks(bStr string, chunkSize int) BinaryChuncks {
	strLen := utf8.RuneCountInString(bStr)

	chunkCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunkSize++
	}

	res := make(BinaryChuncks, 0, chunkCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunck(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunck(lastChunk))
	}
	return res
}

func (bsc BinaryChuncks) Join() string {
	var buf strings.Builder

	for _, bc := range bsc {
		buf.WriteString(string(bc))
	}

	return buf.String()
}

func (bcs BinaryChuncks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))

	for _, bs := range bcs {
		res = append(res, bs.Byte())
	}

	return res
}

func (bs BinaryChunck) Byte() byte {
	num, err := strconv.ParseUint(string(bs), 2, chunksSize)

	if err != nil {
		panic("cant oparse: " + err.Error())
	}

	return byte(num)
}
