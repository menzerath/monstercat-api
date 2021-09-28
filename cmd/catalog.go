package cmd

import (
	"fmt"
	"os"

	"github.com/menzerath/monstercat-api/v2/monstercat"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	catalogSearch string
	catalogType   string
	catalogLimit  int
	catalogOffset int
)

func init() {
	catalogCmd.Flags().StringVarP(&catalogSearch, "search", "s", "", "search query")
	catalogCmd.Flags().StringVarP(&catalogType, "type", "t", "", "type of release")
	catalogCmd.Flags().IntVar(&catalogLimit, "limit", 10, "limit number of catalog")
	catalogCmd.Flags().IntVar(&catalogOffset, "offset", 0, "offset number of catalog")

	rootCmd.AddCommand(catalogCmd)
}

var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "catalog returns a set of the most recent Monstercat catalog items",
	Run: func(cmd *cobra.Command, args []string) {
		catalog, err := monstercat.NewClient().Catalog(catalogSearch, catalogType, catalogLimit, catalogOffset)
		if err != nil {
			fmt.Printf("error fetching catalog: %s", err)
			os.Exit(1)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Title", "Artist", "Album", "Catalog ID", "Type", "Release Date"})
		table.SetCaption(true, fmt.Sprintf("%d of %d results", len(catalog.Data), catalog.Total))

		for _, item := range catalog.Data {
			title := item.Title
			if item.Version != "" {
				title = fmt.Sprintf("%s (%s)", item.Title, item.Version)
			}

			table.Append([]string{title, item.ArtistsTitle, item.Release.Title, item.Release.CatalogID, string(item.Release.Type), item.DebutDate.Format("2006-01-02")})
		}
		table.Render()
	},
}
