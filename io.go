package mtx

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToJSON(what any) string {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	enc.Encode(what)
	return buf.String()
}

func ToSource(what SourceWriteable) string {
	buf := new(bytes.Buffer)
	what.SourceOn(buf)
	return buf.String()
}

type SourceWriteable interface {
	SourceOn(io.Writer)
}
