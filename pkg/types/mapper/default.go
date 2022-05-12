package mapper

import (
	"github.com/iwilltry42/norman/v3/pkg/data"
	"github.com/iwilltry42/norman/v3/pkg/types"
)

type DefaultMapper struct {
	Field string
}

func (d DefaultMapper) FromInternal(data data.Object) {
}

func (d DefaultMapper) ToInternal(data data.Object) error {
	return nil
}

func (d DefaultMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	if d.Field != "" {
		return ValidateField(d.Field, schema)
	}
	return nil
}
