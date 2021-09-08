package scalar

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalTimestamp if the type referenced in .gqlgen.yml is a function that returns a marshaller we can use it to encode and decode
// onto any existing go type.
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

// UnmarshalTimestamp is only required if the scalar appears as an input. The raw values have already been decoded
// from json into int/float64/bool/nil/map[string]interface/[]interface
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case int64:
		return time.Unix(v, 0), nil
	case json.Number:
		vv, err := v.Int64()
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(vv, 0), nil
	}

	return time.Time{}, errors.New("time should be a unix timestamp")
}
