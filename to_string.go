package sqlex

import (
	"strconv"
	"time"
)

func ToString(col ColumnTypeInfo, val interface{}) string {
	if val == nil {
		return ""
	}

	kind := getKindFromColumn(col)

	if isStringKind(kind) {
		return val.(string)
	}

	if isSliceKind(kind) {
		return string(val.([]byte))
	}

	if isIntKind(kind) {
		return strconv.FormatInt(castToInt64(val), 10)
	}

	if isUintKind(kind) {
		return strconv.FormatUint(castToUint64(val), 10)
	}

	if isFloatKind(kind) {
		fval := castToFloat64(val)
		return strconv.FormatFloat(fval, 'f', 6, 64)
	}

	if isBoolKind(kind) {
		if val.(bool) {
			return "true"
		} else {
			return "false"
		}
	}

	if isTimeKind(kind, col) {
		return val.(time.Time).Format(time.RFC3339)
	}

	return ""
}
