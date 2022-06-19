package core

import (
	"encoding/json"
	"os"
)

func JSONOut(what any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(what)
}
