package config

import "github.com/tracker-tv/tmdb-ids-producer/internal/models"

type Config struct {
	BaseURL   string      `env:"TTV_BASE_URL,required" envDefault:"https://files.tmdb.org/p/exports"`
	Type      models.Type `env:"TTV_TYPE,required"`
	OutputDir string      `env:"TTV_OUTPUT_DIR,required" envDefault:"./downloads"`
}
