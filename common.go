package storable

import (
	"strings"
)

type Field struct {
	bson string
	typ  string
}

// NewField return a new Field instance.
func NewField(bson, typ string) Field {
	return Field{bson, typ}
}

// Type returns the type of this Field.
func (f Field) Type() string {
	return f.typ
}

// String returns a string with a valid representation to be used on mgo.
func (f Field) String() string {
	return f.bson
}

type Map struct {
	bson string
	typ  string
}

// NewMap return a new Map instance.
func NewMap(bson, typ string) Map {
	return Map{bson, typ}
}

// Type returns the type used by this Map.
func (f Map) Type() string {
	return f.typ
}

// Key returns a Field for the specific map key.
func (f Map) Key(key string) Field {
	bson := strings.Replace(f.bson, "[map]", key, -1)
	return NewField(bson, f.typ)
}

var (
	IdField = NewField("_id", "bson.ObjectId")
)

type Dir int

const (
	Asc  Dir = 1
	Desc Dir = -1
)

type Sort []FieldSort

type FieldSort struct {
	F Field
	D Dir
}

// String returns a representation of Sort compatible with the format of mgo.
func (s Sort) String() string {
	var fields []string
	for _, fs := range s {
		f := ""
		if fs.D == Desc {
			f += "-"
		}

		f += fs.F.String()

		fields = append(fields, f)
	}

	return strings.Join(fields, ",")
}

// IsEmpty returns if this sort map is empty or not
func (s Sort) IsEmpty() bool {
	return len(s) == 0
}
