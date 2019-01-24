package crypto

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHashString(t *testing.T) {
	s := "life is always this hard"
	hash, err := HashString(s)
	if err != nil {
		t.Error("generator error", err)
	}
	fmt.Println(hash)
}

func TestMatchHashString(t *testing.T) {
	s := "life is always this hard"
	hash, _ := HashString(s)
	b := MatchHashString(s, hash)
	fmt.Println(hash)
	if !reflect.DeepEqual(b, true) {
		t.Error("not match")
	}
}
