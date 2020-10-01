package test 

import (
	"testing"
	"github.com/sivaramsajeev/go-starter/dirs"
)

func TestHello(t *testing.T) {
    want := "Addon printed"
    if got := dirs.AddonPrint(); got != want {
        t.Errorf("dirs.AddonPrint() = %q, want %q", got, want)
    }
}





