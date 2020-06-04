// Provides pipelines for to map things to another form
package mappers

import (
	"github.com/liturgiko/doxa/pkg/ages/ares"
	"github.com/liturgiko/doxa/pkg/models"
	"github.com/liturgiko/doxa/pkg/utils/ltstring"
	"time"
)

// Converts a stream of LineParts into Ltx structs
func Lp2Lt(in <-chan *ares.LineParts) <-chan *models.Ltx {
	out := make(chan *models.Ltx)
	go func() {
		for lineParts := range in {
			var ltx models.Ltx
			ltx.ID = ltstring.ToId(
				lineParts.Language,
				lineParts.Country,
				lineParts.Realm,
				lineParts.Topic,
				lineParts.Key,
			)
			ltx.Library = ltstring.ToDomain(lineParts.Language, lineParts.Country, lineParts.Realm)
			ltx.Topic = lineParts.Topic
			ltx.Key = lineParts.Key

			if lineParts.HasValue {
				ltx.Value = lineParts.Value
				ltx.NWP = ltstring.ToNwp(lineParts.Value)
				ltx.NNP = ltstring.ToNnp(ltx.NWP)
			}
			if lineParts.HasComment {
				ltx.Comment = lineParts.Comment
			}
			if lineParts.IsRedirect {
				ltx.Redirect = lineParts.Redirect
			}
			timestamp := time.Now().UTC().String()
			ltx.CreatedWhen = timestamp
			ltx.ModifiedWhen = timestamp
			out <- &ltx
		}
		close(out)
	}()
	return out
}


