// Code generated from LML.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // LML

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LMLParser.
type LMLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LMLParser#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by LMLParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by LMLParser#propertyBlock.
	VisitPropertyBlock(ctx *PropertyBlockContext) interface{}

	// Visit a parse tree produced by LMLParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by LMLParser#dowName.
	VisitDowName(ctx *DowNameContext) interface{}

	// Visit a parse tree produced by LMLParser#insert.
	VisitInsert(ctx *InsertContext) interface{}

	// Visit a parse tree produced by LMLParser#para.
	VisitPara(ctx *ParaContext) interface{}

	// Visit a parse tree produced by LMLParser#pspan.
	VisitPspan(ctx *PspanContext) interface{}

	// Visit a parse tree produced by LMLParser#span.
	VisitSpan(ctx *SpanContext) interface{}

	// Visit a parse tree produced by LMLParser#media.
	VisitMedia(ctx *MediaContext) interface{}

	// Visit a parse tree produced by LMLParser#nid.
	VisitNid(ctx *NidContext) interface{}

	// Visit a parse tree produced by LMLParser#monthName.
	VisitMonthName(ctx *MonthNameContext) interface{}

	// Visit a parse tree produced by LMLParser#position.
	VisitPosition(ctx *PositionContext) interface{}

	// Visit a parse tree produced by LMLParser#positionType.
	VisitPositionType(ctx *PositionTypeContext) interface{}

	// Visit a parse tree produced by LMLParser#directive.
	VisitDirective(ctx *DirectiveContext) interface{}

	// Visit a parse tree produced by LMLParser#lookup.
	VisitLookup(ctx *LookupContext) interface{}

	// Visit a parse tree produced by LMLParser#rid.
	VisitRid(ctx *RidContext) interface{}

	// Visit a parse tree produced by LMLParser#override.
	VisitOverride(ctx *OverrideContext) interface{}

	// Visit a parse tree produced by LMLParser#overrideMode.
	VisitOverrideMode(ctx *OverrideModeContext) interface{}

	// Visit a parse tree produced by LMLParser#overrideDay.
	VisitOverrideDay(ctx *OverrideDayContext) interface{}

	// Visit a parse tree produced by LMLParser#sid.
	VisitSid(ctx *SidContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplCalendar.
	VisitTmplCalendar(ctx *TmplCalendarContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplHtmlCss.
	VisitTmplHtmlCss(ctx *TmplHtmlCssContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplPdfCss.
	VisitTmplPdfCss(ctx *TmplPdfCssContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplDay.
	VisitTmplDay(ctx *TmplDayContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplID.
	VisitTmplID(ctx *TmplIDContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplMonth.
	VisitTmplMonth(ctx *TmplMonthContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplPageHeaderEven.
	VisitTmplPageHeaderEven(ctx *TmplPageHeaderEvenContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplPageFooterEven.
	VisitTmplPageFooterEven(ctx *TmplPageFooterEvenContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplPageHeaderOdd.
	VisitTmplPageHeaderOdd(ctx *TmplPageHeaderOddContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplPageFooterOdd.
	VisitTmplPageFooterOdd(ctx *TmplPageFooterOddContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplPageHeader.
	VisitTmplPageHeader(ctx *TmplPageHeaderContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplPageFooter.
	VisitTmplPageFooter(ctx *TmplPageFooterContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplPageNumber.
	VisitTmplPageNumber(ctx *TmplPageNumberContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplStatus.
	VisitTmplStatus(ctx *TmplStatusContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplTitle.
	VisitTmplTitle(ctx *TmplTitleContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplType.
	VisitTmplType(ctx *TmplTypeContext) interface{}

	// Visit a parse tree produced by LMLParser#tmplYear.
	VisitTmplYear(ctx *TmplYearContext) interface{}

	// Visit a parse tree produced by LMLParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LMLParser#switchBlock.
	VisitSwitchBlock(ctx *SwitchBlockContext) interface{}

	// Visit a parse tree produced by LMLParser#switchBlockStatementGroup.
	VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{}

	// Visit a parse tree produced by LMLParser#switchLabel.
	VisitSwitchLabel(ctx *SwitchLabelContext) interface{}

	// Visit a parse tree produced by LMLParser#integerExpression.
	VisitIntegerExpression(ctx *IntegerExpressionContext) interface{}

	// Visit a parse tree produced by LMLParser#integerList.
	VisitIntegerList(ctx *IntegerListContext) interface{}

	// Visit a parse tree produced by LMLParser#dowExpression.
	VisitDowExpression(ctx *DowExpressionContext) interface{}

	// Visit a parse tree produced by LMLParser#dowList.
	VisitDowList(ctx *DowListContext) interface{}

	// Visit a parse tree produced by LMLParser#ldpInt.
	VisitLdpInt(ctx *LdpIntContext) interface{}

	// Visit a parse tree produced by LMLParser#monthDay.
	VisitMonthDay(ctx *MonthDayContext) interface{}

	// Visit a parse tree produced by LMLParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}
}
