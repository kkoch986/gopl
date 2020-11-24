package raw

import (
	"encoding/json"
	"io"

	"github.com/kkoch986/gopl/ast"
)

// Writes statement lists to a raw format which can be incrementally indexed when the file is loaded into another script.

func Serialize(sl []ast.Statement, w io.Writer) error {
	for _, v := range sl {
		json, err := json.Marshal(v)
		if err != nil {
			return err
		}
		_, err = w.Write(json)
		if err != nil {
			return err
		}
	}
	return nil
}
