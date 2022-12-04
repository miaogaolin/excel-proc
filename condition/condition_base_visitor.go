// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package condition // Condition
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseConditionVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseConditionVisitor) VisitGetNumExpr(ctx *GetNumExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConditionVisitor) VisitParent(ctx *ParentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConditionVisitor) VisitGetStringExpr(ctx *GetStringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConditionVisitor) VisitGetArrayExpr(ctx *GetArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConditionVisitor) VisitAndOr(ctx *AndOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConditionVisitor) VisitConditionKey(ctx *ConditionKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConditionVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConditionVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}
