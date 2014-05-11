package stemmer

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func Test_stem(t *testing.T) {

	//read vocabulary
	fVoc, err := os.Open("voc.txt")
	if err != nil {
		panic(err)
	}
	vocab := bufio.NewReader(fVoc)
	//read output
	fOut, err := os.Open("output.txt")
	if err != nil {
		panic(err)
	}
	output := bufio.NewReader(fOut)

	for word, errV := vocab.ReadSlice('\n'); errV == nil; word, errV = vocab.ReadSlice('\n') {
		stem, errO := output.ReadSlice('\n')
		if errO != nil {
			panic(err)
		}

		sWord := strings.TrimSpace(string(word))
		sStem := strings.TrimSpace(string(stem))
		stemRes := Stem(sWord)

		if stemRes != string(sStem) {
			t.Error(
				"For", sWord,
				"expected", sStem,
				"got", stemRes,
			)
		}
	}
}
