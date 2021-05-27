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

	words := regexp.MustCompile(`\w+`).FindAllString(sentence, -1) //split the sentence into a slice of words
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
#### BenchmarkReverseSentence(b *testing.B)
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


#### tests - test cases
```
type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want string
}{
	{
		name: "testing single word",
		args: args{
			s: "hello",
		},
		want: "hello",
	},
	{
		name: "testing empty",
		args: args{
			s: "",
		},
		want: "",
	}, {
		name: "testing hello world",
		args: args{
			s: "hello world",
		},
		want: "world hello",
	},
	{
		name: "testing 10 words sentence",
		args: args{
			s: "today is the first day of the rest of my life",
		},
		want: "life my of rest the of day first the is today",
	}, {
		name: "testing 13 words sentence",
		args: args{
			s: "Below is the corresponding mapping UML diagram with the example given above",
		},
		want: "above given example the with diagram UML mapping corresponding the is Below",
	},
	{
		name: "testing numbers",
		args: args{
			s: "1 2 3 4 5 6 7 8 9 0",
		},
		want: "0 9 8 7 6 5 4 3 2 1",
	},
}
```
