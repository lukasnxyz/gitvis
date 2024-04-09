package draw

import (
	"testing"
	"bytes"
)

func TestPrintMsg(t *testing.T) {
	t.Run("same output as input", func(t *testing.T) {
		buff := &bytes.Buffer{}
		PrintMsg(buff, "ready to go")

		got := buff.String()
		want := "ready to go"

		if got != want {
			t.Errorf("got %s, wanted %s", got, want)
		}
	})
}
