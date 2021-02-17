package sqlex

import (
	"math"
	"strconv"
	"time"
)

func ToInt(col ColumnTypeInfo, val interface{}) int64 {
	if val == nil {
		return 0
	}

	kind := getKindFromColumn(col)

	if isIntKind(kind) {
		return castToInt64(val)
	}

	if isUintKind(kind) {
		convInt := castToUint64(val)

		if convInt > math.MaxInt64 {
			return 0
		}

		return int64(castToUint64(val))
	}

	if isStringKind(kind) {
		conv, err := strconv.ParseInt(val.(string), 10, 64)

		if err != nil {
			return 0
		}

		return conv
	}

	if isSliceKind(kind) {
		conv, err := strconv.ParseInt(string(val.([]byte)), 10, 64)

		if err != nil {
			return 0
		}

		return conv
	}

	if isFloatKind(kind) {
		fval := castToFloat64(val)
		return int64(fval)
	}

	if isBoolKind(kind) && val.(bool) == true {
		return int64(1)
	}

	if isBoolKind(kind) && val.(bool) == false {
		return int64(0)
	}

	if isTimeKind(kind, col) {
		return val.(time.Time).Unix()
	}

	return 0
}
