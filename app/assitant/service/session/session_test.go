package session

import "testing"

func TestEncryptUserTokenImpl(t *testing.T) {

	key16 := "79244226452948404D635166546A576E5A7234743777217A25432A462D4A614E"

	truth := UserToken{
		UserUID:   "UID",
		UserType:  "SUPER IDOT",
		IssueTime: "today",
		Apps:      []string{"app1", "app2"},
	}

	cipher := truth.EncryptImpl(key16)

	result, err := DecryptUserTokenImpl(key16, cipher)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if truth.UserUID != result.UserUID || truth.UserType != result.UserType || truth.IssueTime != result.IssueTime || len(truth.Apps) != len(result.Apps) {
		t.Fatalf("不一样. expect %v, got %v", truth, result)
	}

	for i := range truth.Apps {
		if truth.Apps[i] != result.Apps[i] {
			t.Fatalf("不一样. expect %v, got %v", truth, result)
		}
	}
}
