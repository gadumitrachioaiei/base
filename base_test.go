package base

import (
	"encoding/base32"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	result, err := NewDecoder().Decode("ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%q %d\n", result, len(result))
}

var result []byte

func BenchmarkDecode(b *testing.B) {
	d := NewDecoder()
	var r []byte
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, err = d.Decode("ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY")
		if err != nil {
			b.Fatal(err)
		}
	}
	result = r
}

func BenchmarkStdDecode(b *testing.B) {
	var r []byte
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, err = base32.StdEncoding.DecodeString("ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY=")
		if err != nil {
			b.Fatal(err)
		}
	}
	result = r
}
