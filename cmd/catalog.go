package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/menzerath/monstercat-api/v2/monstercat"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	catalogSearch            string
	catalogBrands            []string
	catalogGenres            []string
	catalogTypes             []string
	catalogTags              []string
	catalogIncludeGold       bool
	catalogIncludeUnreleased bool
	catalogSort              string
	catalogLimit             int
	catalogOffset            int
)

func init() {
	catalogCmd.Flags().StringVarP(&catalogSearch, "search", "s", "", "search query (title, album, artist, ...)")
	catalogCmd.Flags().StringArrayVarP(&catalogBrands, "brand", "b", []string{}, "brand identifier (1 = Uncaged, 2 = Instinct, 3 = CotW, 4 = Silk, 5 = Silk Showcase)")
	catalogCmd.Flags().StringArrayVarP(&catalogGenres, "genre", "g", []string{}, "genre (dubstep, acoustic, ...)")
	catalogCmd.Flags().StringArrayVarP(&catalogTypes, "type", "t", []string{}, "release type (Single, EP, Album)")
	catalogCmd.Flags().StringArrayVar(&catalogTags, "tag", []string{}, "tags (chill, badass, ...)")
	catalogCmd.Flags().BoolVar(&catalogIncludeGold, "gold", true, "include gold releases")
	catalogCmd.Flags().BoolVar(&catalogIncludeUnreleased, "unreleased", true, "include unreleased releases")
	catalogCmd.Flags().StringVar(&catalogSort, "sort", "-date", "sort by")
	catalogCmd.Flags().IntVar(&catalogLimit, "limit", 10, "limit number of catalog")
	catalogCmd.Flags().IntVar(&catalogOffset, "offset", 0, "offset number of catalog")

	rootCmd.AddCommand(catalogCmd)
}

var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "catalog returns a set of the most recent Monstercat catalog items",
	Run: func(cmd *cobra.Command, args []string) {
		options := make([]monstercat.BrowseOption, 0)
		if catalogSearch != "" {
			options = append(options, monstercat.WithSearch(catalogSearch))
		}
		for _, brand := range catalogBrands {
			brandID, err := strconv.Atoi(brand)
			if err != nil {
				fmt.Printf("%s is not a valid brand id\n", brand)
				os.Exit(1)
			}
			options = append(options, monstercat.WithBrand(brandID))
		}
		for _, genre := range catalogGenres {
			options = append(options, monstercat.WithGenre(genre))
		}
		for _, releaseType := range catalogTypes {
			options = append(options, monstercat.WithReleaseType(releaseType))
		}
		for _, tag := range catalogTags {
			options = append(options, monstercat.WithTag(tag))
		}
		options = append(
			options,
			monstercat.IncludeGold(catalogIncludeGold),
			monstercat.IncludeUnreleased(catalogIncludeUnreleased),
			monstercat.WithSort(catalogSort),
			monstercat.WithLimit(catalogLimit),
			monstercat.WithOffset(catalogOffset),
		)

		catalog, err := monstercat.NewClient().BrowseCatalog(options...)
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
