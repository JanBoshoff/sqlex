package sqlex

import ()

func ToUint(col ColumnTypeInfo, val interface{}) uint64 {
	if val == nil {
		return 0
	}

	/*colType := col.ScanType()
	kind := colType.Kind()

	if kind == reflect.Int ||
		kind == reflect.Int64 ||
		kind == reflect.Int32 ||
		kind == reflect.Int16 ||
		kind == reflect.Int8 {

		return castToInt64(val)
	}

	if kind == reflect.Uint ||
		kind == reflect.Uint64 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint8 {

		convInt := castToUint64(val)

		if convInt > math.MaxInt64 {
			return 0
		}

		return int64(castToUint64(val))
	}

	if kind == reflect.String {
		conv, err := strconv.ParseInt(val.(string), 10, 64)

		if err != nil {
			return 0
		}

		return conv
	}

	if kind == reflect.Slice {
		conv, err := strconv.ParseInt(string(val.([]byte)), 10, 64)

		if err != nil {
			return 0
		}

		return conv
	}

	if kind == reflect.Float64 || kind == reflect.Float32 {
		fval := castToFloat64(val)
		return int64(fval)
	}

	if kind == reflect.Bool && val.(bool) == true {
		return int64(1)
	}

	if kind == reflect.Bool && val.(bool) == false {
		return int64(0)
	}
	*/

	return 0
}
