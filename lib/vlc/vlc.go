package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) string {
	str = prepateText(str)

	chunck := spliteByChuncks(encodeBin(str), chunksSize)

	return chunck.ToHex().ToString()
}

func Decode(encodedText string) string {
	bString := NewHexChunks(encodedText).ToBinary().Join()

	dTree := getEncodingTable().DecodingTree()

	return exportText(dTree.Decode(bString))
}

// spliteByChancks разбивает текст на чанки
// 10001,10010101

// Метод prepareText преобразует текст в нижний регистр
// изменяет текст: ! + нижний регистр
// i.g.: My name is Ted -> !my name is !ted
func prepateText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

// exportText is oppsite to prepateText
func exportText(str string) string {
	var buf strings.Builder

	var isCapital bool

	for _, ch := range str {

		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false

			continue
		}
		if ch == '!' {
			isCapital = true

			continue
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("uncnow character: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		'a': "01100001",
		'b': "01100010",
		'c': "01100011",
		'd': "01100100",
		'e': "01100101",
		'f': "01100110",
		'g': "01100111",
		'h': "01101000",
		'i': "01101001",
		'j': "01101010",
		'k': "01101011",
		'l': "01101100",
		'm': "01101101",
		'n': "01101110",
		'o': "01101111",
		'p': "01110000",
		'q': "01110001",
		'r': "01110010",
		's': "01110011",
		't': "01110100",
		'u': "01110101",
		'v': "01110110",
		'w': "01110111",
		'x': "01111000",
		'y': "01111001",
		'z': "01111010",
		'A': "01000001",
		'B': "01000010",
		'C': "01000011",
		'D': "01000100",
		'E': "01000101",
		'F': "01000110",
		'G': "01000111",
		'H': "01001000",
		'I': "01001001",
		'J': "01001010",
		'K': "01001011",
		'L': "01001100",
		'M': "01001101",
		'N': "01001110",
		'O': "01001111",
		'P': "01010000",
		'Q': "01010001",
		'R': "01010010",
		'S': "01010011",
		'T': "01010100",
		'U': "01010101",
		'V': "01010110",
		'W': "01010111",
		'X': "01011000",
		'Y': "01011001",
		'Z': "01011010",
		'!': "00100001",
	}
}
