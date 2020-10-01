package auth 

import (
	"testing"
)

func TestHello(t *testing.T) {
    want := "Gopher authorized!!!"
    if got := GetAuth(); got != want {
        t.Errorf("GetAuth() = %q, want %q", got, want)
    }
}


