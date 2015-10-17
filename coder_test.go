package bitsy

import (
	"reflect"
	"testing"
)

var (
	url       = 1337
	enc       = "pjz2p"
	alphabet  = "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"
	blockSize = 24
	minLength = 5
)

// TestNew verifies that the returned instance is of the proper type.
func TestNew(t *testing.T) {
	c := New()
	if reflect.TypeOf(c).String() != "*bitsy.Coder" {
		t.Error("New returned incorrect type")
	}
}

func TestEncode(t *testing.T) {
	c := New()
	expected := enc
	if result := c.Encode(url, 5); result != expected {
		t.Fatalf("Encode: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}
}

func TestDecode(t *testing.T) {
	c := New()
	expected := url
	if result := c.Decode(enc); result != expected {
		t.Fatalf("Decode: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}
}

func Test_encode(t *testing.T) {
	c := New()
	expected := 15482880
	if result := c.encode(567); result != expected {
		t.Fatalf("Encode: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}
}

func Test_decode(t *testing.T) {
	c := New()
	expected := 6635520
	if result := c.decode(678); result != expected {
		t.Fatalf("decode: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}

}

func Test_enbase(t *testing.T) {
	tiny := New()
	expected := "mmmcw"
	if result := tiny.enbase(169, minLength); result != expected {
		t.Fatalf("Enbase: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}

}

func Test_debase(t *testing.T) {
	tiny := New()
	expected := 22154543
	if result := tiny.debase("toshd"); result != expected {
		t.Fatalf("Debase: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}

}

func TestMask(t *testing.T) {
	var in, out = blockSize, 16777215
	if x := mask(in); x != out {
		t.Errorf("getMask(%v) = %v, want %v", in, x, out)
	}
}

func TestReverseString(t *testing.T) {
	const in, out = "1337", "7331"
	if x := reverseString(in); x != out {
		t.Errorf("reverseString(%v) = %v, want %v", in, x, out)
	}
}
