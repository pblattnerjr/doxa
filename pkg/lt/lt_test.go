package lt

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	home := os.Getenv("home")
	dbPath := os.Getenv("dbPath")
	site := os.Getenv("site")
	template := os.Getenv("template")
	domains := []string{"gr_gr_cog", "en_us_dedes", "spa_ga_"}
	GenerateFromTemplate(home, dbPath, template, site, domains)
}
