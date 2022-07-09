package autocomplete

import (
	"context"

	"github.com/gary23w/metasearch_api/internal/models"
)

type Service interface {
	models.Provider
	AutoComplete(ctx context.Context, text string) ([]string, error)
}
