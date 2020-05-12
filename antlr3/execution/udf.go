package execution

type UDF func(args... interface{}) (interface{}, error)

var udfMapping = make(map[string]UDF)

func GetUDF(functionName string) (UDF, bool) {
	udf, ok := udfMapping[functionName]
	return udf, ok
}
