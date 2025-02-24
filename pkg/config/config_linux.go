package config

// GetPlatformDefaultConfig gets the defaults for the platform
func GetPlatformDefaultConfig() OSConfig {
	return OSConfig{
		EditCommand:         ``,
		EditCommandTemplate: `{{editor}} {{filename}}`,
		OpenCommand:         `sh -c "xdg-open {{filename}} >/dev/null"`,
		OpenLinkCommand:     `sh -c "xdg-open {{link}} >/dev/null"`,
	}
}
