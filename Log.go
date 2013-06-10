//=======================================//
// Author: Austin Mueller                //
//=======================================//

package logger

import (
    "runtime"
    "reflect"
    "strconv"

    "appengine"
)

func Log(appEngineContext appengine.Context, message string, varsMap map[string](interface{})) {
	varsString := ""
	if varsMap != nil {
		for varName, value := range varsMap {
			varsString += varName + ": " + interfaceToString(value) + "\n"
		}
		appEngineContext.Infof(message + "\nVARS:\n" + varsString)
		return
	}
	
	appEngineContext.Infof(message)
}

func LogError(appEngineContext appengine.Context, message string, varsMap map[string](interface{})) {
	pointer, fileName, lineNum, _ := runtime.Caller(1)
 	funcName := runtime.FuncForPC(pointer).Name()
	placementTag := fileName + ": " + funcName + ": " + strconv.FormatInt(int64(lineNum), 10)
	varsString := ""
	if varsMap != nil {
		for varName, value := range varsMap {
			varsString += varName + ": " + interfaceToString(value) + "\n"
		}
		appEngineContext.Errorf(placementTag + "\n\t" + message + "\nVARS:\n" + varsString)
		return
	}
	appEngineContext.Errorf(placementTag + "\n\t" + message)
}

func interfaceToString(i interface{}) string {
	iValue := reflect.ValueOf(i)
	iType := reflect.TypeOf(i)
	kind := iValue.Kind()
	switch kind {
		case reflect.Invalid:
			return "INVALID"
		case reflect.Bool:
			return strconv.FormatBool(i.(bool))
		case reflect.Int:
			return strconv.FormatInt(int64(i.(int)), 10)
		case reflect.Int8:
			return strconv.FormatInt(int64(i.(int8)), 10)
		case reflect.Int16:
			return strconv.FormatInt(int64(i.(int16)), 10)
		case reflect.Int32:
			return strconv.FormatInt(int64(i.(int32)), 10)
		case reflect.Int64:
			return strconv.FormatInt(i.(int64), 10)
		case reflect.Uint:
			return strconv.FormatUint(uint64(i.(uint)), 10)
		case reflect.Uint8:
			return strconv.FormatUint(uint64(i.(uint8)), 10)
		case reflect.Uint16:
			return strconv.FormatUint(uint64(i.(uint16)), 10)
		case reflect.Uint32:
			return strconv.FormatUint(uint64(i.(uint32)), 10)
		case reflect.Uint64:
			return strconv.FormatUint(i.(uint64), 10)
		case reflect.Uintptr:
			return "UINT POINTER"
		case reflect.Float32:
			return strconv.FormatFloat(float64(i.(float32)), 'f', 10, 32)
		case reflect.Float64:
			return strconv.FormatFloat(i.(float64), 'f', 10, 64)
		case reflect.Complex64:
			return "COMPLEX64"
		case reflect.Complex128:
			return "COMPLEX128"
		case reflect.Array:
			return "ARRAY"
		case reflect.Chan:
			return "CHANNEL"
		case reflect.Func:
			return "FUNCTION"
		case reflect.Interface:
			return "INTERFACE"
		case reflect.Map:
			return "MAP"
		case reflect.Ptr:
			return "POINTER"
		case reflect.Slice:
			return "SLICE"
		case reflect.String:
			return i.(string)
		case reflect.Struct:
			numFields := iValue.NumField()
			structStr := "\n"
			for i := 0; i < numFields; i ++ {
				if iType.Field(i).Name != "ProfilePicBlobKey" && iType.Field(i).Name != "BackgroundPicBlobKey" {
					structStr += "\t" + iType.Field(i).Name + ": " + interfaceToString(iValue.Field(i).Interface()) + "\n"
				}
			}
			return structStr
		case reflect.UnsafePointer:
			return "UNSAFE POINTER"
	}

	return ""
}