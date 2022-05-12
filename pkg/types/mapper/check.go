package mapper

import (
	"fmt"

	"github.com/iwilltry42/norman/v3/pkg/types"
)

func ValidateField(field string, schema *types.Schema) error {
	if _, ok := schema.ResourceFields[field]; !ok {
		return fmt.Errorf("field %s missing on schema %s", field, schema.ID)
	}

	return nil
}
