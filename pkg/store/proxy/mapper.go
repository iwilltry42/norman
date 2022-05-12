package proxy

import (
	"github.com/iwilltry42/norman/v3/pkg/data"
	"github.com/iwilltry42/norman/v3/pkg/types"
)

type AddAPIVersionKind struct {
	APIVersion string
	Kind       string
	Next       types.Mapper
}

func (d AddAPIVersionKind) FromInternal(data data.Object) {
	if d.Next != nil {
		d.Next.FromInternal(data)
	}
}

func (d AddAPIVersionKind) ToInternal(data data.Object) error {
	if d.Next != nil {
		if err := d.Next.ToInternal(data); err != nil {
			return err
		}
	}

	data.Set("apiVersion", d.APIVersion)
	data.Set("kind", d.Kind)
	return nil
}

func (d AddAPIVersionKind) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
