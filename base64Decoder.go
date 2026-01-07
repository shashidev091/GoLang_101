package main

import (
	"fmt"
	"os"
)

var base64Table = [256]byte{
	'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7,
	'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15,
	'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
	'Y': 24, 'Z': 25,
	'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31, 'g': 32, 'h': 33,
	'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39, 'o': 40, 'p': 41,
	'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47, 'w': 48, 'x': 49,
	'y': 50, 'z': 51,
	'0': 52, '1': 53, '2': 54, '3': 55, '4': 56, '5': 57, '6': 58,
	'7': 59, '8': 60, '9': 61,
	'+': 62, '/': 63,
}

func decodeBase64Raw(input string) []byte {
	output := make([]byte, 0, len(input)*3/4)

	var buffer uint32
	var bits uint

	for i := 0; i < len(input); i++ {
		c := input[i]

		if c == '=' {
			break
		}

		val := base64Table[c]
		buffer = (buffer << 6) | uint32(val)
		bits += 6

		if bits >= 8 {
			bits -= 8
			output = append(output, byte(buffer>>bits))
			buffer &= (1<<bits - 1)
		}
	}

	return output
}

func Base64ToJpeg(base64Str string) {
	// Remove data URI prefix manually (no strings package)
	prefix := "data:image/jpeg;base64,"
	if len(base64Str) > len(prefix) && base64Str[:len(prefix)] == prefix {
		base64Str = base64Str[len(prefix):]
	}

	imageBytes := decodeBase64Raw(base64Str)

	file, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(imageBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println("JPEG written successfully: output.jpg")
}
