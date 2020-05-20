package atem2lml

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/emirpasic/gods/trees/btree"
	"github.com/liturgiko/doxa/pkg/enums/statuses"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
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
	Sections    *btree.Tree
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

const pathSeparator = string(os.PathSeparator)

func ParseTemplate(fileIn, dirOut, id string, tTree, sTree *btree.Tree) error {
	var mt MetaTemplate
	var err error
	// get the name of the file
	filename := filepath.Base(fileIn)
	filename = filename[0 : len(filename)-5]
	// get the internal sections tree for this template
	if v, ok := tTree.Get(filename); ok {
		mt.Sections = v.(TemplateNode).Sections
	} else {
		return errors.New(fmt.Sprintf("could not find TemplateNode for %s\n", filename))
	}
	// set up stack for templates.  We will create one for each section encountered
	templates := arraystack.New()
	mt.Title = filename
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
	internalPaths := arraystack.New()
	currentInternalPath := filename

	inHead := false
	currentState = LineStart

	var lines []string
	var lineNbr int

	if lines, err = ltfile.GetFileLines(fileIn); err != nil {return err}
	for i, line := range lines {
		lineNbr = i + 1
		text := strings.ReplaceAll(line, "-", "_")
		var s scanner.Scanner
		s.Init(strings.NewReader(text))
		s.Whitespace = 1<<':' | 1<<'-' | 1<<'\t' | 1<<' ' // treat these as white space
		lineSB := strings.Builder{}
		identifierSB := strings.Builder{}
		currentState = LineStart

		for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
			token := s.TokenText()

			switch token {
			case "Actor":
				lineSB.WriteString(fmt.Sprintf("p.actor "))
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
				lineSB.WriteString(fmt.Sprintf("p.dialog "))
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
				sectionId := identifierSB.String()
				if strings.HasSuffix(sectionId,".") {
					// TODO: figure out why it sometimes ends with dot
					sectionId = sectionId[:len(sectionId)-1]
				}
				identifierSB.Reset()
				identifierSB.WriteString(resolvePath(sectionId, currentInternalPath, &mt, tTree, sTree))
				if identifierSB.Len() == 0 {
					fmt.Printf("Could not resolve section path %s\n",sectionId)
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
				lineSB.WriteString(fmt.Sprintf("p.reading "))
				currentState = GotEndReading
			case "End_Rubric":
				currentState = GotEndRubric
			case "End_Section":
				insertPath := mt.ID
				if mt.Writer != nil {
					mt.Writer.Flush()
				}
				if mt.File != nil {
					mt.File.Close()
				}
				if v, ok := internalPaths.Pop(); ok {
					currentInternalPath = v.(string)
				} else {
					fmt.Println("Expected value, but stack empty")
				}
				if v, ok := templates.Pop(); ok {
					mt = v.(MetaTemplate)
					mt.Writer.WriteString(fmt.Sprintf("insert \"%s\"\n", insertPath))
					// TODO: remove next line for production
					mt.Writer.Flush()
				} else {
					fmt.Println("Expected value, but stack empty")
				}
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
			case "End_Title", "End_title":
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
					mt.Writer.WriteString(fmt.Sprintf("id: \"%s\"\n", mt.ID))
					internalPaths.Push(currentInternalPath)
					currentInternalPath = currentInternalPath + "." + token
					currentState = Neutral
				} else {
					currentState = GotFooterCommemoration
				}
			case "date":
				currentState = GotFooterDate
			case "lookup":
				if inHead {
					currentState = GotFooterLookup
				} else {
					if currentState != GotTemplateName {
						lineSB.WriteString(token)
					}
				}
			case "pageNbr":
				currentState = GotFooterPageNumber
			case "title":
				lineSB.WriteString(token)
				//if currentState == GotSection {
				//	mt.ID = mt.ID + pathSeparator + token
				//	mt.TemplateDir = dirOut
				//	currentState = Neutral
				//	mt.setFilePath()
				//	if mt.File, err = os.Create(mt.FilePath); err != nil {
				//		return err
				//	}
				//	mt.setWriter()
				//	mt.Writer.WriteString(fmt.Sprintf("id: %s\n", mt.ID))
				//	internalPaths.Push(currentInternalPath)
				//	currentInternalPath = currentInternalPath + "." + token
				//	currentState = Neutral
				//} else {
				//	currentState = GotFooterTitle
				//}
			case "text":
				lineSB.WriteString(token)
				//if currentState == GotSection {
				//	mt.ID = mt.ID + pathSeparator + token
				//	mt.TemplateDir = dirOut
				//	currentState = Neutral
				//	mt.setFilePath()
				//	if mt.File, err = os.Create(mt.FilePath); err != nil {
				//		return err
				//	}
				//	mt.setWriter()
				//	mt.Writer.WriteString(fmt.Sprintf("id: %s\n", mt.ID))
				//	internalPaths.Push(currentInternalPath)
				//	currentInternalPath = currentInternalPath + "." + token
				//	currentState = Neutral
				//} else {
				//	currentState = GotHeaderFooterText
				//}
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
				{
					lineSB.WriteString(fmt.Sprintf("p "))
					currentState = GotPara
				}
			case "Passthrough-Html":
				currentState = GotPassthroughHTML
			case "Preface":
				currentState = GotPreface
			case "Reading":
				{
					lineSB.WriteString(fmt.Sprintf("p.reading "))
					currentState = GotReading
				}
			case "rid":
				lineSB.WriteString(fmt.Sprintf("rid "))
				currentState = GotRid
			case "role":
				currentState = GotRole
			case "Rubric":
				lineSB.WriteString(fmt.Sprintf("p.rubric "))
				currentState = GotRubric
			case "Section":
				currentState = GotSection
				tmpMt := MetaTemplate{}
				tmpMt.TemplateDir = dirOut
				tmpMt.ID = mt.ID // once we get the section name, we will append it to ID
				tmpMt.Type = BLOCK
				tmpMt.Imports = mt.Imports
				tmpMt.Sections = mt.Sections
				tmpMt.Status = mt.Status
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
				lineSB.WriteString(fmt.Sprintf("sid "))
				currentState = GotSid
			case "Status":
				lineSB.WriteString("status: ")
				currentState = GotStatus
			case "Sub-Title":
				currentState = GotSubTitle
			case "Switch-Version":
				currentState = GotSwitchVersion
			case "Template_Commemoration":
				currentState = GotTemplateCommemoration
			case "Template":
				lineSB.WriteString("id: \"")
				currentState = GotTemplateName
			case "Draft":
				if currentState == GotStatus {
					mt.Status = DRAFT
					lineSB.WriteString(strings.ToLower(statuses.Draft.String()))
				} else {
					lineSB.WriteString(token)
				}
			case "Final":
				if currentState == GotStatus {
					mt.Status = FINAL
					lineSB.WriteString(strings.ToLower(statuses.Final.String()))
				} else {
					lineSB.WriteString(token)
				}
			case "NA":
				if currentState == GotStatus {
					mt.Status = NA
					lineSB.WriteString(strings.ToLower(statuses.NA.String()))
				} else {
					lineSB.WriteString(token)
				}
			case "Review":
				if currentState == GotStatus {
					mt.Status = REVIEW
					lineSB.WriteString(strings.ToLower(statuses.Review.String()))
				} else {
					lineSB.WriteString(token)
				}
			case "Template_Title":
				currentState = GotTemplateTitle
			case "Title":
				{
					lineSB.WriteString(fmt.Sprintf("p.title "))
					currentState = GotTitle
				}
			case "use":
				currentState = GotUse
			case "Verse":
				lineSB.WriteString(fmt.Sprintf("p.verse "))
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
					case GotAmpersand:
						{
							report(mt.ID, lineNbr, GotAmpersand.stateName())
						}
					case GotAtAllLiturgicalDayProperties:
						{
							report(mt.ID, lineNbr, GotAtAllLiturgicalDayProperties.stateName())
						}
					case GotAtDay:
						{
							report(mt.ID, lineNbr, GotAtDay.stateName())
						}
					case GotAtDayOfMonth:
						{
							report(mt.ID, lineNbr, GotAtDayOfMonth.stateName())
						}
					case GotAtDayOfMovableCycle:
						{
							report(mt.ID, lineNbr, GotAtDayOfMovableCycle.stateName())
						}
					case GotAtDayOfPeriod:
						{
							report(mt.ID, lineNbr, GotAtDayOfPeriod.stateName())
						}
					case GotAtDayOfWeekAsNumber:
						{
							report(mt.ID, lineNbr, GotAtDayOfWeekAsNumber.stateName())
						}
					case GotAtDayOfWeekAsText:
						{
							report(mt.ID, lineNbr, GotAtDayOfWeekAsText.stateName())
						}
					case GotAtEothinon:
						{
							report(mt.ID, lineNbr, GotAtEothinon.stateName())
						}
					case GotAtLukanCycleElapsedDays:
						{
							report(mt.ID, lineNbr, GotAtLukanCycleElapsedDays.stateName())
						}
					case GotAtLukanCycleStartDate:
						{
							report(mt.ID, lineNbr, GotAtLukanCycleStartDate.stateName())
						}
					case GotAtLukanCycleWeek:
						{
							report(mt.ID, lineNbr, GotAtLukanCycleWeek.stateName())
						}
					case GotAtLukanCycleWeekDay:
						{
							report(mt.ID, lineNbr, GotAtLukanCycleWeekDay.stateName())
						}
					case GotAtMode:
						{
							report(mt.ID, lineNbr, GotAtMode.stateName())
						}
					case GotAtModeOfWeek:
						{
							report(mt.ID, lineNbr, GotAtModeOfWeek.stateName())
						}
					case GotAtNameOfPeriod:
						{
							report(mt.ID, lineNbr, GotAtNameOfPeriod.stateName())
						}
					case GotAtServiceDate:
						{
							report(mt.ID, lineNbr, GotAtServiceDate.stateName())
						}
					case GotAtServiceYear:
						{
							report(mt.ID, lineNbr, GotAtServiceYear.stateName())
						}
					case GotAtSundayAfterElevationCrossDate:
						{
							report(mt.ID, lineNbr, GotAtSundayAfterElevationCrossDate.stateName())
						}
					case GotAtSundaysBeforeTriodion:
						{
							report(mt.ID, lineNbr, GotAtSundaysBeforeTriodion.stateName())
						}
					case GotAtVer:
						{
							report(mt.ID, lineNbr, GotAtVer.stateName())
						}
					case GotBreak:
						{
							report(mt.ID, lineNbr, GotBreak.stateName())
						}
					case GotBreakTypeLine:
						{
							report(mt.ID, lineNbr, GotBreakTypeLine.stateName())
						}
					case GotBreakTypePage:
						{
							report(mt.ID, lineNbr, GotBreakTypePage.stateName())
						}
					case GotBlock:
						{
							report(mt.ID, lineNbr, GotBlock.stateName())
						}
					case GotColon:
						{
							report(mt.ID, lineNbr, GotColon.stateName())
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
								report(mt.ID, lineNbr, GotDay.stateName())
							}
						}
					case GotDialog:
					case GotEndActor:
						{
							report(mt.ID, lineNbr, GotEndActor.stateName())
						}
					case GotEndBreak:
						{
							report(mt.ID, lineNbr, GotEndBreak.stateName())
						}
					case GotEndBlock:
						{
							report(mt.ID, lineNbr, GotEndBlock.stateName())
						}
					case GotEndDate:
						{
							report(mt.ID, lineNbr, GotEndDate.stateName())
						}
					case GotEndDialog:
						{
							report(mt.ID, lineNbr, GotEndDialog.stateName())
						}
					case GotEndEndPageHeaderEven:
						{
							report(mt.ID, lineNbr, GotEndEndPageHeaderEven.stateName())
						}
					case GotEndEndPageHeaderOdd:
						{
							report(mt.ID, lineNbr, GotEndEndPageHeaderOdd.stateName())
						}
					case GotEndHymn:
						{
							report(mt.ID, lineNbr, GotEndHymn.stateName())
						}
					case GotEndInsert:
						{
							report(mt.ID, lineNbr, GotEndInsert.stateName())
						}
					case GotEndMedia:
						{
							report(mt.ID, lineNbr, GotEndMedia.stateName())
						}
					case GotEndPageFooterEven:
						{
							report(mt.ID, lineNbr, GotEndPageFooterEven.stateName())
						}
					case GotEndPageFooterOdd:
						{
							report(mt.ID, lineNbr, GotEndPageFooterOdd.stateName())
						}
					case GotEndPara:
						{
							report(mt.ID, lineNbr, GotEndPara.stateName())
						}
					case GotEndPassthroughHTML:
						{
							report(mt.ID, lineNbr, GotEndPassthroughHTML.stateName())
						}
					case GotEndPreface:
						{
							report(mt.ID, lineNbr, GotEndPreface.stateName())
						}
					case GotEndReading:
						{
							report(mt.ID, lineNbr, GotEndReading.stateName())
						}
					case GotEndRubric:
						{
							report(mt.ID, lineNbr, GotEndRubric.stateName())
						}
					case GotEndSection:
						{
							report(mt.ID, lineNbr, GotEndSection.stateName())
						}
					case GotEndSetMcDay:
						{
							report(mt.ID, lineNbr, GotEndSetMcDay.stateName())
						}
					case GotEndSetPageNumber:
						{
							report(mt.ID, lineNbr, GotEndSetPageNumber.stateName())
						}
					case GotEndSubTitle:
						{
							report(mt.ID, lineNbr, GotEndSubTitle.stateName())
						}
					case GotEndSwitchVersion:
						{
							report(mt.ID, lineNbr, GotEndSwitchVersion.stateName())
						}
					case GotEndTemplateCommemoration:
						{
							report(mt.ID, lineNbr, GotEndTemplateCommemoration.stateName())
						}
					case GotEndTitle:
						{
							report(mt.ID, lineNbr, GotEndTitle.stateName())
						}
					case GotEndVerse:
						{
							report(mt.ID, lineNbr, GotEndVerse.stateName())
						}
					case GotEndWhen:
						{
							report(mt.ID, lineNbr, GotEndWhen.stateName())
						}
					case GotEqualSign:
						{
							report(mt.ID, lineNbr, GotEqualSign.stateName())
						}
					case GotFooterColumnCenter:
						{
							report(mt.ID, lineNbr, GotFooterColumnCenter.stateName())
						}
					case GotFooterColumnLeft:
						{
							report(mt.ID, lineNbr, GotFooterColumnLeft.stateName())
						}
					case GotFooterColumnRight:
						{
							report(mt.ID, lineNbr, GotFooterColumnRight.stateName())
						}
					case GotFooterCommemoration:
						{
							report(mt.ID, lineNbr, GotFooterCommemoration.stateName())
						}
					case GotFooterDate:
						{
							report(mt.ID, lineNbr, GotFooterDate.stateName())
						}
					case GotFooterLookup:
						{
							report(mt.ID, lineNbr, GotFooterLookup.stateName())
						}
					case GotFooterPageNumber:
						{
							report(mt.ID, lineNbr, GotFooterPageNumber.stateName())
						}
					case GotFooterTitle:
						{
							report(mt.ID, lineNbr, GotFooterTitle.stateName())
						}
					case GotHeaderFooterText:
						{
							report(mt.ID, lineNbr, GotHeaderFooterText.stateName())
						}
					case GotHymn:
					case GotImport:
						{
							switch token {
							case ".":
								if s.Peek() == '*' {
									mt.addImport(identifierSB.String())
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
										identifierSB.WriteString(parts[0])
										sb.WriteString(identifierSB.String())
										mt.addImport(identifierSB.String())
									} else {
										if identifierSB.Len() > 0 { // append the period
											identifierSB.WriteString(token)
										}
									}
								}
							case "*": // ignore
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
									identifierSB.WriteString(parts[0])
									sb.WriteString(identifierSB.String())
									sb.WriteString(pathSeparator)
									mt.addImport(sb.String())
									identifierSB.Reset()
								} else {
									identifierSB.WriteString(token)
								}
							}
						}
					case GotInsertPreface:
						{
							report(mt.ID, lineNbr, GotInsertPreface.stateName())
						}
					case GotInsertSection:
						{
							identifierSB.WriteString(token)
						}
					case GotInsertTemplate:
						{
							lineSB.WriteString(token)
						}
					case GotMedia:
						{
							report(mt.ID, lineNbr, GotMedia.stateName())
						}
					case GotMediaOff:
						{
							report(mt.ID, lineNbr, GotMediaOff.stateName())
						}
					case GotMonth:
						if inHead {
							var i int
							if i, err := strconv.Atoi(token); err != nil {
								fmt.Printf("month=%d, type: %T\n", i, i)
							}
							mt.Month = i
						} else {
							report(mt.ID, lineNbr, GotMonth.stateName())
						}
					case GotOtherwise:
						{
							report(mt.ID, lineNbr, GotOtherwise.stateName())
						}
					case GotPageFooterEven:
						{
							report(mt.ID, lineNbr, GotPageFooterEven.stateName())
						}
					case GotPageFooterOdd:
						{
							report(mt.ID, lineNbr, GotPageFooterOdd.stateName())
						}
					case GotPageHeaderEven:
						{
							report(mt.ID, lineNbr, GotPageHeaderEven.stateName())
						}
					case GotPageHeaderOdd:
						{
							report(mt.ID, lineNbr, GotPageHeaderOdd.stateName())
						}
					case GotPara:
					case GotPassthroughHTML:
						{
							report(mt.ID, lineNbr, GotPassthroughHTML.stateName())
						}
					case GotPreface:
						{
							report(mt.ID, lineNbr, GotPreface.stateName())
						}
					case GotReading:
						{
							lineSB.WriteString(fmt.Sprintf("p.reading "))
						}
					case GotRid:
						{
							report(mt.ID, lineNbr, GotRid.stateName())
						}
					case GotRole:
						{
							report(mt.ID, lineNbr, GotRole.stateName())
						}
					case GotRubric:
						{
							report(mt.ID, lineNbr, GotRubric.stateName())
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
							mt.Writer.WriteString(fmt.Sprintf("id: \"%s\"\n", mt.ID))
							internalPaths.Push(currentInternalPath)
							currentInternalPath = currentInternalPath + "." + token
							currentState = Neutral
						}
					case GotSetDate:
						{
							report(mt.ID, lineNbr, GotSetDate.stateName())
						}
					case GotSetMcDay:
						{
							report(mt.ID, lineNbr, GotSetMcDay.stateName())
						}
					case GotSetPageNumber:
						{
							report(mt.ID, lineNbr, GotSetPageNumber.stateName())
						}
					case GotSid:
						{
							lineSB.WriteString(token)
						}
					case GotStatus:
						{
							var status string
							switch token {
							case "NA":
								{
								status = statuses.NA.String()
								}
							case "DRAFT":
								{
									status = statuses.Draft.String()
								}
							case "REVIEW":
								{
									status = statuses.Review.String()
								}
							case "FINAL":
								{
									status = statuses.Final.String()
								}
							}
							lineSB.WriteString(strings.ToLower(status))
						}
					case GotSubTitle:
						{
							report(mt.ID, lineNbr, GotSubTitle.stateName())
						}
					case GotSwitchVersion:
						{
							report(mt.ID, lineNbr, GotSwitchVersion.stateName())
						}
					case GotTemplateCommemoration:
						{
							report(mt.ID, lineNbr, GotTemplateCommemoration.stateName())
						}
					case GotTemplateName:
						{
						lineSB.WriteString(token)
						}
					case GotTemplateStatusDraft:
						{
							report(mt.ID, lineNbr, GotTemplateStatusDraft.stateName())
						}
					case GotTemplateStatusFinal:
						{
							report(mt.ID, lineNbr, GotTemplateStatusFinal.stateName())
						}
					case GotTemplateStatusNA:
						{
							report(mt.ID, lineNbr, GotTemplateStatusNA.stateName())
						}
					case GotTemplateStatusReview:
						{
							report(mt.ID, lineNbr, GotTemplateStatusReview.stateName())
						}
					case GotTemplateTitle:
						{
							report(mt.ID, lineNbr, GotTemplateTitle.stateName())
						}
					case GotTitle:
					case GotUse:
						{
							report(mt.ID, lineNbr, GotUse.stateName())
						}
					case GotVerse:
						{
							report(mt.ID, lineNbr, GotVerse.stateName())
						}
					case GotWhenDateIs:
						{
							report(mt.ID, lineNbr, GotWhenDateIs.stateName())
						}
					case GotWhenExists:
						{
							report(mt.ID, lineNbr, GotWhenExists.stateName())
						}
					case GotWhenLukanCycleDayIs:
						{
							report(mt.ID, lineNbr, GotWhenLukanCycleDayIs.stateName())
						}
					case GotWhenModeOfWeekIs:
						{
							report(mt.ID, lineNbr, GotWhenModeOfWeekIs.stateName())
						}
					case GotWhenMovableCycleDayIs:
						{
							report(mt.ID, lineNbr, GotWhenMovableCycleDayIs.stateName())
						}
					case GotWhenNameOfDayIs:
						{
							report(mt.ID, lineNbr, GotWhenNameOfDayIs.stateName())
						}
					case GotWhenSundayAfterElevationOfCross:
						{
							report(mt.ID, lineNbr, GotWhenSundayAfterElevationOfCross.stateName())
						}
					case GotWhenSundaysBeforeTriodionIs:
						{
							report(mt.ID, lineNbr, GotWhenSundaysBeforeTriodionIs.stateName())
						}
					case GotYear:
						if inHead {
							var i int
							if i, err := strconv.Atoi(token); err != nil {
								fmt.Printf("year=%d, type: %T\n", i, i)
							}
							mt.Year = i
						} else {
							report(mt.ID, lineNbr, GotYear.stateName())
						}
					}
				}
			}
		}
		if lineSB.Len() > 0 || identifierSB.Len() > 0 {
			switch currentState {
			case GotImport: // lml does not have scope imports
				//{
				//	if identifierSB.Len() > 0 {
				//		theImport := identifierSB.String()
				//		if strings.HasPrefix(theImport,".") && strings.HasSuffix(theImport, ".") {
				//		 // ignore
				//		} else {
				//			switch theImport {
				//			case "iTags", "bTags", "roles", "rules": // ignore
				//			default:
				//				{
				//					lineSB.WriteString("import \"")
				//					lineSB.WriteString(identifierSB.String())
				//					lineSB.WriteString("\"")
				//				}
				//			}
				//		}
				//		identifierSB.Reset()
				//	}
				//}
			case GotInsertSection:
			case GotStatus:
				{
					lineSB.WriteString("\n")
				}
			case GotTemplateName:
				{
					lineSB.WriteString("\"\n")
				}
			}
			if lineSB.Len() > 0 {
				lineOut := lineSB.String()
				mt.Writer.WriteString(fmt.Sprintf("%s\n", lineOut))
				// TODO: remove next line for production
				mt.Writer.Flush()
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
// resolvePath attempts to resolve the path in the following order:
// 1. path as is, using templates tree
// 2. path as is, using sections tree
// 3. path scoped to containing sections, using sections tree
// 4. path scoped used imports, and reading sections tree.
func resolvePath(pathToResolve, currentInternalPath string, mt *MetaTemplate, tTree, sTree *btree.Tree) string {
	// see if the path to resolve is actually a template name
	if v, ok := tTree.Get(pathToResolve); ok {
		return v.(TemplateNode).ID
	} else if v, ok := sTree.Get(pathToResolve); ok { // check if path exists in sections tree

		return v.(SectionNode).Path
	} else {
		// try  to resolve path within scope of containing sections within the current template
		parts := strings.Split(currentInternalPath,".")
		l := len(parts)
		for i, _ := range parts {
			sectionId := strings.Join(parts[0:l-i],".") + "." + pathToResolve
			if v, ok := mt.Sections.Get(sectionId); ok {
				return v.(string)
				break
			}
		}
		// if we get here, we have not resolved the path yet.
		// try to resolve the path using the imports
		for _ , impt := range mt.Imports {
			sectionId := impt + "." + pathToResolve
			if v, ok := sTree.Get(sectionId); ok {
				return v.(SectionNode).Path
				break
			}
		}
	}
	return ""
}
func report(file string, line int, state string) {
	fmt.Printf("\r%s %d Unhandled: %s", file, line, state)
}