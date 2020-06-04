// Code generated from LML.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // LML

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLMLListener is a complete listener for a parse tree produced by LMLParser.
type BaseLMLListener struct{}

var _ LMLListener = &BaseLMLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLMLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLMLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLMLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLMLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTemplate is called when production template is entered.
func (s *BaseLMLListener) EnterTemplate(ctx *TemplateContext) {}

// ExitTemplate is called when production template is exited.
func (s *BaseLMLListener) ExitTemplate(ctx *TemplateContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseLMLListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseLMLListener) ExitProperty(ctx *PropertyContext) {}

// EnterPropertyBlock is called when production propertyBlock is entered.
func (s *BaseLMLListener) EnterPropertyBlock(ctx *PropertyBlockContext) {}

// ExitPropertyBlock is called when production propertyBlock is exited.
func (s *BaseLMLListener) ExitPropertyBlock(ctx *PropertyBlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseLMLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseLMLListener) ExitStatement(ctx *StatementContext) {}

// EnterDowName is called when production dowName is entered.
func (s *BaseLMLListener) EnterDowName(ctx *DowNameContext) {}

// ExitDowName is called when production dowName is exited.
func (s *BaseLMLListener) ExitDowName(ctx *DowNameContext) {}

// EnterInsert is called when production insert is entered.
func (s *BaseLMLListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production insert is exited.
func (s *BaseLMLListener) ExitInsert(ctx *InsertContext) {}

// EnterPara is called when production para is entered.
func (s *BaseLMLListener) EnterPara(ctx *ParaContext) {}

// ExitPara is called when production para is exited.
func (s *BaseLMLListener) ExitPara(ctx *ParaContext) {}

// EnterSpan is called when production span is entered.
func (s *BaseLMLListener) EnterSpan(ctx *SpanContext) {}

// ExitSpan is called when production span is exited.
func (s *BaseLMLListener) ExitSpan(ctx *SpanContext) {}

// EnterMedia is called when production media is entered.
func (s *BaseLMLListener) EnterMedia(ctx *MediaContext) {}

// ExitMedia is called when production media is exited.
func (s *BaseLMLListener) ExitMedia(ctx *MediaContext) {}

// EnterNid is called when production nid is entered.
func (s *BaseLMLListener) EnterNid(ctx *NidContext) {}

// ExitNid is called when production nid is exited.
func (s *BaseLMLListener) ExitNid(ctx *NidContext) {}

// EnterMonthName is called when production monthName is entered.
func (s *BaseLMLListener) EnterMonthName(ctx *MonthNameContext) {}

// ExitMonthName is called when production monthName is exited.
func (s *BaseLMLListener) ExitMonthName(ctx *MonthNameContext) {}

// EnterPosition is called when production position is entered.
func (s *BaseLMLListener) EnterPosition(ctx *PositionContext) {}

// ExitPosition is called when production position is exited.
func (s *BaseLMLListener) ExitPosition(ctx *PositionContext) {}

// EnterPositionType is called when production positionType is entered.
func (s *BaseLMLListener) EnterPositionType(ctx *PositionTypeContext) {}

// ExitPositionType is called when production positionType is exited.
func (s *BaseLMLListener) ExitPositionType(ctx *PositionTypeContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseLMLListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseLMLListener) ExitDirective(ctx *DirectiveContext) {}

// EnterLookup is called when production lookup is entered.
func (s *BaseLMLListener) EnterLookup(ctx *LookupContext) {}

// ExitLookup is called when production lookup is exited.
func (s *BaseLMLListener) ExitLookup(ctx *LookupContext) {}

// EnterRid is called when production rid is entered.
func (s *BaseLMLListener) EnterRid(ctx *RidContext) {}

// ExitRid is called when production rid is exited.
func (s *BaseLMLListener) ExitRid(ctx *RidContext) {}

// EnterOverride is called when production override is entered.
func (s *BaseLMLListener) EnterOverride(ctx *OverrideContext) {}

// ExitOverride is called when production override is exited.
func (s *BaseLMLListener) ExitOverride(ctx *OverrideContext) {}

// EnterOverrideMode is called when production overrideMode is entered.
func (s *BaseLMLListener) EnterOverrideMode(ctx *OverrideModeContext) {}

// ExitOverrideMode is called when production overrideMode is exited.
func (s *BaseLMLListener) ExitOverrideMode(ctx *OverrideModeContext) {}

// EnterOverrideDay is called when production overrideDay is entered.
func (s *BaseLMLListener) EnterOverrideDay(ctx *OverrideDayContext) {}

// ExitOverrideDay is called when production overrideDay is exited.
func (s *BaseLMLListener) ExitOverrideDay(ctx *OverrideDayContext) {}

// EnterSid is called when production sid is entered.
func (s *BaseLMLListener) EnterSid(ctx *SidContext) {}

// ExitSid is called when production sid is exited.
func (s *BaseLMLListener) ExitSid(ctx *SidContext) {}

// EnterTmplCalendar is called when production tmplCalendar is entered.
func (s *BaseLMLListener) EnterTmplCalendar(ctx *TmplCalendarContext) {}

// ExitTmplCalendar is called when production tmplCalendar is exited.
func (s *BaseLMLListener) ExitTmplCalendar(ctx *TmplCalendarContext) {}

// EnterTmplHtmlCss is called when production tmplHtmlCss is entered.
func (s *BaseLMLListener) EnterTmplHtmlCss(ctx *TmplHtmlCssContext) {}

// ExitTmplHtmlCss is called when production tmplHtmlCss is exited.
func (s *BaseLMLListener) ExitTmplHtmlCss(ctx *TmplHtmlCssContext) {}

// EnterTmplPdfCss is called when production tmplPdfCss is entered.
func (s *BaseLMLListener) EnterTmplPdfCss(ctx *TmplPdfCssContext) {}

// ExitTmplPdfCss is called when production tmplPdfCss is exited.
func (s *BaseLMLListener) ExitTmplPdfCss(ctx *TmplPdfCssContext) {}

// EnterTmplDay is called when production tmplDay is entered.
func (s *BaseLMLListener) EnterTmplDay(ctx *TmplDayContext) {}

// ExitTmplDay is called when production tmplDay is exited.
func (s *BaseLMLListener) ExitTmplDay(ctx *TmplDayContext) {}

// EnterTmplID is called when production tmplID is entered.
func (s *BaseLMLListener) EnterTmplID(ctx *TmplIDContext) {}

// ExitTmplID is called when production tmplID is exited.
func (s *BaseLMLListener) ExitTmplID(ctx *TmplIDContext) {}

// EnterTmplMonth is called when production tmplMonth is entered.
func (s *BaseLMLListener) EnterTmplMonth(ctx *TmplMonthContext) {}

// ExitTmplMonth is called when production tmplMonth is exited.
func (s *BaseLMLListener) ExitTmplMonth(ctx *TmplMonthContext) {}

// EnterTmplPageHeaderEven is called when production tmplPageHeaderEven is entered.
func (s *BaseLMLListener) EnterTmplPageHeaderEven(ctx *TmplPageHeaderEvenContext) {}

// ExitTmplPageHeaderEven is called when production tmplPageHeaderEven is exited.
func (s *BaseLMLListener) ExitTmplPageHeaderEven(ctx *TmplPageHeaderEvenContext) {}

// EnterTmplPageFooterEven is called when production tmplPageFooterEven is entered.
func (s *BaseLMLListener) EnterTmplPageFooterEven(ctx *TmplPageFooterEvenContext) {}

// ExitTmplPageFooterEven is called when production tmplPageFooterEven is exited.
func (s *BaseLMLListener) ExitTmplPageFooterEven(ctx *TmplPageFooterEvenContext) {}

// EnterTmplPageHeaderOdd is called when production tmplPageHeaderOdd is entered.
func (s *BaseLMLListener) EnterTmplPageHeaderOdd(ctx *TmplPageHeaderOddContext) {}

// ExitTmplPageHeaderOdd is called when production tmplPageHeaderOdd is exited.
func (s *BaseLMLListener) ExitTmplPageHeaderOdd(ctx *TmplPageHeaderOddContext) {}

// EnterTmplPageFooterOdd is called when production tmplPageFooterOdd is entered.
func (s *BaseLMLListener) EnterTmplPageFooterOdd(ctx *TmplPageFooterOddContext) {}

// ExitTmplPageFooterOdd is called when production tmplPageFooterOdd is exited.
func (s *BaseLMLListener) ExitTmplPageFooterOdd(ctx *TmplPageFooterOddContext) {}

// EnterTmplPageHeader is called when production tmplPageHeader is entered.
func (s *BaseLMLListener) EnterTmplPageHeader(ctx *TmplPageHeaderContext) {}

// ExitTmplPageHeader is called when production tmplPageHeader is exited.
func (s *BaseLMLListener) ExitTmplPageHeader(ctx *TmplPageHeaderContext) {}

// EnterTmplPageFooter is called when production tmplPageFooter is entered.
func (s *BaseLMLListener) EnterTmplPageFooter(ctx *TmplPageFooterContext) {}

// ExitTmplPageFooter is called when production tmplPageFooter is exited.
func (s *BaseLMLListener) ExitTmplPageFooter(ctx *TmplPageFooterContext) {}

// EnterTmplPageNumber is called when production tmplPageNumber is entered.
func (s *BaseLMLListener) EnterTmplPageNumber(ctx *TmplPageNumberContext) {}

// ExitTmplPageNumber is called when production tmplPageNumber is exited.
func (s *BaseLMLListener) ExitTmplPageNumber(ctx *TmplPageNumberContext) {}

// EnterTmplStatus is called when production tmplStatus is entered.
func (s *BaseLMLListener) EnterTmplStatus(ctx *TmplStatusContext) {}

// ExitTmplStatus is called when production tmplStatus is exited.
func (s *BaseLMLListener) ExitTmplStatus(ctx *TmplStatusContext) {}

// EnterTmplTitle is called when production tmplTitle is entered.
func (s *BaseLMLListener) EnterTmplTitle(ctx *TmplTitleContext) {}

// ExitTmplTitle is called when production tmplTitle is exited.
func (s *BaseLMLListener) ExitTmplTitle(ctx *TmplTitleContext) {}

// EnterTmplType is called when production tmplType is entered.
func (s *BaseLMLListener) EnterTmplType(ctx *TmplTypeContext) {}

// ExitTmplType is called when production tmplType is exited.
func (s *BaseLMLListener) ExitTmplType(ctx *TmplTypeContext) {}

// EnterTmplYear is called when production tmplYear is entered.
func (s *BaseLMLListener) EnterTmplYear(ctx *TmplYearContext) {}

// ExitTmplYear is called when production tmplYear is exited.
func (s *BaseLMLListener) ExitTmplYear(ctx *TmplYearContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLMLListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLMLListener) ExitBlock(ctx *BlockContext) {}

// EnterSwitchBlock is called when production switchBlock is entered.
func (s *BaseLMLListener) EnterSwitchBlock(ctx *SwitchBlockContext) {}

// ExitSwitchBlock is called when production switchBlock is exited.
func (s *BaseLMLListener) ExitSwitchBlock(ctx *SwitchBlockContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BaseLMLListener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BaseLMLListener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BaseLMLListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BaseLMLListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterIntegerExpression is called when production integerExpression is entered.
func (s *BaseLMLListener) EnterIntegerExpression(ctx *IntegerExpressionContext) {}

// ExitIntegerExpression is called when production integerExpression is exited.
func (s *BaseLMLListener) ExitIntegerExpression(ctx *IntegerExpressionContext) {}

// EnterIntegerList is called when production integerList is entered.
func (s *BaseLMLListener) EnterIntegerList(ctx *IntegerListContext) {}

// ExitIntegerList is called when production integerList is exited.
func (s *BaseLMLListener) ExitIntegerList(ctx *IntegerListContext) {}

// EnterDowExpression is called when production dowExpression is entered.
func (s *BaseLMLListener) EnterDowExpression(ctx *DowExpressionContext) {}

// ExitDowExpression is called when production dowExpression is exited.
func (s *BaseLMLListener) ExitDowExpression(ctx *DowExpressionContext) {}

// EnterDowList is called when production dowList is entered.
func (s *BaseLMLListener) EnterDowList(ctx *DowListContext) {}

// ExitDowList is called when production dowList is exited.
func (s *BaseLMLListener) ExitDowList(ctx *DowListContext) {}

// EnterLdpInt is called when production ldpInt is entered.
func (s *BaseLMLListener) EnterLdpInt(ctx *LdpIntContext) {}

// ExitLdpInt is called when production ldpInt is exited.
func (s *BaseLMLListener) ExitLdpInt(ctx *LdpIntContext) {}

// EnterMonthDay is called when production monthDay is entered.
func (s *BaseLMLListener) EnterMonthDay(ctx *MonthDayContext) {}

// ExitMonthDay is called when production monthDay is exited.
func (s *BaseLMLListener) ExitMonthDay(ctx *MonthDayContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLMLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLMLListener) ExitExpression(ctx *ExpressionContext) {}
