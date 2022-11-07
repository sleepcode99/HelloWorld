package test

import (
	"main/arimatika"
	"testing"
)
// function testing golang
func TestAdd(t *testing.T) {

  testTable := []struct{
    a, b int
    hasil int
  }{
    {a: 1, b: 2, hasil:3},
    {a: -1, b: 3, hasil:2},
    {a: 4, b: 4, hasil:8},
    {a: 5, b: 2, hasil:7},
    {a: -5, b: 3, hasil:-2},
  }

  for _, tableDrive := range testTable {
      result := arimatika.Tambah(tableDrive.a, tableDrive.b)
      if result != tableDrive.hasil {
        t.Logf("Get result %d, but expected %d", result, tableDrive.hasil)
        t.Fail()
    }
  }
}