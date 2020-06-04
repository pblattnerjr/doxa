package ldp

import (
	"fmt"
	"github.com/liturgiko/doxa/pkg/enums/calendarTypes"
	"testing"
	"time"
)

type Data struct {
	testYear      int
	expectedYear  int
	expectedMonth time.Month
	expectedDay   int
}

var triodionStartDates = []Data{
	{2019, 2019, time.February, 17},
	{2020, 2020, time.February, 9},
	{2021, 2021, time.February, 21},
	{2022, 2022, time.February, 13},
	{2023, 2023, time.February, 5},
}
var palmSundayDates = []Data{
	{2019, 2019, time.April, 21},
	{2020, 2020, time.April, 12},
	{2021, 2021, time.May, 25},
	{2022, 2022, time.April, 17},
	{2023, 2023, time.April, 9},
}
var paschaDates = []Data{
	{2019, 2019, time.April, 28},
	{2020, 2020, time.April, 19},
	{2021, 2021, time.May, 2},
	{2022, 2022, time.April, 24},
	{2023, 2023, time.April, 16},
}
var thomasSundayDates = []Data{
	{2019, 2018, time.April, 15},
	{2020, 2019, time.May, 5},
	{2021, 2020, time.April, 26},
	{2022, 2021, time.May, 9},
	{2023, 2022, time.May, 1},
}
var pentecostDates = []Data{
	{2019, 2019, time.June, 16},
	{2020, 2020, time.June, 7},
	{2021, 2021, time.June, 20},
	{2022, 2022, time.June, 12},
	{2023, 2023, time.June, 4},
}

type TopicData struct {
	testYear      int
	testMonth     int
	testDay       int
	topic         string
	expectedTopic string
}

var topics = []TopicData{
	// test days since Sunday After Last Elevation of the Cross
	{2007, 9, 16, "le.go.lu.*", "le.go.lu.d364"},
	// mode tests
	{2020, 3, 21, "oc.*", "oc.m6.d7"},
	{2020, 3, 22, "oc.*", "oc.m7.d1"},
	{2020, 3, 23, "oc.*", "oc.m7.d2"},
	// eothinon tests
	{2020, 3, 22, "le.go.eo.*", "le.go.eo.w07"},
	// calendar month and day tests
	{2020, 3, 22, "me.*", "me.m03.d22"},
}

type ModeData struct {
	testYear     int
	testMonth    int
	testDay      int
	expectedMode int
}

/*
Mode 1
     2
     3
     4
     5 aka Plagal of the First
     6 aka Plagal of the Second
     7 aka Grave
     8 aka Plagal of the Fourth
*/
var modes = []ModeData{
	{2020, 3, 21, 6},
	{2020, 3, 22, 7},
	{2020, 3, 23, 7},
}

// Test liturgical dates from the perspective of a
// date that occurs before the start of the triodion
func TestNewLDPBeforeTriodion(t *testing.T) {
	for i, d := range paschaDates {
		ldp, err := NewLDPYMD(d.testYear, 1, 6, calendarTypes.Gregorian)
		if err != nil {
			t.Error(err.Error())
		}
		p := ldp.TriodionStartDateThisYear
		if p.Month() != triodionStartDates[i].expectedMonth {
			t.Error(fmt.Sprintf("triodion start date: expected month %s, got %s", triodionStartDates[i].expectedMonth, p.Month()))
		}
		if p.Day() != triodionStartDates[i].expectedDay {
			t.Error(fmt.Sprintf("triodion start date: expected day %d, got %d", triodionStartDates[i].expectedDay, p.Day()))
		}
		p = ldp.PaschaDateThisYear
		if p.Month() != d.expectedMonth {
			t.Error(fmt.Sprintf("pascha starte date: expected month %s, got %s", d.expectedMonth, p.Month()))
		}
		if p.Day() != d.expectedDay {
			t.Error(fmt.Sprintf("pascha starte date: expected day %d, got %d", d.expectedDay, p.Day()))
		}
		p = ldp.PentecostDate
		if p.Month() != pentecostDates[i].expectedMonth {
			t.Error(fmt.Sprintf("pascha starte date: expected month %s, got %s", pentecostDates[i].expectedMonth, p.Month()))
		}
		if p.Day() != pentecostDates[i].expectedDay {
			t.Error(fmt.Sprintf("pascha starte date: expected day %d, got %d", pentecostDates[i].expectedDay, p.Day()))
		}
	}
}

// Test liturgical dates from the perspective of a
// date that occurs after the end of the triodion
func TestNewLDPAfterTriodion(t *testing.T) {
	for i, d := range paschaDates {
		ldp, err := NewLDPYMD(d.testYear, 1, 6, calendarTypes.Gregorian)
		if err != nil {
			t.Error(err.Error())
		}
		p := ldp.TriodionStartDateThisYear
		if p.Month() != triodionStartDates[i].expectedMonth {
			t.Error(fmt.Sprintf("triodion start date: expected month %s, got %s", triodionStartDates[i].expectedMonth, p.Month()))
		}
		if p.Day() != triodionStartDates[i].expectedDay {
			t.Error(fmt.Sprintf("triodion start date: expected day %d, got %d", triodionStartDates[i].expectedDay, p.Day()))
		}
		p = ldp.PaschaDateThisYear
		if p.Month() != d.expectedMonth {
			t.Error(fmt.Sprintf("pascha starte date: expected month %s, got %s", d.expectedMonth, p.Month()))
		}
		if p.Day() != d.expectedDay {
			t.Error(fmt.Sprintf("pascha starte date: expected day %d, got %d", d.expectedDay, p.Day()))
		}
		p = ldp.ThomasSundayDate
		if p.Month() != thomasSundayDates[i].expectedMonth {
			t.Error(fmt.Sprintf("Sunday of Thomas %d: expected month %s, got %s", d.expectedYear, thomasSundayDates[i].expectedMonth, p.Month()))
		}
		if p.Day() != thomasSundayDates[i].expectedDay {
			t.Error(fmt.Sprintf("Sunday of Thomas %d: expected day %d, got %d", d.expectedYear, thomasSundayDates[i].expectedDay, p.Day()))
		}
		p = ldp.PentecostDate
		if p.Month() != pentecostDates[i].expectedMonth {
			t.Error(fmt.Sprintf("pascha starte date: expected month %s, got %s", pentecostDates[i].expectedMonth, p.Month()))
		}
		if p.Day() != pentecostDates[i].expectedDay {
			t.Error(fmt.Sprintf("pascha starte date: expected day %d, got %d", pentecostDates[i].expectedDay, p.Day()))
		}
	}
}

// Test for new LDP where we do not specify the year.
// We test two conditions:
// 1) the date is not before time.Now
// 2) the date is before time.Now, in which case the year is incremented by 1
func TestNewLDPForMonthAndDay(t *testing.T) {
	now := time.Now()
	ldp, err := NewLDPMD(10, 23, calendarTypes.Gregorian)
	if err != nil {
		t.Error(err.Error())
	}
	if ldp.TheDay.Year() != now.Year() {
		t.Error(fmt.Sprintf("Expected %d, got %d", now.Year(), ldp.TheDay.Year()))
	}
	if ldp.StrMonth() != "10" {
		t.Error(fmt.Sprintf("Expected month 10, got %s", ldp.StrMonth()))
	}
	if ldp.StrDay() != "23" {
		t.Error(fmt.Sprintf("Expected day 23, got %s", ldp.StrDay()))
	}
	ldp, err = NewLDPMD(01, 23, calendarTypes.Gregorian)
	if err != nil {
		t.Error(err.Error())
	}
	if ldp.TheDay.Year() != now.Year()+1 {
		t.Error(fmt.Sprintf("Expected %d, got %d", now.Year()+1, ldp.TheDay.Year()))
	}
	if ldp.StrMonth() != "1" {
		t.Error(fmt.Sprintf("Expected month 1, got %s", ldp.StrMonth()))
	}
	if ldp.StrDay() != "23" {
		t.Error(fmt.Sprintf("Expected day 23, got %s", ldp.StrDay()))
	}
}

func TestComputeDayOfPascha(t *testing.T) {
	for _, d := range paschaDates {
		p := ComputeDayOfPascha(d.testYear, calendarTypes.Gregorian)
		if p.Month() != d.expectedMonth {
			t.Error(fmt.Sprintf("Expected month %s, got %s", d.expectedMonth, p.Month()))
		}
		if p.Day() != d.expectedDay {
			t.Error(fmt.Sprintf("Expected day %d, got %d", d.expectedDay, p.Day()))
		}
	}
}

type DateData struct {
	Year, Month, Day int
	ExpectedWeekDay time.Weekday
}
func TestNewDate(t *testing.T) {
	var data = []DateData{
		{2006,9,14,time.Thursday},
		{2007,9,14,time.Friday},
		{2008,9,14,time.Sunday},
		{2009,9,14,time.Monday},
		{2010,9,14, time.Tuesday},
		{2011,9,14, time.Wednesday},
		{2012,9,14, time.Friday},
		{2013,9,14, time.Saturday},
		{2014,9,14, time.Sunday},
		{2015,9,14, time.Monday},
		{2016,9,14, time.Wednesday},
		{2017,9,14, time.Thursday},
		{2018,9,14, time.Friday},
		{2019,9,14, time.Saturday},
		{2020,9,14, time.Monday},
		{2021,9,14, time.Tuesday},
		{2022,9,14, time.Wednesday},
		{2023,9,14, time.Thursday},
		{2024,9,14, time.Saturday},
		{2025,9,14, time.Sunday},
	}
	for i, td := range data {
		d := NewDate(td.Year, td.Month, td.Day)
		if d.Weekday() != td.ExpectedWeekDay {
			t.Error(fmt.Sprintf("case %d: expected weekday %s, got %s", i, td.ExpectedWeekDay, d.Weekday()))
		}
	}
}
func TestLDP_RelativeTopic(t *testing.T) {
	for _, d := range topics {
		ldp, err := NewLDPYMD(d.testYear, d.testMonth, d.testDay, calendarTypes.Gregorian)
		if err != nil {
			t.Error(err.Error())
		}
		rt := ldp.RelativeTopic(d.topic)
		if rt != d.expectedTopic {
			t.Error(fmt.Sprintf("expected %s, got %s", d.expectedTopic, rt))
		}
	}
}
func TestLDP_GetModeOfWeek(t *testing.T) {
	for _, d := range modes {
		ldp, err := NewLDPYMD(d.testYear, d.testMonth, d.testDay, calendarTypes.Gregorian)
		if err != nil {
			t.Error(err.Error())
		}
		mode := ldp.ModeOfWeek
		if mode != d.expectedMode {
			t.Error(fmt.Sprintf("expected %d, got %d", d.expectedMode, mode))
		}
	}
}
func TestElevationOfCross(t *testing.T) {
	var data = []ElevationData{
		{2007,9,15,
			2006,9,14,
			2006,9,17,
			2006,9,18,
			363,363,
			"Sat", 52},
		{2007,9,16,
			2007,9,14,
			2007,9,16,
			2006,9,18,
			0,364,
			"Sun", 53},
		{2007,9,17,
			2007,9,14,
			2007,9,16,
			2007,9,17,
			1,1,
			"Mon", 1},
		{2008,9,20,
			2007,9,14,
			2007,9,16,
			2007,9,17,
			370,370,
			"Sat", 53},
		{2008,9,21,
			2008,9,14,
			2008,9,21,
			2007,9,17,
			0,371,
			"Sun", 54},
		{2008,9,22,
			2008,9,14,
			2008,9,21,
			2008,9,22 ,
			1,1,
			"Mon", 1},
		{2009,9,19,
			2008,9,14,
			2008,9,21,
			2008,9,22 ,
			363,363,
			"Sat", 52},
		{2009,9,20,
			2009,9,14,
			2009,9,20,
			2008,9,22 ,
			0,364,
			"Sun", 53},
		{2009,9,21,
			2009,9,14,
			2009,9,20,
			2009,9,21 ,
			1,1,
			"Mon", 1},
		{2010,9,18,
			2009,9,14,
			2009,9,20,
			2009,9,21 ,
			363,363,
			"Sat", 52},
		{2010,9,19,
			2010,9,14,
			2010,9,19,
			2009,9,21 ,
			0,364,
			"Sun", 53},
		{2010,9,20,
			2010,9,14,
			2010,9,19,
			2010,9,20,
			1,1,
			"Mon", 1},
		{2011,9,17,
			2010,9,14,
			2010,9,19,
			2010,9,20,
			363,363,
			"Sat", 52},
		{2011,9,18,
			2011,9,14,
			2011,9,18,
			2010,9,20,
			0,364,
			"Sun", 53},
		{2011,9,19,
			2011,9,14,
			2011,9,18,
			2011,9,19,
			1,1,
			"Mon", 1},
		{2012,9,15,
			2011,9,14,
			2011,9,18,
			2011,9,19,
			363,363,
			"Sat", 52},
		{2012,9,16,
			2012,9,14,
			2012,9,16,
			2011,9,19,
			0,364,
			"Sun", 53},
		{2012,9,17,
			2012,9,14,
			2012,9,16,
			2012,9,17,
			1,1,
			"Mon", 1},
	}
	for i, d := range data {
		ldp, err := NewLDPYMD(d.LiturgicalYear, d.LiturgicalMonth, d.LiturgicalDay, calendarTypes.Gregorian)
		if err != nil {
			t.Error(err.Error())
		}
		e := ldp.NewElevationData()
		fmt.Println(e.String())
		if d.LiturgicalYear != e.LiturgicalYear {
			t.Error(fmt.Sprintf("case %d: expected %d for LiturgicalYear, got %d", i, d.LiturgicalYear, e.LiturgicalYear))
		}
		if d.LiturgicalMonth != e.LiturgicalMonth {
			t.Error(fmt.Sprintf("case %d: expected %d for LiturgicalMonth, got %d", i, d.LiturgicalMonth, e.LiturgicalMonth))
		}
		if d.LiturgicalDay != e.LiturgicalDay {
			t.Error(fmt.Sprintf("case %d: expected %d for LiturgicalDay, got %d", i, d.LiturgicalDay, e.LiturgicalDay))
		}
		// elevation date
		if d.ElevationYear != e.ElevationYear {
			t.Error(fmt.Sprintf("case %d: expected %d for ElevationYear, got %d", i, d.ElevationYear, e.ElevationYear))
		}
		if d.ElevationMonth != e.ElevationMonth {
			t.Error(fmt.Sprintf("case %d: expected %d for ElevationMonth, got %d", i, d.ElevationMonth, e.ElevationMonth))
		}
		if d.ElevationDay != e.ElevationDay {
			t.Error(fmt.Sprintf("case %d: expected %d for ElevationDay, got %d", i, d.ElevationDay, e.ElevationDay))
		}
		// Sunday after elevation date
		if d.SundayAfterYear != e.SundayAfterYear {
			t.Error(fmt.Sprintf("case %d: expected %d for SundayAfterYear, got %d", i, d.SundayAfterYear, e.SundayAfterYear))
		}
		if d.SundayAfterMonth != e.SundayAfterMonth {
			t.Error(fmt.Sprintf("case %d: expected %d for SundayAfterMonth, got %d", i, d.SundayAfterMonth, e.SundayAfterMonth))
		}
		if d.SundayAfterDay != e.SundayAfterDay {
			t.Error(fmt.Sprintf("case %d: expected %d for SundayAfterDay, got %d", i, d.SundayAfterDay, e.SundayAfterDay))
		}
		// Lukan Cycle start date
		if d.LukanCycleStartYear != e.LukanCycleStartYear {
			t.Error(fmt.Sprintf("case %d: expected %d for LukanCycleStartYear, got %d", i, d.LukanCycleStartYear, e.LukanCycleStartYear))
		}
		if d.LukanCycleStartMonth != e.LukanCycleStartMonth {
			t.Error(fmt.Sprintf("case %d: expected %d for LukanCycleStartMonth, got %d", i, d.LukanCycleStartMonth, e.LukanCycleStartMonth))
		}
		if d.LukanCycleStartDay != e.LukanCycleStartDay {
			t.Error(fmt.Sprintf("case %d: expected %d for LukanCycleStartDay, got %d", i, d.LukanCycleStartDay, e.LukanCycleStartDay))
		}
		if d.ElapsedDays != e.ElapsedDays {
			t.Error(fmt.Sprintf("case %d: expected %d for ElapsedDays, got %d", i, d.ElapsedDays, e.ElapsedDays))
		}
		if d.LukanCycleDayNbr != e.LukanCycleDayNbr {
			t.Error(fmt.Sprintf("case %d: expected %d for LukanCycleDayNbr, got %d", i, d.LukanCycleDayNbr, e.LukanCycleDayNbr))
		}
		if d.LukanCycleDayName != e.LukanCycleDayName {
			t.Error(fmt.Sprintf("case %d: expected %s for LukanCycleDayName, got %s", i, d.LukanCycleDayName, e.LukanCycleDayName))
		}
		if d.LukanCycleWeekNbr != e.LukanCycleWeekNbr {
			t.Error(fmt.Sprintf("case %d: expected %d for LukanCycleWeekNbr, got %d", i, d.LukanCycleWeekNbr, e.LukanCycleWeekNbr))
		}
	}
}
