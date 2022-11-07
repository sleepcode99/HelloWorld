package test

import (
	"main/arimatika"
	"testing"
)

func TestMin(t *testing.T) {
  result := arimatika.Pengurangan(4,5)
  if result != 1 {
    t.Fail()
  }
}