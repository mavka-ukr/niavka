package compiler

import (
	"fmt"
	"github.com/llir/llvm/ir/types"
)

// Variable in scope
type Variable struct {
	Name  string
	Type  VariableType
	Scope *Scope
}

type VariableType interface {
	IsBuiltIn() bool
	AsIr() types.Type
}

type PrimitiveType string

func (t PrimitiveType) IsBuiltIn() bool {
	return true
}

const (
	Boolean PrimitiveType = "логічне"
	Int8    PrimitiveType = "ц8"
	Int16   PrimitiveType = "ц16"
	Int32   PrimitiveType = "ц32"
	Int64   PrimitiveType = "ц64"
	UInt8   PrimitiveType = "н8"
	UInt16  PrimitiveType = "н16"
	UInt32  PrimitiveType = "н32"
	UInt64  PrimitiveType = "н64"
	Float16 PrimitiveType = "д16"
	Float32 PrimitiveType = "д32"
	Float64 PrimitiveType = "д64"
	Str     PrimitiveType = "рядок"
	Void    PrimitiveType = "ніщо"
)

type NoSuchBuiltinTypeError string

func (e NoSuchBuiltinTypeError) Error() string {
	return fmt.Sprintf("no such built-in type: %q", e)
}

func (e NoSuchBuiltinTypeError) String() string {
	return e.Error()
}

func PrimitiveFromName(s string) (PrimitiveType, error) {
	switch b := PrimitiveType(s); b {
	case Boolean, Int8, Int16, Int32, Int64, UInt8, UInt16, UInt32, UInt64, Float16, Float32, Float64, Str, Void:
		return b, nil
	default:
		return "", NoSuchBuiltinTypeError(s)
	}
}

func (t PrimitiveType) AsIr() types.Type {
	switch t {
	case Boolean:
		return types.I1
	case Int8, UInt8:
		return types.I8
	case Int16, UInt16:
		return types.I16
	case Int32, UInt32:
		return types.I32
	case Int64, UInt64:
		return types.I64
	case Float16:
		return types.Half
	case Float32:
		return types.Float
	case Float64:
		return types.Double
	case Str:
		return types.NewStruct(types.I8Ptr, types.I64)
	case Void:
		return types.Void
	default:
		return nil
	}
}
