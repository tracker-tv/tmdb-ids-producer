package downloader

import (
	"testing"
	"time"
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
}
