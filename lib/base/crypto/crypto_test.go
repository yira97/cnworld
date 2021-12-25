package crypto

import "testing"

func TestGCM(t *testing.T) {
	key16 := "79244226452948404D635166546A576E5A7234743777217A25432A462D4A614E"
	truth := []byte("hello world")
	cip := NewGCM_encrypt(key16, truth)

	result, err := NewGCM_decrypt(key16, cip)
	if err != nil {
		t.Fatal(err)
	}

	if string(result) != string(truth) {
		t.Fatalf("不一样. Got %s, want %s", result, truth)
	}
}
