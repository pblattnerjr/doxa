/*
Provides utilities for parsing lines from ares files from the AGES Liturgical Workbench (ALWB) system. The acronymn `ares` means, `AGES resource.` The name of an ares file indicates its subject (called a topic) and its version (called a domain).
*/
package ares

import (
	"errors"
	"github.com/liturgiko/doxa/pkg/utils/ltstring"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	ErrFileMissingAres = errors.New("ares: file missing .ares extension")
	ErrFileMissingTopic = errors.New("ares: filename missing topic or domain")
	ErrLineMissingEqualSign = errors.New("ares: missing equal sign between key and value")
	ErrRedirectInvalid = errors.New("ares: invalid redirect value")
	ErrValueMissingQuote = errors.New("ares: value missing initial or final quote")
)

type LineParts struct {
	Language       string
	Country        string
	Realm          string
	Topic          string
	Key            string
	Value          string
	Redirect       string
	Comment        string
	IsAresId       bool
	IsBlank        bool
	IsCommentedOut bool
	IsRedirect     bool
	HasComment     bool
	HasValue       bool
	LineNbr        int
}

type FilenameParts struct {
	Topic    string
	Language string
	Country  string
	Realm    string
}

// Parse the line into the fields of a LineParts struct.
// We include the language, country, realm, and topic
// so they are available to downstream channels
func ParseLine(filenameParts FilenameParts, line string) (LineParts, error) {
	var result LineParts
	result.Language = filenameParts.Language
	result.Country = filenameParts.Country
	result.Realm = filenameParts.Realm
	result.Topic = filenameParts.Topic
	result.IsAresId = strings.HasPrefix(line, "A_")
	result.IsBlank = len(strings.TrimSpace(line)) < 1
	result.IsCommentedOut = strings.HasPrefix(line, "//")
	if result.IsBlank || result.IsAresId || result.IsCommentedOut {
		// fall thru
	} else {
		parts := strings.Split(line, "=")
		if len(parts) > 1 {
			result.Key = strings.TrimSpace(parts[0])
			result.Value = strings.TrimSpace(parts[1])
			if strings.Index(result.Value, "//") > -1 {
				result.HasComment = true
				valueParts := strings.Split(result.Value, "//")
				result.Value = strings.TrimSpace(valueParts[0])
				result.Comment = strings.TrimSpace(valueParts[1])
			}
			result.IsRedirect = !strings.HasPrefix(result.Value, "\"")
			if result.IsRedirect {
				result.Redirect, _ = ToRedirectId(result.Value)
				result.Value = ""
			} else {
				if strings.HasPrefix(result.Value, "\"") && strings.HasSuffix(result.Value, "\"") {
					result.Value, _ = strconv.Unquote(result.Value)
					result.HasValue = true
				} else {
					return result, ErrValueMissingQuote
				}
			}
		} else {
			return result, ErrLineMissingEqualSign
		}
	}
	return result, nil
}

// Returns from the filename the topic, ISO country code, ISO language code, and realm.
func ParseAresFileName(f string) (FilenameParts, error) {
	var result FilenameParts
	_, filename := filepath.Split(f)
	if strings.HasSuffix(filename, ".ares") {
		parts := strings.Split(filename[:len(filename)-5], "_")
		if len(parts) == 4 {
			result.Topic = parts[0]
			result.Language = strings.ToLower(parts[1])
			result.Country = strings.ToLower(parts[2])
			result.Realm = strings.ToLower(parts[3])
		} else {
			return result, ErrFileMissingTopic
		}
	} else {
		return result, ErrFileMissingAres
	}
	return result, nil
}

// Converts value from this form: properties_en_UK_lash.media.key
// To this form: en_us_lash~properties~media.key
func ToRedirectId(value string) (string, error) {
	var lang, country, realm, topic, key string
	domainParts := strings.Split(strings.TrimSpace(value), "_")
	if len(domainParts) == 4 {
		topic = domainParts[0]
		lang = domainParts[1]
		country = domainParts[2]
		keyParts := strings.Split(domainParts[3], ".")
		realm = keyParts[0]
		key = domainParts[3][len(realm)+1:]
		return ltstring.ToId(lang, country, realm, topic, key), nil
	} else {
		return value, ErrRedirectInvalid
	}
}

