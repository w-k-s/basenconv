package basenconv

import (
  "testing"
)

func TestParseBase62(t *testing.T) {
	assert(t,ParseBase62("0"),uint64(0))
	assert(t,ParseBase62("z"),uint64(61))
    assert(t,ParseBase62("5Ca"),uint64(20000))
}

func TestParseHex(t *testing.T){
	assert(t,ParseHex("0"),uint64(0))
	assert(t,ParseHex("A"),uint64(10))
    assert(t,ParseHex("F"),uint64(15))
}

func TestParseBinary(t *testing.T){
	assert(t,ParseBinary("0"),uint64(0))
	assert(t,ParseBinary("11"),uint64(3))
    assert(t,ParseBinary("1000001"),uint64(65))
}
