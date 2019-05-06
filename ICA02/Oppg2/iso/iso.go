package iso

import (
	"fmt"
)

// Oppg b
const ascii = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f" +
	"\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f" +
	"\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf" +
	"\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf" +
	"\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf" +
	"\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf" +
	"\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xde\xee\xef" +
	"\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"

	// IterateOverExtendedASCIIStringLiteral For Løkke
func IterateOverExtendedASCIIStringLiteral() {
	for i := 128; i < 255; i++ {
		fmt.Printf("%X, %q, %b \n", i, i, i)
	}
}

//IterateOverExtendedASCIIStringLiteralB skriver ut tegn til stringen tekst
func IterateOverExtendedASCIIStringLiteralB(tekst string) {
	for i := 0; i < len(tekst); i++ {
		fmt.Printf("%X, %q, %b \n", tekst[i], tekst[i], tekst[i])
	}
}

// GetExtendedASCIIStringLiteral Denne funksjonen henter konstanten extended ascii
func GetExtendedASCIIStringLiteral() string {
	return ascii
}

// GreetingExtendedASCII bla bla
func GreetingExtendedASCII() {
	fmt.Printf("Salut, ça va °-) Ça coute €50")
	fmt.Printf("Salut, ça va °-) Κοστίζει €50")
	fmt.Printf("Salut, ça va °-) Κοστίζει €50 Forstår du?")
}

// ASCIISkrivUtASCII asdasd
func ASCIISkrivUtASCII() {
	for i := 128; i <= 255; i++ {
		fmt.Printf("%X, %q, %b \n", i, i, i)
	}
}
