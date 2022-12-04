// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package condition // Condition
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ConditionParser struct {
	*antlr.BaseParser
}

var conditionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func conditionParserInit() {
	staticData := &conditionParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'['", "','", "']'", "'=~'", "'!~'",
		"'=='", "'!='", "'>'", "'<'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "LIKE", "NOTLIKE", "EQ", "NOTEQ", "GT",
		"LT", "IN", "NOTIN", "NOT", "AND", "OR", "COL", "ID", "COMMENT", "NL",
		"FLOAT", "DEC", "STRING", "WS",
	}
	staticData.ruleNames = []string{
		"expr", "conditionKey", "array", "number",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 65, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 3, 0, 26, 8, 0, 1, 0, 1, 0, 1, 0, 5, 0, 31, 8, 0,
		10, 0, 12, 0, 34, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		5, 2, 44, 8, 2, 10, 2, 12, 2, 47, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5,
		2, 54, 8, 2, 10, 2, 12, 2, 57, 9, 2, 1, 2, 1, 2, 3, 2, 61, 8, 2, 1, 3,
		1, 3, 1, 3, 0, 1, 0, 4, 0, 2, 4, 6, 0, 5, 1, 0, 14, 15, 1, 0, 8, 11, 1,
		0, 10, 13, 1, 0, 17, 18, 1, 0, 23, 24, 67, 0, 25, 1, 0, 0, 0, 2, 35, 1,
		0, 0, 0, 4, 60, 1, 0, 0, 0, 6, 62, 1, 0, 0, 0, 8, 9, 6, 0, -1, 0, 9, 10,
		3, 2, 1, 0, 10, 11, 7, 0, 0, 0, 11, 12, 3, 4, 2, 0, 12, 26, 1, 0, 0, 0,
		13, 14, 3, 2, 1, 0, 14, 15, 7, 1, 0, 0, 15, 16, 5, 25, 0, 0, 16, 26, 1,
		0, 0, 0, 17, 18, 3, 2, 1, 0, 18, 19, 7, 2, 0, 0, 19, 20, 3, 6, 3, 0, 20,
		26, 1, 0, 0, 0, 21, 22, 5, 1, 0, 0, 22, 23, 3, 0, 0, 0, 23, 24, 5, 2, 0,
		0, 24, 26, 1, 0, 0, 0, 25, 8, 1, 0, 0, 0, 25, 13, 1, 0, 0, 0, 25, 17, 1,
		0, 0, 0, 25, 21, 1, 0, 0, 0, 26, 32, 1, 0, 0, 0, 27, 28, 10, 5, 0, 0, 28,
		29, 7, 3, 0, 0, 29, 31, 3, 0, 0, 6, 30, 27, 1, 0, 0, 0, 31, 34, 1, 0, 0,
		0, 32, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 1, 1, 0, 0, 0, 34, 32, 1,
		0, 0, 0, 35, 36, 5, 3, 0, 0, 36, 37, 5, 19, 0, 0, 37, 38, 5, 4, 0, 0, 38,
		3, 1, 0, 0, 0, 39, 40, 5, 5, 0, 0, 40, 45, 5, 25, 0, 0, 41, 42, 5, 6, 0,
		0, 42, 44, 5, 25, 0, 0, 43, 41, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43,
		1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 48, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0,
		48, 61, 5, 7, 0, 0, 49, 50, 5, 5, 0, 0, 50, 55, 3, 6, 3, 0, 51, 52, 5,
		6, 0, 0, 52, 54, 3, 6, 3, 0, 53, 51, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55,
		53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 58, 1, 0, 0, 0, 57, 55, 1, 0, 0,
		0, 58, 59, 5, 7, 0, 0, 59, 61, 1, 0, 0, 0, 60, 39, 1, 0, 0, 0, 60, 49,
		1, 0, 0, 0, 61, 5, 1, 0, 0, 0, 62, 63, 7, 4, 0, 0, 63, 7, 1, 0, 0, 0, 5,
		25, 32, 45, 55, 60,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ConditionParserInit initializes any static state used to implement ConditionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConditionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConditionParserInit() {
	staticData := &conditionParserStaticData
	staticData.once.Do(conditionParserInit)
}

// NewConditionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConditionParser(input antlr.TokenStream) *ConditionParser {
	ConditionParserInit()
	this := new(ConditionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &conditionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// ConditionParser tokens.
const (
	ConditionParserEOF     = antlr.TokenEOF
	ConditionParserT__0    = 1
	ConditionParserT__1    = 2
	ConditionParserT__2    = 3
	ConditionParserT__3    = 4
	ConditionParserT__4    = 5
	ConditionParserT__5    = 6
	ConditionParserT__6    = 7
	ConditionParserLIKE    = 8
	ConditionParserNOTLIKE = 9
	ConditionParserEQ      = 10
	ConditionParserNOTEQ   = 11
	ConditionParserGT      = 12
	ConditionParserLT      = 13
	ConditionParserIN      = 14
	ConditionParserNOTIN   = 15
	ConditionParserNOT     = 16
	ConditionParserAND     = 17
	ConditionParserOR      = 18
	ConditionParserCOL     = 19
	ConditionParserID      = 20
	ConditionParserCOMMENT = 21
	ConditionParserNL      = 22
	ConditionParserFLOAT   = 23
	ConditionParserDEC     = 24
	ConditionParserSTRING  = 25
	ConditionParserWS      = 26
)

// ConditionParser rules.
const (
	ConditionParserRULE_expr         = 0
	ConditionParserRULE_conditionKey = 1
	ConditionParserRULE_array        = 2
	ConditionParserRULE_number       = 3
)

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConditionParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConditionParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GetNumExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewGetNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetNumExprContext {
	var p = new(GetNumExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GetNumExprContext) GetOp() antlr.Token { return s.op }

func (s *GetNumExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *GetNumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetNumExprContext) ConditionKey() IConditionKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionKeyContext)
}

func (s *GetNumExprContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *GetNumExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(ConditionParserEQ, 0)
}

func (s *GetNumExprContext) GT() antlr.TerminalNode {
	return s.GetToken(ConditionParserGT, 0)
}

func (s *GetNumExprContext) LT() antlr.TerminalNode {
	return s.GetToken(ConditionParserLT, 0)
}

func (s *GetNumExprContext) NOTEQ() antlr.TerminalNode {
	return s.GetToken(ConditionParserNOTEQ, 0)
}

func (s *GetNumExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConditionVisitor:
		return t.VisitGetNumExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentContext struct {
	*ExprContext
}

func NewParentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentContext {
	var p = new(ParentContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConditionVisitor:
		return t.VisitParent(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetStringExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewGetStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetStringExprContext {
	var p = new(GetStringExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GetStringExprContext) GetOp() antlr.Token { return s.op }

func (s *GetStringExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *GetStringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetStringExprContext) ConditionKey() IConditionKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionKeyContext)
}

func (s *GetStringExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(ConditionParserSTRING, 0)
}

func (s *GetStringExprContext) LIKE() antlr.TerminalNode {
	return s.GetToken(ConditionParserLIKE, 0)
}

func (s *GetStringExprContext) NOTLIKE() antlr.TerminalNode {
	return s.GetToken(ConditionParserNOTLIKE, 0)
}

func (s *GetStringExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(ConditionParserEQ, 0)
}

func (s *GetStringExprContext) NOTEQ() antlr.TerminalNode {
	return s.GetToken(ConditionParserNOTEQ, 0)
}

func (s *GetStringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConditionVisitor:
		return t.VisitGetStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetArrayExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewGetArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetArrayExprContext {
	var p = new(GetArrayExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GetArrayExprContext) GetOp() antlr.Token { return s.op }

func (s *GetArrayExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *GetArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetArrayExprContext) ConditionKey() IConditionKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionKeyContext)
}

func (s *GetArrayExprContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *GetArrayExprContext) IN() antlr.TerminalNode {
	return s.GetToken(ConditionParserIN, 0)
}

func (s *GetArrayExprContext) NOTIN() antlr.TerminalNode {
	return s.GetToken(ConditionParserNOTIN, 0)
}

func (s *GetArrayExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConditionVisitor:
		return t.VisitGetArrayExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndOrContext struct {
	*ExprContext
	op antlr.Token
}

func NewAndOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndOrContext {
	var p = new(AndOrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndOrContext) GetOp() antlr.Token { return s.op }

func (s *AndOrContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndOrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndOrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndOrContext) AND() antlr.TerminalNode {
	return s.GetToken(ConditionParserAND, 0)
}

func (s *AndOrContext) OR() antlr.TerminalNode {
	return s.GetToken(ConditionParserOR, 0)
}

func (s *AndOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConditionVisitor:
		return t.VisitAndOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConditionParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ConditionParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, ConditionParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGetArrayExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(9)
			p.ConditionKey()
		}
		{
			p.SetState(10)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*GetArrayExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ConditionParserIN || _la == ConditionParserNOTIN) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*GetArrayExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(11)
			p.Array()
		}

	case 2:
		localctx = NewGetStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)
			p.ConditionKey()
		}
		{
			p.SetState(14)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*GetStringExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3840) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*GetStringExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(15)
			p.Match(ConditionParserSTRING)
		}

	case 3:
		localctx = NewGetNumExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.ConditionKey()
		}
		{
			p.SetState(18)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*GetNumExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*GetNumExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(19)
			p.Number()
		}

	case 4:
		localctx = NewParentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)
			p.Match(ConditionParserT__0)
		}
		{
			p.SetState(22)
			p.expr(0)
		}
		{
			p.SetState(23)
			p.Match(ConditionParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndOrContext(p, NewExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, ConditionParserRULE_expr)
			p.SetState(27)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(28)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AndOrContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ConditionParserAND || _la == ConditionParserOR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AndOrContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(29)
				p.expr(6)
			}

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IConditionKeyContext is an interface to support dynamic dispatch.
type IConditionKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionKeyContext differentiates from other interfaces.
	IsConditionKeyContext()
}

type ConditionKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionKeyContext() *ConditionKeyContext {
	var p = new(ConditionKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConditionParserRULE_conditionKey
	return p
}

func (*ConditionKeyContext) IsConditionKeyContext() {}

func NewConditionKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionKeyContext {
	var p = new(ConditionKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConditionParserRULE_conditionKey

	return p
}

func (s *ConditionKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionKeyContext) COL() antlr.TerminalNode {
	return s.GetToken(ConditionParserCOL, 0)
}

func (s *ConditionKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConditionVisitor:
		return t.VisitConditionKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConditionParser) ConditionKey() (localctx IConditionKeyContext) {
	this := p
	_ = this

	localctx = NewConditionKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConditionParserRULE_conditionKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(ConditionParserT__2)
	}
	{
		p.SetState(36)
		p.Match(ConditionParserCOL)
	}
	{
		p.SetState(37)
		p.Match(ConditionParserT__3)
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConditionParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConditionParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ConditionParserSTRING)
}

func (s *ArrayContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ConditionParserSTRING, i)
}

func (s *ArrayContext) AllNumber() []INumberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberContext); ok {
			len++
		}
	}

	tst := make([]INumberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberContext); ok {
			tst[i] = t.(INumberContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Number(i int) INumberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConditionVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConditionParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConditionParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Match(ConditionParserT__4)
		}
		{
			p.SetState(40)
			p.Match(ConditionParserSTRING)
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConditionParserT__5 {
			{
				p.SetState(41)
				p.Match(ConditionParserT__5)
			}
			{
				p.SetState(42)
				p.Match(ConditionParserSTRING)
			}

			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(48)
			p.Match(ConditionParserT__6)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(ConditionParserT__4)
		}
		{
			p.SetState(50)
			p.Number()
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConditionParserT__5 {
			{
				p.SetState(51)
				p.Match(ConditionParserT__5)
			}
			{
				p.SetState(52)
				p.Number()
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(58)
			p.Match(ConditionParserT__6)
		}

	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConditionParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConditionParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ConditionParserFLOAT, 0)
}

func (s *NumberContext) DEC() antlr.TerminalNode {
	return s.GetToken(ConditionParserDEC, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConditionVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConditionParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConditionParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ConditionParserFLOAT || _la == ConditionParserDEC) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *ConditionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ConditionParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
