package nlp

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	var source = "gr_gr_cog~tr.d015~trMA.LaudsGlory.text"
	var text = "Προκαθάρωμεν ἑαυτούς, ἀδελφοί, τῇ βασιλίδι τῶν ἀρετῶν· ἰδοὺ γὰρ παραγέγονε, πλοῦτον ἡμῖν ἀγαθῶν κομίζουσα, τῶν παθῶν κατευνάζει τὰ οἰδήματα, καὶ τῷ Δεσπότῃ καταλλάττει τοὺς πταίσαντας· διὸ μετ' εὐφροσύνης ταύτην ὑποδεξώμεθα, βοῶντες Χριστῷ τῷ Θεῷ, ὁ ἀναστὰς ἐκ τῶν νεκρῶν, ἀκατακρίτους ἡμᾶς διαφύλαξον, δοξολογοῦντάς σε τὸν μόνον ἀναμάρτητον."
	for _, token := range Tokens(source, text, true, "gr") {
		fmt.Printf("%s\t%s\n",token.Position, token.Text)
	}
}
func TestFrequency_Count(t *testing.T) {
	var source = "gr_gr_cog~tr.d015~trMA.LaudsGlory.text"
	var text = "Προκαθάρωμεν ἑαυτούς, ἀδελφοί, τῇ βασιλίδι τῶν ἀρετῶν· ἰδοὺ γὰρ παραγέγονε, πλοῦτον ἡμῖν ἀγαθῶν κομίζουσα, τῶν παθῶν κατευνάζει τὰ οἰδήματα, καὶ τῷ Δεσπότῃ καταλλάττει τοὺς πταίσαντας· διὸ μετ' εὐφροσύνης ταύτην ὑποδεξώμεθα, βοῶντες Χριστῷ τῷ Θεῷ, ὁ ἀναστὰς ἐκ τῶν νεκρῶν, ἀκατακρίτους ἡμᾶς διαφύλαξον, δοξολογοῦντάς σε τὸν μόνον ἀναμάρτητον."
	var freq Frequency
	freq.Map = make(map[string]int)

	freq.Count(source, text, true, "gr")
	for token, count := range freq.Map {
		fmt.Printf("%s\t%d\n",token,count)
	}
}