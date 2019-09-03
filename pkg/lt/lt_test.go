package lt

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	home := os.Getenv("home")
	site := os.Getenv("site")
	template := os.Getenv("template")
	domains := []string{"gr_gr_cog", "en_us_dedes", "spa_ga_"}
	Generate(home, template, site, domains)
}
