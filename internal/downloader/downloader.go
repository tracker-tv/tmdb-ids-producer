package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/tracker-tv/tmdb-ids-producer/internal/config"
	"github.com/tracker-tv/tmdb-ids-producer/internal/models"
)

type Downloader struct {
	Client    *http.Client
	Now       time.Time
	BaseURL   string
	OutputDir string
	MediaType models.Type
}

func New(client *http.Client, cfg *config.Config) *Downloader {
	if client == nil {
		client = http.DefaultClient
		client.Timeout = 30 * time.Second
	}
	return &Downloader{
		Client:    client,
		Now:       time.Now(),
		BaseURL:   cfg.BaseURL,
		OutputDir: cfg.OutputDir,
		MediaType: cfg.Type,
	}
}

func (d *Downloader) Download() (string, error) {
	n := d.Now
	filename := getFileName(n, d.MediaType)
	url := fmt.Sprintf("%s/%s", d.BaseURL, filename)

	res, err := d.Client.Get(url)
	if err != nil {
		return "", fmt.Errorf("error downloading file: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download file: %s", res.Status)
	}

	if err := os.MkdirAll(d.OutputDir, 0o755); err != nil {
		return "", fmt.Errorf("error creating tmp directory: %w", err)
	}

	output, err := os.Create(fmt.Sprintf("%s/%s", d.OutputDir, filename))
	if err != nil {
		return "", fmt.Errorf("error creating file: %w", err)
	}
	defer output.Close()

	if _, err := io.Copy(output, res.Body); err != nil {
		return "", fmt.Errorf("error saving file: %w", err)
	}

	return output.Name(), nil
}

func getFileName(t time.Time, mediaType models.Type) string {
	return fmt.Sprintf(
		"%s_ids_%02d_%02d_%d.json.gz",
		mediaType,
		t.Month(), t.Day(), t.Year(),
	)
}
