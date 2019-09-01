// Calculates the liturgical properties for the specified date
// Ported from AGES Initiatives Java version.  The original
// liturgical day properties was written by John Holder of
// St. Catherine Greek Orthodox Church in Denver, Colorado.
// The golang version was ported by Michael Colburn, OCMC.
package ldp

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Triodion: 1st day: Sunday of Publican and Pharisee.  9 weeks before Pascha.
// 1st three Sundays precede Great Lent.
// Pascha: Day 1 of Pentecostarion.
// All-Saints: Last Day of Pentecostarion.
// Apostles' Fast: Monday after All Saints, up to and including Jun29, ApostlesPeter&Paul
//
// Thomas Sunday: eight-tone cycle begins w/ Tone 1, ends fri. 6th week Lent.
// Sunday of All-Saints: Eothinon cycle begins with Eothinon 1  (tones of week pl.4)
// Eothinon Cycle runs thru 5th Sunday of Lent (Sunday before Palm Sunday)
//

// LDP is the main struct for liturgical day properties
type LDP struct {
	TheDay                                   time.Time
	AllSaintsDateLastYear                    time.Time
	AllSaintsDateThisYear                    time.Time
	DayOfSeason                              int    // return 1..70 (0 if no day set). Valid only when isPentecostarion or isTriodion.
	DayOfWeek                                string // for debugging output
	DayOfWeekOverride                        string // for debugging output
	DaysSinceStartLastLukanCycle             int
	DaysSinceStartOfTriodion                 int
	DaysSinceSundayAfterLastElevationOfCross int
	DaysUntilStartOfTriodion                 int // Used to control lectionary and commemorations
	ElevationOfCrossDateLast                 time.Time
	EothinonNumber                           int // 0..11. Valid values for 11 week cycle, only valid on Sundays!!!!
	GreatLentStartDate                       time.Time
	IsDaysOfLuke                             bool
	IsFriday                                 bool
	IsMonday                                 bool
	IsPascha                                 bool
	IsPentecostarion                         bool
	IsSaturday                               bool
	IsSunday                                 bool
	IsThursday                               bool
	IsTriodion                               bool
	IsTuesday                                bool
	IsWednesday                              bool
	LazarusSaturdayNextDate                  time.Time
	ModeOfWeek                               int // return 0..8
	ModeOfWeekOverride                       int
	NbrDayOfMonth                            string
	NbrDayOfWeek                             string
	NbrDayOfWeekOverride                     string
	NbrMonth                                 string
	NumberOfSundaysBeforeStartOfTriodion     int
	originalYear                             int
	originalMonth                            int
	originalDay                              int
	originalDayOfSeason                      int
	PalmSundayDate                           time.Time
	PalmSundayNextDate                       time.Time
	PaschaDateLast                           time.Time
	PaschaDateLastYear                       time.Time
	PaschaDateNext                           time.Time
	PaschaDateThisYear                       time.Time
	PentecostDate                            time.Time
	StartDateOfLukanCycleLast                time.Time
	SundayAfterElevationOfCrossDateLast      time.Time
	TheDayBefore                             time.Time
	ThomasSundayDate                         time.Time
	TriodionStartDateLast                    time.Time
	TriodionStartDateLastYear                time.Time
	TriodionStartDateNextYear                time.Time
	TriodionStartDateThisYear                time.Time
}

func (ldp *LDP) StrYear() string {
	return fmt.Sprintf("%d", ldp.TheDay.Year())
}
func (ldp *LDP) StrMonth() string {
	return fmt.Sprintf("%d", ldp.TheDay.Month())
}
func (ldp *LDP) StrDay() string {
	return fmt.Sprintf("%d", ldp.TheDay.Day())
}
func (ldp *LDP) WeekDay() string {
	return fmt.Sprintf("%d", ldp.TheDay.Weekday())
}

// used as the format to create a date
const ShortForm = "2019-10-25"

func validateYMD(year, month, day int) error {
	if year < 1583 {
		return errors.New("year cannot be less than 1583")
	}
	if month < 1 || month > 12 {
		return errors.New("month must be between 1 and 12")
	}
	if day < 1 || day > 31 {
		return errors.New("day must be between 1 and 31")
	}
	return nil
}
// Creates a new LDP initialized to the specified date
func NewLDPYMD(year, month, day int) (LDP, error) {
	var ldp LDP
	if err := validateYMD(year, month, day); err != nil {return ldp, err}

	today := time.Now()
	t := NewDate(year, month, day)
	// if the date is before today, do it for next year by default.
	if t.Before(today) {
		t = NewDate(year+1, month, day)
	}
	ldp.TheDay = t
	ldp.reinitializeOriginalDateTrackers()
	ldp.SetLiturgicalPropertiesByDate(t.Year())
	ldp.SetYesterday()
	return ldp, nil
}

// Returns a new LDP initialized to the specified month and day.  The year is set to the current one.
func NewLDPMD(month, day int) (LDP, error) {
	ldp, err := NewLDPYMD(time.Now().Year(), month, day)
	return ldp, err
}

// Returns a new LDP initialized for today's date
func NewLDP() (LDP, error) {
	now := time.Now()
	ldp, err := NewLDPYMD(now.Year(), int(now.Month()), now.Day())
	return ldp, err
}

func NewDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

var GreekMonths = []string{"Ἰανουαρίου",
	"Φεβρουαρίου",
	"Μαρτίου",
	"Ἀπριλίου",
	"Μαΐου",
	"Ἰουνίου",
	"Ἰουλίου",
	"Αὐγούστου",
	"Σεπτεμβρίου",
	"Ὀκτωβρίου",
	"Νοεμβρίου",
	"Δεκεμβρίου",
}
var GreekWeekDays = []string{
	"τοῦ Σαββάτου",
	"τῆς Κυριακῆς",
	"τῆς Δευτέρας",
	"τῆς Τρίτης",
	"τῆς Τετάρτης",
	"τῆς Πέμπτης",
	"τῆς Παρασκευῆς",
}
var GreekMonthDays = []string{
	"αʹ",
	"βʹ",
	"γʹ",
	"δʹ",
	"εʹ",
	"Ϛʹ",
	"ζʹ",
	"ηʹ",
	"θʹ",
	"ιʹ",
	"ιαʹ",
	"ιβʹ",
	"ιγʹ",
	"ιδʹ",
	"ιεʹ",
	"ιϚʹ",
	"ιζʹ",
	"ιηʹ",
	"ιθʹ",
	"κʹ",
	"καʹ",
	"κβʹ",
	"κγʹ",
	"κδʹ",
	"κεʹ",
	"κϚʹ",
	"κζʹ",
	"κηʹ",
	"κθʹ",
	"λʹ",
	"λαʹ",
}
var GreekMap = map[string]string{
	"Ιανουάριος":  "Ἰανουαρίου",
	"Φεβρουάριος": "Φεβρουαρίου",
	"Μάρτιος":     "Μαρτίου",
	"Απρίλιος":    "Ἀπριλίου",
	"Μάϊος":       "Μαΐου",
	"Ιούνιος":     "Ἰουνίου",
	"Ιούλιος":     "Ἰουλίου",
	"Αύγουστος":   "Αὐγούστου",
	"Σεπτέμβριος": "Σεπτεμβρίου",
	"Οκτώβριος":   "Ὀκτωβρίου",
	"Νοέμβριος":   "Νοεμβρίου",
	"Δεκέμβριος":  "Δεκεμβρίου",
	"Σάββατο":     "τοῦ Σαββάτου",
	"Κυριακή":     "τῆς Κυριακῆς",
	"Δευτέρα":     "τῆς Δευτέρας",
	"Τρίτη":       "τῆς Τρίτης",
	"Τετάρτη":     "τῆς Τετάρτης",
	"Πέμπτη":      "τῆς Πέμπτης",
	"Παρασκευή":   "τῆς Παρασκευῆς",
	"1":           "αʹ",
	"2":           "βʹ",
	"3":           "γʹ",
	"4":           "δʹ",
	"5":           "εʹ",
	"6":           "Ϛʹ",
	"7":           "ζʹv",
	"8":           "ηʹ",
	"9":           "θʹ",
	"10":          "ιʹ",
	"11":          "ιαʹ",
	"12":          "ιβʹ",
	"13":          "ιγʹ",
	"14":          "ιδʹ",
	"15":          "ιεʹ",
	"16":          "ιϚʹ",
	"17":          "ιζʹ",
	"18":          "ιηʹ",
	"19":          "ιθʹ",
	"20":          "κʹ",
	"21":          "καʹ",
	"22":          "κβʹ",
	"23":          "κγʹ",
	"24":          "κδʹ",
	"25":          "κεʹ",
	"26":          "κϚʹ",
	"27":          "κζʹ",
	"28":          "κηʹ",
	"29":          "κθʹ",
	"30":          "λʹ",
	"31":          "λαʹ",
}

func (ldp *LDP) reinitializeOriginalDateTrackers() {
	ldp.originalYear = -1
	ldp.originalMonth = -1
	ldp.originalDay = -1
	ldp.originalDayOfSeason = -1
}
func (ldp LDP) setDateTo(year, month, day int) {
	ldp.TheDay = NewDate(year, month, day)
	//  setLiturgicalPropertiesByDate(theYear)
	ldp.setYesterday()

	// if not already set, save the date values
	if ldp.originalYear == -1 {
		ldp.originalYear = year
	}
	if ldp.originalMonth == -1 {
		ldp.originalMonth = month
	}
	if ldp.originalDay == -1 {
		ldp.originalDay = day
	}
	if ldp.originalDayOfSeason == -1 {
		ldp.originalDayOfSeason = ldp.DayOfSeason
	}
}
func (ldp *LDP) ResetDate() {}
func (ldp *LDP) setYesterday() {
	ldp.TheDayBefore = ldp.TheDay.AddDate(0, 0, -1)
}
func (ldp *LDP) TimeDelta(dateFrom time.Time, days int) time.Time { return dateFrom }

func (ldp *LDP) SetLiturgicalPropertiesByDate(year int) {
	ldp.SetVariablesToDefaults()
	useGregorianCalendar := true
	ldp.PaschaDateLastYear = ComputeDayOfPascha(year-1, useGregorianCalendar)
	ldp.PaschaDateThisYear = ComputeDayOfPascha(year, useGregorianCalendar)
	ldp.PaschaDateLast = ldp.lastPaschaDate()
	ldp.PaschaDateNext = ldp.nextPaschaDate()
	// 10 weeks before Pascha (inclusive), Starts with the Sunday of Publican and Pharisee
	ldp.TriodionStartDateThisYear = ldp.PaschaDateThisYear.AddDate(0, 0, -(10*7))
	ldp.TriodionStartDateLastYear = ldp.PaschaDateLastYear.AddDate(0,0, -(10*7) )
	ldp.TriodionStartDateNextYear = ldp.PaschaDateNext.AddDate(0,0, -(10*7) )
	ldp.setDateLastTriodionStart()

	ldp.PalmSundayDate = ldp.PaschaDateThisYear.AddDate(0,0,-7 )
	ldp.PentecostDate = ldp.PaschaDateThisYear.AddDate(0,0, 49 )
	ldp.AllSaintsDateThisYear = ldp.PaschaDateThisYear.AddDate(0,0, 56 )
	ldp.AllSaintsDateLastYear = ldp.PaschaDateLastYear.AddDate(0,0, 56 )
	// Pentecost starts  with Pascha and ends with All Saints, which is the day before the beginning
	// of the Apostle's Fast.
	if ldp.TheDay.Equal(ldp.PaschaDateThisYear) ||
		ldp.TheDay.Equal(ldp.AllSaintsDateThisYear) ||
		(ldp.TheDay.After(ldp.PaschaDateThisYear) && ldp.TheDay.Before(ldp.AllSaintsDateThisYear)) {
		ldp.IsPentecostarion = true
	} else {
		ldp.IsPentecostarion = false
	}

	if ldp.TheDay.Equal(ldp.TriodionStartDateThisYear) ||
		(ldp.TheDay.After(ldp.TriodionStartDateThisYear) && ldp.TheDay.Before(ldp.PaschaDateThisYear)) {
		ldp.IsTriodion = true
	}

	// Clean Monday, 7 weeks + a day before Pascha
	ldp.GreatLentStartDate = NewDate(ldp.PaschaDateThisYear.Year(), 0,-(7*7)+1 )
	ldp.PalmSundayNextDate = NewDate(ldp.PaschaDateNext.Year(), 0,-7 )
	ldp.ThomasSundayDate = NewDate(ldp.PaschaDateLast.Year(), 0,7 )
	ldp.LazarusSaturdayNextDate = NewDate(ldp.PaschaDateNext.Year(), 0,-8 )

	ldp.SetDayOfSeason()
	ldp.setDaysSinceStartOfLastTriodion()
	ldp.SetDayOfWeek()
	ldp.setEothinonNumber()
	ldp.setModeOfWeek()
	ldp.setNbrDayOfMonth(ldp.TheDay.Day())
	ldp.setNbrMonth(int(ldp.TheDay.Month()+1))

	ldp.setDateFirstSundayAfterElevationOfCross()
	ldp.setDaysSinceSundayAfterLastElevationOfCross()
	ldp.SetDateStartLukanCycle()
	ldp.setDaysSinceStartLukanCycleLast()
	ldp.ElevationOfCrossDateLast = NewDate(ldp.SundayAfterElevationOfCrossDateLast.Year(),
		int(ldp.SundayAfterElevationOfCrossDateLast.Month()), 14)
	ldp.SetNumberOfSundaysBeforeStartOfTriodionOnJan15()
}

func (ldp *LDP) SetNumberOfSundaysBeforeStartOfTriodionOnJan15() {
	jan15 := NewDate(ldp.TriodionStartDateThisYear.Year(),0,15)
	diffMillis := DiffMillis(ldp.TriodionStartDateThisYear, jan15)
	// Get difference in days, add 1 to be 1-index based instead of zero.
	ldp.DaysUntilStartOfTriodion = int( diffMillis / (24*60*60*1000))
	ldp.NumberOfSundaysBeforeStartOfTriodion = ldp.DaysUntilStartOfTriodion / 7
}

func (ldp *LDP) SetNumberOfSundaysBeforeStartOfTriodion() {
	diffMillis := DiffMillis(ldp.TriodionStartDateThisYear, ldp.TheDay)
	ldp.DaysUntilStartOfTriodion = int( diffMillis / (24*60*60*1000))
	if ldp.DaysUntilStartOfTriodion < 0 {
		ldp.DaysUntilStartOfTriodion = 0
		ldp.NumberOfSundaysBeforeStartOfTriodion = 0
	} else {
		ldp.NumberOfSundaysBeforeStartOfTriodion = ldp.DaysUntilStartOfTriodion / 7
	}
}

func (ldp *LDP) GetMonthOfSundayAfterElevationOfCross() int { return 0 }

func (ldp *LDP) getDayOfSundayAfterElevationOfCross() int { return 0 }

func (ldp *LDP) setDateLastTriodionStart() {
	if ldp.TheDay.Before(ldp.TriodionStartDateThisYear) {
		ldp.TriodionStartDateLast = ldp.TriodionStartDateLastYear
	} else {
		ldp.TriodionStartDateLast = ldp.TriodionStartDateThisYear
	}
}

func (ldp *LDP) setDateFirstSundayAfterElevationOfCross() {
	firstSundayAfterElevationThisYear := computeSundayAfterElevationOfCross(NewDate(ldp.TheDay.Year(), 8, 14))
	firstSundayAfterElevationLastYear := computeSundayAfterElevationOfCross(NewDate(ldp.TheDay.Year()-1, 8, 14))
	if ldp.TheDay.Before(firstSundayAfterElevationThisYear) {
		ldp.SundayAfterElevationOfCrossDateLast = firstSundayAfterElevationLastYear
	} else {
		ldp.SundayAfterElevationOfCrossDateLast = firstSundayAfterElevationThisYear
	}
}

func computeSundayAfterElevationOfCross(date time.Time) time.Time { return date }

func (ldp *LDP) SetDateStartLukanCycle() {
	firstSundayAfterElevationThisYear := computeSundayAfterElevationOfCross(NewDate(ldp.TheDay.Year(), 8, 14))
	firstSundayAfterElevationLastYear := computeSundayAfterElevationOfCross(NewDate(ldp.TheDay.Year()-1, 8, 14))
	startLukanCycleThisYear := firstSundayAfterElevationThisYear.AddDate(0, 0, 1)
	startLukanCycleLastYear := firstSundayAfterElevationLastYear.AddDate(0, 0, 1)
	if ldp.TheDay.Before(startLukanCycleThisYear) {
		ldp.StartDateOfLukanCycleLast = startLukanCycleLastYear
	} else {
		ldp.StartDateOfLukanCycleLast = startLukanCycleThisYear
	}
}

// pass in the year and receive the month and day of Pascha.
func ComputeDayOfPascha(year int, isUserCalendarGregorian bool) time.Time {
	var month, day, r19, r7, r4, n1, n2, n3, cent int
	r19 = year % 19
	r7 = year % 7
	r4 = year % 4
	// This is a formula by Gauss for the number of days after 21-Mar.
	n1 = (19*r19 + 16) % 30
	n2 = (2*r4 + 4*r7 + 6*n1) % 7
	n3 = n1 + n2
	if isUserCalendarGregorian {
		// Then adjust day onto the Gregorian Calendar (only valid from 1583 onwards)
		cent = year / 100
		n3 += cent - cent/4 - 2
	}
	if n3 > 40 {
		month = 5
		day = n3 - 40
	} else if n3 > 10 {
		month = 4
		day = n3 - 10
	} else {
		month = 3
		day = n3 + 21
	}
	// month is zero-indexed (0=Jan) up to this point to support this API.
	return NewDate(year, month, day)
}
func (ldp *LDP) GetModeOfWeek() int {
	if ldp.ModeOfWeekOverride > 0 {
		return ldp.ModeOfWeekOverride
	} else {
		return ldp.ModeOfWeek
	}
}

func (ldp *LDP) SetYesterday() {
	ldp.TheDayBefore = ldp.TheDay
	ldp.TheDayBefore.AddDate(0, 0, -1)
}

// Sometimes it is necessary to temporarily override the mode of the week
// It is important to clear the override after using it
func (ldp *LDP) setModeOfTheWeekOverride(mode string) {
	m, err := strconv.Atoi(mode)
	if err != nil {
		m = 0
	}
	ldp.ModeOfWeekOverride = m
}

// TODO: figure out a way to see if this works properly when one date is within daylight savings and the other is not
func DiffMillis(d1, d2 time.Time) int64 {
	diff := d1.Sub(d2)
	return diff.Nanoseconds() / 1000000
}

func (ldp *LDP) setModeOfWeek() {
	// Thomas Sunday: eight-tone cycle begins w/ Tone 1, ends Fri. 6th week Lent (day before Lazarus Sat.)
	diffMillis := DiffMillis(ldp.TheDay, ldp.ThomasSundayDate)
	// Get difference in weeks, then mod 8 to get cycle number, and add 1 to use 1-based indexes.
	ldp.ModeOfWeek = (int)((diffMillis/(7*24*60*60*1000))%8 + 1)
	if ldp.IsPentecostarion {
		// override for Pascha through the Saturday before the Sunday of Thomas
		switch ldp.DayOfSeason {
		case 1:
			{
				ldp.ModeOfWeek = 1
			}
		case 2:
			{
				ldp.ModeOfWeek = 2
			}
		case 3:
			{
				ldp.ModeOfWeek = 3
			}
		case 4:
			{
				ldp.ModeOfWeek = 4
			}
		case 5:
			{
				ldp.ModeOfWeek = 5
			}
		case 6:
			{
				ldp.ModeOfWeek = 6
			}
		case 7:
			{
				ldp.ModeOfWeek = 8
			} // note that it skips 7
		}
	}
}
func (ldp *LDP) setEothinonNumber() {
	if ldp.IsSunday {
		var diffMillis int64
		if ldp.TheDay.Before(ldp.AllSaintsDateThisYear) {
			diffMillis = DiffMillis(ldp.TheDay, ldp.AllSaintsDateLastYear)
		} else {
			diffMillis = DiffMillis(ldp.TheDay, ldp.AllSaintsDateThisYear)
		}
		ldp.EothinonNumber = (int)(diffMillis/(7*24*60*60*1000))%11 + 1
	} else {
		ldp.EothinonNumber = 0
	}
}

func (ldp *LDP) SetDayOfSeason() {
	if ldp.IsTriodion || ldp.IsPentecostarion {
		// Get difference in milliseconds
		diffMillis := DiffMillis(ldp.TheDay, ldp.TriodionStartDateThisYear)
		// Get difference in days, add 1 to be 1-index based instead of zero.
		ldp.DayOfSeason = (int)(diffMillis/(24*60*60*1000)) + 1
	} else { // movable cycle starts with day 1 of Triodion and continues through the year
		ldp.DayOfSeason = 0
	}
}

func (ldp *LDP) OverrideMovableCycleDay(d int) {
	if d == 0 {
		// zero means reset back to original day of the season
		ldp.DayOfSeason = ldp.originalDayOfSeason
		ldp.DaysSinceStartOfTriodion = ldp.originalDayOfSeason
	} else {
		// override to the specified day
		ldp.DayOfSeason = d
		ldp.DaysSinceStartOfTriodion = d
	}
}

func (ldp *LDP) setDaysSinceStartOfLastTriodion() {
	diffMillis := DiffMillis(ldp.TheDay, ldp.TriodionStartDateLast)
	// Get difference in days, add 1 to be 1-index based instead of zero.
	ldp.DaysSinceStartOfTriodion = int(diffMillis/(24*60*60*1000)) + 1
}

func (ldp *LDP) setDaysSinceSundayAfterLastElevationOfCross() {
	diffMillis := DiffMillis(ldp.TheDay, ldp.SundayAfterElevationOfCrossDateLast)
	// Get difference in days, add 1 to be 1-index based instead of zero.
	ldp.DaysSinceSundayAfterLastElevationOfCross = int(diffMillis / (24 * 60 * 60 * 1000))
}

func (ldp *LDP) setDaysSinceStartLukanCycleLast() {
	diffMillis := DiffMillis(ldp.TheDay, ldp.StartDateOfLukanCycleLast)
	// Get difference in days, add 1 to be 1-index based instead of zero.
	ldp.DaysSinceStartLastLukanCycle = (int(diffMillis / (24 * 60 * 60 * 1000))) + 1
}

// Returns the number of weeks elapsed in the Lukan cycle.
func (ldp *LDP) getWeekOfLukanCycle() int {
	if ldp.DaysSinceStartLastLukanCycle < 8 {
		return 1
	} else {
		return (ldp.DaysSinceStartLastLukanCycle / 7) + 1
	}
}

// Returns the movable day.  It will be -1 if the
// days since the start of the triodion are not between
// 69 and 128 exclusive.
func (ldp *LDP) pentecostarionDayToMovableDay() int {
	if ldp.DaysSinceStartOfTriodion > 69 && ldp.DaysSinceStartOfTriodion < 128 {
		return ldp.DaysSinceStartOfTriodion - 70
	} else {
		return -1
	}
}

func (ldp *LDP) SetDayOfWeek() {
	dow := ldp.TheDay.Weekday()
	switch dow {
	case time.Sunday:
		ldp.IsSunday = true
		ldp.DayOfWeek = "Sun"
		ldp.NbrDayOfWeek = "1"
	case time.Monday:
		ldp.IsMonday = true
		ldp.DayOfWeek = "Mon"
		ldp.NbrDayOfWeek = "2"
	case time.Tuesday:
		ldp.IsTuesday = true
		ldp.DayOfWeek = "Tue"
		ldp.NbrDayOfWeek = "3"
	case time.Wednesday:
		ldp.IsWednesday = true
		ldp.DayOfWeek = "Wed"
		ldp.NbrDayOfWeek = "4"
	case time.Thursday:
		ldp.IsThursday = true
		ldp.DayOfWeek = "Thu"
		ldp.NbrDayOfWeek = "5"
	case time.Friday:
		ldp.IsFriday = true
		ldp.DayOfWeek = "Fri"
		ldp.NbrDayOfWeek = "6"
	case time.Saturday:
		ldp.IsSaturday = true
		ldp.DayOfWeek = "Sat"
		ldp.NbrDayOfWeek = "7"
	}
}

func (ldp *LDP) ComputeSundayAfterElevationOfCross(elevationOfCross time.Time) time.Time {
	dayOffset := 0
	switch elevationOfCross.Weekday() {
	case time.Sunday:
		dayOffset = 7
	case time.Monday:
		dayOffset = 6
	case time.Tuesday:
		dayOffset = 5
	case time.Wednesday:
		dayOffset = 4
	case time.Thursday:
		dayOffset = 3
	case time.Friday:
		dayOffset = 2
	case time.Saturday:
		dayOffset = 1
	}
	return NewDate(elevationOfCross.Year(), 8, 14+dayOffset)
}

// Returns the month as an integer, such that 1 = January
func (ldp *LDP) getIntMonth() int { return int(ldp.TheDay.Month()) }

// Set the string form of the month, with a leading zero if > 10
func (ldp *LDP) setNbrMonth(month int) {
	ldp.NbrMonth = fmt.Sprintf("%02d", month)
}

// get the day of the month as an integer
func (ldp *LDP) getIntDayOfMonth() int {
	i, _ := strconv.Atoi(ldp.NbrDayOfMonth)
	return i
}

// Set the string form of the day of the month, with a leading zero if > 10
func (ldp *LDP) setNbrDayOfMonth(dayOfMonth int) {
	ldp.NbrDayOfMonth = fmt.Sprintf("%02d", dayOfMonth)
}

// If the Number day of the week has been overridden,
// returns the ldp.NbrDayOfWeekOverride
// otherwise, returns ldp.NbrDayOfWeek
func (ldp *LDP) getNbrDayOfWeek() string {
	if ldp.NbrDayOfWeekOverride == "" {
		return ldp.NbrDayOfWeek
	} else {
		return ldp.NbrDayOfWeekOverride
	}
}

func (ldp *LDP) getIntWeekOfLent() int {
	result := 0
	daysSinceStart := ldp.DaysSinceStartOfTriodion
	if daysSinceStart >= 23 && daysSinceStart <= 29 {
		result = 1
	} else if daysSinceStart >= 30 && daysSinceStart <= 36 {
		result = 2
	} else if daysSinceStart >= 37 && daysSinceStart <= 43 {
		result = 3
	} else if daysSinceStart >= 44 && daysSinceStart <= 50 {
		result = 4
	} else if daysSinceStart >= 51 && daysSinceStart <= 57 {
		result = 5
	} else if daysSinceStart >= 58 && daysSinceStart <= 64 {
		result = 6
	} else if daysSinceStart >= 65 && daysSinceStart <= 70 {
		result = 7
	}
	return result
}

func (ldp *LDP) getIntDayOfWeek() int {
	s, _ := strconv.Atoi(ldp.getNbrDayOfWeek())
	return s
}

func (ldp *LDP) setNbrDayOfWeek(intDayOfWeek int) {
	ldp.NbrDayOfWeek = fmt.Sprintf("%d", intDayOfWeek)
}

// the format if day is "D1", "D2", etc.
func (ldp *LDP) setNbrDayOfWeekOverride(day string) {
	if day == "" {
		ldp.NbrDayOfWeekOverride = ""
	} else {
		ldp.NbrDayOfWeekOverride = day[:1]
	}
}

// Because the date can be reset for an instance of this class,
// it is necessary to reset certain variables to their default value.
// Otherwise, there value can carry over erroneously to a new date.
func (ldp *LDP) SetVariablesToDefaults() {
	ldp.ModeOfWeek = 0
	ldp.ModeOfWeekOverride = 0
	ldp.EothinonNumber = 0
	ldp.DayOfSeason = 0
	ldp.DaysSinceStartOfTriodion = 0
	ldp.DaysSinceSundayAfterLastElevationOfCross = 0
	ldp.DaysSinceStartLastLukanCycle = 0
	ldp.NumberOfSundaysBeforeStartOfTriodion = 0
	ldp.IsPentecostarion = false
	ldp.IsTriodion = false
	ldp.IsPascha = false
	ldp.IsDaysOfLuke = false
	ldp.IsSunday = false
	ldp.IsMonday = false
	ldp.IsTuesday = false
	ldp.IsWednesday = false
	ldp.IsThursday = false
	ldp.IsFriday = false
	ldp.IsSaturday = false
}
// if pascha has not occurred this year, returns pascha for
// the current year.  Otherwise, returns pascha for next year
func (ldp *LDP) nextPaschaDate() time.Time {
	thisYear := ComputeDayOfPascha(ldp.TheDay.Year(), true)
	nextYear := ComputeDayOfPascha(ldp.TheDay.Year()+1, true);
	if thisYear.After(ldp.TheDay) {
		return thisYear
	} else {
		return nextYear
	}
}
func (ldp *LDP) lastPaschaDate() time.Time {
	lastYear := ComputeDayOfPascha(ldp.TheDay.Year()-1, true)
	thisYear := ComputeDayOfPascha(ldp.TheDay.Year(), true)
	if thisYear.Before(ldp.TheDay)|| thisYear.Equal(ldp.TheDay) {
		return thisYear
	} else {
		return lastYear
	}
}

func (ldp *LDP) daysInMonth(month int) int {
	if month == 2 {
		return 28
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	} else {
		return 31
	}
}
func (ldp *LDP) F() {}
