// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package g4 // Config
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

type ConfigParser struct {
	*antlr.BaseParser
}

var configParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func configParserInit() {
	staticData := &configParserStaticData
	staticData.literalNames = []string{
		"", "' '", "':'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "ID", "COMMENT", "NL", "FLOAT", "DEC", "STRING", "OTHER",
	}
	staticData.ruleNames = []string{
		"stat", "sections", "section", "fields", "field", "fieldName", "fieldValue",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 98, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 5, 0, 16, 8, 0, 10, 0, 12, 0, 19, 9, 0, 1,
		0, 1, 0, 5, 0, 23, 8, 0, 10, 0, 12, 0, 26, 9, 0, 1, 0, 1, 0, 1, 0, 5, 0,
		31, 8, 0, 10, 0, 12, 0, 34, 9, 0, 1, 0, 3, 0, 37, 8, 0, 1, 1, 5, 1, 40,
		8, 1, 10, 1, 12, 1, 43, 9, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1,
		50, 9, 1, 4, 1, 52, 8, 1, 11, 1, 12, 1, 53, 1, 2, 4, 2, 57, 8, 2, 11, 2,
		12, 2, 58, 1, 2, 1, 2, 1, 2, 3, 2, 64, 8, 2, 1, 3, 1, 3, 4, 3, 68, 8, 3,
		11, 3, 12, 3, 69, 4, 3, 72, 8, 3, 11, 3, 12, 3, 73, 1, 4, 1, 4, 5, 4, 78,
		8, 4, 10, 4, 12, 4, 81, 9, 4, 1, 4, 1, 4, 5, 4, 85, 8, 4, 10, 4, 12, 4,
		88, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 94, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		58, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 1, 1, 0, 6, 8, 104, 0, 36, 1, 0, 0,
		0, 2, 51, 1, 0, 0, 0, 4, 56, 1, 0, 0, 0, 6, 71, 1, 0, 0, 0, 8, 75, 1, 0,
		0, 0, 10, 91, 1, 0, 0, 0, 12, 95, 1, 0, 0, 0, 14, 16, 5, 5, 0, 0, 15, 14,
		1, 0, 0, 0, 16, 19, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0,
		18, 20, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 20, 24, 3, 6, 3, 0, 21, 23, 5,
		5, 0, 0, 22, 21, 1, 0, 0, 0, 23, 26, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24,
		25, 1, 0, 0, 0, 25, 27, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 27, 28, 3, 2, 1,
		0, 28, 37, 1, 0, 0, 0, 29, 31, 5, 5, 0, 0, 30, 29, 1, 0, 0, 0, 31, 34,
		1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 35, 1, 0, 0, 0,
		34, 32, 1, 0, 0, 0, 35, 37, 3, 2, 1, 0, 36, 17, 1, 0, 0, 0, 36, 32, 1,
		0, 0, 0, 37, 1, 1, 0, 0, 0, 38, 40, 5, 5, 0, 0, 39, 38, 1, 0, 0, 0, 40,
		43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 44, 1, 0, 0,
		0, 43, 41, 1, 0, 0, 0, 44, 48, 3, 4, 2, 0, 45, 47, 5, 5, 0, 0, 46, 45,
		1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0,
		49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 41, 1, 0, 0, 0, 52, 53, 1,
		0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 3, 1, 0, 0, 0, 55,
		57, 9, 0, 0, 0, 56, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0,
		0, 58, 56, 1, 0, 0, 0, 59, 63, 1, 0, 0, 0, 60, 61, 5, 5, 0, 0, 61, 64,
		5, 5, 0, 0, 62, 64, 5, 0, 0, 1, 63, 60, 1, 0, 0, 0, 63, 62, 1, 0, 0, 0,
		64, 5, 1, 0, 0, 0, 65, 67, 3, 8, 4, 0, 66, 68, 5, 5, 0, 0, 67, 66, 1, 0,
		0, 0, 68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72,
		1, 0, 0, 0, 71, 65, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0,
		73, 74, 1, 0, 0, 0, 74, 7, 1, 0, 0, 0, 75, 79, 3, 10, 5, 0, 76, 78, 5,
		1, 0, 0, 77, 76, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79,
		80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 86, 5, 2, 0,
		0, 83, 85, 5, 1, 0, 0, 84, 83, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84,
		1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 89, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0,
		89, 90, 3, 12, 6, 0, 90, 9, 1, 0, 0, 0, 91, 93, 5, 3, 0, 0, 92, 94, 5,
		7, 0, 0, 93, 92, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 11, 1, 0, 0, 0, 95,
		96, 7, 0, 0, 0, 96, 13, 1, 0, 0, 0, 14, 17, 24, 32, 36, 41, 48, 53, 58,
		63, 69, 73, 79, 86, 93,
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

// ConfigParserInit initializes any static state used to implement ConfigParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConfigParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConfigParserInit() {
	staticData := &configParserStaticData
	staticData.once.Do(configParserInit)
}

// NewConfigParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConfigParser(input antlr.TokenStream) *ConfigParser {
	ConfigParserInit()
	this := new(ConfigParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &configParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// ConfigParser tokens.
const (
	ConfigParserEOF     = antlr.TokenEOF
	ConfigParserT__0    = 1
	ConfigParserT__1    = 2
	ConfigParserID      = 3
	ConfigParserCOMMENT = 4
	ConfigParserNL      = 5
	ConfigParserFLOAT   = 6
	ConfigParserDEC     = 7
	ConfigParserSTRING  = 8
	ConfigParserOTHER   = 9
)

// ConfigParser rules.
const (
	ConfigParserRULE_stat       = 0
	ConfigParserRULE_sections   = 1
	ConfigParserRULE_section    = 2
	ConfigParserRULE_fields     = 3
	ConfigParserRULE_field      = 4
	ConfigParserRULE_fieldName  = 5
	ConfigParserRULE_fieldValue = 6
)

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyFrom(ctx *StatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldStatContext struct {
	*StatContext
}

func NewFieldStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldStatContext {
	var p = new(FieldStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *FieldStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldStatContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *FieldStatContext) Sections() ISectionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectionsContext)
}

func (s *FieldStatContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserNL)
}

func (s *FieldStatContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserNL, i)
}

func (s *FieldStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConfigVisitor:
		return t.VisitFieldStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotFieldStatContext struct {
	*StatContext
}

func NewNotFieldStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotFieldStatContext {
	var p = new(NotFieldStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *NotFieldStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotFieldStatContext) Sections() ISectionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectionsContext)
}

func (s *NotFieldStatContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserNL)
}

func (s *NotFieldStatContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserNL, i)
}

func (s *NotFieldStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConfigVisitor:
		return t.VisitNotFieldStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConfigParser) Stat() (localctx IStatContext) {
	this := p
	_ = this

	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConfigParserRULE_stat)
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

	var _alt int

	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFieldStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConfigParserNL {
			{
				p.SetState(14)
				p.Match(ConfigParserNL)
			}

			p.SetState(19)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(20)
			p.Fields()
		}
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(21)
					p.Match(ConfigParserNL)
				}

			}
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}
		{
			p.SetState(27)
			p.Sections()
		}

	case 2:
		localctx = NewNotFieldStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(29)
					p.Match(ConfigParserNL)
				}

			}
			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
		}
		{
			p.SetState(35)
			p.Sections()
		}

	}

	return localctx
}

// ISectionsContext is an interface to support dynamic dispatch.
type ISectionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionsContext differentiates from other interfaces.
	IsSectionsContext()
}

type SectionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionsContext() *SectionsContext {
	var p = new(SectionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_sections
	return p
}

func (*SectionsContext) IsSectionsContext() {}

func NewSectionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionsContext {
	var p = new(SectionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_sections

	return p
}

func (s *SectionsContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionsContext) AllSection() []ISectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISectionContext); ok {
			len++
		}
	}

	tst := make([]ISectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISectionContext); ok {
			tst[i] = t.(ISectionContext)
			i++
		}
	}

	return tst
}

func (s *SectionsContext) Section(i int) ISectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionContext); ok {
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

	return t.(ISectionContext)
}

func (s *SectionsContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserNL)
}

func (s *SectionsContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserNL, i)
}

func (s *SectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConfigVisitor:
		return t.VisitSections(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConfigParser) Sections() (localctx ISectionsContext) {
	this := p
	_ = this

	localctx = NewSectionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConfigParserRULE_sections)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1022) != 0 {
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(38)
					p.Match(ConfigParserNL)
				}

			}
			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}
		{
			p.SetState(44)
			p.Section()
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(45)
					p.Match(ConfigParserNL)
				}

			}
			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) CopyFrom(ctx *SectionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GetSectionContext struct {
	*SectionContext
}

func NewGetSectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetSectionContext {
	var p = new(GetSectionContext)

	p.SectionContext = NewEmptySectionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SectionContext))

	return p
}

func (s *GetSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetSectionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserNL)
}

func (s *GetSectionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserNL, i)
}

func (s *GetSectionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConfigParserEOF, 0)
}

func (s *GetSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConfigVisitor:
		return t.VisitGetSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConfigParser) Section() (localctx ISectionContext) {
	this := p
	_ = this

	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConfigParserRULE_section)

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

	var _alt int

	localctx = NewGetSectionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			p.SetState(55)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConfigParserNL:
		{
			p.SetState(60)
			p.Match(ConfigParserNL)
		}
		{
			p.SetState(61)
			p.Match(ConfigParserNL)
		}

	case ConfigParserEOF:
		{
			p.SetState(62)
			p.Match(ConfigParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *FieldsContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserNL)
}

func (s *FieldsContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserNL, i)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConfigVisitor:
		return t.VisitFields(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConfigParser) Fields() (localctx IFieldsContext) {
	this := p
	_ = this

	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConfigParserRULE_fields)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(65)
				p.Field()
			}
			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(66)
						p.Match(ConfigParserNL)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(69)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) FieldValue() IFieldValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldValueContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConfigVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConfigParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConfigParserRULE_field)
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
		p.SetState(75)
		p.FieldName()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConfigParserT__0 {
		{
			p.SetState(76)
			p.Match(ConfigParserT__0)
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(ConfigParserT__1)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConfigParserT__0 {
		{
			p.SetState(83)
			p.Match(ConfigParserT__0)
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
		p.FieldValue()
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) ID() antlr.TerminalNode {
	return s.GetToken(ConfigParserID, 0)
}

func (s *FieldNameContext) DEC() antlr.TerminalNode {
	return s.GetToken(ConfigParserDEC, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConfigVisitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConfigParser) FieldName() (localctx IFieldNameContext) {
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConfigParserRULE_fieldName)
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
		p.SetState(91)
		p.Match(ConfigParserID)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ConfigParserDEC {
		{
			p.SetState(92)
			p.Match(ConfigParserDEC)
		}

	}

	return localctx
}

// IFieldValueContext is an interface to support dynamic dispatch.
type IFieldValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldValueContext differentiates from other interfaces.
	IsFieldValueContext()
}

type FieldValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldValueContext() *FieldValueContext {
	var p = new(FieldValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_fieldValue
	return p
}

func (*FieldValueContext) IsFieldValueContext() {}

func NewFieldValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldValueContext {
	var p = new(FieldValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_fieldValue

	return p
}

func (s *FieldValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ConfigParserSTRING, 0)
}

func (s *FieldValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ConfigParserFLOAT, 0)
}

func (s *FieldValueContext) DEC() antlr.TerminalNode {
	return s.GetToken(ConfigParserDEC, 0)
}

func (s *FieldValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConfigVisitor:
		return t.VisitFieldValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConfigParser) FieldValue() (localctx IFieldValueContext) {
	this := p
	_ = this

	localctx = NewFieldValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConfigParserRULE_fieldValue)
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
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&448) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
