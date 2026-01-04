package main

import (
	"fmt"

	"github.com/tracker-tv/tmdb-ids-producer/internal/config"
	"github.com/tracker-tv/tmdb-ids-producer/internal/downloader"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	fmt.Printf("Config: %+v\n", cfg)

	d := downloader.New(nil, cfg)

	filename, err := d.Download()
	if err != nil {
		panic(fmt.Sprintf("failed to download file: %v", err))
	}

	fmt.Printf("Downloaded file: %s\n", filename)
}
