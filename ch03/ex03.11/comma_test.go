package ch03_test

import (
	"gopl/ch03"
	"testing"
)

func TestCommaSimpleNumber(t *testing.T) {
	tests := map[string]struct {
		num string
		want string
	}{
		"empty string": { num: "", want: "" },
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
			got := ch03.Commify(tc.num)
			if got != tc.want {
				t.Fatalf("ch03.Comma: expected %q but got %q\n", tc.want, got)
			}
		})
	}
}

func TestCommaSignedNumber(t *testing.T) {
	tests := map[string]struct {
		num string
		want string
	}{
		"empty string": { num: "", want: "" },
		"four-digit positive number" : { num: "+1234", want: "+1,234" },
		"four-digit negative number" : { num: "-1234", want: "-1,234" },
		"five-digit postive number" : { num: "+12345", want: "+12,345" },
		"five-digit negative number" : { num: "-12345", want: "-12,345" },
		"six-digit positive number" : { num: "+123456", want: "+123,456" },
		"six-digit negative number" : { num: "-123456", want: "-123,456" },
		"seven-digit positive number" : { num: "+1234567", want: "+1,234,567" },
		"seven-digit negative number" : { num: "-1234567", want: "-1,234,567" },
		"eight-digit positive number" : { num: "+12345678", want: "+12,345,678" },
		"eight-digit negative number" : { num: "-12345678", want: "-12,345,678" },
		"nine-digit positive number" : { num: "+123456789", want: "+123,456,789" },
		"nine-digit negative number" : { num: "-123456789", want: "-123,456,789" },
		"ten-digit positive number" : { num: "+1234567890", want: "+1,234,567,890" },
		"ten-digit negative number" : { num: "-1234567890", want: "-1,234,567,890" },
		"eleven-digit positive number" : { num: "+12345678901", want: "+12,345,678,901" },
		"eleven-digit negative number" : { num: "-12345678901", want: "-12,345,678,901" },
		"twelve-digit positive number" : { num: "+123456789012", want: "+123,456,789,012" },
		"twelve-digit negative number" : { num: "-123456789012", want: "-123,456,789,012" },
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ch03.Commify(tc.num)
			if got != tc.want {
				t.Fatalf("ch03.Comma: expected %q but got %q\n", tc.want, got)
			}
		})
	}
}

func TestCommaFractionalPartNumber(t *testing.T) {
	tests := map[string]struct {
		num string
		want string
	}{
		"empty string": { num: "", want: "" },
		"one-digit" : { num: "1.012", want: "1.012" },
		"two-digit" : { num: "12.123", want: "12.123" },
		"three-digit" : { num: "123.456", want: "123.456" },
		"four-digit" : { num: "1234.789", want: "1,234.789" },
		"five-digit" : { num: "12345.012", want: "12,345.012" },
		"six-digit" : { num: "123456.345", want: "123,456.345" },
		"seven-digit" : { num: "1234567.678", want: "1,234,567.678" },
		"eight-digit" : { num: "12345678.9", want: "12,345,678.9" },
		"nine-digit" : { num: "123456789.0", want: "123,456,789.0" },
		"ten-digit" : { num: "1234567890.01", want: "1,234,567,890.01" },
		"eleven-digit" : { num: "12345678901.234", want: "12,345,678,901.234" },
		"twelve-digit" : { num: "123456789012.5", want: "123,456,789,012.5" },
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ch03.Commify(tc.num)
			if got != tc.want {
				t.Fatalf("ch03.Comma: expected %q but got %q\n", tc.want, got)
			}
		})
	}
}

func TestCommaFractionalPartSignedNumber(t *testing.T) {
	tests := map[string]struct {
		num string
		want string
	}{
		"empty string": { num: "", want: "" },
		"one-digit" : { num: "-1.012", want: "-1.012" },
		"two-digit" : { num: "+12.123", want: "+12.123" },
		"three-digit" : { num: "+123.456", want: "+123.456" },
		"four-digit" : { num: "-1234.789", want: "-1,234.789" },
		"five-digit" : { num: "-12345.012", want: "-12,345.012" },
		"six-digit" : { num: "+123456.345", want: "+123,456.345" },
		"seven-digit" : { num: "-1234567.678", want: "-1,234,567.678" },
		"eight-digit" : { num: "+12345678.9", want: "+12,345,678.9" },
		"nine-digit" : { num: "-123456789.0", want: "-123,456,789.0" },
		"ten-digit" : { num: "+1234567890.01", want: "+1,234,567,890.01" },
		"eleven-digit" : { num: "-12345678901.234", want: "-12,345,678,901.234" },
		"twelve-digit" : { num: "+123456789012.5", want: "+123,456,789,012.5" },
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ch03.Commify(tc.num)
			if got != tc.want {
				t.Fatalf("ch03.Comma: expected %q but got %q\n", tc.want, got)
			}
		})
	}
}
