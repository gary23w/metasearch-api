package providers

import (
	"context"

	"github.com/gary23w/metasearch_api/internal/models"
)

type Provider = models.Provider

type ProviderFunc func(ctx context.Context) (Provider, error)

var registry = make(map[string]ProviderFunc)

func Register(name string, fnc ProviderFunc) {
	registry[name] = fnc
}

func List() []ProviderFunc {
	var arr []ProviderFunc
	for _, fnc := range registry {
		arr = append(arr, fnc)
	}
	return arr
}
