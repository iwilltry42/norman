package subscribe

import (
	"net/http"

	"github.com/iwilltry42/norman/v3/pkg/types"
)

func Register(schemas *types.Schemas) {
	schemas.MustImportAndCustomize(Subscribe{}, func(schema *types.Schema) {
		schema.CollectionMethods = []string{http.MethodGet}
		schema.ResourceMethods = []string{}
		schema.ListHandler = Handler
		schema.PluralName = "subscribe"
	})
}
