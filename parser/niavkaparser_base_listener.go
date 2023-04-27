// Code generated from NiavkaParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // NiavkaParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseNiavkaParserListener is a complete listener for a parse tree produced by NiavkaParser.
type BaseNiavkaParserListener struct{}

var _ NiavkaParserListener = &BaseNiavkaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNiavkaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNiavkaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNiavkaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNiavkaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseNiavkaParserListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseNiavkaParserListener) ExitFile(ctx *FileContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseNiavkaParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseNiavkaParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterProgram_element is called when production program_element is entered.
func (s *BaseNiavkaParserListener) EnterProgram_element(ctx *Program_elementContext) {}

// ExitProgram_element is called when production program_element is exited.
func (s *BaseNiavkaParserListener) ExitProgram_element(ctx *Program_elementContext) {}

// EnterDiia is called when production diia is entered.
func (s *BaseNiavkaParserListener) EnterDiia(ctx *DiiaContext) {}

// ExitDiia is called when production diia is exited.
func (s *BaseNiavkaParserListener) ExitDiia(ctx *DiiaContext) {}

// EnterArithmetic_mul is called when production arithmetic_mul is entered.
func (s *BaseNiavkaParserListener) EnterArithmetic_mul(ctx *Arithmetic_mulContext) {}

// ExitArithmetic_mul is called when production arithmetic_mul is exited.
func (s *BaseNiavkaParserListener) ExitArithmetic_mul(ctx *Arithmetic_mulContext) {}

// EnterCall is called when production call is entered.
func (s *BaseNiavkaParserListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseNiavkaParserListener) ExitCall(ctx *CallContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseNiavkaParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseNiavkaParserListener) ExitNumber(ctx *NumberContext) {}

// EnterChain is called when production chain is entered.
func (s *BaseNiavkaParserListener) EnterChain(ctx *ChainContext) {}

// ExitChain is called when production chain is exited.
func (s *BaseNiavkaParserListener) ExitChain(ctx *ChainContext) {}

// EnterArithmetic_add is called when production arithmetic_add is entered.
func (s *BaseNiavkaParserListener) EnterArithmetic_add(ctx *Arithmetic_addContext) {}

// ExitArithmetic_add is called when production arithmetic_add is exited.
func (s *BaseNiavkaParserListener) ExitArithmetic_add(ctx *Arithmetic_addContext) {}

// EnterString is called when production string is entered.
func (s *BaseNiavkaParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseNiavkaParserListener) ExitString(ctx *StringContext) {}

// EnterId is called when production id is entered.
func (s *BaseNiavkaParserListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseNiavkaParserListener) ExitId(ctx *IdContext) {}

// EnterNested is called when production nested is entered.
func (s *BaseNiavkaParserListener) EnterNested(ctx *NestedContext) {}

// ExitNested is called when production nested is exited.
func (s *BaseNiavkaParserListener) ExitNested(ctx *NestedContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseNiavkaParserListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseNiavkaParserListener) ExitAssign(ctx *AssignContext) {}

// EnterAssign_value is called when production assign_value is entered.
func (s *BaseNiavkaParserListener) EnterAssign_value(ctx *Assign_valueContext) {}

// ExitAssign_value is called when production assign_value is exited.
func (s *BaseNiavkaParserListener) ExitAssign_value(ctx *Assign_valueContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseNiavkaParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseNiavkaParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterType_value is called when production type_value is entered.
func (s *BaseNiavkaParserListener) EnterType_value(ctx *Type_valueContext) {}

// ExitType_value is called when production type_value is exited.
func (s *BaseNiavkaParserListener) ExitType_value(ctx *Type_valueContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseNiavkaParserListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseNiavkaParserListener) ExitArgs(ctx *ArgsContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseNiavkaParserListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseNiavkaParserListener) ExitArg(ctx *ArgContext) {}

// EnterParams is called when production params is entered.
func (s *BaseNiavkaParserListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseNiavkaParserListener) ExitParams(ctx *ParamsContext) {}

// EnterParam is called when production param is entered.
func (s *BaseNiavkaParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseNiavkaParserListener) ExitParam(ctx *ParamContext) {}

// EnterBody is called when production body is entered.
func (s *BaseNiavkaParserListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseNiavkaParserListener) ExitBody(ctx *BodyContext) {}

// EnterBody_element is called when production body_element is entered.
func (s *BaseNiavkaParserListener) EnterBody_element(ctx *Body_elementContext) {}

// ExitBody_element is called when production body_element is exited.
func (s *BaseNiavkaParserListener) ExitBody_element(ctx *Body_elementContext) {}

// EnterArithmetic_op_mul is called when production arithmetic_op_mul is entered.
func (s *BaseNiavkaParserListener) EnterArithmetic_op_mul(ctx *Arithmetic_op_mulContext) {}

// ExitArithmetic_op_mul is called when production arithmetic_op_mul is exited.
func (s *BaseNiavkaParserListener) ExitArithmetic_op_mul(ctx *Arithmetic_op_mulContext) {}

// EnterArithmetic_op_add is called when production arithmetic_op_add is entered.
func (s *BaseNiavkaParserListener) EnterArithmetic_op_add(ctx *Arithmetic_op_addContext) {}

// ExitArithmetic_op_add is called when production arithmetic_op_add is exited.
func (s *BaseNiavkaParserListener) ExitArithmetic_op_add(ctx *Arithmetic_op_addContext) {}

// EnterNl is called when production nl is entered.
func (s *BaseNiavkaParserListener) EnterNl(ctx *NlContext) {}

// ExitNl is called when production nl is exited.
func (s *BaseNiavkaParserListener) ExitNl(ctx *NlContext) {}

// EnterNls is called when production nls is entered.
func (s *BaseNiavkaParserListener) EnterNls(ctx *NlsContext) {}

// ExitNls is called when production nls is exited.
func (s *BaseNiavkaParserListener) ExitNls(ctx *NlsContext) {}
