package basenconv

import (
  "testing"
)

func TestFormatBase62(t *testing.T) {
	assert(t,FormatBase62(0),"0")
	assert(t,FormatBase62(61),"z")
    assert(t,FormatBase62(20000),"5Ca")
}

func TestFormatHex(t *testing.T){
	assert(t,FormatHex(0),"0")
	assert(t,FormatHex(10),"A")
    assert(t,FormatHex(15),"F")
}

func TestFormatBinary(t *testing.T){
	assert(t,FormatBinary(0),"0")
	assert(t,FormatBinary(3),"11")
    assert(t,FormatBinary(65),"1000001")
}

func assert(t *testing.T,got interface{},want interface{}){
	if got != want {
       t.Errorf("result was incorrect, got: %v, want: %v.", got, want)
    }
}