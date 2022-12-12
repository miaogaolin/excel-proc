package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/miaogaolin/excel-proc/config/g4"
	"github.com/miaogaolin/excel-proc/utils"
	"github.com/pkg/errors"
)

type ConfigSection struct {
	Template      string
	Condition     string
	ConditionLine int
}
type Config struct {
	Fields   map[string]string
	Sections []ConfigSection
}

func ReadConfig(path string) (*Config, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	visitor := new(configVisitor)

	// 新家一个 CharStream
	input := antlr.NewInputStream(string(content))

	lexer := g4.NewConfigLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(visitor)

	// 新建一个此法符号的缓冲区，用于存储此法分析器将生成的词法符号
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	// 新建一个语法分析器，处理词法符号缓冲区中的内容
	parser := g4.NewConfigParser(tokens)
	parser.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	parser.RemoveErrorListeners()
	parser.AddErrorListener(visitor)

	tree := parser.Stat()

	res := visitor.Visit(tree)

	var errStr string
	// parser error
	for _, v := range visitor.Errors {
		errStr += v.Error() + utils.LineSeperator()
	}

	if errStr != "" {
		return nil, errors.New(errStr)
	}

	if v, ok := res.(error); ok {
		return nil, v
	}
	return res.(*Config), nil
}

type configVisitor struct {
	g4.BaseConfigVisitor
	antlr.DefaultErrorListener
	fields   map[string]string
	sections []ConfigSection
	Errors   []error
}

func (c *configVisitor) Visit(tree antlr.ParseTree) interface{} {
	if c.fields == nil {
		c.fields = make(map[string]string)
	}
	switch tree.(type) {
	case *g4.FieldValueContext:
		return tree.Accept(c)
	default:
		tree.Accept(c)
	}

	return &Config{
		Fields:   c.fields,
		Sections: c.sections,
	}
}

func (c *configVisitor) VisitField(ctx *g4.FieldContext) interface{} {
	name := ctx.FieldName().GetText()
	val := c.Visit(ctx.FieldValue()).(string)
	c.fields[name] = val
	return nil
}

func (c *configVisitor) VisitFieldStat(ctx *g4.FieldStatContext) interface{} {
	c.Visit(ctx.Sections())
	c.Visit(ctx.Fields())
	return nil
}

func (c *configVisitor) VisitNotFieldStat(ctx *g4.NotFieldStatContext) interface{} {
	return c.Visit(ctx.Sections())
}

func (c *configVisitor) VisitSections(ctx *g4.SectionsContext) interface{} {
	for _, v := range ctx.AllSection() {
		c.Visit(v)
	}
	return nil
}

func (c *configVisitor) VisitGetSection(ctx *g4.GetSectionContext) interface{} {
	trimText := utils.LineSeperator()
	if ctx.EOF() != nil {
		trimText += ctx.EOF().GetText()
	}
	txt := strings.Trim(ctx.GetText(), trimText)
	var (
		conditionEnd int
		condition    string
	)
	// condition
	if strings.HasPrefix(txt, ";") {
		conditionEnd = strings.Index(txt, utils.LineSeperator())
		condition = txt[1:conditionEnd]
	}

	// template
	tpl := strings.Trim(txt[conditionEnd:], utils.LineSeperator())

	if tpl == "" {
		return fmt.Errorf("line:%d,column:%d, template is empty", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	c.sections = append(c.sections, ConfigSection{
		Condition:     condition,
		Template:      tpl,
		ConditionLine: ctx.GetStart().GetLine(),
	})
	return nil
}

func (c *configVisitor) VisitFields(ctx *g4.FieldsContext) interface{} {
	for _, v := range ctx.AllField() {
		c.Visit(v)
	}
	return nil
}

func (c *configVisitor) VisitFieldValue(ctx *g4.FieldValueContext) interface{} {
	if ctx.STRING() != nil {
		txt := ctx.STRING().GetText()
		if strings.HasPrefix(txt, "'") {
			return strings.Trim(txt, "'")
		}
		return strings.Trim(txt, `"`)
	}
	return ctx.GetText()
}

func (c *configVisitor) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, fmt.Errorf("line:%d, column:%d, %s", line, column, msg))
}
