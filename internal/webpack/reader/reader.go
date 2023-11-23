package reader

import (
	"errors"
	"go-web-base/internal/webpack/reader/manifest"
	"go-web-base/internal/webpack/reader/stats"
)

// Read assets asset manifest
func Read(plugin, host, fsPath, webPath string, isDev bool) (map[string][]string, error) {
	//log.Println("read", plugin, isDev)
	if plugin == "stats" {
		return stats.Read(isDev, host, fsPath, webPath)
	} else if plugin == "manifest" {
		return manifest.Read(fsPath)
	} else {
		return map[string][]string{}, errors.New("go-webpack: bad plugin type")
	}

}
