package cmd

import (
	"fmt"
	"os"

	"github.com/menzerath/monstercat-api/monstercat"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	releasesSearch string
	releasesType   string
	releasesLimit  int
	releasesOffset int
)

func init() {
	releasesCmd.Flags().StringVarP(&releasesSearch, "search", "s", "", "search query")
	releasesCmd.Flags().StringVarP(&releasesType, "type", "t", "", "type of release")
	releasesCmd.Flags().IntVar(&releasesLimit, "limit", 10, "limit number of releases")
	releasesCmd.Flags().IntVar(&releasesOffset, "offset", 0, "offset number of releases")

	rootCmd.AddCommand(releasesCmd)
}

var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "releases returns a set of the most recent Monstercat releases",
	Run: func(cmd *cobra.Command, args []string) {
		releases, err := monstercat.NewClient().Releases(releasesSearch, releasesType, releasesLimit, releasesOffset)
		if err != nil {
			fmt.Printf("error fetching releases: %s", err)
			os.Exit(1)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Catalog ID", "Title", "Artist", "Type", "Release Date"})
		table.SetCaption(true, fmt.Sprintf("%d of %d results", len(releases.Results), releases.Total))

		for _, release := range releases.Results {
			table.Append([]string{release.CatalogID, release.Title, release.Artist, string(release.Type), release.ReleaseDate.Format("2006-01-02")})
		}
		table.Render()
	},
}
