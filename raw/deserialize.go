package raw

import (
	"encoding/json"
	"io"

	"github.com/kkoch986/gopl/ast"
)

type rawStatement struct {
	S ast.Statement
}

func (rs *rawStatement) UnmarshalJSON(b []byte) error {
	statement, err := ast.UnmarshalJSONTerm(b)
	// TODO: check that the terms type is rule|fact|query
	if err != nil {
		return err
	}
	rs.S = statement
	return nil
}

func Deserialize(r io.Reader, out chan<- ast.Statement) {
	decoder := json.NewDecoder(r)
	for {
		var s rawStatement
		err := decoder.Decode(&s)
		if err == io.EOF {
			break
		} else if err != nil {
			// TODO: handle other kinds of errors
			panic(err)
		}
		out <- s.S
	}
	close(out)
}
