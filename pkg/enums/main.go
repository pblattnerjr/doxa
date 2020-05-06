package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	const in = `	GotActor // Actor
	GotAmpersand // @
	GotAtAllLiturgicalDayProperties // All_Liturgical_Day_Properties
	GotAtDay // day
	GotAtDayOfMonth // Day_of_Month
	GotAtDayOfMovableCycle // Day_of_Movable_Cycle
	GotAtDayOfPeriod // Day_of_Period
	GotAtDayOfWeekAsNumber // Day_of_Week_As_Number
	GotAtDayOfWeekAsText // Day_of_Week_As_Text
	GotAtEothinon // Eothinon
	GotAtLukanCycleElapsedDays // Lukan_Cycle_Elapsed_Days
	GotAtLukanCycleStartDate // Lukan_Cycle_Start_Date'
	GotAtLukanCycleWeek // Lukan_Cycle_Week
	GotAtLukanCycleWeekDay // Lukan_Cycle_Week_Day
	GotAtMode // mode
	GotAtModeOfWeek // Mode_of_Week
	GotAtNameOfPeriod // Name_of_Period
	GotAtServiceDate // Service_Date
	GotAtServiceYear // Service_Year
	GotAtSundayAfterElevationCrossDate // Sunday_After_Elevation_Cross_Date
	GotAtSundaysBeforeTriodion // Sundays_Before_Triodion
	GotAtVer // ver
	GotBreak // Break
	GotBreakTypeLine // line
	GotBreakTypePage // page
	GotBlock // bTag
	GotColon // :
	GotDay // day
	GotDialog // Dialog
	GotEndActor // End-Actor
	GotEndBreak // End_Break
	GotEndBlock // End-bTag
	GotEndDate // End_Date
	GotEndDialog // End-Dialog
	GotEndEndPageHeaderEven // End_Page_Header_Even
	GotEndEndPageHeaderOdd // End_Page_Header_Odd
	GotEndHymn // End-Hymn
	GotEndInsert // End-Insert
	GotEndMedia // End-Media
	GotEndPageFooterEven // End_Page_Footer_Even
	GotEndPageFooterOdd // End_Page_Footer_Odd
	GotEndPara // End-Para
	GotEndPassthroughHTML // END-Passthrough-Html
	GotEndPreface // End-Preface
	GotEndReading // End-Reading
	GotEndRubric // End-Rubric
	GotEndSection // End-Section
	GotEndSetMcDay // End_mcDay
	GotEndSetPageNumber // End_Set_Page_Number
	GotEndSubTitle // End-Sub-Title
	GotEndSwitchVersion // End-Switch-Version
	GotEndTemplateCommemoration // End_Template_Commemoration
	GotEndTitle // End-Title || End_Title
	GotEndVerse // End-Verse
	GotEndWhen // end-when
	GotEqualSign // =
	GotFooterColumnCenter // center
	GotFooterColumnLeft // left
	GotFooterColumnRight // right
	GotFooterCommemoration // commemoration
	GotFooterDate // date
	GotFooterLookup // lookup
	GotFooterPageNumber // pageNbr
	GotFooterTitle // title
	GotHeaderFooterText // text
	GotHymn // Hymn
	GotImport // import
	GotInsertPreface // Insert_preface - unused
	GotInsertSection // Insert_section
	GotInsertTemplate // Insert_template
	GotMedia // Media
	GotMediaOff // media-off
	GotMonth // month
	GotNid
	GotOtherwise // otherwise
	GotPageFooterEven // Page_Footer_Even
	GotPageFooterOdd // Page_Footer_Odd
	GotPageHeaderEven // Page_Header_Even
	GotPageHeaderOdd // Page_Header_Odd
	GotPara // Para
	GotPassthroughHTML // Passthrough-Html
	GotPreface // Preface
	GotReading // Reading
	GotRid // rid
	GotRole // role
	GotRubric // Rubric
	GotSection // Section
	GotSetDate // Set_Date
	GotSetMcDay // Set_mcDay
	GotSetPageNumber // Set_Page_Number
	GotSid // sid
	GotStatus // Status
	GotSubTitle // Sub-Title
	GotSwitchVersion // Switch-Version
	GotTemplateCommemoration // Template_Commemoration
	GotTemplateName // Template
	GotTemplateStatusDraft // Draft
	GotTemplateStatusFinal // Final
	GotTemplateStatusNA // NA
	GotTemplateStatusReview // Review
	GotTemplateTitle // Template_Title
	GotTitle // Title
	GotUse // use
	GotVerse // Verse
	GotWhenDateIs // when-date-is
	GotWhenExists // when-exists
	GotWhenLukanCycleDayIs // when-Lukan-Cycle-Day-is
	GotWhenModeOfWeekIs // when-mode-of-week-is
	GotWhenMovableCycleDayIs // when-movable-cycle-day-is
	GotWhenNameOfDayIs // when-name-of-day-is
	GotWhenSundayAfterElevationOfCross // when-Sunday-after-Elevation-of-Cross-is
	GotWhenSundaysBeforeTriodionIs // when-sundays-before-triodion-is
	GotYear // year
`
type Enum struct {
	Name, Value string
}
var enums []Enum
var lines = strings.Split(in, "\n")
for _, line := range lines {
	parts := strings.Split(line, "//")
	if len(parts) == 2 {
		enums = append(enums, Enum{strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])})
	}
}

const SWITCH = `
	switch token {
	{{range .}}
		case "{{.Value}}": currentState = {{.Name}}
    {{end}}
        default: {
			switch currentState {
				{{range .}}
				case {{.Name}}: 
					{
						fmt.Printf("Unhandled: %s\n",{{.Name}}.stateName())
					}
				{{end}}
            }
        }
    }
`
t := template.Must(template.New("letter").Parse(SWITCH))
err := t.Execute(os.Stdout, enums)
if err != nil {
	fmt.Println(err)
}
const NAMES = `
	switch s {
	{{range .}}
	case {{.Name}}:
		return "{{.Name}}"
	{{end}}
    case default:
		return "UNKNOWN"
    }
`
	t2 := template.Must(template.New("letter").Parse(NAMES))
	err = t2.Execute(os.Stdout, enums)
	if err != nil {
		fmt.Println(err)
	}
}
