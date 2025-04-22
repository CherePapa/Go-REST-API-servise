package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChuncks []BinaryChunck

type BinaryChunck string

type HexChuncks []HexChunck

type HexChunck string

type encodingTable map[rune]string

const chunksSize = 8

const hexChuncksSeparation = " "

func NewHexChunks(str string) HexChuncks {
	parts := strings.Split(str, hexChuncksSeparation)

	res := make(HexChuncks, 0, len(parts))

	for _, part := range parts {
		res = append(res, HexChunck(part))
	}

	return res
}

func (hcs HexChuncks) ToString() string {

	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))

	for _, hc := range hcs[1:] {
		buf.WriteString(hexChuncksSeparation)
		buf.WriteString(string(hc))

	}

	return buf.String()
}

// Join joins chunks into one line and returns as string
func (bsc BinaryChuncks) Join() string {
	var buf strings.Builder

	for _, bc := range bsc {
		buf.WriteString(string(bc))
	}

	return buf.String()
}

func (hcs HexChuncks) ToBinary() BinaryChuncks {
	res := make(BinaryChuncks, 0, len(hcs))

	for _, chunk := range hcs {
		bChunk := chunk.ToBinary()

		res = append(res, bChunk)
	}

	return res
}

func (hs HexChunck) ToBinary() BinaryChunck {
	num, err := strconv.ParseUint(string(hs), 16, chunksSize)

	if err != nil {
		panic("cant parse binary chunk: " + err.Error())
	}

	res := fmt.Sprint("%08b", num)

	return BinaryChunck(res)
}

func (bcs BinaryChuncks) ToHex() HexChuncks {
	res := make(HexChuncks, 0, len(bcs))

	for _, chunck := range bcs {
		hexChunck := chunck.ToHex()
		res = append(res, hexChunck)
	}

	return res
}

func (bs BinaryChunck) ToHex() HexChunck {
	num, err := strconv.ParseUint(string(bs), 2, chunksSize)
	if err != nil {
		panic("cant parse binary chunk: " + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%X", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunck(res)
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
