# interview_golang_v
A simple pre-interview question


## file structure  
- reverse_sentence.go  
- reverse_sentence_test.go
- Makefile

## reverse_sentence.go     

This file contains the ReverseSentence function; this function reverses the order of words in a given sentence, clause, phrase or any alignment of words seperated by spaces.   
  
It does this by splitting the sentence into words using an empty space " " to mark a complete word.   

The array of words are concatenated in reverse order with space between each words to form a sentence.  

The newly formed sentenced is trimmed of leading and trailing spaces.
```

//ReverseSentence reverses the words in a given sentence
func ReverseSentence(sentence string) (reversedSentence string) {

	if sentence == "" { //return sentence if empty
		return sentence
	}

	words := strings.Split(sentence, " ") //split the sentence into a slice of words
	for _, word := range words {
		reversedSentence = word + " " + reversedSentence
	}

	return strings.TrimSpace(reversedSentence)
}
```

## reverse_sentence_test.go

#### TestReverseSentence(t *testing.T)   

Table driven test was employed here. <b>tests</b> is an array of test cases; for each test case, a "test go routine" is spun.


```
func TestReverseSentence(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revsentence.ReverseSentence(tt.args.s); got != tt.want {
				t.Errorf("ReverseSentence() = %v, want %v", got, tt.want)
			}
		})
	}

}

```

```
func BenchmarkReverseSentence(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				revsentence.ReverseSentence(tt.args.s)
			}
		})
	}
}
```
