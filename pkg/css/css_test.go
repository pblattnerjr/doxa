package css

import (
	"os"
	"testing"
)

func TestWriteCss(t *testing.T) {
	err := WriteCss(os.Getenv("path"), os.Getenv("filename"))
	if err != nil {
		t.Error("error:", err.Error())
	}
}
