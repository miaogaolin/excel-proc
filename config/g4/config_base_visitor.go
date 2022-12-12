// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package g4 // Config
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseConfigVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseConfigVisitor) VisitFieldStat(ctx *FieldStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConfigVisitor) VisitNotFieldStat(ctx *NotFieldStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConfigVisitor) VisitSections(ctx *SectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConfigVisitor) VisitGetSection(ctx *GetSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConfigVisitor) VisitFields(ctx *FieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConfigVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConfigVisitor) VisitFieldName(ctx *FieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConfigVisitor) VisitFieldValue(ctx *FieldValueContext) interface{} {
	return v.VisitChildren(ctx)
}
