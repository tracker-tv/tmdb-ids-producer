package downloader

import (
	"compress/gzip"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tracker-tv/tmdb-ids-producer/internal/config"
	"github.com/tracker-tv/tmdb-ids-producer/internal/models"
)

func TestGetFileName(t *testing.T) {
	now := time.Date(2026, time.January, 4, 0, 0, 0, 0, time.UTC)
	filename := getFileName(now, "movie")
	expected := "movie_ids_01_04_2026.json.gz"
	if filename != expected {
		t.Errorf("expected %s, got %s", expected, filename)
	}
}

func TestDownload(t *testing.T) {
	testdataFile := "../../testdata/movie_ids_12_30_2025.json"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile(testdataFile)
		if err != nil {
			t.Fatalf("failed to read testdata file: %v", err)
		}

		w.Header().Set("Content-Type", "application/gzip")
		gzipWriter := gzip.NewWriter(w)
		defer gzipWriter.Close()

		if _, err := gzipWriter.Write(data); err != nil {
			t.Fatalf("failed to write gzipped response: %v", err)
		}
	}))
	defer server.Close()

	cfg := &config.Config{
		BaseURL:   server.URL,
		OutputDir: t.TempDir(),
		Type:      models.Movie,
	}

	now := time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC)
	downloader := New(nil, cfg)
	downloader.Now = now

	filePath, err := downloader.Download()
	if err != nil {
		t.Fatalf("Download failed: %v", err)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Downloaded file does not exist: %s", filePath)
	}

	expectedFilename := "movie_ids_12_30_2025.json.gz"
	if filepath.Base(filePath) != expectedFilename {
		t.Errorf("expected filename %s, got %s", expectedFilename, filepath.Base(filePath))
	}
}
