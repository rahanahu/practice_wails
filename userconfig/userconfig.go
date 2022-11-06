package userconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const appConfigFileName string = "testApp/test.json"

type UserConfig struct {
	Dir string `json:"directory"`
	ctx *context.Context
}

func NewUserConfig(c *context.Context) *UserConfig {
	return &UserConfig{ctx: c}
}

func (a *UserConfig) SaveUserConfig() error {
	w, err := configFileWriteHandler()
	if err != nil {
		return err
	}

	if err := a.saveconfig(w); err != nil {
		return err
	}
	return nil
}

func (a *UserConfig) saveconfig(w io.Writer) error {
	s, err := json.Marshal(a)
	if err != nil {
		return err
	}
	if _, err := w.Write(s); err != nil {
		return err
	}
	return nil
}

func configFilePath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, appConfigFileName), nil
}

func configFileWriteHandler() (io.Writer, error) {
	path, err := configFilePath()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return nil, err
	}
	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %w", err)
	}

	return f, nil
}

func (a *UserConfig) SelectDirectory() (string, error) {

	o := runtime.OpenDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            "",
		Title:                      "Select Directory",
		Filters:                    nil,
		ShowHiddenFiles:            true,
		CanCreateDirectories:       true,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: true,
	}
	if s, err := runtime.OpenDirectoryDialog(*a.ctx, o); err != nil {
		return s, err
	} else {
		a.Dir = s
		a.SaveUserConfig()
		return s, nil
	}

}

func (a *UserConfig) LoadUserConfig() (*UserConfig, error) {
	path, err := configFilePath()
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var u UserConfig
	if err := json.Unmarshal(bs, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
