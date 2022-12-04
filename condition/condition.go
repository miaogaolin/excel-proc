package condition

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/pkg/errors"
)

var (
	ErrParseData   = errors.New("parse data error")
	ErrNotFoundCol = errors.New("not field")
)

type condition struct {
	BaseConditionVisitor

	data map[string]interface{}
	Err  error
}

func Validate(data map[string]interface{}, conditionExpr string) (bool, error) {
	conError := NewConditionError()
	// 新家一个 CharStream
	input := antlr.NewInputStream(conditionExpr)

	lexer := NewConditionLexer(input)
	lexer.RemoveErrorListeners()

	// 新建一个此法符号的缓冲区，用于存储此法分析器将生成的词法符号
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	// 新建一个语法分析器，处理词法符号缓冲区中的内容
	parser := NewConditionParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(conError)

	tree := parser.Expr()

	visitor := new(condition)
	visitor.data = data
	res := visitor.Visit(tree)

	var err error
	// parser error
	for _, v := range conError.Errors {
		if err == nil {
			err = v
			continue
		}
		err = errors.Wrap(err, v.Error())
	}

	if err != nil {
		return false, err
	}
	if visitor.Err != nil {
		return false, visitor.Err
	}

	return res.(bool), nil
}

func (p *condition) Visit(tree antlr.ParseTree) interface{} {
	switch d := tree.(type) {
	case *GetNumExprContext:
		return d.Accept(p)
	case *ExprContext:
		return d.Accept(p)
	case *GetArrayExprContext:
		return d.Accept(p)
	case *GetStringExprContext:
		return d.Accept(p)
	case *AndOrContext:
		return d.Accept(p)
	case *ConditionKeyContext:
		return d.Accept(p)
	case *ParentContext:
		return d.Accept(p)
	}
	return nil
}

func (p *condition) VisitGetNumExpr(ctx *GetNumExprContext) interface{} {
	keySource := ctx.ConditionKey().GetText()
	v := p.Visit(ctx.ConditionKey())
	if v == nil {
		return nil
	}
	colVal, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
	if err != nil {
		p.Err = errors.Wrap(ErrParseData, fmt.Sprintf("%v %v", keySource, err))
		return false
	}

	rightData := strings.Trim(ctx.Number().GetText(), `"`)
	data, err := strconv.ParseFloat(rightData, 64)
	if err != nil {
		p.Err = errors.Wrap(ErrParseData, fmt.Sprintf("%v %v", keySource, err))
		return false
	}
	code := ctx.GetOp().GetText()

	switch code {
	case ">":
		if colVal > data {
			return true
		}
	case "<":
		if colVal < data {
			return true
		}
	case "==":
		if colVal == data {
			return true
		}
	case "!=":
		if colVal != data {
			return true
		}
	}
	return false
}

func (p *condition) VisitParent(ctx *ParentContext) interface{} {
	return p.Visit(ctx.Expr())
}

func (p *condition) VisitGetStringExpr(ctx *GetStringExprContext) interface{} {
	v := p.Visit(ctx.ConditionKey())
	if v == nil {
		return nil
	}
	colVal := fmt.Sprintf("%v", v)

	data := strings.Trim(ctx.STRING().GetText(), `"`)

	code := ctx.GetOp().GetText()

	switch code {
	case "=~":
		if strings.Contains(colVal, data) {
			return true
		}
	case "!~":
		if !strings.Contains(colVal, data) {
			return true
		}
	case "==":
		if colVal == data {
			return true
		}
	case "!=":
		if colVal != data {
			return true
		}
	}
	return false
}

func (p *condition) VisitGetArrayExpr(ctx *GetArrayExprContext) interface{} {
	keySource := ctx.ConditionKey().GetText()
	colVal := p.Visit(ctx.ConditionKey())
	if colVal == nil {
		return false
	}

	var data []interface{}
	err := json.Unmarshal([]byte(ctx.Array().GetText()), &data)
	if err != nil {
		p.Err = errors.Wrap(ErrParseData, fmt.Sprintf("%v %v", keySource, err))
		return false
	}

	code := strings.ToLower(
		strings.ReplaceAll(ctx.GetOp().GetText(), " ", ""),
	)

	switch code {
	case "in":
		for _, v := range data {
			if fmt.Sprintf("%v", colVal) == fmt.Sprintf("%v", v) {
				return true
			}
		}
	case "notin":
		res := true
		for _, v := range data {
			if fmt.Sprintf("%v", colVal) == fmt.Sprintf("%v", v) {
				res = false
				break
			}
		}
		return res
	}

	return false
}

func (p *condition) VisitAndOr(ctx *AndOrContext) interface{} {
	expr0 := p.Visit(ctx.Expr(0))
	expr1 := p.Visit(ctx.Expr(1))
	if expr0 == nil {
		p.Err = fmt.Errorf("err,%s", ctx.Expr(0).GetText())
		return nil
	}

	if expr1 == nil {
		p.Err = fmt.Errorf("err,%s", ctx.Expr(1).GetText())
		return nil
	}

	op := strings.ToLower(ctx.GetOp().GetText())
	switch op {
	case "and":
		return expr0.(bool) && expr1.(bool)
	case "or":
		return expr0.(bool) || expr1.(bool)
	}
	return false
}

func (p *condition) VisitConditionKey(ctx *ConditionKeyContext) interface{} {
	txt := ctx.COL().GetText()
	v, ok := p.data[txt]
	if !ok {
		column := ctx.GetStart().GetColumn()
		p.Err = errors.Wrap(ErrNotFoundCol, fmt.Sprintf("column:%d, {%s} isn't exist", column, txt))
		return nil
	}
	return v
}

func (p *condition) VisitArray(ctx *ArrayContext) interface{} {
	return nil
}

func (p *condition) VisitNumber(ctx *NumberContext) interface{} {
	return nil
}
