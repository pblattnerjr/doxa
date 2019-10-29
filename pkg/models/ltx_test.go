package models

import "testing"

var sharedLtex Ltx

func init() {
	sharedLtex.ID = "gr_gr_cog~actors~Priest"
	sharedLtex.Value = "ΙΕΡΕΥΣ"
}

// Test Ltx.ToId() for well formed ID
// Happy Path (HP) test
func TestLtxToAresHP(t *testing.T) {

	var domain Domain
	domain.Realm = "cog"
	domain.Country = "gr"
	domain.Language = "gr"

	var id Id
	id.Domain = domain
	id.Topic = "actors"
	id.Key = "Priest"

	testId, err := sharedLtex.ToId()
	if err != nil {
		t.Error(err.Error())
	}
	if testId.Language != domain.Language {
		t.Errorf("Error: expected %s, got %s", domain.Language, testId.Language)
	}
	if testId.Country != domain.Country {
		t.Errorf("Error: expected %s, got %s", domain.Country, testId.Country)
	}
	if testId.Realm != domain.Realm {
		t.Errorf("Error: expected %s, got %s", domain.Realm, testId.Realm)
	}
	if testId.Topic != id.Topic {
		t.Errorf("Error: expected %s, got %s", id.Topic, testId.Topic)
	}
	if testId.Key != id.Key {
		t.Errorf("Error: expected %s, got %s", id.Key, testId.Key)
	}
}
