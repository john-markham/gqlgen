package model

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/john-markham/gqlgen/_examples/scalars/external"
	"github.com/john-markham/gqlgen/graphql"
)

type Banned bool

func (b Banned) MarshalGQL(w io.Writer) {
	if b {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func (b *Banned) UnmarshalGQL(v any) error {
	switch v := v.(type) {
	case string:
		*b = Banned(strings.EqualFold(v, "true"))
		return nil
	case bool:
		*b = Banned(v)
		return nil
	default:
		return fmt.Errorf("%T is not a bool", v)
	}
}

type User struct {
	ID       external.ObjectID
	Name     string
	Created  time.Time  // direct binding to builtin types with external Marshal/Unmarshal methods
	Modified *time.Time // direct binding to builtin types with external Marshal/Unmarshal methods
	ValPrefs Prefs      // external un/marshal that act on pointers
	PtrPrefs *Prefs
	IsBanned Banned
	Address  Address
	Tier     Tier
}

// Point is serialized as a simple array, eg [1, 2]
type Point struct {
	X int
	Y int
}

func (p *Point) UnmarshalGQL(v any) error {
	pointStr, ok := v.(string)
	if !ok {
		return errors.New("points must be strings")
	}

	parts := strings.Split(pointStr, ",")

	if len(parts) != 2 {
		return errors.New("points must have 2 parts")
	}

	var err error
	if p.X, err = strconv.Atoi(parts[0]); err != nil {
		return err
	}
	if p.Y, err = strconv.Atoi(parts[1]); err != nil {
		return err
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (p Point) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `"%d,%d"`, p.X, p.Y)
}

// if the type referenced in .gqlgen.yml is a function that returns a marshaller we can use it to encode and decode
// onto any existing go type.
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

// Unmarshal{Typename} is only required if the scalar appears as an input. The raw values have already been decoded
// from json into int/float64/bool/nil/map[string]interface/[]interface
func UnmarshalTimestamp(v any) (time.Time, error) {
	if tmpStr, ok := v.(int64); ok {
		return time.Unix(tmpStr, 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}

// Lets redefine the base ID type to use an id from an external library
func MarshalID(id external.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("=%d=", id)))
	})
}

// And the same for the unmarshaler
func UnmarshalID(v any) (external.ObjectID, error) {
	str, ok := v.(string)
	if !ok {
		return 0, errors.New("ids must be strings")
	}
	i, err := strconv.Atoi(str[1 : len(str)-1])
	return external.ObjectID(i), err
}

type SearchArgs struct {
	Location     *Point
	CreatedAfter *time.Time
	IsBanned     Banned
}

// A custom enum that uses integers to represent the values in memory but serialize as string for graphql
type Tier uint

const (
	TierA Tier = iota
	TierB Tier = iota
	TierC Tier = iota
)

func TierForStr(str string) (Tier, error) {
	switch str {
	case "A":
		return TierA, nil
	case "B":
		return TierB, nil
	case "C":
		return TierC, nil
	default:
		return 0, fmt.Errorf("%s is not a valid Tier", str)
	}
}

func (e Tier) IsValid() bool {
	switch e {
	case TierA, TierB, TierC:
		return true
	}
	return false
}

func (e Tier) String() string {
	switch e {
	case TierA:
		return "A"
	case TierB:
		return "B"
	case TierC:
		return "C"
	default:
		panic("invalid enum value")
	}
}

func (e *Tier) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("enums must be strings")
	}

	var err error
	*e, err = TierForStr(str)
	return err
}

func (e Tier) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Prefs struct {
	DarkMode bool
}

func MarshalPreferences(p *Prefs) graphql.Marshaler {
	return graphql.MarshalBoolean(p.DarkMode)
}

func UnmarshalPreferences(v any) (*Prefs, error) {
	tmp, err := graphql.UnmarshalBoolean(v)
	if err != nil {
		return nil, err
	}
	return &Prefs{DarkMode: tmp}, nil
}
