// Provides string utilities for liturgical text processing
package ltstring

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"time"
	"unicode"
)

// Converts text to a Normalized form with no punctuation. It will be lowercase with no diacritics or punctuation.
func ToNnp(text string) string {
	rmvAccents := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	rmvPunct := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Punct)), norm.NFC)
	result, _, _ := transform.String(rmvAccents, text)
	result, _, _ = transform.String(rmvPunct, result)
	result = strings.ToLower(result)
	return result
}

// Converts text to a Normalized form with punctuation. It will be lowercase with no diacritics. Punctuation is preserved.
func ToNwp(text string) string {
	rmvAccents := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(rmvAccents, text)
	result = strings.ToLower(result)
	return result
}

// Create canonical OLW ID from supplied parts
func ToId(lang string, country string, realm string, topic string, key string) string {
	return ToDomain(lang, country, realm) + "~" + topic + "~" + key
}

// Create canonical OLW domain from parts
func ToDomain(language string, country string, realm string) string {
	return strings.ToLower(language + "_" + country + "_" + realm)
}

// Returns a timestamp that is standardized for Liturgical Text Software (LTS).
// Always returns UTC.
// Example: 2019-08-14 12:50:29.275002 +0000 UTC
func Timestamp() string {
	return 	time.Now().UTC().String()
}

