package mygomod

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	InferRunParamListPattern   = "^\\[.*\\]$"
	InferRunParamBoolPattern   = "(^[tT]rue$|^[fF]alse$)"
	InferRunParamNumberPattern = "^\\d+$"
	InferRunParamFloatPattern  = "^\\d+.\\d+$"
)

func SliceReflect(value any) {
	var paramValue string
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		var builder strings.Builder
		list := reflect.ValueOf(value)
		for idx := 0; idx < list.Len(); idx++ {
			if idx != 0 {
				builder.WriteString(",")
			}
			builder.WriteString(fmt.Sprintf("%v", list.Index(idx)))
		}
		paramValue = builder.String()
	default:
		paramValue = fmt.Sprintf("%v", value)
	}
	fmt.Println(paramValue)
}

func SliceReflectTest() {
	SliceReflect(12345)
	SliceReflect("hello world")
	SliceReflect([]int{1, 2, 3})
	SliceReflect([]string{"one", "two", "three"})
}

func convRunParamStrValue(str string) interface{} {
	matchTyp := matchRunParamStrPattern(str)
	if matchTyp != "list" {
		return parseElemByType(str, matchTyp)
	}
	str = strings.TrimPrefix(str, "[")
	str = strings.TrimSuffix(str, "]")
	strList := strings.Split(str, ",")
	var nMatchType string
	for _, nstr := range strList {
		nMatchType = matchRunParamStrPattern(nstr)
		if nMatchType == "float" {
			break
		}
	}
	if nMatchType == "list" {
		fmt.Println("nested list is not support, conv to string")
		nMatchType = "string"
	}
	nList := []interface{}{}
	for _, nstr := range strList {
		nList = append(nList, parseElemByType(nstr, nMatchType))
	}
	return nList
}

func parseElemByType(str, typ string) interface{} {
	switch typ {
	case "bool":
		switch str {
		case "true", "True":
			return true
		case "false", "False":
			return false
		}
		fmt.Println("convRunParamStrValue not bool")
	case "number":
		intVal, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Println("convRunParamStrValue strconv.ParseInt", err)
		}
		return intVal
	case "float":
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("convRunParamStrValue strconv.ParseFloat", err)
		}
		return floatVal
	default:
		// default is string
	}
	return str
}

func matchRunParamStrPattern(str string) string {
	patterns := map[string]string{
		"list":   InferRunParamListPattern,
		"bool":   InferRunParamBoolPattern,
		"number": InferRunParamNumberPattern,
		"float":  InferRunParamFloatPattern,
	}
	matchTyp := "string"
	for typ, pattern := range patterns {
		validator := regexp.MustCompile(pattern)
		if !validator.MatchString(str) {
			continue
		}
		matchTyp = typ
		break
	}
	fmt.Println("valid", matchTyp, "\tstr:", str)
	return matchTyp
}

func ConvertStrValueTest() {
	anyList := []string{
		// "[123.675,116.28,103.53]",
		// "[11,22,3]",
		// "[one,two_dup-v2.1,app-det_v1.0.yaml]",
		"33",
		"400",
		"400",
		"33.25",
		// "true",
		// "False",
		// "rgb",
		// "../class_names.txt",
		// "111,222,333",
		// "[aaaa true",
		"[]",
	}
	for _, nany := range anyList {
		fmt.Println(convRunParamStrValue(nany))
		fmt.Println("================================")
	}
}
