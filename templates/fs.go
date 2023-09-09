package templates

import "embed"

//go:embed *
var static embed.FS

func Load(file string) ([]byte, error) {
	return static.ReadFile(file + ".stub")
}
