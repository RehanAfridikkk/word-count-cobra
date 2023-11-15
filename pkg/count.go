package pkg


type CountsResult struct {
	LineCount        int
	WordsCount       int
	VowelsCount      int
	PunctuationCount int
}

func Counts(chunk []byte, results chan<- CountsResult) {
    lineCount := 0
    wordsCount := 0
    vowelsCount := 0
    punctuationCount := 0

    inWord := false

    for _, b := range chunk {
        switch {
        case b == '\n':
            lineCount++
        case b == ' ' || b == '\t':
            if inWord {
                wordsCount++
                inWord = false
            }
        default:
            inWord = true

            if isVowel(b) {
                vowelsCount++
            }

            if isPunctuation(b) {
                punctuationCount++
            }
        }
    }

    if inWord {
        wordsCount++
    }

    results <- CountsResult{LineCount: lineCount, WordsCount: wordsCount, VowelsCount: vowelsCount, PunctuationCount: punctuationCount}
}


func isVowel(b byte) bool {
	vowels := "AEIOUaeiou"
	return byteInSlice(b, []byte(vowels))
}

func isPunctuation(b byte) bool {
	punctuation := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	return byteInSlice(b, []byte(punctuation))
}

func byteInSlice(b byte, slice []byte) bool {
	for _, value := range slice {
		if b == value {
			return true
		}
	}
	return false
}
