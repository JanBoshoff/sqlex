package sqlex

import (
	"database/sql"
	"reflect"
	"time"
)

type QueryResult interface {
	ColumnTypes() ([]*sql.ColumnType, error)
	Scan(dest ...interface{}) error
}

type ColumnTypeInfo interface {
	ScanType() reflect.Type
}

func getKindFromColumn(col ColumnTypeInfo) reflect.Kind {
	colType := col.ScanType()
	return colType.Kind()
}

func castToInt64(v interface{}) int64 {
	switch v := v.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return int64(v)
	}

	return 0
}

func castToUint64(v interface{}) uint64 {
	switch v := v.(type) {
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return uint64(v)
	}

	return 0
}

func castToFloat64(v interface{}) float64 {
	switch v := v.(type) {
	case float64:
		return float64(v)
	case float32:
		return float64(v)
	}

	return 0
}

func isIntKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return true
	}

	return false
}

func isUintKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return true
	}

	return false
}

func isFloatKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Float64, reflect.Float32:
		return true
	}

	return false
}

func isStringKind(kind reflect.Kind) bool {
	return kind == reflect.String
}

func isSliceKind(kind reflect.Kind) bool {
	return kind == reflect.Slice
}

func isBoolKind(kind reflect.Kind) bool {
	return kind == reflect.Bool
}

func isTimeKind(kind reflect.Kind, col ColumnTypeInfo) bool {
	if kind != reflect.Struct {
		return false
	}

	colType := col.ScanType()

	if colType != reflect.TypeOf(time.Now()) {
		return false
	}

	return true
}
