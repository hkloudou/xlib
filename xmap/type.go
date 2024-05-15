package xmap

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Type is Result type
type Type int

const (
	// Null is a null json value
	Null Type = iota
	// False is a json false boolean
	False
	// Number is json number
	Number
	// String is a json string
	String
	// True is a json true boolean
	True
	// JSON is a raw block of JSON
	JSON
)

// Result represents a json value that is returned from Get().
type Result struct {
	raw any
	// exists bool
}

func (m Result) Exists() bool {
	return m.raw != nil
}

func (m Result) String(defs ...string) string {
	switch v := m.raw.(type) {
	case bool:
		return strings.ToLower(fmt.Sprintf("%v", v))
	case string:
		return v
	case int:
	case float64:
		return fmt.Sprintf("%v", v)
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return ""
}

func (m Result) Bool() bool {
	switch v := m.raw.(type) {
	case bool:
		return v
	case string:
		b, _ := strconv.ParseBool(strings.ToLower(v))
		return b
	case int:
	case float64:
		return v != 0
	}

	return false
}

func (m Result) Int(defs ...int64) int64 {
	switch v := m.raw.(type) {
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		n, ok := parseInt(v)
		if !ok {
			break
		}
		return n
	case float64:
		i, ok := safeInt(v)
		if ok {
			return i
		}
		i, ok = parseInt(m.String())
		if ok {
			return i
		}
		return int64(v)
	case int:
		return int64(v)
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

// Uint returns an unsigned integer representation.
func (m Result) Uint(defs ...uint64) uint64 {
	switch v := m.raw.(type) {
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		n, ok := parseUint(v)
		if !ok {
			break
		}
		return n
	case float64:
		i, ok := safeInt(v)
		if ok && i >= 0 {
			return uint64(i)
		}
		u, ok := parseUint(fmt.Sprintf("%v", v))
		if ok {
			return u
		}
		return uint64(v)
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

// Float returns an float64 representation.
func (m Result) Float(defs ...float64) float64 {
	switch v := m.raw.(type) {
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			break
		}
		return n
	case float64:
		return v
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

// Time returns a time.Time representation.
func (m Result) Time() time.Time {
	res, _ := time.Parse(time.RFC3339, m.String())
	return res
}

// safeInt validates a given JSON number
// ensures it lies within the minimum and maximum representable JSON numbers
func safeInt(f float64) (n int64, ok bool) {
	// https://tc39.es/ecma262/#sec-number.min_safe_integer
	// https://tc39.es/ecma262/#sec-number.max_safe_integer
	if f < -9007199254740991 || f > 9007199254740991 {
		return 0, false
	}
	return int64(f), true
}

func parseInt(s string) (n int64, ok bool) {
	var i int
	var sign bool
	if len(s) > 0 && s[0] == '-' {
		sign = true
		i++
	}
	if i == len(s) {
		return 0, false
	}
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int64(s[i]-'0')
		} else {
			return 0, false
		}
	}
	if sign {
		return n * -1, true
	}
	return n, true
}

func parseUint(s string) (n uint64, ok bool) {
	var i int
	if i == len(s) {
		return 0, false
	}
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + uint64(s[i]-'0')
		} else {
			return 0, false
		}
	}
	return n, true
}

// String returns a string representation of the value.
// func (t Result) String() string {
// 	switch t.Type {
// 	default:
// 		return ""
// 	case False:
// 		return "false"
// 	case Number:
// 		if len(t.Raw) == 0 {
// 			// calculated result
// 			return strconv.FormatFloat(t.Num, 'f', -1, 64)
// 		}
// 		var i int
// 		if t.Raw[0] == '-' {
// 			i++
// 		}
// 		for ; i < len(t.Raw); i++ {
// 			if t.Raw[i] < '0' || t.Raw[i] > '9' {
// 				return strconv.FormatFloat(t.Num, 'f', -1, 64)
// 			}
// 		}
// 		return t.Raw
// 	case String:
// 		return t.Str
// 	case JSON:
// 		return t.Raw
// 	case True:
// 		return "true"
// 	}
// }
