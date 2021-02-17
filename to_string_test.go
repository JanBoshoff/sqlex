package sqlex_test

import (
	"testing"
	"time"

	"github.com/JanBoshoff/sqlex"
)

/**
 * ToString String and Slice Input Tests
 */

func TestToStringWithNilInput(t *testing.T) {
	expected := ""
	col := NewColumnTypeMock("")
	result := sqlex.ToString(col, nil)

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithEmptyStringInput(t *testing.T) {
	expected := ""
	col := NewColumnTypeMock("")
	result := sqlex.ToString(col, "")

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithStringInput(t *testing.T) {
	expected := "test"
	col := NewColumnTypeMock("")
	result := sqlex.ToString(col, "test")

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithEmptySliceInput(t *testing.T) {
	expected := ""
	col := NewColumnTypeMock([]byte("string"))
	result := sqlex.ToString(col, []byte(""))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithSliceInput(t *testing.T) {
	expected := "test"
	col := NewColumnTypeMock([]byte("string"))
	result := sqlex.ToString(col, []byte("test"))

	if result != expected {
		failTest(t, expected, result)
	}
}

/**
 * ToString Integer Input Tests
 */

func TestToStringWithIntInput(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(1)
	result := sqlex.ToString(col, 1)

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithInt64Input(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(int64(1))
	result := sqlex.ToString(col, int64(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithInt32Input(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(int32(1))
	result := sqlex.ToString(col, int32(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithInt16Input(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(int16(1))
	result := sqlex.ToString(col, int16(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithInt8Input(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(int8(1))
	result := sqlex.ToString(col, int8(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithUIntInput(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(uint(1))
	result := sqlex.ToString(col, uint(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithUInt64Input(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(uint64(1))
	result := sqlex.ToString(col, uint64(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithUInt32Input(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(uint32(1))
	result := sqlex.ToString(col, uint32(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithUInt16Input(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(uint16(1))
	result := sqlex.ToString(col, uint16(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithUInt8Input(t *testing.T) {
	expected := "1"
	col := NewColumnTypeMock(uint8(1))
	result := sqlex.ToString(col, uint8(1))

	if result != expected {
		failTest(t, expected, result)
	}
}

/**
 * ToString Float Input Tests
 */

func TestToStringWithFloat64Input(t *testing.T) {
	expected := "1.100000"
	col := NewColumnTypeMock(float64(1.1))
	result := sqlex.ToString(col, float64(1.1))

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithFloat32Input(t *testing.T) {
	expected := "1.100000"
	col := NewColumnTypeMock(float32(1.1))
	result := sqlex.ToString(col, float32(1.1))

	if result != expected {
		failTest(t, expected, result)
	}
}

/**
 * ToString Bool Input Tests
 */

func TestToStringWithBoolTrueInput(t *testing.T) {
	expected := "true"
	col := NewColumnTypeMock(true)
	result := sqlex.ToString(col, true)

	if result != expected {
		failTest(t, expected, result)
	}
}

func TestToStringWithBoolFalseInput(t *testing.T) {
	expected := "false"
	col := NewColumnTypeMock(false)
	result := sqlex.ToString(col, false)

	if result != expected {
		failTest(t, expected, result)
	}
}

/**
 * ToString Bool Input Tests
 */

func TestToStringWithTimeInput(t *testing.T) {
	now := time.Now()

	expected := now.Format(time.RFC3339)
	col := NewColumnTypeMock(now)
	result := sqlex.ToString(col, now)

	if result != expected {
		failTest(t, expected, result)
	}
}
