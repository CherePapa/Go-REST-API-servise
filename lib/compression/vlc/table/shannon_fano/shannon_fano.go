package shannon_fano

import (
	vlc "archiver/lib/compression/vlc/table"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Generator struct{}

type charStat map[rune]int

type encdonigTable map[rune]code

type code struct {
	Char    rune
	Quality int
	Bits    uint32
	Size    int
}

func NewGenetator() Generator {
	return Generator{}
}

func (g Generator) NewTable(text string) vlc.EncodingTable {
	return build(newCharStat(text)).Export()
}

func (et encdonigTable) Export() map[rune]string {
	res := make(map[rune]string)

	for k, v := range et {
		byteStr := fmt.Sprintf("%b", v.Bits)

		if lenDiff := v.Size - len(byteStr); lenDiff > 0 {
			byteStr = strings.Repeat("0", lenDiff) + byteStr
		}

		res[k] = byteStr
	}

	return res
}

func build(stat charStat) encdonigTable {
	codes := make([]code, 0, len(stat))

	for ch, qty := range stat {
		codes = append(codes, code{
			Char:    ch,
			Quality: qty,
		})
	}

	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quality != codes[j].Quality {
			return codes[i].Quality > codes[j].Quality
		}

		return codes[i].Char < codes[j].Char
	})

	assignCodes(codes)

	res := make(encdonigTable)

	for _, code := range codes {
		res[code.Char] = code
	}

	return res
}

func assignCodes(codes []code) {
	if len(codes) < 2 {
		return
	}

	divider := bestDeviderPosition(codes)

	for i := 0; i < len(codes); i++ {
		codes[i].Bits <<= 1
		codes[i].Size++

		if i >= divider {
			codes[i].Bits |= 1

		}
	}
	assignCodes(codes[:divider])
	assignCodes(codes[divider:])
}

func bestDeviderPosition(codes []code) int {

	total := 0
	for _, code := range codes {
		total += code.Quality
	}
	left := 0
	prevDiff := math.MaxInt
	bestPosition := 0

	for i := 0; i < len(codes)-1; i++ {
		rigth := total - left

		diff := abs(rigth - left)

		if diff > prevDiff {
			break
		}

		prevDiff = diff
		bestPosition = i + 1
	}

	return bestPosition
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func newCharStat(text string) charStat {
	res := make(charStat)

	for _, ch := range text {
		res[ch]++
	}

	return res
}
