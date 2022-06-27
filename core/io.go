package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func JSONOut(what any) {
	fmt.Fprintln(os.Stdout, ToJSON(what))
}

func ToJSON(what any) string {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	enc.Encode(what)
	return buf.String()
}
