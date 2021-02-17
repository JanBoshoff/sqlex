package sqlex_test

import (
	"testing"

	"github.com/JanBoshoff/sqlex"
)

/**
 * ToUint Nil Inputs
 */

func TestToUintWithNilInput(t *testing.T) {
	expected := uint64(0)
	col := NewColumnTypeMock(uint64(1))
	result := sqlex.ToUint(col, nil)

	if result != expected {
		failTest(t, expected, result)
	}
}
