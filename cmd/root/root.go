package root

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	api "github.com/gary23w/metasearch_api/api"
	enginesearch "github.com/gary23w/metasearch_api/internal/engine"
	_ "github.com/gary23w/metasearch_api/internal/providers/all"
	"github.com/gary23w/metasearch_api/internal/search"
)

var system = &cobra.Command{
	Use:   "metasearch_api",
	Short: "runs a metasearch rest api engine",
}

func init() {
	cmdQuery := &cobra.Command{
		Use:     "query [search query]",
		Short:   "search for a query",
		Aliases: []string{"qu", "q"},
		RunE: func(cmd *cobra.Command, args []string) error {
			qu := strings.Join(args, " ")
			ctx := context.Background()
			s, err := enginesearch.NewEngine(ctx)
			if err != nil {
				return err
			}
			limit, _ := cmd.Flags().GetInt("limit")

			it := s.Search(ctx, search.Request{
				Query: qu,
			})
			defer it.Close()
			for i := 0; i < limit && it.Next(ctx); i++ {
				r := it.Result()
				fmt.Printf("%s - %q (%T)\n\n", r.GetURL(), r.GetTitle(), r)
			}
			if err := it.Err(); err != nil {
				return err
			}
			tok := it.Token()
			fmt.Println("\n\ntoken:", hex.EncodeToString([]byte(tok)))
			if err := it.Err(); err != nil {
				return err
			}
			return nil
		},
	}
	cmdQuery.Flags().IntP("limit", "n", 10, "limit the number of results")
	system.AddCommand(cmdQuery)

	cmdAutoc := &cobra.Command{
		Use:     "complete [search query]",
		Short:   "autocomplete for a query",
		Aliases: []string{"ac"},
		RunE: func(cmd *cobra.Command, args []string) error {
			qu := strings.Join(args, " ")
			ctx := context.Background()
			s, err := enginesearch.NewEngine(ctx)
			if err != nil {
				return err
			}
			list, err := s.AutoComplete(ctx, qu)
			if err != nil {
				return err
			}
			for _, r := range list {
				fmt.Println(r)
			}
			return nil
		},
	}
	system.AddCommand(cmdAutoc)

	cmdAPI := &cobra.Command{
		Use:     "api start",
		Short:   "start the api",
		Aliases: []string{"api"},
		RunE: func(cmd *cobra.Command, args []string) error {
			port := "8080"
			//cloud provider support
			if os.Getenv("PORT") != "" {
				port = os.Getenv("PORT")
			}
			api.Start(port)
			return nil
		},
	}
	system.AddCommand(cmdAPI)
}

func Execute() {
	if err := system.Execute(); err != nil {
		log.Fatal(err)
	}
}
