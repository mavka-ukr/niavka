// Code generated from NiavkaParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // NiavkaParser
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

type NiavkaParser struct {
	*antlr.BaseParser
}

var niavkaparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func niavkaparserParserInit() {
	staticData := &niavkaparserParserStaticData
	staticData.literalNames = []string{
		"", "'\\u043A\\u0456\\u043D\\u0435\\u0446\\u044C'", "'\\u0434\\u0456\\u044F'",
		"", "", "'('", "')'", "'['", "']'", "','", "'='", "'+'", "'-'", "'*'",
		"'/'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "END", "DIIA", "SKIP_SPACES", "NL", "OPEN_PAREN", "CLOSE_PAREN",
		"OPEN_ARRAY", "CLOSE_ARRAY", "COMMA", "ASSIGN", "PLUS", "MINUS", "MUL",
		"DIV", "DOT", "ID", "NUMBER", "INTEGER", "FLOAT", "STRING",
	}
	staticData.ruleNames = []string{
		"file", "program", "program_element", "diia", "value", "assign", "assign_value",
		"identifier", "type_value", "args", "arg", "params", "param", "body",
		"body_element", "arithmetic_op_mul", "arithmetic_op_add", "nl", "nls",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 176, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 46, 8, 1, 10, 1, 12, 1, 49, 9, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 55, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 62, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 68, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 74, 8,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3,
		4, 87, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 103, 8, 4, 1, 4, 5, 4, 106, 8, 4, 10, 4,
		12, 4, 109, 9, 4, 1, 5, 1, 5, 3, 5, 113, 8, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 127, 8, 9, 10, 9,
		12, 9, 130, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 5, 11, 142, 8, 11, 10, 11, 12, 11, 145, 9, 11, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 154, 8, 13, 10, 13, 12, 13,
		157, 9, 13, 1, 14, 1, 14, 1, 14, 3, 14, 162, 8, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 5, 18, 171, 8, 18, 10, 18, 12, 18, 174, 9,
		18, 1, 18, 0, 1, 8, 19, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 0, 2, 1, 0, 13, 14, 1, 0, 11, 12, 178, 0, 38, 1,
		0, 0, 0, 2, 41, 1, 0, 0, 0, 4, 54, 1, 0, 0, 0, 6, 56, 1, 0, 0, 0, 8, 86,
		1, 0, 0, 0, 10, 110, 1, 0, 0, 0, 12, 117, 1, 0, 0, 0, 14, 119, 1, 0, 0,
		0, 16, 121, 1, 0, 0, 0, 18, 123, 1, 0, 0, 0, 20, 131, 1, 0, 0, 0, 22, 135,
		1, 0, 0, 0, 24, 146, 1, 0, 0, 0, 26, 149, 1, 0, 0, 0, 28, 161, 1, 0, 0,
		0, 30, 163, 1, 0, 0, 0, 32, 165, 1, 0, 0, 0, 34, 167, 1, 0, 0, 0, 36, 172,
		1, 0, 0, 0, 38, 39, 3, 2, 1, 0, 39, 40, 5, 0, 0, 1, 40, 1, 1, 0, 0, 0,
		41, 47, 3, 4, 2, 0, 42, 43, 3, 34, 17, 0, 43, 44, 3, 4, 2, 0, 44, 46, 1,
		0, 0, 0, 45, 42, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47,
		48, 1, 0, 0, 0, 48, 3, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 55, 3, 6, 3,
		0, 51, 55, 3, 8, 4, 0, 52, 55, 3, 10, 5, 0, 53, 55, 3, 36, 18, 0, 54, 50,
		1, 0, 0, 0, 54, 51, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0,
		55, 5, 1, 0, 0, 0, 56, 57, 5, 2, 0, 0, 57, 58, 3, 14, 7, 0, 58, 59, 5,
		5, 0, 0, 59, 61, 3, 36, 18, 0, 60, 62, 3, 22, 11, 0, 61, 60, 1, 0, 0, 0,
		61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 3, 36, 18, 0, 64, 65, 1,
		0, 0, 0, 65, 67, 5, 6, 0, 0, 66, 68, 3, 16, 8, 0, 67, 66, 1, 0, 0, 0, 67,
		68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 73, 3, 34, 17, 0, 70, 71, 3, 26,
		13, 0, 71, 72, 3, 34, 17, 0, 72, 74, 1, 0, 0, 0, 73, 70, 1, 0, 0, 0, 73,
		74, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 3, 36, 18, 0, 76, 77, 5, 1,
		0, 0, 77, 7, 1, 0, 0, 0, 78, 79, 6, 4, -1, 0, 79, 87, 5, 17, 0, 0, 80,
		87, 5, 20, 0, 0, 81, 87, 3, 14, 7, 0, 82, 83, 5, 5, 0, 0, 83, 84, 3, 8,
		4, 0, 84, 85, 5, 6, 0, 0, 85, 87, 1, 0, 0, 0, 86, 78, 1, 0, 0, 0, 86, 80,
		1, 0, 0, 0, 86, 81, 1, 0, 0, 0, 86, 82, 1, 0, 0, 0, 87, 107, 1, 0, 0, 0,
		88, 89, 10, 2, 0, 0, 89, 90, 3, 30, 15, 0, 90, 91, 3, 8, 4, 3, 91, 106,
		1, 0, 0, 0, 92, 93, 10, 1, 0, 0, 93, 94, 3, 32, 16, 0, 94, 95, 3, 8, 4,
		2, 95, 106, 1, 0, 0, 0, 96, 97, 10, 5, 0, 0, 97, 98, 5, 15, 0, 0, 98, 106,
		3, 14, 7, 0, 99, 100, 10, 4, 0, 0, 100, 102, 5, 5, 0, 0, 101, 103, 3, 18,
		9, 0, 102, 101, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0,
		104, 106, 5, 6, 0, 0, 105, 88, 1, 0, 0, 0, 105, 92, 1, 0, 0, 0, 105, 96,
		1, 0, 0, 0, 105, 99, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0,
		0, 0, 107, 108, 1, 0, 0, 0, 108, 9, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110,
		112, 3, 14, 7, 0, 111, 113, 3, 16, 8, 0, 112, 111, 1, 0, 0, 0, 112, 113,
		1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 5, 10, 0, 0, 115, 116, 3, 12,
		6, 0, 116, 11, 1, 0, 0, 0, 117, 118, 3, 8, 4, 0, 118, 13, 1, 0, 0, 0, 119,
		120, 5, 16, 0, 0, 120, 15, 1, 0, 0, 0, 121, 122, 3, 14, 7, 0, 122, 17,
		1, 0, 0, 0, 123, 128, 3, 20, 10, 0, 124, 125, 5, 9, 0, 0, 125, 127, 3,
		20, 10, 0, 126, 124, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0,
		0, 0, 128, 129, 1, 0, 0, 0, 129, 19, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0,
		131, 132, 3, 36, 18, 0, 132, 133, 3, 8, 4, 0, 133, 134, 3, 36, 18, 0, 134,
		21, 1, 0, 0, 0, 135, 143, 3, 24, 12, 0, 136, 137, 3, 36, 18, 0, 137, 138,
		5, 9, 0, 0, 138, 139, 3, 36, 18, 0, 139, 140, 3, 24, 12, 0, 140, 142, 1,
		0, 0, 0, 141, 136, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0,
		0, 143, 144, 1, 0, 0, 0, 144, 23, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146,
		147, 3, 14, 7, 0, 147, 148, 3, 16, 8, 0, 148, 25, 1, 0, 0, 0, 149, 155,
		3, 28, 14, 0, 150, 151, 3, 34, 17, 0, 151, 152, 3, 28, 14, 0, 152, 154,
		1, 0, 0, 0, 153, 150, 1, 0, 0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0,
		0, 0, 155, 156, 1, 0, 0, 0, 156, 27, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0,
		158, 162, 3, 8, 4, 0, 159, 162, 3, 10, 5, 0, 160, 162, 3, 36, 18, 0, 161,
		158, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 29, 1,
		0, 0, 0, 163, 164, 7, 0, 0, 0, 164, 31, 1, 0, 0, 0, 165, 166, 7, 1, 0,
		0, 166, 33, 1, 0, 0, 0, 167, 168, 5, 4, 0, 0, 168, 35, 1, 0, 0, 0, 169,
		171, 3, 34, 17, 0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170,
		1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 37, 1, 0, 0, 0, 174, 172, 1, 0,
		0, 0, 15, 47, 54, 61, 67, 73, 86, 102, 105, 107, 112, 128, 143, 155, 161,
		172,
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

// NiavkaParserInit initializes any static state used to implement NiavkaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNiavkaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NiavkaParserInit() {
	staticData := &niavkaparserParserStaticData
	staticData.once.Do(niavkaparserParserInit)
}

// NewNiavkaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNiavkaParser(input antlr.TokenStream) *NiavkaParser {
	NiavkaParserInit()
	this := new(NiavkaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &niavkaparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "NiavkaParser.g4"

	return this
}

// NiavkaParser tokens.
const (
	NiavkaParserEOF         = antlr.TokenEOF
	NiavkaParserEND         = 1
	NiavkaParserDIIA        = 2
	NiavkaParserSKIP_SPACES = 3
	NiavkaParserNL          = 4
	NiavkaParserOPEN_PAREN  = 5
	NiavkaParserCLOSE_PAREN = 6
	NiavkaParserOPEN_ARRAY  = 7
	NiavkaParserCLOSE_ARRAY = 8
	NiavkaParserCOMMA       = 9
	NiavkaParserASSIGN      = 10
	NiavkaParserPLUS        = 11
	NiavkaParserMINUS       = 12
	NiavkaParserMUL         = 13
	NiavkaParserDIV         = 14
	NiavkaParserDOT         = 15
	NiavkaParserID          = 16
	NiavkaParserNUMBER      = 17
	NiavkaParserINTEGER     = 18
	NiavkaParserFLOAT       = 19
	NiavkaParserSTRING      = 20
)

// NiavkaParser rules.
const (
	NiavkaParserRULE_file              = 0
	NiavkaParserRULE_program           = 1
	NiavkaParserRULE_program_element   = 2
	NiavkaParserRULE_diia              = 3
	NiavkaParserRULE_value             = 4
	NiavkaParserRULE_assign            = 5
	NiavkaParserRULE_assign_value      = 6
	NiavkaParserRULE_identifier        = 7
	NiavkaParserRULE_type_value        = 8
	NiavkaParserRULE_args              = 9
	NiavkaParserRULE_arg               = 10
	NiavkaParserRULE_params            = 11
	NiavkaParserRULE_param             = 12
	NiavkaParserRULE_body              = 13
	NiavkaParserRULE_body_element      = 14
	NiavkaParserRULE_arithmetic_op_mul = 15
	NiavkaParserRULE_arithmetic_op_add = 16
	NiavkaParserRULE_nl                = 17
	NiavkaParserRULE_nls               = 18
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF_program returns the f_program rule contexts.
	GetF_program() IProgramContext

	// SetF_program sets the f_program rule contexts.
	SetF_program(IProgramContext)

	// Getter signatures
	EOF() antlr.TerminalNode
	Program() IProgramContext

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	f_program IProgramContext
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) GetF_program() IProgramContext { return s.f_program }

func (s *FileContext) SetF_program(v IProgramContext) { s.f_program = v }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(NiavkaParserEOF, 0)
}

func (s *FileContext) Program() IProgramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) File() (localctx IFileContext) {
	this := p
	_ = this

	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NiavkaParserRULE_file)

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
		p.SetState(38)

		var _x = p.Program()

		localctx.(*FileContext).f_program = _x
	}
	{
		p.SetState(39)
		p.Match(NiavkaParserEOF)
	}

	return localctx
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllProgram_element() []IProgram_elementContext
	Program_element(i int) IProgram_elementContext
	AllNl() []INlContext
	Nl(i int) INlContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllProgram_element() []IProgram_elementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgram_elementContext); ok {
			len++
		}
	}

	tst := make([]IProgram_elementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgram_elementContext); ok {
			tst[i] = t.(IProgram_elementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Program_element(i int) IProgram_elementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgram_elementContext); ok {
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

	return t.(IProgram_elementContext)
}

func (s *ProgramContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
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

	return t.(INlContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NiavkaParserRULE_program)
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
		p.SetState(41)
		p.Program_element()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NiavkaParserNL {
		{
			p.SetState(42)
			p.Nl()
		}
		{
			p.SetState(43)
			p.Program_element()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IProgram_elementContext is an interface to support dynamic dispatch.
type IProgram_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Diia() IDiiaContext
	Value() IValueContext
	Assign() IAssignContext
	Nls() INlsContext

	// IsProgram_elementContext differentiates from other interfaces.
	IsProgram_elementContext()
}

type Program_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgram_elementContext() *Program_elementContext {
	var p = new(Program_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_program_element
	return p
}

func (*Program_elementContext) IsProgram_elementContext() {}

func NewProgram_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Program_elementContext {
	var p = new(Program_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_program_element

	return p
}

func (s *Program_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Program_elementContext) Diia() IDiiaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiiaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiiaContext)
}

func (s *Program_elementContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Program_elementContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *Program_elementContext) Nls() INlsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlsContext)
}

func (s *Program_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Program_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Program_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterProgram_element(s)
	}
}

func (s *Program_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitProgram_element(s)
	}
}

func (s *Program_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitProgram_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Program_element() (localctx IProgram_elementContext) {
	this := p
	_ = this

	localctx = NewProgram_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NiavkaParserRULE_program_element)

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

	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Diia()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.value(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.Assign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(53)
			p.Nls()
		}

	}

	return localctx
}

// IDiiaContext is an interface to support dynamic dispatch.
type IDiiaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetD_name returns the d_name rule contexts.
	GetD_name() IIdentifierContext

	// GetD_params returns the d_params rule contexts.
	GetD_params() IParamsContext

	// GetD_type returns the d_type rule contexts.
	GetD_type() IType_valueContext

	// GetD_body returns the d_body rule contexts.
	GetD_body() IBodyContext

	// SetD_name sets the d_name rule contexts.
	SetD_name(IIdentifierContext)

	// SetD_params sets the d_params rule contexts.
	SetD_params(IParamsContext)

	// SetD_type sets the d_type rule contexts.
	SetD_type(IType_valueContext)

	// SetD_body sets the d_body rule contexts.
	SetD_body(IBodyContext)

	// Getter signatures
	DIIA() antlr.TerminalNode
	OPEN_PAREN() antlr.TerminalNode
	CLOSE_PAREN() antlr.TerminalNode
	AllNl() []INlContext
	Nl(i int) INlContext
	AllNls() []INlsContext
	Nls(i int) INlsContext
	END() antlr.TerminalNode
	Identifier() IIdentifierContext
	Type_value() IType_valueContext
	Body() IBodyContext
	Params() IParamsContext

	// IsDiiaContext differentiates from other interfaces.
	IsDiiaContext()
}

type DiiaContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	d_name   IIdentifierContext
	d_params IParamsContext
	d_type   IType_valueContext
	d_body   IBodyContext
}

func NewEmptyDiiaContext() *DiiaContext {
	var p = new(DiiaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_diia
	return p
}

func (*DiiaContext) IsDiiaContext() {}

func NewDiiaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiiaContext {
	var p = new(DiiaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_diia

	return p
}

func (s *DiiaContext) GetParser() antlr.Parser { return s.parser }

func (s *DiiaContext) GetD_name() IIdentifierContext { return s.d_name }

func (s *DiiaContext) GetD_params() IParamsContext { return s.d_params }

func (s *DiiaContext) GetD_type() IType_valueContext { return s.d_type }

func (s *DiiaContext) GetD_body() IBodyContext { return s.d_body }

func (s *DiiaContext) SetD_name(v IIdentifierContext) { s.d_name = v }

func (s *DiiaContext) SetD_params(v IParamsContext) { s.d_params = v }

func (s *DiiaContext) SetD_type(v IType_valueContext) { s.d_type = v }

func (s *DiiaContext) SetD_body(v IBodyContext) { s.d_body = v }

func (s *DiiaContext) DIIA() antlr.TerminalNode {
	return s.GetToken(NiavkaParserDIIA, 0)
}

func (s *DiiaContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(NiavkaParserOPEN_PAREN, 0)
}

func (s *DiiaContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(NiavkaParserCLOSE_PAREN, 0)
}

func (s *DiiaContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *DiiaContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
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

	return t.(INlContext)
}

func (s *DiiaContext) AllNls() []INlsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlsContext); ok {
			len++
		}
	}

	tst := make([]INlsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlsContext); ok {
			tst[i] = t.(INlsContext)
			i++
		}
	}

	return tst
}

func (s *DiiaContext) Nls(i int) INlsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlsContext); ok {
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

	return t.(INlsContext)
}

func (s *DiiaContext) END() antlr.TerminalNode {
	return s.GetToken(NiavkaParserEND, 0)
}

func (s *DiiaContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DiiaContext) Type_value() IType_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_valueContext)
}

func (s *DiiaContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *DiiaContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *DiiaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiiaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiiaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterDiia(s)
	}
}

func (s *DiiaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitDiia(s)
	}
}

func (s *DiiaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitDiia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Diia() (localctx IDiiaContext) {
	this := p
	_ = this

	localctx = NewDiiaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NiavkaParserRULE_diia)
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
		p.SetState(56)
		p.Match(NiavkaParserDIIA)
	}
	{
		p.SetState(57)

		var _x = p.Identifier()

		localctx.(*DiiaContext).d_name = _x
	}
	{
		p.SetState(58)
		p.Match(NiavkaParserOPEN_PAREN)
	}

	{
		p.SetState(59)
		p.Nls()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NiavkaParserID {
		{
			p.SetState(60)

			var _x = p.Params()

			localctx.(*DiiaContext).d_params = _x
		}

	}
	{
		p.SetState(63)
		p.Nls()
	}

	{
		p.SetState(65)
		p.Match(NiavkaParserCLOSE_PAREN)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NiavkaParserID {
		{
			p.SetState(66)

			var _x = p.Type_value()

			localctx.(*DiiaContext).d_type = _x
		}

	}
	{
		p.SetState(69)
		p.Nl()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(70)

			var _x = p.Body()

			localctx.(*DiiaContext).d_body = _x
		}
		{
			p.SetState(71)
			p.Nl()
		}

	}
	{
		p.SetState(75)
		p.Nls()
	}
	{
		p.SetState(76)
		p.Match(NiavkaParserEND)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Arithmetic_mulContext struct {
	*ValueContext
	a_left      IValueContext
	a_operation IArithmetic_op_mulContext
	a_right     IValueContext
}

func NewArithmetic_mulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Arithmetic_mulContext {
	var p = new(Arithmetic_mulContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *Arithmetic_mulContext) GetA_left() IValueContext { return s.a_left }

func (s *Arithmetic_mulContext) GetA_operation() IArithmetic_op_mulContext { return s.a_operation }

func (s *Arithmetic_mulContext) GetA_right() IValueContext { return s.a_right }

func (s *Arithmetic_mulContext) SetA_left(v IValueContext) { s.a_left = v }

func (s *Arithmetic_mulContext) SetA_operation(v IArithmetic_op_mulContext) { s.a_operation = v }

func (s *Arithmetic_mulContext) SetA_right(v IValueContext) { s.a_right = v }

func (s *Arithmetic_mulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithmetic_mulContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *Arithmetic_mulContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *Arithmetic_mulContext) Arithmetic_op_mul() IArithmetic_op_mulContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_op_mulContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmetic_op_mulContext)
}

func (s *Arithmetic_mulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterArithmetic_mul(s)
	}
}

func (s *Arithmetic_mulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitArithmetic_mul(s)
	}
}

func (s *Arithmetic_mulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitArithmetic_mul(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallContext struct {
	*ValueContext
	c_value IValueContext
	c_args  IArgsContext
}

func NewCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallContext {
	var p = new(CallContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *CallContext) GetC_value() IValueContext { return s.c_value }

func (s *CallContext) GetC_args() IArgsContext { return s.c_args }

func (s *CallContext) SetC_value(v IValueContext) { s.c_value = v }

func (s *CallContext) SetC_args(v IArgsContext) { s.c_args = v }

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(NiavkaParserOPEN_PAREN, 0)
}

func (s *CallContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(NiavkaParserCLOSE_PAREN, 0)
}

func (s *CallContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CallContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitCall(s)
	}
}

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	*ValueContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NiavkaParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type ChainContext struct {
	*ValueContext
	c_left  IValueContext
	c_right IIdentifierContext
}

func NewChainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChainContext {
	var p = new(ChainContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ChainContext) GetC_left() IValueContext { return s.c_left }

func (s *ChainContext) GetC_right() IIdentifierContext { return s.c_right }

func (s *ChainContext) SetC_left(v IValueContext) { s.c_left = v }

func (s *ChainContext) SetC_right(v IIdentifierContext) { s.c_right = v }

func (s *ChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainContext) DOT() antlr.TerminalNode {
	return s.GetToken(NiavkaParserDOT, 0)
}

func (s *ChainContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ChainContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterChain(s)
	}
}

func (s *ChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitChain(s)
	}
}

func (s *ChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitChain(s)

	default:
		return t.VisitChildren(s)
	}
}

type Arithmetic_addContext struct {
	*ValueContext
	a_left      IValueContext
	a_operation IArithmetic_op_addContext
	a_right     IValueContext
}

func NewArithmetic_addContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Arithmetic_addContext {
	var p = new(Arithmetic_addContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *Arithmetic_addContext) GetA_left() IValueContext { return s.a_left }

func (s *Arithmetic_addContext) GetA_operation() IArithmetic_op_addContext { return s.a_operation }

func (s *Arithmetic_addContext) GetA_right() IValueContext { return s.a_right }

func (s *Arithmetic_addContext) SetA_left(v IValueContext) { s.a_left = v }

func (s *Arithmetic_addContext) SetA_operation(v IArithmetic_op_addContext) { s.a_operation = v }

func (s *Arithmetic_addContext) SetA_right(v IValueContext) { s.a_right = v }

func (s *Arithmetic_addContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithmetic_addContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *Arithmetic_addContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *Arithmetic_addContext) Arithmetic_op_add() IArithmetic_op_addContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_op_addContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmetic_op_addContext)
}

func (s *Arithmetic_addContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterArithmetic_add(s)
	}
}

func (s *Arithmetic_addContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitArithmetic_add(s)
	}
}

func (s *Arithmetic_addContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitArithmetic_add(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*ValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(NiavkaParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	*ValueContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedContext struct {
	*ValueContext
	n_value IValueContext
}

func NewNestedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedContext {
	var p = new(NestedContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NestedContext) GetN_value() IValueContext { return s.n_value }

func (s *NestedContext) SetN_value(v IValueContext) { s.n_value = v }

func (s *NestedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(NiavkaParserOPEN_PAREN, 0)
}

func (s *NestedContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(NiavkaParserCLOSE_PAREN, 0)
}

func (s *NestedContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *NestedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterNested(s)
	}
}

func (s *NestedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitNested(s)
	}
}

func (s *NestedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitNested(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Value() (localctx IValueContext) {
	return p.value(0)
}

func (p *NiavkaParser) value(_p int) (localctx IValueContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewValueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, NiavkaParserRULE_value, _p)
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NiavkaParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(79)
			p.Match(NiavkaParserNUMBER)
		}

	case NiavkaParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)
			p.Match(NiavkaParserSTRING)
		}

	case NiavkaParserID:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(81)
			p.Identifier()
		}

	case NiavkaParserOPEN_PAREN:
		localctx = NewNestedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(82)
			p.Match(NiavkaParserOPEN_PAREN)
		}
		{
			p.SetState(83)

			var _x = p.value(0)

			localctx.(*NestedContext).n_value = _x
		}
		{
			p.SetState(84)
			p.Match(NiavkaParserCLOSE_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmetic_mulContext(p, NewValueContext(p, _parentctx, _parentState))
				localctx.(*Arithmetic_mulContext).a_left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, NiavkaParserRULE_value)
				p.SetState(88)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(89)

					var _x = p.Arithmetic_op_mul()

					localctx.(*Arithmetic_mulContext).a_operation = _x
				}
				{
					p.SetState(90)

					var _x = p.value(3)

					localctx.(*Arithmetic_mulContext).a_right = _x
				}

			case 2:
				localctx = NewArithmetic_addContext(p, NewValueContext(p, _parentctx, _parentState))
				localctx.(*Arithmetic_addContext).a_left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, NiavkaParserRULE_value)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(93)

					var _x = p.Arithmetic_op_add()

					localctx.(*Arithmetic_addContext).a_operation = _x
				}
				{
					p.SetState(94)

					var _x = p.value(2)

					localctx.(*Arithmetic_addContext).a_right = _x
				}

			case 3:
				localctx = NewChainContext(p, NewValueContext(p, _parentctx, _parentState))
				localctx.(*ChainContext).c_left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, NiavkaParserRULE_value)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(97)
					p.Match(NiavkaParserDOT)
				}
				{
					p.SetState(98)

					var _x = p.Identifier()

					localctx.(*ChainContext).c_right = _x
				}

			case 4:
				localctx = NewCallContext(p, NewValueContext(p, _parentctx, _parentState))
				localctx.(*CallContext).c_value = _prevctx

				p.PushNewRecursionContext(localctx, _startState, NiavkaParserRULE_value)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(100)
					p.Match(NiavkaParserOPEN_PAREN)
				}
				p.SetState(102)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1245232) != 0 {
					{
						p.SetState(101)

						var _x = p.Args()

						localctx.(*CallContext).c_args = _x
					}

				}
				{
					p.SetState(104)
					p.Match(NiavkaParserCLOSE_PAREN)
				}

			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetA_left returns the a_left rule contexts.
	GetA_left() IIdentifierContext

	// GetA_type returns the a_type rule contexts.
	GetA_type() IType_valueContext

	// GetA_value returns the a_value rule contexts.
	GetA_value() IAssign_valueContext

	// SetA_left sets the a_left rule contexts.
	SetA_left(IIdentifierContext)

	// SetA_type sets the a_type rule contexts.
	SetA_type(IType_valueContext)

	// SetA_value sets the a_value rule contexts.
	SetA_value(IAssign_valueContext)

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Identifier() IIdentifierContext
	Assign_value() IAssign_valueContext
	Type_value() IType_valueContext

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	a_left  IIdentifierContext
	a_type  IType_valueContext
	a_value IAssign_valueContext
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) GetA_left() IIdentifierContext { return s.a_left }

func (s *AssignContext) GetA_type() IType_valueContext { return s.a_type }

func (s *AssignContext) GetA_value() IAssign_valueContext { return s.a_value }

func (s *AssignContext) SetA_left(v IIdentifierContext) { s.a_left = v }

func (s *AssignContext) SetA_type(v IType_valueContext) { s.a_type = v }

func (s *AssignContext) SetA_value(v IAssign_valueContext) { s.a_value = v }

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(NiavkaParserASSIGN, 0)
}

func (s *AssignContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssignContext) Assign_value() IAssign_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_valueContext)
}

func (s *AssignContext) Type_value() IType_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_valueContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NiavkaParserRULE_assign)
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
		p.SetState(110)

		var _x = p.Identifier()

		localctx.(*AssignContext).a_left = _x
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NiavkaParserID {
		{
			p.SetState(111)

			var _x = p.Type_value()

			localctx.(*AssignContext).a_type = _x
		}

	}
	{
		p.SetState(114)
		p.Match(NiavkaParserASSIGN)
	}
	{
		p.SetState(115)

		var _x = p.Assign_value()

		localctx.(*AssignContext).a_value = _x
	}

	return localctx
}

// IAssign_valueContext is an interface to support dynamic dispatch.
type IAssign_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext

	// IsAssign_valueContext differentiates from other interfaces.
	IsAssign_valueContext()
}

type Assign_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_valueContext() *Assign_valueContext {
	var p = new(Assign_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_assign_value
	return p
}

func (*Assign_valueContext) IsAssign_valueContext() {}

func NewAssign_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_valueContext {
	var p = new(Assign_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_assign_value

	return p
}

func (s *Assign_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_valueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Assign_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterAssign_value(s)
	}
}

func (s *Assign_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitAssign_value(s)
	}
}

func (s *Assign_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitAssign_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Assign_value() (localctx IAssign_valueContext) {
	this := p
	_ = this

	localctx = NewAssign_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NiavkaParserRULE_assign_value)

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
		p.SetState(117)
		p.value(0)
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(NiavkaParserID, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NiavkaParserRULE_identifier)

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
		p.SetState(119)
		p.Match(NiavkaParserID)
	}

	return localctx
}

// IType_valueContext is an interface to support dynamic dispatch.
type IType_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTv_single returns the tv_single rule contexts.
	GetTv_single() IIdentifierContext

	// SetTv_single sets the tv_single rule contexts.
	SetTv_single(IIdentifierContext)

	// Getter signatures
	Identifier() IIdentifierContext

	// IsType_valueContext differentiates from other interfaces.
	IsType_valueContext()
}

type Type_valueContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	tv_single IIdentifierContext
}

func NewEmptyType_valueContext() *Type_valueContext {
	var p = new(Type_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_type_value
	return p
}

func (*Type_valueContext) IsType_valueContext() {}

func NewType_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_valueContext {
	var p = new(Type_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_type_value

	return p
}

func (s *Type_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_valueContext) GetTv_single() IIdentifierContext { return s.tv_single }

func (s *Type_valueContext) SetTv_single(v IIdentifierContext) { s.tv_single = v }

func (s *Type_valueContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Type_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterType_value(s)
	}
}

func (s *Type_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitType_value(s)
	}
}

func (s *Type_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitType_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Type_value() (localctx IType_valueContext) {
	this := p
	_ = this

	localctx = NewType_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NiavkaParserRULE_type_value)

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
		p.SetState(121)

		var _x = p.Identifier()

		localctx.(*Type_valueContext).tv_single = _x
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArg() []IArgContext
	Arg(i int) IArgContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
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

	return t.(IArgContext)
}

func (s *ArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NiavkaParserCOMMA)
}

func (s *ArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NiavkaParserCOMMA, i)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Args() (localctx IArgsContext) {
	this := p
	_ = this

	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NiavkaParserRULE_args)
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
		p.SetState(123)
		p.Arg()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NiavkaParserCOMMA {
		{
			p.SetState(124)
			p.Match(NiavkaParserCOMMA)
		}
		{
			p.SetState(125)
			p.Arg()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetA_value returns the a_value rule contexts.
	GetA_value() IValueContext

	// SetA_value sets the a_value rule contexts.
	SetA_value(IValueContext)

	// Getter signatures
	AllNls() []INlsContext
	Nls(i int) INlsContext
	Value() IValueContext

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	a_value IValueContext
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) GetA_value() IValueContext { return s.a_value }

func (s *ArgContext) SetA_value(v IValueContext) { s.a_value = v }

func (s *ArgContext) AllNls() []INlsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlsContext); ok {
			len++
		}
	}

	tst := make([]INlsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlsContext); ok {
			tst[i] = t.(INlsContext)
			i++
		}
	}

	return tst
}

func (s *ArgContext) Nls(i int) INlsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlsContext); ok {
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

	return t.(INlsContext)
}

func (s *ArgContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitArg(s)
	}
}

func (s *ArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Arg() (localctx IArgContext) {
	this := p
	_ = this

	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NiavkaParserRULE_arg)

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
		p.SetState(131)
		p.Nls()
	}
	{
		p.SetState(132)

		var _x = p.value(0)

		localctx.(*ArgContext).a_value = _x
	}
	{
		p.SetState(133)
		p.Nls()
	}

	return localctx
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllNls() []INlsContext
	Nls(i int) INlsContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
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

	return t.(IParamContext)
}

func (s *ParamsContext) AllNls() []INlsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlsContext); ok {
			len++
		}
	}

	tst := make([]INlsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlsContext); ok {
			tst[i] = t.(INlsContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Nls(i int) INlsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlsContext); ok {
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

	return t.(INlsContext)
}

func (s *ParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(NiavkaParserCOMMA)
}

func (s *ParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(NiavkaParserCOMMA, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Params() (localctx IParamsContext) {
	this := p
	_ = this

	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NiavkaParserRULE_params)

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
	{
		p.SetState(135)
		p.Param()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(136)
				p.Nls()
			}
			{
				p.SetState(137)
				p.Match(NiavkaParserCOMMA)
			}
			{
				p.SetState(138)
				p.Nls()
			}
			{
				p.SetState(139)
				p.Param()
			}

		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP_name returns the p_name rule contexts.
	GetP_name() IIdentifierContext

	// GetP_type returns the p_type rule contexts.
	GetP_type() IType_valueContext

	// SetP_name sets the p_name rule contexts.
	SetP_name(IIdentifierContext)

	// SetP_type sets the p_type rule contexts.
	SetP_type(IType_valueContext)

	// Getter signatures
	Identifier() IIdentifierContext
	Type_value() IType_valueContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	p_name IIdentifierContext
	p_type IType_valueContext
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) GetP_name() IIdentifierContext { return s.p_name }

func (s *ParamContext) GetP_type() IType_valueContext { return s.p_type }

func (s *ParamContext) SetP_name(v IIdentifierContext) { s.p_name = v }

func (s *ParamContext) SetP_type(v IType_valueContext) { s.p_type = v }

func (s *ParamContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ParamContext) Type_value() IType_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_valueContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Param() (localctx IParamContext) {
	this := p
	_ = this

	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NiavkaParserRULE_param)

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
		p.SetState(146)

		var _x = p.Identifier()

		localctx.(*ParamContext).p_name = _x
	}
	{
		p.SetState(147)

		var _x = p.Type_value()

		localctx.(*ParamContext).p_type = _x
	}

	return localctx
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBody_element() []IBody_elementContext
	Body_element(i int) IBody_elementContext
	AllNl() []INlContext
	Nl(i int) INlContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) AllBody_element() []IBody_elementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBody_elementContext); ok {
			len++
		}
	}

	tst := make([]IBody_elementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBody_elementContext); ok {
			tst[i] = t.(IBody_elementContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Body_element(i int) IBody_elementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBody_elementContext); ok {
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

	return t.(IBody_elementContext)
}

func (s *BodyContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
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

	return t.(INlContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Body() (localctx IBodyContext) {
	this := p
	_ = this

	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NiavkaParserRULE_body)

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
	{
		p.SetState(149)
		p.Body_element()
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(150)
				p.Nl()
			}
			{
				p.SetState(151)
				p.Body_element()
			}

		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IBody_elementContext is an interface to support dynamic dispatch.
type IBody_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	Assign() IAssignContext
	Nls() INlsContext

	// IsBody_elementContext differentiates from other interfaces.
	IsBody_elementContext()
}

type Body_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBody_elementContext() *Body_elementContext {
	var p = new(Body_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_body_element
	return p
}

func (*Body_elementContext) IsBody_elementContext() {}

func NewBody_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Body_elementContext {
	var p = new(Body_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_body_element

	return p
}

func (s *Body_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Body_elementContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Body_elementContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *Body_elementContext) Nls() INlsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INlsContext)
}

func (s *Body_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Body_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Body_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterBody_element(s)
	}
}

func (s *Body_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitBody_element(s)
	}
}

func (s *Body_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitBody_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Body_element() (localctx IBody_elementContext) {
	this := p
	_ = this

	localctx = NewBody_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NiavkaParserRULE_body_element)

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

	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.value(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Assign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Nls()
		}

	}

	return localctx
}

// IArithmetic_op_mulContext is an interface to support dynamic dispatch.
type IArithmetic_op_mulContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode

	// IsArithmetic_op_mulContext differentiates from other interfaces.
	IsArithmetic_op_mulContext()
}

type Arithmetic_op_mulContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmetic_op_mulContext() *Arithmetic_op_mulContext {
	var p = new(Arithmetic_op_mulContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_arithmetic_op_mul
	return p
}

func (*Arithmetic_op_mulContext) IsArithmetic_op_mulContext() {}

func NewArithmetic_op_mulContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithmetic_op_mulContext {
	var p = new(Arithmetic_op_mulContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_arithmetic_op_mul

	return p
}

func (s *Arithmetic_op_mulContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithmetic_op_mulContext) MUL() antlr.TerminalNode {
	return s.GetToken(NiavkaParserMUL, 0)
}

func (s *Arithmetic_op_mulContext) DIV() antlr.TerminalNode {
	return s.GetToken(NiavkaParserDIV, 0)
}

func (s *Arithmetic_op_mulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithmetic_op_mulContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arithmetic_op_mulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterArithmetic_op_mul(s)
	}
}

func (s *Arithmetic_op_mulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitArithmetic_op_mul(s)
	}
}

func (s *Arithmetic_op_mulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitArithmetic_op_mul(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Arithmetic_op_mul() (localctx IArithmetic_op_mulContext) {
	this := p
	_ = this

	localctx = NewArithmetic_op_mulContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NiavkaParserRULE_arithmetic_op_mul)
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
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NiavkaParserMUL || _la == NiavkaParserDIV) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArithmetic_op_addContext is an interface to support dynamic dispatch.
type IArithmetic_op_addContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsArithmetic_op_addContext differentiates from other interfaces.
	IsArithmetic_op_addContext()
}

type Arithmetic_op_addContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmetic_op_addContext() *Arithmetic_op_addContext {
	var p = new(Arithmetic_op_addContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_arithmetic_op_add
	return p
}

func (*Arithmetic_op_addContext) IsArithmetic_op_addContext() {}

func NewArithmetic_op_addContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithmetic_op_addContext {
	var p = new(Arithmetic_op_addContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_arithmetic_op_add

	return p
}

func (s *Arithmetic_op_addContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithmetic_op_addContext) PLUS() antlr.TerminalNode {
	return s.GetToken(NiavkaParserPLUS, 0)
}

func (s *Arithmetic_op_addContext) MINUS() antlr.TerminalNode {
	return s.GetToken(NiavkaParserMINUS, 0)
}

func (s *Arithmetic_op_addContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithmetic_op_addContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arithmetic_op_addContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterArithmetic_op_add(s)
	}
}

func (s *Arithmetic_op_addContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitArithmetic_op_add(s)
	}
}

func (s *Arithmetic_op_addContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitArithmetic_op_add(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Arithmetic_op_add() (localctx IArithmetic_op_addContext) {
	this := p
	_ = this

	localctx = NewArithmetic_op_addContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NiavkaParserRULE_arithmetic_op_add)
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
		p.SetState(165)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NiavkaParserPLUS || _la == NiavkaParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INlContext is an interface to support dynamic dispatch.
type INlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NL() antlr.TerminalNode

	// IsNlContext differentiates from other interfaces.
	IsNlContext()
}

type NlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNlContext() *NlContext {
	var p = new(NlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_nl
	return p
}

func (*NlContext) IsNlContext() {}

func NewNlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NlContext {
	var p = new(NlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_nl

	return p
}

func (s *NlContext) GetParser() antlr.Parser { return s.parser }

func (s *NlContext) NL() antlr.TerminalNode {
	return s.GetToken(NiavkaParserNL, 0)
}

func (s *NlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterNl(s)
	}
}

func (s *NlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitNl(s)
	}
}

func (s *NlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitNl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Nl() (localctx INlContext) {
	this := p
	_ = this

	localctx = NewNlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NiavkaParserRULE_nl)

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
		p.SetState(167)
		p.Match(NiavkaParserNL)
	}

	return localctx
}

// INlsContext is an interface to support dynamic dispatch.
type INlsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNl() []INlContext
	Nl(i int) INlContext

	// IsNlsContext differentiates from other interfaces.
	IsNlsContext()
}

type NlsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNlsContext() *NlsContext {
	var p = new(NlsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NiavkaParserRULE_nls
	return p
}

func (*NlsContext) IsNlsContext() {}

func NewNlsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NlsContext {
	var p = new(NlsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NiavkaParserRULE_nls

	return p
}

func (s *NlsContext) GetParser() antlr.Parser { return s.parser }

func (s *NlsContext) AllNl() []INlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INlContext); ok {
			len++
		}
	}

	tst := make([]INlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INlContext); ok {
			tst[i] = t.(INlContext)
			i++
		}
	}

	return tst
}

func (s *NlsContext) Nl(i int) INlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INlContext); ok {
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

	return t.(INlContext)
}

func (s *NlsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NlsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NlsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.EnterNls(s)
	}
}

func (s *NlsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NiavkaParserListener); ok {
		listenerT.ExitNls(s)
	}
}

func (s *NlsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NiavkaParserVisitor:
		return t.VisitNls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NiavkaParser) Nls() (localctx INlsContext) {
	this := p
	_ = this

	localctx = NewNlsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, NiavkaParserRULE_nls)

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
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(169)
				p.Nl()
			}

		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

func (p *NiavkaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ValueContext = nil
		if localctx != nil {
			t = localctx.(*ValueContext)
		}
		return p.Value_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NiavkaParser) Value_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
