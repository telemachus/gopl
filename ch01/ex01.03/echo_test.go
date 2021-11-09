package echo_test

import (
	"echo"
	"testing"
)

var args = make([]string, 0, 3000)
var ss = []string{"fizz", "buzz", "fizzbuzz"}

func init() {
	for i := 0; i < 1000; i++ {
		for j := 0; j < len(ss); j++ {
			args = append(args, ss[j])
		}
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo.Concat(args[:])
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo.Join(args[:])
	}
}
