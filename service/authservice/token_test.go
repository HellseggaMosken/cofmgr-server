package authservice

import (
	"fmt"
	"testing"
)

func TestTokenValid(t *testing.T) {
	id := "idid-idid-id"
	isAdmin := true
	token := CreateToken(id, isAdmin)
	fmt.Println("Token:", token)

	_isAdmin, _id, expired, valid := ParseToken(token)
	if !valid {
		t.Error("token should be valid")
		return
	}
	if expired {
		t.Error("claims should not be expired")
		return
	}
	if isAdmin != _isAdmin {
		t.Errorf("wrong claims isAdmin: %t, expect %t", _isAdmin, isAdmin)
		return
	}
	if id != _id {
		t.Errorf("wrong claims id: %s, expect %s", _id, id)
		return
	}
}

func TestTokenInValid(t *testing.T) {
	id := "idid-idid-id"
	isAdmin := true
	token := CreateToken(id, isAdmin)
	token = token + "1"

	_, _, _, valid := ParseToken(token)
	if valid {
		t.Error("token should be not valid")
		return
	}
}
