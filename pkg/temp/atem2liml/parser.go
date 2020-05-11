package atem2lml

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"github.com/liturgiko/doxa/pkg/utils/stack"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/scanner"
)

type LookupMethod int

const (
	NID LookupMethod = iota // No ID
	RID                     // Relative ID
	SID                     // Specific ID
)

type Status int

const (
	NA Status = iota // No Status
	DRAFT
	REVIEW
	FINAL
)

type Type int

const (
	BOOK Type = iota
	BLOCK
	SERVICE
)

type TokenError struct {
	Token    string
	Filename string
	Line     int
	Column   int
	Error    error
}

func (t *TokenError) Set(s *scanner.Scanner, msg string) {
	currentState = ERROR
	t.Token = s.TokenText()
	t.Filename = s.Filename
	t.Line = s.Line
	t.Column = s.Column
	t.Error = errors.New(msg)

}
func (t *TokenError) String() string {
	return fmt.Sprintf("%s:%d:%d - %s - error: %s\n", t.Filename, t.Line, t.Column, t.Token, t.Error)
}

type Directive struct {
	Name LookupMethod
	Arg  string
}

type Span struct {
	Class      string
	Id         string
	Value      string
	Directives []Directive
	Error      TokenError
}

func (s *Span) addDirective(d Directive) {
	s.Directives = append(s.Directives, d)
}

func newSpan() Span {
	var span Span
	span.Class = "Bk"
	return span
}

type Row struct {
	Cells []Cell
}

func (r *Row) addCell(c Cell) {
	r.Cells = append(r.Cells, c)
}

type Cell struct {
	Class string
	Col   string
	Spans []Span
}

func (c *Cell) addSpans(s Span, libraries []string) {
	var sb strings.Builder
	for _, library := range libraries {
		var span Span
		span.Class = s.Class

		sb.WriteString(library)
		sb.WriteString("/")
		sb.WriteString(s.Id)
		span.Id = sb.String()
		sb.Reset()

		span.Value = getValue(span.Id)
		span.Error = s.Error
		span.Directives = s.Directives
		c.Spans = append(c.Spans, span)
	}
}

type Align int

const (
	Left Align = iota
	Center
	Right
)

type HeaderFooter struct {
	Alignment Align
}
type PageType int

const (
	Even PageType = iota
	Odd
)

type Header struct {
	PageType PageType
}

// After parsing a file, the data will all be held here
type MetaTemplate struct {
	TemplateDir string
	ID          string
	Imports     []string
	Type        Type
	Status      Status
	Title       string
	Day         int
	Month       int
	Year        int
	Rows        []Row
	FilePath    string
	File        *os.File
	Writer      *bufio.Writer
}

func (m *MetaTemplate) addImport(i string) {
	m.Imports = append(m.Imports, i)
}

func (m *MetaTemplate) addRow(r Row) {
	m.Rows = append(m.Rows, r)
}
func (m *MetaTemplate) setFilePath() error {
	var err error
	m.FilePath = m.TemplateDir + m.ID + ".lml"
	err = ltfile.CreateDirs(filepath.Dir(m.FilePath))
	return err
}
func (m *MetaTemplate) setWriter() {
	m.Writer = bufio.NewWriter(m.File)
}

var mt MetaTemplate

const pathSeparator = string(os.PathSeparator)

func ParseTemplate(fileIn, dirOut, id string, sectionMap map[string]string) error {
	// get the name of the file
	filename := filepath.Base(fileIn)
	filename = filename[0 : len(filename)-5]
	// read file in as a string
	content, err := ioutil.ReadFile(fileIn)
	if err != nil {
		fmt.Println(err.Error())
	}
	text := string(content)
	// Problem: text/scanner treats hyphen as a word separator.  Change to underscore to get around this.
	text = strings.ReplaceAll(text, "-", "_")

	// set up stack for templates.  We will create one for each section encountered
	templates := stack.New()
	var mt MetaTemplate
	mt.ID = id
	mt.TemplateDir = dirOut
	mt.Type = BOOK // we will override this if encounter a Set_Date

	mt.setFilePath()
	if mt.File, err = os.Create(mt.FilePath); err != nil {
		return err
	}
	defer mt.File.Close()
	mt.setWriter()

	// set up stack for keeping track of internal section paths
	// we need this in order to resolve insert_section that refers to an internal section in this file
	internalPaths := stack.New()
	currentInternalPath := filename

	inHead := false
	importSb := strings.Builder{}

	var s scanner.Scanner
	s.Init(strings.NewReader(text))
	s.Whitespace = 1<<':' | 1<<'-' | 1<<'\t' | 1<<'\n' | 1<<' ' // treat these as white space
	lineSB := strings.Builder{}
	currentState = LineStart

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		token := s.TokenText()
		switch token {
		case "\r":
			continue
		case "Actor":
			currentState = GotActor
		case "@":
			currentState = GotAmpersand
		case "All_Liturgical_Day_Properties":
			currentState = GotAtAllLiturgicalDayProperties
		case "day":
			{
				if currentState == GotAmpersand {
					currentState = GotAtDay
				} else {
					currentState = GotDay
				}
			}
		case "Day_of_Month":
			currentState = GotAtDayOfMonth
		case "Day_of_Movable_Cycle":
			currentState = GotAtDayOfMovableCycle
		case "Day_of_Period":
			currentState = GotAtDayOfPeriod
		case "Day_of_Week_As_Number":
			currentState = GotAtDayOfWeekAsNumber
		case "Day_of_Week_As_Text":
			currentState = GotAtDayOfWeekAsText
		case "Eothinon":
			currentState = GotAtEothinon
		case "Lukan_Cycle_Elapsed_Days":
			currentState = GotAtLukanCycleElapsedDays
		case "Lukan_Cycle_Start_Date'":
			currentState = GotAtLukanCycleStartDate
		case "Lukan_Cycle_Week":
			currentState = GotAtLukanCycleWeek
		case "Lukan_Cycle_Week_Day":
			currentState = GotAtLukanCycleWeekDay
		case "mode":
			currentState = GotAtMode
		case "Mode_of_Week":
			currentState = GotAtModeOfWeek
		case "Name_of_Period":
			currentState = GotAtNameOfPeriod
		case "Service_Date":
			currentState = GotAtServiceDate
		case "Service_Year":
			currentState = GotAtServiceYear
		case "Sunday_After_Elevation_Cross_Date":
			currentState = GotAtSundayAfterElevationCrossDate
		case "Sundays_Before_Triodion":
			currentState = GotAtSundaysBeforeTriodion
		case "ver":
			currentState = GotAtVer
		case "Break":
			currentState = GotBreak
		case "line":
			currentState = GotBreakTypeLine
		case "page":
			currentState = GotBreakTypePage
		case "bTag":
			currentState = GotBlock
		case ":":
			currentState = GotColon
		case "Dialog":
			currentState = GotDialog
		case "End_Actor":
			currentState = GotEndActor
		case "End_Break":
			currentState = GotEndBreak
		case "End_bTag":
			currentState = GotEndBlock
		case "End_Date":
			currentState = GotEndDate
		case "End_Dialog":
			currentState = GotEndDialog
		case "End_Page_Header_Even":
			currentState = GotEndEndPageHeaderEven
		case "End_Page_Header_Odd":
			currentState = GotEndEndPageHeaderOdd
		case "End_Head":
			currentState = GotEndHead
			inHead = false
		case "End_Hymn":
			currentState = GotEndHymn
		case "End_Insert":
			sectionId := lineSB.String()
			resolvedPath := resolveSectionPath(sectionId, currentInternalPath, sectionMap)
			if len(resolvedPath) > 0 {
				mt.Writer.WriteString(fmt.Sprintf("insert %s\n", resolvedPath))
				mt.Writer.Flush()
				lineSB.Reset()
				currentState = Neutral
			} else {
				fmt.Printf("Could not resolve section path %s\n",sectionId)
				currentState = Neutral
			}
		case "End_Media":
			currentState = GotEndMedia
		case "End_Page_Footer_Even":
			currentState = GotEndPageFooterEven
		case "End_Page_Footer_Odd":
			currentState = GotEndPageFooterOdd
		case "End_Para":
			currentState = GotEndPara
		case "END-Passthrough-Html":
			currentState = GotEndPassthroughHTML
		case "End_Preface":
			currentState = GotEndPreface
		case "End_Reading":
			currentState = GotEndReading
		case "End_Rubric":
			currentState = GotEndRubric
		case "End_Section":
			if mt.Writer == nil || mt.File == nil {
				fmt.Println(fileIn)
				fmt.Println(templates.Peek().(MetaTemplate).ID)
			}
			mt.Writer.Flush()
			mt.File.Close()
			mt = templates.Pop().(MetaTemplate)
			currentInternalPath = internalPaths.Pop().(string)
			currentState = Neutral
		case "End_mcDay":
			currentState = GotEndSetMcDay
		case "End_Set_Page_Number":
			currentState = GotEndSetPageNumber
		case "End_Sub-Title":
			currentState = GotEndSubTitle
		case "End_Switch-Version":
			currentState = GotEndSwitchVersion
		case "End_Template":
			mt.Writer.Flush()
			mt.File.Close()
		case "End_Template_Commemoration":
			currentState = GotEndTemplateCommemoration
		case "End_Title || End_Title":
			currentState = GotEndTitle
		case "End_Verse":
			currentState = GotEndVerse
		case "end-when":
			currentState = GotEndWhen
		case "=":
			currentState = GotEqualSign
		case "center":
			currentState = GotFooterColumnCenter
		case "left":
			currentState = GotFooterColumnLeft
		case "right":
			currentState = GotFooterColumnRight
		case "commemoration":
			if currentState == GotSection {
				mt.ID = mt.ID + pathSeparator + token
				mt.TemplateDir = dirOut
				currentState = Neutral
				mt.setFilePath()
				if mt.File, err = os.Create(mt.FilePath); err != nil {
					return err
				}
				mt.setWriter()
				mt.Writer.WriteString(fmt.Sprintf("id: %s\n", mt.ID))
				internalPaths.Push(currentInternalPath)
				currentInternalPath = currentInternalPath + "." + token
				currentState = Neutral
			} else {
				currentState = GotFooterCommemoration
			}
		case "date":
			currentState = GotFooterDate
		case "lookup":
			currentState = GotFooterLookup
		case "pageNbr":
			currentState = GotFooterPageNumber
		case "title":
			if currentState == GotSection {
				mt.ID = mt.ID + pathSeparator + token
				mt.TemplateDir = dirOut
				currentState = Neutral
				mt.setFilePath()
				if mt.File, err = os.Create(mt.FilePath); err != nil {
					return err
				}
				mt.setWriter()
				mt.Writer.WriteString(fmt.Sprintf("id: %s\n", mt.ID))
				internalPaths.Push(currentInternalPath)
				currentInternalPath = currentInternalPath + "." + token
				currentState = Neutral
			} else {
				currentState = GotFooterTitle
			}
		case "text":
			if currentState == GotSection {
				mt.ID = mt.ID + pathSeparator + token
				mt.TemplateDir = dirOut
				currentState = Neutral
				mt.setFilePath()
				if mt.File, err = os.Create(mt.FilePath); err != nil {
					return err
				}
				mt.setWriter()
				mt.Writer.WriteString(fmt.Sprintf("id: %s\n", mt.ID))
				internalPaths.Push(currentInternalPath)
				currentInternalPath = currentInternalPath + "." + token
				currentState = Neutral
			} else {
				currentState = GotHeaderFooterText
			}
		case "Head":
			currentState = GotHead
			inHead = true
		case "Hymn":
			currentState = GotHymn
		case "import":
			currentState = GotImport
		case "Insert_preface - unused":
			currentState = GotInsertPreface
		case "Insert_section":
			currentState = GotInsertSection
		case "Insert_template":
			currentState = GotInsertTemplate
		case "Media":
			currentState = GotMedia
		case "media-off":
			currentState = GotMediaOff
		case "month":
			currentState = GotMonth
		case "otherwise":
			currentState = GotOtherwise
		case "Page_Footer_Even":
			currentState = GotPageFooterEven
		case "Page_Footer_Odd":
			currentState = GotPageFooterOdd
		case "Page_Header_Even":
			currentState = GotPageHeaderEven
		case "Page_Header_Odd":
			currentState = GotPageHeaderOdd
		case "Para":
			currentState = GotPara
		case "Passthrough-Html":
			currentState = GotPassthroughHTML
		case "Preface":
			currentState = GotPreface
		case "Reading":
			currentState = GotReading
		case "rid":
			currentState = GotRid
		case "role":
			currentState = GotRole
		case "Rubric":
			currentState = GotRubric
		case "Section":
			currentState = GotSection
			tmpMt := MetaTemplate{}
			tmpMt.TemplateDir = dirOut
			tmpMt.ID = mt.ID // once we get the section name, we will append it to ID
			tmpMt.Type = BLOCK
			templates.Push(mt)
			mt = tmpMt
		case "Set_Date":
			currentState = GotSetDate
			mt.Type = SERVICE
		case "Set_mcDay":
			currentState = GotSetMcDay
		case "Set_Page_Number":
			currentState = GotSetPageNumber
		case "sid":
			currentState = GotSid
		case "Status":
			currentState = GotStatus
		case "Sub-Title":
			currentState = GotSubTitle
		case "Switch-Version":
			currentState = GotSwitchVersion
		case "Template_Commemoration":
			currentState = GotTemplateCommemoration
		case "Template":
			currentState = GotTemplateName
		case "Draft":
			mt.Status = DRAFT
			currentState = Neutral
		case "Final":
			mt.Status = FINAL
			currentState = Neutral
		case "NA":
			mt.Status = NA
			currentState = Neutral
		case "Review":
			mt.Status = REVIEW
			currentState = Neutral
		case "Template_Title":
			currentState = GotTemplateTitle
		case "Title":
			currentState = GotTitle
		case "use":
			currentState = GotUse
		case "Verse":
			currentState = GotVerse
		case "when-date-is":
			currentState = GotWhenDateIs
		case "when-exists":
			currentState = GotWhenExists
		case "when-Lukan-Cycle-Day-is":
			currentState = GotWhenLukanCycleDayIs
		case "when-mode-of-week-is":
			currentState = GotWhenModeOfWeekIs
		case "when-movable-cycle-day-is":
			currentState = GotWhenMovableCycleDayIs
		case "when-name-of-day-is":
			currentState = GotWhenNameOfDayIs
		case "when-Sunday-after-Elevation-of-Cross-is":
			currentState = GotWhenSundayAfterElevationOfCross
		case "when-sundays-before-triodion-is":
			currentState = GotWhenSundaysBeforeTriodionIs
		case "year":
			currentState = GotYear
		default:
			{
				switch currentState {
				case GotActor:
					{
						fmt.Printf("Unhandled: %s\n", GotActor.stateName())
					}
				case GotAmpersand:
					{
						fmt.Printf("Unhandled: %s\n", GotAmpersand.stateName())
					}
				case GotAtAllLiturgicalDayProperties:
					{
						fmt.Printf("Unhandled: %s\n", GotAtAllLiturgicalDayProperties.stateName())
					}
				case GotAtDay:
					{
						fmt.Printf("Unhandled: %s\n", GotAtDay.stateName())
					}
				case GotAtDayOfMonth:
					{
						fmt.Printf("Unhandled: %s\n", GotAtDayOfMonth.stateName())
					}
				case GotAtDayOfMovableCycle:
					{
						fmt.Printf("Unhandled: %s\n", GotAtDayOfMovableCycle.stateName())
					}
				case GotAtDayOfPeriod:
					{
						fmt.Printf("Unhandled: %s\n", GotAtDayOfPeriod.stateName())
					}
				case GotAtDayOfWeekAsNumber:
					{
						fmt.Printf("Unhandled: %s\n", GotAtDayOfWeekAsNumber.stateName())
					}
				case GotAtDayOfWeekAsText:
					{
						fmt.Printf("Unhandled: %s\n", GotAtDayOfWeekAsText.stateName())
					}
				case GotAtEothinon:
					{
						fmt.Printf("Unhandled: %s\n", GotAtEothinon.stateName())
					}
				case GotAtLukanCycleElapsedDays:
					{
						fmt.Printf("Unhandled: %s\n", GotAtLukanCycleElapsedDays.stateName())
					}
				case GotAtLukanCycleStartDate:
					{
						fmt.Printf("Unhandled: %s\n", GotAtLukanCycleStartDate.stateName())
					}
				case GotAtLukanCycleWeek:
					{
						fmt.Printf("Unhandled: %s\n", GotAtLukanCycleWeek.stateName())
					}
				case GotAtLukanCycleWeekDay:
					{
						fmt.Printf("Unhandled: %s\n", GotAtLukanCycleWeekDay.stateName())
					}
				case GotAtMode:
					{
						fmt.Printf("Unhandled: %s\n", GotAtMode.stateName())
					}
				case GotAtModeOfWeek:
					{
						fmt.Printf("Unhandled: %s\n", GotAtModeOfWeek.stateName())
					}
				case GotAtNameOfPeriod:
					{
						fmt.Printf("Unhandled: %s\n", GotAtNameOfPeriod.stateName())
					}
				case GotAtServiceDate:
					{
						fmt.Printf("Unhandled: %s\n", GotAtServiceDate.stateName())
					}
				case GotAtServiceYear:
					{
						fmt.Printf("Unhandled: %s\n", GotAtServiceYear.stateName())
					}
				case GotAtSundayAfterElevationCrossDate:
					{
						fmt.Printf("Unhandled: %s\n", GotAtSundayAfterElevationCrossDate.stateName())
					}
				case GotAtSundaysBeforeTriodion:
					{
						fmt.Printf("Unhandled: %s\n", GotAtSundaysBeforeTriodion.stateName())
					}
				case GotAtVer:
					{
						fmt.Printf("Unhandled: %s\n", GotAtVer.stateName())
					}
				case GotBreak:
					{
						fmt.Printf("Unhandled: %s\n", GotBreak.stateName())
					}
				case GotBreakTypeLine:
					{
						fmt.Printf("Unhandled: %s\n", GotBreakTypeLine.stateName())
					}
				case GotBreakTypePage:
					{
						fmt.Printf("Unhandled: %s\n", GotBreakTypePage.stateName())
					}
				case GotBlock:
					{
						fmt.Printf("Unhandled: %s\n", GotBlock.stateName())
					}
				case GotColon:
					{
						fmt.Printf("Unhandled: %s\n", GotColon.stateName())
					}
				case GotDay:
					{
						if inHead {
							var i int
							if i, err := strconv.Atoi(token); err != nil {
								fmt.Printf("day=%d, type: %T\n", i, i)
							}
							mt.Day = i
						} else {
							fmt.Printf("Unhandled: %s\n", GotDay.stateName())
						}
					}
				case GotDialog:
					{
						fmt.Printf("Unhandled: %s\n", GotDialog.stateName())
					}
				case GotEndActor:
					{
						fmt.Printf("Unhandled: %s\n", GotEndActor.stateName())
					}
				case GotEndBreak:
					{
						fmt.Printf("Unhandled: %s\n", GotEndBreak.stateName())
					}
				case GotEndBlock:
					{
						fmt.Printf("Unhandled: %s\n", GotEndBlock.stateName())
					}
				case GotEndDate:
					{
						fmt.Printf("Unhandled: %s\n", GotEndDate.stateName())
					}
				case GotEndDialog:
					{
						fmt.Printf("Unhandled: %s\n", GotEndDialog.stateName())
					}
				case GotEndEndPageHeaderEven:
					{
						fmt.Printf("Unhandled: %s\n", GotEndEndPageHeaderEven.stateName())
					}
				case GotEndEndPageHeaderOdd:
					{
						fmt.Printf("Unhandled: %s\n", GotEndEndPageHeaderOdd.stateName())
					}
				case GotEndHymn:
					{
						fmt.Printf("Unhandled: %s\n", GotEndHymn.stateName())
					}
				case GotEndInsert:
					{
						fmt.Printf("Unhandled: %s\n", GotEndInsert.stateName())
					}
				case GotEndMedia:
					{
						fmt.Printf("Unhandled: %s\n", GotEndMedia.stateName())
					}
				case GotEndPageFooterEven:
					{
						fmt.Printf("Unhandled: %s\n", GotEndPageFooterEven.stateName())
					}
				case GotEndPageFooterOdd:
					{
						fmt.Printf("Unhandled: %s\n", GotEndPageFooterOdd.stateName())
					}
				case GotEndPara:
					{
						fmt.Printf("Unhandled: %s\n", GotEndPara.stateName())
					}
				case GotEndPassthroughHTML:
					{
						fmt.Printf("Unhandled: %s\n", GotEndPassthroughHTML.stateName())
					}
				case GotEndPreface:
					{
						fmt.Printf("Unhandled: %s\n", GotEndPreface.stateName())
					}
				case GotEndReading:
					{
						fmt.Printf("Unhandled: %s\n", GotEndReading.stateName())
					}
				case GotEndRubric:
					{
						fmt.Printf("Unhandled: %s\n", GotEndRubric.stateName())
					}
				case GotEndSection:
					{
						fmt.Printf("Unhandled: %s\n", GotEndSection.stateName())
					}
				case GotEndSetMcDay:
					{
						fmt.Printf("Unhandled: %s\n", GotEndSetMcDay.stateName())
					}
				case GotEndSetPageNumber:
					{
						fmt.Printf("Unhandled: %s\n", GotEndSetPageNumber.stateName())
					}
				case GotEndSubTitle:
					{
						fmt.Printf("Unhandled: %s\n", GotEndSubTitle.stateName())
					}
				case GotEndSwitchVersion:
					{
						fmt.Printf("Unhandled: %s\n", GotEndSwitchVersion.stateName())
					}
				case GotEndTemplateCommemoration:
					{
						fmt.Printf("Unhandled: %s\n", GotEndTemplateCommemoration.stateName())
					}
				case GotEndTitle:
					{
						fmt.Printf("Unhandled: %s\n", GotEndTitle.stateName())
					}
				case GotEndVerse:
					{
						fmt.Printf("Unhandled: %s\n", GotEndVerse.stateName())
					}
				case GotEndWhen:
					{
						fmt.Printf("Unhandled: %s\n", GotEndWhen.stateName())
					}
				case GotEqualSign:
					{
						fmt.Printf("Unhandled: %s\n", GotEqualSign.stateName())
					}
				case GotFooterColumnCenter:
					{
						fmt.Printf("Unhandled: %s\n", GotFooterColumnCenter.stateName())
					}
				case GotFooterColumnLeft:
					{
						fmt.Printf("Unhandled: %s\n", GotFooterColumnLeft.stateName())
					}
				case GotFooterColumnRight:
					{
						fmt.Printf("Unhandled: %s\n", GotFooterColumnRight.stateName())
					}
				case GotFooterCommemoration:
					{
						fmt.Printf("Unhandled: %s\n", GotFooterCommemoration.stateName())
					}
				case GotFooterDate:
					{
						fmt.Printf("Unhandled: %s\n", GotFooterDate.stateName())
					}
				case GotFooterLookup:
					{
						fmt.Printf("Unhandled: %s\n", GotFooterLookup.stateName())
					}
				case GotFooterPageNumber:
					{
						fmt.Printf("Unhandled: %s\n", GotFooterPageNumber.stateName())
					}
				case GotFooterTitle:
					{
						fmt.Printf("Unhandled: %s\n", GotFooterTitle.stateName())
					}
				case GotHeaderFooterText:
					{
						fmt.Printf("Unhandled: %s\n", GotHeaderFooterText.stateName())
					}
				case GotHymn:
					{
						fmt.Printf("Unhandled: %s\n", GotHymn.stateName())
					}
				case GotImport:
					{
						switch token {
						case ".":
							if s.Peek() == '*' {
								mt.addImport(importSb.String())
								importSb.Reset()
								currentState = Neutral
							} else {
								parts := strings.Split(token, "_")
								if len(parts) == 4 {
									sb := strings.Builder{}
									sb.WriteString(parts[1])
									sb.WriteString("_")
									sb.WriteString(strings.ToLower(parts[2]))
									sb.WriteString("_")
									sb.WriteString(parts[3])
									sb.WriteString(pathSeparator)
									importSb.WriteString(parts[0])
									sb.WriteString(importSb.String())
									mt.addImport(importSb.String())
									importSb.Reset()
									currentState = Neutral
								} else {
									importSb.WriteString(token)
								}
							}
						default:
							parts := strings.Split(token, "_")
							if len(parts) == 4 {
								sb := strings.Builder{}
								sb.WriteString(parts[1])
								sb.WriteString("_")
								sb.WriteString(strings.ToLower(parts[2]))
								sb.WriteString("_")
								sb.WriteString(parts[3])
								sb.WriteString(pathSeparator)
								importSb.WriteString(parts[0])
								sb.WriteString(importSb.String())
								sb.WriteString(pathSeparator)
								mt.addImport(sb.String())
								importSb.Reset()
								currentState = Neutral
							} else {
								importSb.WriteString(token)
							}
						}
					}
				case GotInsertPreface:
					{
						fmt.Printf("Unhandled: %s\n", GotInsertPreface.stateName())
					}
				case GotInsertSection:
					{
						// TODO: get path to section being inserted
						// TODO: handle tab indentation
						lineSB.WriteString(token)
					}
				case GotInsertTemplate:
					{
						fmt.Printf("Unhandled: %s\n", GotInsertTemplate.stateName())
					}
				case GotMedia:
					{
						fmt.Printf("Unhandled: %s\n", GotMedia.stateName())
					}
				case GotMediaOff:
					{
						fmt.Printf("Unhandled: %s\n", GotMediaOff.stateName())
					}
				case GotMonth:
					if inHead {
						var i int
						if i, err := strconv.Atoi(token); err != nil {
							fmt.Printf("month=%d, type: %T\n", i, i)
						}
						mt.Month = i
					} else {
						fmt.Printf("Unhandled: %s\n", GotMonth.stateName())
					}
				case GotOtherwise:
					{
						fmt.Printf("Unhandled: %s\n", GotOtherwise.stateName())
					}
				case GotPageFooterEven:
					{
						fmt.Printf("Unhandled: %s\n", GotPageFooterEven.stateName())
					}
				case GotPageFooterOdd:
					{
						fmt.Printf("Unhandled: %s\n", GotPageFooterOdd.stateName())
					}
				case GotPageHeaderEven:
					{
						fmt.Printf("Unhandled: %s\n", GotPageHeaderEven.stateName())
					}
				case GotPageHeaderOdd:
					{
						fmt.Printf("Unhandled: %s\n", GotPageHeaderOdd.stateName())
					}
				case GotPara:
					{
						fmt.Printf("Unhandled: %s\n", GotPara.stateName())
					}
				case GotPassthroughHTML:
					{
						fmt.Printf("Unhandled: %s\n", GotPassthroughHTML.stateName())
					}
				case GotPreface:
					{
						fmt.Printf("Unhandled: %s\n", GotPreface.stateName())
					}
				case GotReading:
					{
						fmt.Printf("Unhandled: %s\n", GotReading.stateName())
					}
				case GotRid:
					{
						fmt.Printf("Unhandled: %s\n", GotRid.stateName())
					}
				case GotRole:
					{
						fmt.Printf("Unhandled: %s\n", GotRole.stateName())
					}
				case GotRubric:
					{
						fmt.Printf("Unhandled: %s\n", GotRubric.stateName())
					}
				case GotSection:
					{
						mt.ID = mt.ID + pathSeparator + token
						mt.TemplateDir = dirOut
						currentState = Neutral
						mt.setFilePath()
						if mt.File, err = os.Create(mt.FilePath); err != nil {
							return err
						}
						mt.setWriter()
						mt.Writer.WriteString(fmt.Sprintf("id: %s\n", mt.ID))
						internalPaths.Push(currentInternalPath)
						currentInternalPath = currentInternalPath + "." + token
						currentState = Neutral
					}
				case GotSetDate:
					{
						fmt.Printf("Unhandled: %s\n", GotSetDate.stateName())
					}
				case GotSetMcDay:
					{
						fmt.Printf("Unhandled: %s\n", GotSetMcDay.stateName())
					}
				case GotSetPageNumber:
					{
						fmt.Printf("Unhandled: %s\n", GotSetPageNumber.stateName())
					}
				case GotSid:
					{
						fmt.Printf("Unhandled: %s\n", GotSid.stateName())
					}
				case GotStatus:
					{
						fmt.Printf("Unhandled: %s\n", GotStatus.stateName())
					}
				case GotSubTitle:
					{
						fmt.Printf("Unhandled: %s\n", GotSubTitle.stateName())
					}
				case GotSwitchVersion:
					{
						fmt.Printf("Unhandled: %s\n", GotSwitchVersion.stateName())
					}
				case GotTemplateCommemoration:
					{
						fmt.Printf("Unhandled: %s\n", GotTemplateCommemoration.stateName())
					}
				case GotTemplateName:
					{
						mt.Title = token
						mt.Writer.WriteString(fmt.Sprintf("id: \"%s\"\n", mt.ID))
						currentState = Neutral
					}
				case GotTemplateStatusDraft:
					{
						fmt.Printf("Unhandled: %s\n", GotTemplateStatusDraft.stateName())
					}
				case GotTemplateStatusFinal:
					{
						fmt.Printf("Unhandled: %s\n", GotTemplateStatusFinal.stateName())
					}
				case GotTemplateStatusNA:
					{
						fmt.Printf("Unhandled: %s\n", GotTemplateStatusNA.stateName())
					}
				case GotTemplateStatusReview:
					{
						fmt.Printf("Unhandled: %s\n", GotTemplateStatusReview.stateName())
					}
				case GotTemplateTitle:
					{
						fmt.Printf("Unhandled: %s\n", GotTemplateTitle.stateName())
					}
				case GotTitle:
					{
						fmt.Printf("Unhandled: %s\n", GotTitle.stateName())
					}
				case GotUse:
					{
						fmt.Printf("Unhandled: %s\n", GotUse.stateName())
					}
				case GotVerse:
					{
						fmt.Printf("Unhandled: %s\n", GotVerse.stateName())
					}
				case GotWhenDateIs:
					{
						fmt.Printf("Unhandled: %s\n", GotWhenDateIs.stateName())
					}
				case GotWhenExists:
					{
						fmt.Printf("Unhandled: %s\n", GotWhenExists.stateName())
					}
				case GotWhenLukanCycleDayIs:
					{
						fmt.Printf("Unhandled: %s\n", GotWhenLukanCycleDayIs.stateName())
					}
				case GotWhenModeOfWeekIs:
					{
						fmt.Printf("Unhandled: %s\n", GotWhenModeOfWeekIs.stateName())
					}
				case GotWhenMovableCycleDayIs:
					{
						fmt.Printf("Unhandled: %s\n", GotWhenMovableCycleDayIs.stateName())
					}
				case GotWhenNameOfDayIs:
					{
						fmt.Printf("Unhandled: %s\n", GotWhenNameOfDayIs.stateName())
					}
				case GotWhenSundayAfterElevationOfCross:
					{
						fmt.Printf("Unhandled: %s\n", GotWhenSundayAfterElevationOfCross.stateName())
					}
				case GotWhenSundaysBeforeTriodionIs:
					{
						fmt.Printf("Unhandled: %s\n", GotWhenSundaysBeforeTriodionIs.stateName())
					}
				case GotYear:
					if inHead {
						var i int
						if i, err := strconv.Atoi(token); err != nil {
							fmt.Printf("year=%d, type: %T\n", i, i)
						}
						mt.Year = i
					} else {
						fmt.Printf("Unhandled: %s\n", GotYear.stateName())
					}
				}
			}
		}
	}
	return nil
}

var currentState state

type state int

const (
	LineStart                          state = iota
	GotActor                                 // Actor
	GotAmpersand                             // @
	GotAtAllLiturgicalDayProperties          // @All_Liturgical_Day_Properties
	GotAtDay                                 // @day
	GotAtDayOfMonth                          // @Day_of_Month
	GotAtDayOfMovableCycle                   // @Day_of_Movable_Cycle
	GotAtDayOfPeriod                         // @Day_of_Period
	GotAtDayOfWeekAsNumber                   // @Day_of_Week_As_Number
	GotAtDayOfWeekAsText                     // @Day_of_Week_As_Text
	GotAtEothinon                            // @Eothinon
	GotAtLukanCycleElapsedDays               // @Lukan_Cycle_Elapsed_Days
	GotAtLukanCycleStartDate                 // @Lukan_Cycle_Start_Date'
	GotAtLukanCycleWeek                      // @Lukan_Cycle_Week
	GotAtLukanCycleWeekDay                   // @Lukan_Cycle_Week_Day
	GotAtMode                                // @mode
	GotAtModeOfWeek                          // @Mode_of_Week
	GotAtNameOfPeriod                        // @Name_of_Period
	GotAtServiceDate                         // @Service_Date
	GotAtServiceYear                         // @Service_Year
	GotAtSundayAfterElevationCrossDate       // @Sunday_After_Elevation_Cross_Date
	GotAtSundaysBeforeTriodion               // @Sundays_Before_Triodion
	GotAtVer                                 // @ver
	GotBreak                                 // Break
	GotBreakTypeLine                         // line
	GotBreakTypePage                         // page
	GotBlock                                 // bTag
	GotColon                                 // :
	GotDay                                   // day
	GotDialog                                // Dialog
	GotEnd                                   // End
	GotEndActor                              // End_Actor
	GotEndBreak                              // End_Break
	GotEndBlock                              // End_bTag
	GotEndDate                               // End_Date
	GotEndDialog                             // End_Dialog
	GotEndEndPageHeaderEven                  // End_Page_Header_Even
	GotEndEndPageHeaderOdd                   // End_Page_Header_Odd
	GotEndHead                               // End_Head
	GotEndHymn                               // End_Hymn
	GotEndInsert                             // End_Insert
	GotEndMedia                              // End_Media
	GotEndPageFooterEven                     // End_Page_Footer_Even
	GotEndPageFooterOdd                      // End_Page_Footer_Odd
	GotEndPara                               // End_Para
	GotEndPassthroughHTML                    // END-Passthrough-Html
	GotEndPreface                            // End_Preface
	GotEndReading                            // End_Reading
	GotEndRubric                             // End_Rubric
	GotEndSection                            // End_Section
	GotEndSetMcDay                           // End_mcDay
	GotEndSetPageNumber                      // End_Set_Page_Number
	GotEndSubTitle                           // End_Sub-Title
	GotEndSwitchVersion                      // End_Switch-Version
	GotEndTemplateCommemoration              // End_Template_Commemoration
	GotEndTitle                              // End_Title || End_Title
	GotEndVerse                              // End_Verse
	GotEndWhen                               // end-when
	GotEqualSign                             // =
	GotFooterColumnCenter                    // center
	GotFooterColumnLeft                      // left
	GotFooterColumnRight                     // right
	GotFooterCommemoration                   // @commemoration
	GotFooterDate                            // @date
	GotFooterLookup                          // @lookup
	GotFooterPageNumber                      // @pageNbr
	GotFooterTitle                           // @title
	GotHeaderFooterText                      // @text
	GotHead                                  // Head
	GotHymn                                  // Hymn
	GotImport                                // import
	GotInsertPreface                         // Insert_preface - unused
	GotInsertSection                         // Insert_section
	GotInsertTemplate                        // Insert_template
	GotMedia                                 // Media
	GotMediaOff                              // media-off
	GotMonth                                 // month
	GotNid
	GotOtherwise                       // otherwise
	GotPageFooterEven                  // Page_Footer_Even
	GotPageFooterOdd                   // Page_Footer_Odd
	GotPageHeaderEven                  // Page_Header_Even
	GotPageHeaderOdd                   // Page_Header_Odd
	GotPara                            // Para
	GotPassthroughHTML                 // Passthrough-Html
	GotPreface                         // Preface
	GotReading                         // Reading
	GotRid                             // rid
	GotRole                            // role
	GotRubric                          // Rubric
	GotSection                         // Section
	GotSetDate                         // Set_Date
	GotSetMcDay                        // Set_mcDay
	GotSetPageNumber                   // Set_Page_Number
	GotSid                             // sid
	GotStatus                          // Status
	GotSubTitle                        // Sub-Title
	GotSwitchVersion                   // Switch-Version
	GotTemplateCommemoration           // Template_Commemoration
	GotTemplateName                    // Template
	GotTemplateStatusDraft             // Draft
	GotTemplateStatusFinal             // Final
	GotTemplateStatusNA                // NA
	GotTemplateStatusReview            // Review
	GotTemplateTitle                   // Template_Title
	GotTitle                           // Title
	GotUse                             // use
	GotVerse                           // Verse
	GotWhenDateIs                      // when-date-is
	GotWhenExists                      // when-exists
	GotWhenLukanCycleDayIs             // when-Lukan-Cycle-Day-is
	GotWhenModeOfWeekIs                // when-mode-of-week-is
	GotWhenMovableCycleDayIs           // when-movable-cycle-day-is
	GotWhenNameOfDayIs                 // when-name-of-day-is
	GotWhenSundayAfterElevationOfCross // when-Sunday-after-Elevation-of-Cross-is
	GotWhenSundaysBeforeTriodionIs     // when-sundays-before-triodion-is
	GotYear                            // year
	Neutral
	ERROR
	END
)

func (s state) stateName() string {
	switch s {
	case LineStart:
		return "LINE_START"
	case GotActor:
		return "GotActor"

	case GotAmpersand:
		return "GotAmpersand"

	case GotAtAllLiturgicalDayProperties:
		return "GotAtAllLiturgicalDayProperties"

	case GotAtDay:
		return "GotAtDay"

	case GotAtDayOfMonth:
		return "GotAtDayOfMonth"

	case GotAtDayOfMovableCycle:
		return "GotAtDayOfMovableCycle"

	case GotAtDayOfPeriod:
		return "GotAtDayOfPeriod"

	case GotAtDayOfWeekAsNumber:
		return "GotAtDayOfWeekAsNumber"

	case GotAtDayOfWeekAsText:
		return "GotAtDayOfWeekAsText"

	case GotAtEothinon:
		return "GotAtEothinon"

	case GotAtLukanCycleElapsedDays:
		return "GotAtLukanCycleElapsedDays"

	case GotAtLukanCycleStartDate:
		return "GotAtLukanCycleStartDate"

	case GotAtLukanCycleWeek:
		return "GotAtLukanCycleWeek"

	case GotAtLukanCycleWeekDay:
		return "GotAtLukanCycleWeekDay"

	case GotAtMode:
		return "GotAtMode"

	case GotAtModeOfWeek:
		return "GotAtModeOfWeek"

	case GotAtNameOfPeriod:
		return "GotAtNameOfPeriod"

	case GotAtServiceDate:
		return "GotAtServiceDate"

	case GotAtServiceYear:
		return "GotAtServiceYear"

	case GotAtSundayAfterElevationCrossDate:
		return "GotAtSundayAfterElevationCrossDate"

	case GotAtSundaysBeforeTriodion:
		return "GotAtSundaysBeforeTriodion"

	case GotAtVer:
		return "GotAtVer"

	case GotBreak:
		return "GotBreak"

	case GotBreakTypeLine:
		return "GotBreakTypeLine"

	case GotBreakTypePage:
		return "GotBreakTypePage"

	case GotBlock:
		return "GotBlock"

	case GotColon:
		return "GotColon"

	case GotDay:
		return "GotDay"

	case GotDialog:
		return "GotDialog"

	case GotEndActor:
		return "GotEndActor"

	case GotEndBreak:
		return "GotEndBreak"

	case GotEndBlock:
		return "GotEndBlock"

	case GotEndDate:
		return "GotEndDate"

	case GotEndDialog:
		return "GotEndDialog"

	case GotEndEndPageHeaderEven:
		return "GotEndEndPageHeaderEven"

	case GotEndEndPageHeaderOdd:
		return "GotEndEndPageHeaderOdd"

	case GotEndHymn:
		return "GotEndHymn"

	case GotEndInsert:
		return "GotEndInsert"

	case GotEndMedia:
		return "GotEndMedia"

	case GotEndPageFooterEven:
		return "GotEndPageFooterEven"

	case GotEndPageFooterOdd:
		return "GotEndPageFooterOdd"

	case GotEndPara:
		return "GotEndPara"

	case GotEndPassthroughHTML:
		return "GotEndPassthroughHTML"

	case GotEndPreface:
		return "GotEndPreface"

	case GotEndReading:
		return "GotEndReading"

	case GotEndRubric:
		return "GotEndRubric"

	case GotEndSection:
		return "GotEndSection"

	case GotEndSetMcDay:
		return "GotEndSetMcDay"

	case GotEndSetPageNumber:
		return "GotEndSetPageNumber"

	case GotEndSubTitle:
		return "GotEndSubTitle"

	case GotEndSwitchVersion:
		return "GotEndSwitchVersion"

	case GotEndTemplateCommemoration:
		return "GotEndTemplateCommemoration"

	case GotEndTitle:
		return "GotEndTitle"

	case GotEndVerse:
		return "GotEndVerse"

	case GotEndWhen:
		return "GotEndWhen"

	case GotEqualSign:
		return "GotEqualSign"

	case GotFooterColumnCenter:
		return "GotFooterColumnCenter"

	case GotFooterColumnLeft:
		return "GotFooterColumnLeft"

	case GotFooterColumnRight:
		return "GotFooterColumnRight"

	case GotFooterCommemoration:
		return "GotFooterCommemoration"

	case GotFooterDate:
		return "GotFooterDate"

	case GotFooterLookup:
		return "GotFooterLookup"

	case GotFooterPageNumber:
		return "GotFooterPageNumber"

	case GotFooterTitle:
		return "GotFooterTitle"

	case GotHeaderFooterText:
		return "GotHeaderFooterText"

	case GotHymn:
		return "GotHymn"

	case GotImport:
		return "GotImport"

	case GotInsertPreface:
		return "GotInsertPreface"

	case GotInsertSection:
		return "GotInsertSection"

	case GotInsertTemplate:
		return "GotInsertTemplate"

	case GotMedia:
		return "GotMedia"

	case GotMediaOff:
		return "GotMediaOff"

	case GotMonth:
		return "GotMonth"

	case GotOtherwise:
		return "GotOtherwise"

	case GotPageFooterEven:
		return "GotPageFooterEven"

	case GotPageFooterOdd:
		return "GotPageFooterOdd"

	case GotPageHeaderEven:
		return "GotPageHeaderEven"

	case GotPageHeaderOdd:
		return "GotPageHeaderOdd"

	case GotPara:
		return "GotPara"

	case GotPassthroughHTML:
		return "GotPassthroughHTML"

	case GotPreface:
		return "GotPreface"

	case GotReading:
		return "GotReading"

	case GotRid:
		return "GotRid"

	case GotRole:
		return "GotRole"

	case GotRubric:
		return "GotRubric"

	case GotSection:
		return "GotSection"

	case GotSetDate:
		return "GotSetDate"

	case GotSetMcDay:
		return "GotSetMcDay"

	case GotSetPageNumber:
		return "GotSetPageNumber"

	case GotSid:
		return "GotSid"

	case GotStatus:
		return "GotStatus"

	case GotSubTitle:
		return "GotSubTitle"

	case GotSwitchVersion:
		return "GotSwitchVersion"

	case GotTemplateCommemoration:
		return "GotTemplateCommemoration"

	case GotTemplateName:
		return "GotTemplateName"

	case GotTemplateStatusDraft:
		return "GotTemplateStatusDraft"

	case GotTemplateStatusFinal:
		return "GotTemplateStatusFinal"

	case GotTemplateStatusNA:
		return "GotTemplateStatusNA"

	case GotTemplateStatusReview:
		return "GotTemplateStatusReview"

	case GotTemplateTitle:
		return "GotTemplateTitle"

	case GotTitle:
		return "GotTitle"

	case GotUse:
		return "GotUse"

	case GotVerse:
		return "GotVerse"

	case GotWhenDateIs:
		return "GotWhenDateIs"

	case GotWhenExists:
		return "GotWhenExists"

	case GotWhenLukanCycleDayIs:
		return "GotWhenLukanCycleDayIs"

	case GotWhenModeOfWeekIs:
		return "GotWhenModeOfWeekIs"

	case GotWhenMovableCycleDayIs:
		return "GotWhenMovableCycleDayIs"

	case GotWhenNameOfDayIs:
		return "GotWhenNameOfDayIs"

	case GotWhenSundayAfterElevationOfCross:
		return "GotWhenSundayAfterElevationOfCross"

	case GotWhenSundaysBeforeTriodionIs:
		return "GotWhenSundaysBeforeTriodionIs"

	case GotYear:
		return "GotYear"
	case ERROR:
		return "ERROR"
	case END:
		return "END"
	default:
		return "UNKNOWN"
	}
}
func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// constants for

// simulate database lookup
func getValue(id string) string {
	switch id {
	case "gr_gr_cog~actors~Deacon":
		return "ΔΙΑΚΟΝΟΣ:"
	case "en_us_goa~actors~Deacon":
		return "Deacon:"
	case "gr_gr_cog~prayers~res08":
		return "Εὐλόγησον, Δέσποτα."
	case "en_us_goa~prayers~res08":
		return "Master, give the blessing."
	case "gr_gr_cog~actors~Priest":
		return "ΙΕΡΕΥΣ"
	case "en_us_goa~actors~Priest":
		return "PRIEST"
	case "gr_gr_cog~prayers~enarxis02":
		return "Εὐλογημένη ἡ Βασιλεία τοῦ Πατρὸς καὶ τοῦ Υἱοῦ καὶ τοῦ Ἁγίου Πνεύματος, νῦν καὶ ἀεί, καὶ εἰς τοὺς αἰῶνας τῶν αἰώνων."
	case "en_us_goa~prayers~enarxis02":
		return "Blessed is the Kingdom of the Father and of the Son and of the Holy Spirit, now and forever and to the ages of ages."
	default:
		return "unknown id " + id
	}
}
func resolveSectionPath(pathToResolve string, currentInternalPath string, sectionMap map[string]string) string {
	if doxaPath, ok := sectionMap[pathToResolve]; ok {
		return doxaPath
	} else {
		parts := strings.Split(currentInternalPath,".")
		l := len(parts)
		for i, _ := range parts {
			sectionId := strings.Join(parts[0:l-i],".") + "." + pathToResolve
			if doxaPath, ok := sectionMap[sectionId]; ok {
				return doxaPath
				break
			}
		}
	}
	return ""
}