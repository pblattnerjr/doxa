package repos

import "testing"

func TestClone(t *testing.T) {
	repos := []string{"https://github.com/liturgiko/testrepo1.git", "https://github.com/liturgiko/testrepo2.git"}
	err := CloneConcurrently("/Users/mac002/canBeRemoved/doxago/git", repos)
	if err != nil {
		t.Error("error:", err.Error())
	}
}
