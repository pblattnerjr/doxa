package ltstring

import (
	"strings"
	"testing"
)

// Test text to Nnp
func TestToNnp(t *testing.T) {
	s := "Εὐλογημένη ἡ Βασιλεία τοῦ Πατρὸς, καὶ τοῦ Υἱοῦ καὶ τοῦ Ἁγίου Πνεύματος."
	e := "ευλογημενη η βασιλεια του πατρος και του υιου και του αγιου πνευματος"
	r := ToNnp(s)
	if strings.Compare(r, e) != 0 {
		t.Error("expected ltstring to match", "\nOriginal: "+s+"\nExpected: "+e+"\nGot:      "+r)
	}
}

// Test text to Nnp
func TestToNwp(t *testing.T) {
	s := "Εὐλογημένη ἡ Βασιλεία τοῦ Πατρὸς, καὶ τοῦ Υἱοῦ καὶ τοῦ Ἁγίου Πνεύματος."
	e := "ευλογημενη η βασιλεια του πατρος, και του υιου και του αγιου πνευματος."
	r := ToNwp(s)
	if strings.Compare(r, e) != 0 {
		t.Error("expected ltstring to match", "\nOriginal: "+s+"\nExpected: "+e+"\nGot:      "+r)
	}
}
