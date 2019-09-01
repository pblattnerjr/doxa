// Provides pipelines for to map things to another form
package mappers

import (
	"github.com/liturgiko/doxa/pkg/models"
	"github.com/liturgiko/doxa/pkg/utils/ares"
	"github.com/liturgiko/doxa/pkg/utils/ltstring"
)

// Converts a stream of LineParts into Ltext structs
func Lp2Lt(in <-chan *ares.LineParts) <-chan *models.Ltext {
	out := make(chan *models.Ltext)
	go func() {
		for lineParts := range in {
			var ltext models.Ltext
			ltext.ID = ltstring.ToId(
				lineParts.Language,
				lineParts.Country,
				lineParts.Realm,
				lineParts.Topic,
				lineParts.Key,
			)
			ltext.Topic = lineParts.Topic
			ltext.Key = lineParts.Key

			if lineParts.HasValue {
				ltext.Value = lineParts.Value
				ltext.NWP = ltstring.ToNwp(lineParts.Value)
				ltext.NNP = ltstring.ToNnp(ltext.NWP)
			}
			if lineParts.HasComment {
				ltext.Comment = lineParts.Comment
			}
			if lineParts.IsRedirect {
				ltext.Redirect = lineParts.Redirect
			}
			out <- &ltext
		}
		close(out)
	}()
	return out
}


