package security

import "testing"

func Test_hashInternal(t *testing.T) {
	method := "pbkdf2:sha256"
	want := "ae409ce6f5eaac4f71e42819b480ebb0547d678d2859622f3b1ede4bc67a2a23"
	actualMethodWant := "pbkdf2:sha256:50000"
	h, m := hashInternal(method, "salt", "password")
	if h != want {
		t.Errorf("hashInternal() = %v, want %v", h, want)
	}
	if m != actualMethodWant {
		t.Errorf("hashInternal() method = %v, want %v", m, actualMethodWant)
	}
}
