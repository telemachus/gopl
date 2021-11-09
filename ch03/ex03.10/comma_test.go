package ch03_test

import (
	"gopl/ch03"
	"testing"
)

func TestComma(t *testing.T) {
	tests := map[string]struct {
		num string
		want string
	}{
		"one-digit" : { num: "1", want: "1" },
		"two-digit" : { num: "12", want: "12" },
		"three-digit" : { num: "123", want: "123" },
		"four-digit" : { num: "1234", want: "1,234" },
		"five-digit" : { num: "12345", want: "12,345" },
		"six-digit" : { num: "123456", want: "123,456" },
		"seven-digit" : { num: "1234567", want: "1,234,567" },
		"eight-digit" : { num: "12345678", want: "12,345,678" },
		"nine-digit" : { num: "123456789", want: "123,456,789" },
		"ten-digit" : { num: "1234567890", want: "1,234,567,890" },
		"eleven-digit" : { num: "12345678901", want: "12,345,678,901" },
		"twelve-digit" : { num: "123456789012", want: "123,456,789,012" },
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ch03.Comma(tc.num)
			if got != tc.want {
				t.Fatalf("ch03.Comma: expected %q but got %q\n", tc.want, got)
			}
		})
	}
}
