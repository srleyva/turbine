package Authentication

import (
	"strings"
	"testing"
)

func TestHashPassword(t *testing.T) {
	actualHash, err := HashPassword("test")
	if err != nil {
		t.Errorf("error returned where not expected: %+v", err)
	}

	if !strings.Contains(actualHash, "$2") {
		t.Error("Hash malformed")
	}
}

func TestChackPasswordHash(t *testing.T) {
	hash, _ := HashPassword("test")
	t.Run("check that func returns true when password is correct", func(t *testing.T) {
		if !CheckPasswordHash("test", hash) {
			t.Error("Function returns false when password is correct")
		}
	})

	t.Run("check that func returns false when password is incorrect", func(t *testing.T) {
		hash, _ := HashPassword("incorrect")
		if CheckPasswordHash("test", hash) {
			t.Error("Function returns true when password is correct")
		}
	})
}
