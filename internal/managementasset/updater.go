// Package managementasset locates the management control panel asset.
//
// a5proxy bundles the panel (management.html) into its image and never
// downloads it at runtime. The remote-management.panel-github-repository and
// disable-auto-update-panel config keys are still parsed but ignored.
package managementasset

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/router-for-me/CLIProxyAPI/v8/internal/config"
	"github.com/router-for-me/CLIProxyAPI/v8/internal/util"
	log "github.com/sirupsen/logrus"
)

const managementAssetName = "management.html"

// ManagementFileName exposes the control panel asset filename.
const ManagementFileName = managementAssetName

// SetCurrentConfig is kept for API compatibility; the bundled panel needs no config.
func SetCurrentConfig(_ *config.Config) {}

// StartAutoUpdater no longer updates anything. It only reports whether the
// bundled management asset exists so a missing panel is visible in the logs.
func StartAutoUpdater(_ context.Context, configFilePath string) {
	path := FilePath(strings.TrimSpace(configFilePath))
	if path == "" {
		log.Debug("management control panel: no static path resolved")
		return
	}
	if _, err := os.Stat(path); err != nil {
		log.Warnf("management control panel not found at %s; /management.html will return 404", path)
		return
	}
	log.Infof("management control panel: serving bundled asset %s", path)
}

// StaticDir resolves the directory that stores the management control panel asset.
func StaticDir(configFilePath string) string {
	if override := strings.TrimSpace(os.Getenv("MANAGEMENT_STATIC_PATH")); override != "" {
		cleaned := filepath.Clean(override)
		if strings.EqualFold(filepath.Base(cleaned), managementAssetName) {
			return filepath.Dir(cleaned)
		}
		return cleaned
	}

	if writable := util.WritablePath(); writable != "" {
		return filepath.Join(writable, "static")
	}

	configFilePath = strings.TrimSpace(configFilePath)
	if configFilePath == "" {
		return ""
	}

	base := filepath.Dir(configFilePath)
	fileInfo, err := os.Stat(configFilePath)
	if err == nil {
		if fileInfo.IsDir() {
			base = configFilePath
		}
	}

	return filepath.Join(base, "static")
}

// FilePath resolves the absolute path to the management control panel asset.
func FilePath(configFilePath string) string {
	if override := strings.TrimSpace(os.Getenv("MANAGEMENT_STATIC_PATH")); override != "" {
		cleaned := filepath.Clean(override)
		if strings.EqualFold(filepath.Base(cleaned), managementAssetName) {
			return cleaned
		}
		return filepath.Join(cleaned, ManagementFileName)
	}

	dir := StaticDir(configFilePath)
	if dir == "" {
		return ""
	}
	return filepath.Join(dir, ManagementFileName)
}
