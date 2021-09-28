package monstercat

import (
	"time"
)

// CatalogItem represents a single release from Monstercat API
type CatalogItem struct {
	ID           string    `json:"Id"`
	ISRC         string    `json:"ISRC"`
	Title        string    `json:"Title"`
	Artists      []Artist  `json:"Artists"`
	ArtistsTitle string    `json:"ArtistsTitle"`
	DebutDate    time.Time `json:"DebutDate"`
	Release      Release   `json:"Release"`
	Tags         []string  `json:"Tags"`
	Version      string    `json:"Version"`

	Brand   string `json:"Brand"`
	BrandID int    `json:"BrandId"`

	BPM            float32 `json:"BPM"`
	Duration       int     `json:"Duration"`
	Explicit       bool    `json:"Explicit"`
	GenrePrimary   string  `json:"GenrePrimary"`
	GenreSecondary string  `json:"GenreSecondary"`

	CreatorFriendly bool `json:"CreatorFriendly"`
	Downloadable    bool `json:"Downloadable"`
	InEarlyAccess   bool `json:"InEarlyAccess"`
	Streamable      bool `json:"Streamable"`
}

// Artist describes a single artist of a CatalogItem
type Artist struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
	Role string `json:"Role"`
	URI  string `json:"URI"`
}

// Release describes which release a CatalogItem belongs to
type Release struct {
	ArtistsTitle string      `json:"ArtistsTitle"`
	CatalogID    string      `json:"CatalogId"`
	Description  string      `json:"Description"`
	ID           string      `json:"Id"`
	ReleaseDate  time.Time   `json:"ReleaseDate"`
	Title        string      `json:"Title"`
	Type         ReleaseType `json:"Type"`
	UPC          string      `json:"UPC"`
	Version      string      `json:"Version"`
}

// ReleaseType describes what kind of release we are looking at
type ReleaseType string

// define all known release types
const (
	ReleaseTypeAlbum       ReleaseType = "Album"
	ReleaseTypeCompilation ReleaseType = "Compilation"
	ReleaseTypeEP          ReleaseType = "EP"
	ReleaseTypeMixes       ReleaseType = "Mixes"
	ReleaseTypePodcast     ReleaseType = "Podcast"
	ReleaseTypeSingle      ReleaseType = "Single"
)
