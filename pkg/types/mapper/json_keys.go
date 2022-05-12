package mapper

import (
	"github.com/iwilltry42/norman/v3/pkg/data"
	"github.com/iwilltry42/norman/v3/pkg/types"
	"github.com/iwilltry42/norman/v3/pkg/types/convert"
)

type JSONKeys struct {
}

func (d JSONKeys) FromInternal(data data.Object) {
}

func (d JSONKeys) ToInternal(data data.Object) error {
	for key, value := range data {
		newKey := convert.ToJSONKey(key)
		if newKey != key {
			data[newKey] = value
			delete(data, key)
		}
	}
	return nil
}

func (d JSONKeys) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
