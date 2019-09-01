package ldp

import (
	"fmt"
	"testing"
	"time"
)

type Data struct {
	testYear      int
	expectedYear  int
	expectedMonth time.Month
	expectedDay   int
}

var paschaDates = []Data{
	{2019, 2019, time.April, 28},
	{2020, 2020, time.April, 19},
	{2021, 2021, time.May, 2},
	{2022, 2022, time.April, 24},
	{2023, 2023, time.April, 16},
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
var pentecostDates = []Data{
	{2019, 2019, time.June, 16},
	{2020, 2020, time.June, 7},
	{2021, 2021, time.June, 20},
	{2022, 2022, time.June, 12},
	{2023, 2023, time.June, 4},
}

func TestNewLDP(t *testing.T) {
	for i, d := range paschaDates {
		ldp, err := NewLDPYMD(d.testYear, 9, 31)
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

// Test for new LDP where we do not specify the year.
// We test two conditions:
// 1) the date is not before time.Now
// 2) the date is before time.Now, in which case the year is incremented by 1
func TestNewLDPForMonthAndDay(t *testing.T) {
	ldp, err := NewLDPMD(10, 23)
	if err != nil {t.Error(err.Error())}
	if ldp.StrYear() != "2019" {
		t.Error(fmt.Sprintf("Expected 2019, got %s", ldp.StrYear()))
	}
	if ldp.StrMonth() != "10" {
		t.Error(fmt.Sprintf("Expected month 10, got %s", ldp.StrMonth()))
	}
	if ldp.StrDay() != "23" {
		t.Error(fmt.Sprintf("Expected day 23, got %s", ldp.StrDay()))
	}
	ldp, err = NewLDPMD(01, 23)
	if err != nil {t.Error(err.Error())}
	if ldp.StrYear() != "2020" {
		t.Error(fmt.Sprintf("Expected 2020, got %s", ldp.StrYear()))
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
		p := ComputeDayOfPascha(d.testYear, true)
		if p.Month() != d.expectedMonth {
			t.Error(fmt.Sprintf("Expected month %s, got %s", d.expectedMonth, p.Month()))
		}
		if p.Day() != d.expectedDay {
			t.Error(fmt.Sprintf("Expected day %d, got %d", d.expectedDay, p.Day()))
		}
	}
}

func TestNewDate(t *testing.T) {
	d := NewDate(2019, 4, 21)
	if d.Month() != time.April {
		t.Error(fmt.Sprintf("Expected month %s, got %s", time.April, d.Month()))
	}
	if d.Day() != 21 {
		t.Error(fmt.Sprintf("Expected day %d, got %d", 21, d.Day()))
	}
}
