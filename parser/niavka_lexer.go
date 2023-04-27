// Code generated from NiavkaLexer.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type NiavkaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var niavkalexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func niavkalexerLexerInit() {
	staticData := &niavkalexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"END", "DIIA", "SKIP_SPACES", "NL", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_ARRAY",
		"CLOSE_ARRAY", "COMMA", "ASSIGN", "PLUS", "MINUS", "MUL", "DIV", "DOT",
		"ID", "NUMBER", "INTEGER", "FLOAT", "STRING", "DIGIT", "ID_START", "ID_CONTINUE",
		"ESCAPE_CHAR",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 133, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 3,
		3, 66, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 94, 8, 15, 10, 15, 12, 15, 97,
		9, 15, 1, 16, 1, 16, 3, 16, 101, 8, 16, 1, 17, 4, 17, 104, 8, 17, 11, 17,
		12, 17, 105, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 115,
		8, 19, 10, 19, 12, 19, 118, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 3, 22, 129, 8, 22, 1, 23, 1, 23, 1, 23, 0, 0,
		24, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 0, 43, 0, 45, 0, 47, 0, 1, 0, 5, 2, 0, 9, 9, 32, 32, 2, 0, 34,
		34, 92, 92, 9, 0, 65, 90, 95, 95, 97, 122, 1028, 1028, 1030, 1031, 1040,
		1103, 1108, 1108, 1110, 1111, 1168, 1169, 2, 0, 48, 57, 700, 700, 9, 0,
		34, 34, 39, 39, 48, 48, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116,
		116, 136, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 1, 49, 1, 0, 0, 0, 3, 56, 1, 0, 0, 0, 5, 60, 1, 0,
		0, 0, 7, 65, 1, 0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 71, 1, 0, 0, 0, 13, 73,
		1, 0, 0, 0, 15, 75, 1, 0, 0, 0, 17, 77, 1, 0, 0, 0, 19, 79, 1, 0, 0, 0,
		21, 81, 1, 0, 0, 0, 23, 83, 1, 0, 0, 0, 25, 85, 1, 0, 0, 0, 27, 87, 1,
		0, 0, 0, 29, 89, 1, 0, 0, 0, 31, 91, 1, 0, 0, 0, 33, 100, 1, 0, 0, 0, 35,
		103, 1, 0, 0, 0, 37, 107, 1, 0, 0, 0, 39, 111, 1, 0, 0, 0, 41, 121, 1,
		0, 0, 0, 43, 123, 1, 0, 0, 0, 45, 128, 1, 0, 0, 0, 47, 130, 1, 0, 0, 0,
		49, 50, 5, 1082, 0, 0, 50, 51, 5, 1110, 0, 0, 51, 52, 5, 1085, 0, 0, 52,
		53, 5, 1077, 0, 0, 53, 54, 5, 1094, 0, 0, 54, 55, 5, 1100, 0, 0, 55, 2,
		1, 0, 0, 0, 56, 57, 5, 1076, 0, 0, 57, 58, 5, 1110, 0, 0, 58, 59, 5, 1103,
		0, 0, 59, 4, 1, 0, 0, 0, 60, 61, 7, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63,
		6, 2, 0, 0, 63, 6, 1, 0, 0, 0, 64, 66, 5, 13, 0, 0, 65, 64, 1, 0, 0, 0,
		65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 5, 10, 0, 0, 68, 8, 1,
		0, 0, 0, 69, 70, 5, 40, 0, 0, 70, 10, 1, 0, 0, 0, 71, 72, 5, 41, 0, 0,
		72, 12, 1, 0, 0, 0, 73, 74, 5, 91, 0, 0, 74, 14, 1, 0, 0, 0, 75, 76, 5,
		93, 0, 0, 76, 16, 1, 0, 0, 0, 77, 78, 5, 44, 0, 0, 78, 18, 1, 0, 0, 0,
		79, 80, 5, 61, 0, 0, 80, 20, 1, 0, 0, 0, 81, 82, 5, 43, 0, 0, 82, 22, 1,
		0, 0, 0, 83, 84, 5, 45, 0, 0, 84, 24, 1, 0, 0, 0, 85, 86, 5, 42, 0, 0,
		86, 26, 1, 0, 0, 0, 87, 88, 5, 47, 0, 0, 88, 28, 1, 0, 0, 0, 89, 90, 5,
		46, 0, 0, 90, 30, 1, 0, 0, 0, 91, 95, 3, 43, 21, 0, 92, 94, 3, 45, 22,
		0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96,
		1, 0, 0, 0, 96, 32, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 101, 3, 37, 18,
		0, 99, 101, 3, 35, 17, 0, 100, 98, 1, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101,
		34, 1, 0, 0, 0, 102, 104, 3, 41, 20, 0, 103, 102, 1, 0, 0, 0, 104, 105,
		1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 36, 1, 0,
		0, 0, 107, 108, 3, 35, 17, 0, 108, 109, 5, 46, 0, 0, 109, 110, 3, 35, 17,
		0, 110, 38, 1, 0, 0, 0, 111, 116, 5, 34, 0, 0, 112, 115, 8, 1, 0, 0, 113,
		115, 3, 47, 23, 0, 114, 112, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118,
		1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 1, 0,
		0, 0, 118, 116, 1, 0, 0, 0, 119, 120, 5, 34, 0, 0, 120, 40, 1, 0, 0, 0,
		121, 122, 2, 48, 57, 0, 122, 42, 1, 0, 0, 0, 123, 124, 7, 2, 0, 0, 124,
		44, 1, 0, 0, 0, 125, 129, 3, 43, 21, 0, 126, 129, 7, 3, 0, 0, 127, 129,
		3, 43, 21, 0, 128, 125, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 127, 1,
		0, 0, 0, 129, 46, 1, 0, 0, 0, 130, 131, 5, 92, 0, 0, 131, 132, 7, 4, 0,
		0, 132, 48, 1, 0, 0, 0, 8, 0, 65, 95, 100, 105, 114, 116, 128, 1, 6, 0,
		0,
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

// NiavkaLexerInit initializes any static state used to implement NiavkaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNiavkaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NiavkaLexerInit() {
	staticData := &niavkalexerLexerStaticData
	staticData.once.Do(niavkalexerLexerInit)
}

// NewNiavkaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNiavkaLexer(input antlr.CharStream) *NiavkaLexer {
	NiavkaLexerInit()
	l := new(NiavkaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &niavkalexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "NiavkaLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NiavkaLexer tokens.
const (
	NiavkaLexerEND         = 1
	NiavkaLexerDIIA        = 2
	NiavkaLexerSKIP_SPACES = 3
	NiavkaLexerNL          = 4
	NiavkaLexerOPEN_PAREN  = 5
	NiavkaLexerCLOSE_PAREN = 6
	NiavkaLexerOPEN_ARRAY  = 7
	NiavkaLexerCLOSE_ARRAY = 8
	NiavkaLexerCOMMA       = 9
	NiavkaLexerASSIGN      = 10
	NiavkaLexerPLUS        = 11
	NiavkaLexerMINUS       = 12
	NiavkaLexerMUL         = 13
	NiavkaLexerDIV         = 14
	NiavkaLexerDOT         = 15
	NiavkaLexerID          = 16
	NiavkaLexerNUMBER      = 17
	NiavkaLexerINTEGER     = 18
	NiavkaLexerFLOAT       = 19
	NiavkaLexerSTRING      = 20
)