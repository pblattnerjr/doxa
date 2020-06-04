// Code generated from LML.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // LML

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LMLListener is a complete listener for a parse tree produced by LMLParser.
type LMLListener interface {
	antlr.ParseTreeListener

	// EnterTemplate is called when entering the template production.
	EnterTemplate(c *TemplateContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterPropertyBlock is called when entering the propertyBlock production.
	EnterPropertyBlock(c *PropertyBlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDowName is called when entering the dowName production.
	EnterDowName(c *DowNameContext)

	// EnterInsert is called when entering the insert production.
	EnterInsert(c *InsertContext)

	// EnterPara is called when entering the para production.
	EnterPara(c *ParaContext)

	// EnterSpan is called when entering the span production.
	EnterSpan(c *SpanContext)

	// EnterMedia is called when entering the media production.
	EnterMedia(c *MediaContext)

	// EnterNid is called when entering the nid production.
	EnterNid(c *NidContext)

	// EnterMonthName is called when entering the monthName production.
	EnterMonthName(c *MonthNameContext)

	// EnterPosition is called when entering the position production.
	EnterPosition(c *PositionContext)

	// EnterPositionType is called when entering the positionType production.
	EnterPositionType(c *PositionTypeContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterLookup is called when entering the lookup production.
	EnterLookup(c *LookupContext)

	// EnterRid is called when entering the rid production.
	EnterRid(c *RidContext)

	// EnterOverride is called when entering the override production.
	EnterOverride(c *OverrideContext)

	// EnterOverrideMode is called when entering the overrideMode production.
	EnterOverrideMode(c *OverrideModeContext)

	// EnterOverrideDay is called when entering the overrideDay production.
	EnterOverrideDay(c *OverrideDayContext)

	// EnterSid is called when entering the sid production.
	EnterSid(c *SidContext)

	// EnterTmplCalendar is called when entering the tmplCalendar production.
	EnterTmplCalendar(c *TmplCalendarContext)

	// EnterTmplHtmlCss is called when entering the tmplHtmlCss production.
	EnterTmplHtmlCss(c *TmplHtmlCssContext)

	// EnterTmplPdfCss is called when entering the tmplPdfCss production.
	EnterTmplPdfCss(c *TmplPdfCssContext)

	// EnterTmplDay is called when entering the tmplDay production.
	EnterTmplDay(c *TmplDayContext)

	// EnterTmplID is called when entering the tmplID production.
	EnterTmplID(c *TmplIDContext)

	// EnterTmplMonth is called when entering the tmplMonth production.
	EnterTmplMonth(c *TmplMonthContext)

	// EnterTmplPageHeaderEven is called when entering the tmplPageHeaderEven production.
	EnterTmplPageHeaderEven(c *TmplPageHeaderEvenContext)

	// EnterTmplPageFooterEven is called when entering the tmplPageFooterEven production.
	EnterTmplPageFooterEven(c *TmplPageFooterEvenContext)

	// EnterTmplPageHeaderOdd is called when entering the tmplPageHeaderOdd production.
	EnterTmplPageHeaderOdd(c *TmplPageHeaderOddContext)

	// EnterTmplPageFooterOdd is called when entering the tmplPageFooterOdd production.
	EnterTmplPageFooterOdd(c *TmplPageFooterOddContext)

	// EnterTmplPageHeader is called when entering the tmplPageHeader production.
	EnterTmplPageHeader(c *TmplPageHeaderContext)

	// EnterTmplPageFooter is called when entering the tmplPageFooter production.
	EnterTmplPageFooter(c *TmplPageFooterContext)

	// EnterTmplPageNumber is called when entering the tmplPageNumber production.
	EnterTmplPageNumber(c *TmplPageNumberContext)

	// EnterTmplStatus is called when entering the tmplStatus production.
	EnterTmplStatus(c *TmplStatusContext)

	// EnterTmplTitle is called when entering the tmplTitle production.
	EnterTmplTitle(c *TmplTitleContext)

	// EnterTmplType is called when entering the tmplType production.
	EnterTmplType(c *TmplTypeContext)

	// EnterTmplYear is called when entering the tmplYear production.
	EnterTmplYear(c *TmplYearContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterSwitchBlock is called when entering the switchBlock production.
	EnterSwitchBlock(c *SwitchBlockContext)

	// EnterSwitchBlockStatementGroup is called when entering the switchBlockStatementGroup production.
	EnterSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// EnterSwitchLabel is called when entering the switchLabel production.
	EnterSwitchLabel(c *SwitchLabelContext)

	// EnterIntegerExpression is called when entering the integerExpression production.
	EnterIntegerExpression(c *IntegerExpressionContext)

	// EnterIntegerList is called when entering the integerList production.
	EnterIntegerList(c *IntegerListContext)

	// EnterDowExpression is called when entering the dowExpression production.
	EnterDowExpression(c *DowExpressionContext)

	// EnterDowList is called when entering the dowList production.
	EnterDowList(c *DowListContext)

	// EnterLdpInt is called when entering the ldpInt production.
	EnterLdpInt(c *LdpIntContext)

	// EnterMonthDay is called when entering the monthDay production.
	EnterMonthDay(c *MonthDayContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitPropertyBlock is called when exiting the propertyBlock production.
	ExitPropertyBlock(c *PropertyBlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDowName is called when exiting the dowName production.
	ExitDowName(c *DowNameContext)

	// ExitInsert is called when exiting the insert production.
	ExitInsert(c *InsertContext)

	// ExitPara is called when exiting the para production.
	ExitPara(c *ParaContext)

	// ExitSpan is called when exiting the span production.
	ExitSpan(c *SpanContext)

	// ExitMedia is called when exiting the media production.
	ExitMedia(c *MediaContext)

	// ExitNid is called when exiting the nid production.
	ExitNid(c *NidContext)

	// ExitMonthName is called when exiting the monthName production.
	ExitMonthName(c *MonthNameContext)

	// ExitPosition is called when exiting the position production.
	ExitPosition(c *PositionContext)

	// ExitPositionType is called when exiting the positionType production.
	ExitPositionType(c *PositionTypeContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitLookup is called when exiting the lookup production.
	ExitLookup(c *LookupContext)

	// ExitRid is called when exiting the rid production.
	ExitRid(c *RidContext)

	// ExitOverride is called when exiting the override production.
	ExitOverride(c *OverrideContext)

	// ExitOverrideMode is called when exiting the overrideMode production.
	ExitOverrideMode(c *OverrideModeContext)

	// ExitOverrideDay is called when exiting the overrideDay production.
	ExitOverrideDay(c *OverrideDayContext)

	// ExitSid is called when exiting the sid production.
	ExitSid(c *SidContext)

	// ExitTmplCalendar is called when exiting the tmplCalendar production.
	ExitTmplCalendar(c *TmplCalendarContext)

	// ExitTmplHtmlCss is called when exiting the tmplHtmlCss production.
	ExitTmplHtmlCss(c *TmplHtmlCssContext)

	// ExitTmplPdfCss is called when exiting the tmplPdfCss production.
	ExitTmplPdfCss(c *TmplPdfCssContext)

	// ExitTmplDay is called when exiting the tmplDay production.
	ExitTmplDay(c *TmplDayContext)

	// ExitTmplID is called when exiting the tmplID production.
	ExitTmplID(c *TmplIDContext)

	// ExitTmplMonth is called when exiting the tmplMonth production.
	ExitTmplMonth(c *TmplMonthContext)

	// ExitTmplPageHeaderEven is called when exiting the tmplPageHeaderEven production.
	ExitTmplPageHeaderEven(c *TmplPageHeaderEvenContext)

	// ExitTmplPageFooterEven is called when exiting the tmplPageFooterEven production.
	ExitTmplPageFooterEven(c *TmplPageFooterEvenContext)

	// ExitTmplPageHeaderOdd is called when exiting the tmplPageHeaderOdd production.
	ExitTmplPageHeaderOdd(c *TmplPageHeaderOddContext)

	// ExitTmplPageFooterOdd is called when exiting the tmplPageFooterOdd production.
	ExitTmplPageFooterOdd(c *TmplPageFooterOddContext)

	// ExitTmplPageHeader is called when exiting the tmplPageHeader production.
	ExitTmplPageHeader(c *TmplPageHeaderContext)

	// ExitTmplPageFooter is called when exiting the tmplPageFooter production.
	ExitTmplPageFooter(c *TmplPageFooterContext)

	// ExitTmplPageNumber is called when exiting the tmplPageNumber production.
	ExitTmplPageNumber(c *TmplPageNumberContext)

	// ExitTmplStatus is called when exiting the tmplStatus production.
	ExitTmplStatus(c *TmplStatusContext)

	// ExitTmplTitle is called when exiting the tmplTitle production.
	ExitTmplTitle(c *TmplTitleContext)

	// ExitTmplType is called when exiting the tmplType production.
	ExitTmplType(c *TmplTypeContext)

	// ExitTmplYear is called when exiting the tmplYear production.
	ExitTmplYear(c *TmplYearContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitSwitchBlock is called when exiting the switchBlock production.
	ExitSwitchBlock(c *SwitchBlockContext)

	// ExitSwitchBlockStatementGroup is called when exiting the switchBlockStatementGroup production.
	ExitSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// ExitSwitchLabel is called when exiting the switchLabel production.
	ExitSwitchLabel(c *SwitchLabelContext)

	// ExitIntegerExpression is called when exiting the integerExpression production.
	ExitIntegerExpression(c *IntegerExpressionContext)

	// ExitIntegerList is called when exiting the integerList production.
	ExitIntegerList(c *IntegerListContext)

	// ExitDowExpression is called when exiting the dowExpression production.
	ExitDowExpression(c *DowExpressionContext)

	// ExitDowList is called when exiting the dowList production.
	ExitDowList(c *DowListContext)

	// ExitLdpInt is called when exiting the ldpInt production.
	ExitLdpInt(c *LdpIntContext)

	// ExitMonthDay is called when exiting the monthDay production.
	ExitMonthDay(c *MonthDayContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
