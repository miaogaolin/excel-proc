// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package condition // Condition
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by ConditionParser.
type ConditionVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ConditionParser#GetNumExpr.
	VisitGetNumExpr(ctx *GetNumExprContext) interface{}

	// Visit a parse tree produced by ConditionParser#Parent.
	VisitParent(ctx *ParentContext) interface{}

	// Visit a parse tree produced by ConditionParser#GetStringExpr.
	VisitGetStringExpr(ctx *GetStringExprContext) interface{}

	// Visit a parse tree produced by ConditionParser#GetArrayExpr.
	VisitGetArrayExpr(ctx *GetArrayExprContext) interface{}

	// Visit a parse tree produced by ConditionParser#AndOr.
	VisitAndOr(ctx *AndOrContext) interface{}

	// Visit a parse tree produced by ConditionParser#conditionKey.
	VisitConditionKey(ctx *ConditionKeyContext) interface{}

	// Visit a parse tree produced by ConditionParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by ConditionParser#number.
	VisitNumber(ctx *NumberContext) interface{}
}
