package draw

import (
	"fmt"
	"io"
)

func PrintMsg(w io.Writer, msg string) {
	fmt.Fprintf(w, msg + "\n")
}
