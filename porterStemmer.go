package stemmer

import (
	"strings"
)

type PorterStemmer interface {
	Stem(string) (string, error)
}

//stem a word
func Stem(origWord string) string {
	if len(origWord) > 2 {
		return step5b(step5a(step4(step3(step2(step1(strings.TrimSpace(origWord)))))))
	}

	return origWord
}

//Step 1 deals with plurals and past participles.
func step1(iw string) (sw string) {
	sw = step1a(iw)
	sw = step1b(sw)
	sw = step1c(sw)
	return
}

func step1a(iw string) string {
	wLen := len(iw)
	if strings.HasSuffix(iw, "sses") || strings.HasSuffix(iw, "ies") {
		return iw[:wLen-2]
	} else if strings.HasSuffix(iw, "ss") {
		return iw
	} else if strings.HasSuffix(iw, "s") {
		return iw[:wLen-1]
	}

	return iw
}

func step1b(iw string) string {
	wLen := len(iw)
	if strings.HasSuffix(iw, "eed") {
		if measure(iw[:wLen-3]) > 0 {
			return iw[:wLen-1]
		}
	} else if strings.HasSuffix(iw, "ed") {
		if hasVowel(iw[:wLen-2]) {
			return step1bInter(iw[:wLen-2])
		}
	} else if strings.HasSuffix(iw, "ing") {
		if hasVowel(iw[:wLen-3]) {
			return step1bInter(iw[:wLen-3])
		}
	}

	return iw
}

func step1bInter(iw string) string {
	wLen := len(iw)
	if strings.HasSuffix(iw, "at") || strings.HasSuffix(iw, "bl") ||
		strings.HasSuffix(iw, "iz") {
		return iw + "e"
	} else if astrD(iw) {
		if iw[wLen-1] != 'l' && iw[wLen-1] != 's' && iw[wLen-1] != 'z' {
			return iw[:wLen-1]
		}
	} else if astrO(iw) {
		if measure(iw) == 1 {
			return iw + "e"
		}
	}
	return iw
}

func step1c(iw string) string {
	wLen := len(iw)
	if strings.HasSuffix(iw, "y") && hasVowel(iw[:wLen-1]) {
		return iw[:wLen-1] + "i"
	}
	return iw
}

func step2(iw string) string {
	wLen := len(iw)
	if strings.HasSuffix(iw, "ational") {
		if measure(iw[:wLen-7]) > 0 {
			return iw[:wLen-7] + "ate"
		}
	} else if strings.HasSuffix(iw, "tional") {
		if measure(iw[:wLen-6]) > 0 {
			return iw[:wLen-2]
		}
	} else if strings.HasSuffix(iw, "enci") || strings.HasSuffix(iw, "anci") {
		if measure(iw[:wLen-4]) > 0 {
			return iw[:wLen-1] + "e"
		}
	} else if strings.HasSuffix(iw, "izer") {
		if measure(iw[:wLen-4]) > 0 {
			return iw[:wLen-4] + "ize"
		}
	} else if strings.HasSuffix(iw, "abli") {
		if measure(iw[:wLen-4]) > 0 {
			return iw[:wLen-4] + "able"
		}
	} else if strings.HasSuffix(iw, "alli") {
		if measure(iw[:wLen-4]) > 0 {
			return iw[:wLen-4] + "al"
		}
	} else if strings.HasSuffix(iw, "entli") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-5] + "ent"
		}
	} else if strings.HasSuffix(iw, "eli") {
		if measure(iw[:wLen-3]) > 0 {
			return iw[:wLen-3] + "e"
		}
	} else if strings.HasSuffix(iw, "ousli") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-5] + "ous"
		}
	} else if strings.HasSuffix(iw, "ization") {
		if measure(iw[:wLen-7]) > 0 {
			return iw[:wLen-7] + "ize"
		}
	} else if strings.HasSuffix(iw, "ation") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-5] + "ate"
		}
	} else if strings.HasSuffix(iw, "ator") {
		if measure(iw[:wLen-4]) > 0 {
			return iw[:wLen-4] + "ate"
		}
	} else if strings.HasSuffix(iw, "alism") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-5] + "al"
		}
	} else if strings.HasSuffix(iw, "iveness") {
		if measure(iw[:wLen-7]) > 0 {
			return iw[:wLen-7] + "ive"
		}
	} else if strings.HasSuffix(iw, "fulness") {
		if measure(iw[:wLen-7]) > 0 {
			return iw[:wLen-7] + "ful"
		}
	} else if strings.HasSuffix(iw, "ousness") {
		if measure(iw[:wLen-7]) > 0 {
			return iw[:wLen-7] + "ous"
		}
	} else if strings.HasSuffix(iw, "aliti") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-5] + "al"
		}
	} else if strings.HasSuffix(iw, "iviti") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-5] + "ive"
		}
	} else if strings.HasSuffix(iw, "biliti") {
		if measure(iw[:wLen-6]) > 0 {
			return iw[:wLen-6] + "ble"
		}

	}

	return iw
}

func step3(iw string) string {
	wLen := len(iw)
	if strings.HasSuffix(iw, "icate") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "ative") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-5]
		}
	} else if strings.HasSuffix(iw, "alize") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "iciti") {
		if measure(iw[:wLen-5]) > 0 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "ical") {
		if measure(iw[:wLen-4]) > 0 {
			return iw[:wLen-2]
		}
	} else if strings.HasSuffix(iw, "ful") {
		if measure(iw[:wLen-3]) > 0 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "ness") {
		if measure(iw[:wLen-4]) > 0 {
			return iw[:wLen-4]
		}
	}
	return iw
}

func step4(iw string) string {
	wLen := len(iw)
	if strings.HasSuffix(iw, "al") {
		if measure(iw[:wLen-2]) > 1 {
			return iw[:wLen-2]
		}
	} else if strings.HasSuffix(iw, "ance") {
		if measure(iw[:wLen-4]) > 1 {
			return iw[:wLen-4]
		}
	} else if strings.HasSuffix(iw, "ence") {
		if measure(iw[:wLen-4]) > 1 {
			return iw[:wLen-4]
		}
	} else if strings.HasSuffix(iw, "er") {
		if measure(iw[:wLen-2]) > 1 {
			return iw[:wLen-2]
		}
	} else if strings.HasSuffix(iw, "ic") {
		if measure(iw[:wLen-2]) > 1 {
			return iw[:wLen-2]
		}
	} else if strings.HasSuffix(iw, "able") {
		if measure(iw[:wLen-4]) > 1 {
			return iw[:wLen-4]
		}
	} else if strings.HasSuffix(iw, "ible") {
		if measure(iw[:wLen-4]) > 1 {
			return iw[:wLen-4]
		}
	} else if strings.HasSuffix(iw, "ant") {
		if measure(iw[:wLen-3]) > 1 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "ement") {
		if measure(iw[:wLen-5]) > 1 {
			return iw[:wLen-5]
		}
	} else if strings.HasSuffix(iw, "ment") {
		if measure(iw[:wLen-4]) > 1 {
			return iw[:wLen-4]
		}
	} else if strings.HasSuffix(iw, "ent") {
		if measure(iw[:wLen-3]) > 1 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "ion") {
		if measure(iw[:wLen-3]) > 1 {
			if wLen > 4 && (iw[wLen-4] == 's' || iw[wLen-4] == 't') {
				return iw[:wLen-3]
			}
		}
	} else if strings.HasSuffix(iw, "ou") {
		if measure(iw[:wLen-2]) > 1 {
			return iw[:wLen-2]
		}
	} else if strings.HasSuffix(iw, "ism") {
		if measure(iw[:wLen-3]) > 1 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "ate") {
		if measure(iw[:wLen-3]) > 1 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "iti") {
		if measure(iw[:wLen-3]) > 1 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "ous") {
		if measure(iw[:wLen-3]) > 1 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "ive") {
		if measure(iw[:wLen-3]) > 1 {
			return iw[:wLen-3]
		}
	} else if strings.HasSuffix(iw, "ize") {
		if measure(iw[:wLen-3]) > 1 {
			return iw[:wLen-3]
		}
	}
	return iw
}

func step5a(iw string) string {
	wLen := len(iw)
	if strings.HasSuffix(iw, "e") && measure(iw[:wLen-1]) > 1 {
		return iw[:wLen-1]
	} else if strings.HasSuffix(iw, "e") && measure(iw[:wLen-1]) == 1 && !astrO(iw[:wLen-1]) {
		return iw[:wLen-1]
	}
	return iw
}

func step5b(iw string) string {
	wLen := len(iw)
	if measure(iw) > 1 && isConsonant(iw, wLen-1) && isConsonant(iw, wLen-2) && iw[wLen-1] == 'l' {
		return iw[:wLen-1]
	}
	return iw
}

//the stem ends cvc, where the second c is not W, X or Y
func astrO(iw string) bool {
	wLen := len(iw) - 1
	if wLen >= 2 && isConsonant(iw, wLen-2) && !isConsonant(iw, wLen-1) && isConsonant(iw, wLen) {
		return iw[wLen] != 'w' && iw[wLen] != 'x' && iw[wLen] != 'y'
	}
	return false
}

//double consonants
func astrD(iw string) bool {
	wLen := len(iw)
	return iw[wLen-1] == iw[wLen-2] && isConsonant(iw, wLen-1)
}

//check if character at index i is a consonant
func isConsonant(w string, i int) bool {
	wLen := len(w)
	if wLen < i || i < 0 {
		return false
	}
	switch w[i] {
	case 'a', 'e', 'i', 'o', 'u':
		return false
	case 'y':
		if i == 0 {
			return true
		} else {
			return i > 0 && !isConsonant(w, i-1)
		}
	default:
		return true
	}
}

//measures the number of consonant sequences
func measure(w string) (val int64) {
	wLen := len(w)

	if wLen <= 0 {
		return
	}

	ptr := 0
	//ignore consonant at start
	for isConsonant(w, ptr) {
		ptr++
		if ptr >= wLen {
			return val
		}
	}

	incVal := false
	//count Vowel-Consonant pair
	for i := ptr; i < wLen; i++ {
		for i < wLen && !isConsonant(w, i) {
			i++
		}
		for i < wLen && isConsonant(w, i) {
			i++
			incVal = true
		}

		if incVal {
			val++
			incVal = false
		}
	}

	return
}

//checks if stem contains a vowel
func hasVowel(str string) bool {
	for i := 0; i < len(str); i++ {
		if !isConsonant(str, i) {
			return true
		}
	}
	return false
}
