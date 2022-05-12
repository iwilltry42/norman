package schemaconvert

import (
	"github.com/iwilltry42/norman/v3/pkg/types"
	convert2 "github.com/iwilltry42/norman/v3/pkg/types/convert"
)

func ToInternal(from interface{}, schema *types.Schema, target interface{}) error {
	data, err := convert2.EncodeToMap(from)
	if err != nil {
		return err
	}
	if err := schema.Mapper.ToInternal(data); err != nil {
		return err
	}
	return convert2.ToObj(data, target)
}