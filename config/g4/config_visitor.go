// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package g4 // Config
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by ConfigParser.
type ConfigVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ConfigParser#FieldStat.
	VisitFieldStat(ctx *FieldStatContext) interface{}

	// Visit a parse tree produced by ConfigParser#NotFieldStat.
	VisitNotFieldStat(ctx *NotFieldStatContext) interface{}

	// Visit a parse tree produced by ConfigParser#sections.
	VisitSections(ctx *SectionsContext) interface{}

	// Visit a parse tree produced by ConfigParser#GetSection.
	VisitGetSection(ctx *GetSectionContext) interface{}

	// Visit a parse tree produced by ConfigParser#fields.
	VisitFields(ctx *FieldsContext) interface{}

	// Visit a parse tree produced by ConfigParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by ConfigParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by ConfigParser#fieldValue.
	VisitFieldValue(ctx *FieldValueContext) interface{}
}
