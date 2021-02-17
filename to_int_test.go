package sqlex_test

import (
	"math"
	"testing"
	"time"

	"github.com/JanBoshoff/sqlex"
)

/**
 * ToInt Nil Inputs
 */

func TestToIntWithNilInput(t *testing.T) {
	expected := int64(0)
	col := NewColumnTypeMock(int64(1))
	result := sqlex.ToInt(col, nil)

	if result != expected {
		failTest(t, expected, result)
	}
}

/**
 * ToInt Int Inputs
 */

func TestToIntWithIntInput(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(1)
	result := sqlex.ToInt(col, 1)

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithInt64Input(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(int64(1))
	result := sqlex.ToInt(col, int64(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithInt32Input(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(int32(1))
	result := sqlex.ToInt(col, int32(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithInt16Input(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(int16(1))
	result := sqlex.ToInt(col, int16(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithInt8Input(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(int8(1))
	result := sqlex.ToInt(col, int8(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithUIntInput(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(1)
	result := sqlex.ToInt(col, 1)

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithUInt64Input(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(uint64(1))
	result := sqlex.ToInt(col, uint64(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithUInt32Input(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(uint32(1))
	result := sqlex.ToInt(col, uint32(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithUInt16Input(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(uint16(1))
	result := sqlex.ToInt(col, uint16(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithUInt8Input(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(uint8(1))
	result := sqlex.ToInt(col, uint8(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithUInt8InputLargerThanInt64(t *testing.T) {
	expected := int64(0)
	col := NewColumnTypeMock(uint64(1))
	result := sqlex.ToInt(col, uint64(math.MaxUint64))

	if result != expected {
		failTest(t, expected, result)
	}
}

/**
 * ToInt String Inputs
 */

func TestToIntWithIntStringInput(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock("")
	result := sqlex.ToInt(col, "1")

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithIntSliceInput(t *testing.T) {
	expected := int64(11)
	col := NewColumnTypeMock([]byte(""))
	result := sqlex.ToInt(col, []byte("11"))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithNonIntStringInput(t *testing.T) {
	expected := int64(0)
	col := NewColumnTypeMock("")
	result := sqlex.ToInt(col, "not an int")

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithNonIntSliceInput(t *testing.T) {
	expected := int64(0)
	col := NewColumnTypeMock([]byte(""))
	result := sqlex.ToInt(col, []byte("not an int"))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithMixedStringInput(t *testing.T) {
	expected := int64(0)
	col := NewColumnTypeMock("")
	result := sqlex.ToInt(col, "123test")

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithMixedSliceInput(t *testing.T) {
	expected := int64(0)
	col := NewColumnTypeMock([]byte(""))
	result := sqlex.ToInt(col, []byte("123test"))

	if result != expected {
		failTest(t, expected, result)
	}
}

/**
 * ToInt Float Inputs
 */

func TestToIntWithFloatInput(t *testing.T) {
	expected := int64(5)
	col := NewColumnTypeMock(float64(5.7))
	result := sqlex.ToInt(col, float64(5.7))

	if result != expected {
		failTest(t, expected, result)
	}
}

/**
 * ToInt Bool Inputs
 */

func TestToIntWithBoolTrueInput(t *testing.T) {
	expected := int64(1)
	col := NewColumnTypeMock(true)
	result := sqlex.ToInt(col, true)

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToIntWithBoolFalseInput(t *testing.T) {
	expected := int64(0)
	col := NewColumnTypeMock(true)
	result := sqlex.ToInt(col, false)

	if result != expected {
		failTest(t, expected, result)
	}
}

/**
 * ToInt Time Inputs
 */

func TestToIntWithTimeInput(t *testing.T) {
	now := time.Now()

	expected := now.Unix()
	col := NewColumnTypeMock(time.Now())
	result := sqlex.ToInt(col, now)

	if result != expected {
		failTest(t, expected, result)
	}
}
