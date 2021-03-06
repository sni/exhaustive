package enumvariants

// Integer, string, and float can be enum types.
// Bool cannot be enum type.
// Only basic types can be enum types.

type UIntEnum uint

const (
	UIntA UIntEnum = 0
	UIntB UIntEnum = 1
)

type StringEnum string

const (
	StringA StringEnum = "stringa"
	StringB StringEnum = "stringb"
	StringC StringEnum = "stringc"
)

type RuneEnum rune

const (
	RuneA RuneEnum = 'a'
)

type ByteEnum byte

const (
	ByteA = ByteEnum('a')
)

type Int32Enum int32

const (
	Int32A Int32Enum = 0
	Int32B Int32Enum = 1
)

type FloatEnum float64

const (
	FloatEnumA FloatEnum = iota
	FloatEnumB
)

type BoolNotEnum bool

const (
	BoolNotEnumA BoolNotEnum = true
	BoolNotEnumB BoolNotEnum = false
)

type NonBasicType S

type S struct{ F int }

var (
	SA NonBasicType = NonBasicType{F: 1}
	SB NonBasicType = NonBasicType{F: 2}
)

const There AcrossBlocksDeclsFiles = 2
