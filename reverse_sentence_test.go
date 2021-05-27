package revsentence_test

import (
	"testing"

	revsentence "github.com/IamNator/interview/reverse_sentence"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want string
}{
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

func TestReverseSentence(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revsentence.ReverseSentence(tt.args.s); got != tt.want {
				t.Errorf("ReverseSentence() = %v, want %v", got, tt.want)
			}
		})
	}

}

func BenchmarkReverseSentence(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				revsentence.ReverseSentence(tt.args.s)
			}
		})
	}
}
