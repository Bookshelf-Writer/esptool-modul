package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type VersionInfo struct {
	Name        string
	Description string
	Version     string
}

func GetVersionInfo() *VersionInfo {
	return &VersionInfo{
		Name:        "Esptool",
		Description: "ESP32 flashing utility written in GoLang",
		Version:     GlobalVersion,
	}
}

func (v *VersionInfo) String() string {
	return fmt.Sprintf("%s\n\n%s\nVersion: %s\n",
		v.Name,
		v.Description,
		v.Version,
	)

}

func versionCommand(jsonOutput bool) error {
	if jsonOutput {
		prettyJson, err := json.MarshalIndent(GetVersionInfo(), "", "  ")
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(prettyJson)
		return err
	} else {
		_, err := fmt.Print(GetVersionInfo().String())
		return err
	}
}
