package validplease

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type attribute struct {
	Name string
	Param string
}

type validError struct {
	Success bool
	Message string
	Filed string
}

var messagesFa = map[string]string {
	"ROLE_NOT_FOUND": "این متد پیدا نشد",
	"MAX_LEN_PARAM_ERROR": "مقدار پارامتر  maxLen صحیح نیست. لطفا یک مقدار عددی وارد نمایید",
	"MIN_LEN_PARAM_ERROR": "مقدار پارامتر  minLen صحیح نیست. لطفا یک مقدار عددی وارد نمایید",
	"VALUE_OF": "مقدار ",
	"LEN_OF": "طول ",
	"REQUIRED": " اجباری است",
	"CHARS": " کاراکتر است",
	"MORE": " بیشتر از ",
	"LOWER": " کمتر از ",
	"INCORRECT": " اشتباه است ",
	"CORRECT_VALUE": " مقادیر قابل قبول: ",
	"EXCEPT": " همه به جز:",
}

var messagesEn = map[string]string {
	"ROLE_NOT_FOUND": "this rule didn't found",
	"MAX_LEN_PARAM_ERROR": "invalid type for maxLen parameter. Please set a valid int",
	"MIN_LEN_PARAM_ERROR": "invalid type for minLen parameter. Please set a valid int",
	"VALUE_OF": "value of ",
	"LEN_OF": "length of ",
	"REQUIRED": " is required",
	"CHARS": " characters",
	"MORE": " is more than ",
	"LOWER": " is lower than ",
	"INCORRECT": " is incorrect ",
	"CORRECT_VALUE": " correct values are:",
	"EXCEPT": " any excludes:",
}

var messages = map[string]string {
}

var funcMap = map[string]interface{} {
	"len": _len,
	"maxLen": maxLen,
	"minLen":  minLen,
	"required": required,
	"in": isIn,
	"notIn": notIn,
	"email": email,
	"ip": ip,
	"ipv6": ipv6,
	"url": url,
}

func _len(filed string, value interface{}, indicator string) validError {
	i, err := strconv.Atoi(indicator)
	if err != nil {
		return  validError{ Success: false, Message: messages["MAX_LEN_PARAM_ERROR"],Filed: filed}
	}
	if len(fmt.Sprintf("%v", value)) > i {
		return  validError{ Success: false, Message: messages["LEN_OF"]+filed+messages["MORE"]+indicator+ messages["CHARS"],Filed: filed}
	} else if len(fmt.Sprintf("%v", value)) < i {
		return  validError{ Success: false, Message: messages["LEN_OF"]+filed+messages["LOWER"]+indicator+ messages["CHARS"],Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func maxLen(filed string, value interface{}, indicator string) validError {
	i, err := strconv.Atoi(indicator)
	if err != nil {
		return  validError{ Success: false, Message: messages["MAX_LEN_PARAM_ERROR"],Filed: filed}
	}
	if len(fmt.Sprintf("%v", value)) > i {
		return  validError{ Success: false, Message: messages["LEN_OF"]+filed+messages["MORE"]+indicator+ messages["CHARS"],Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func minLen(filed string, value interface{}, indicator string) validError {
	i, err := strconv.Atoi(indicator)
	if err != nil {
		return  validError{ Success: false, Message: messages["MIN_LEN_PARAM_ERROR"],Filed: filed}
	}
	if len(fmt.Sprintf("%v", value)) < i {
		return  validError{ Success: false, Message: messages["LEN_OF"]+filed+messages["LOWER"]+indicator+ messages["CHARS"],Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func required(filed string, value interface{}, _ string) validError {
	if len(fmt.Sprintf("%v", value)) == 0 {
		return  validError{ Success: false, Message: messages["VALUE_OF"]+filed+messages["REQUIRED"],Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func isIn(filed string, value interface{}, indicator string) validError {
	if fmt.Sprintf("%v", value) == "" {
		return validError{ Success: true, Message: "",Filed: filed}
	}

	values := strings.Split(indicator[:], "|")
	founded := false
	for _,v := range values {
		if v == fmt.Sprintf("%v", value) {
			founded = true
		}
	}
	if !founded {
		return  validError{ Success: false, Message: messages["VALUE_OF"]+filed+messages["INCORRECT"]+ messages["CORRECT_VALUE"]+indicator,Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func notIn(filed string, value interface{}, indicator string) validError {
	if fmt.Sprintf("%v", value) == "" {
		return validError{ Success: true, Message: "",Filed: filed}
	}

	values := strings.Split(indicator[:], "|")
	founded := true
	for _,v := range values {
		if v == fmt.Sprintf("%v", value) {
			founded = true
		}
	}
	if founded {
		return  validError{ Success: false, Message: messages["VALUE_OF"]+filed+messages["INCORRECT"]+ messages["CORRECT_VALUE"]+messages["EXCEPT"]+indicator,Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func email(filed string, value interface{}, _ string) validError {
	match, _ := regexp.MatchString("^(([^<>()\\[\\]\\\\.,;:\\s@\"]+(\\.[^<>()\\[\\]\\\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$", fmt.Sprintf("%v", value))
	if !match {
		return  validError{ Success: false, Message: messages["VALUE_OF"]+filed+messages["INCORRECT"],Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func url(filed string, value interface{}, _ string) validError {
	match, _ := regexp.MatchString("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)", fmt.Sprintf("%v", value))
	if !match {
		return  validError{ Success: false, Message: messages["VALUE_OF"]+filed+messages["INCORRECT"],Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func ip(filed string, value interface{}, _ string) validError {
	match, _ := regexp.MatchString("^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$", fmt.Sprintf("%v", value))
	if !match {
		return  validError{ Success: false, Message: messages["VALUE_OF"]+filed+messages["INCORRECT"],Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func ipv6(filed string, value interface{}, _ string) validError {
	match, _ := regexp.MatchString("^(?:(?:(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):){6})(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):(?:(?:[0-9a-fA-F]{1,4})))|(?:(?:(?:(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9]))\\.){3}(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9])))))))|(?:(?:::(?:(?:(?:[0-9a-fA-F]{1,4})):){5})(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):(?:(?:[0-9a-fA-F]{1,4})))|(?:(?:(?:(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9]))\\.){3}(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9])))))))|(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})))?::(?:(?:(?:[0-9a-fA-F]{1,4})):){4})(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):(?:(?:[0-9a-fA-F]{1,4})))|(?:(?:(?:(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9]))\\.){3}(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9])))))))|(?:(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):){0,1}(?:(?:[0-9a-fA-F]{1,4})))?::(?:(?:(?:[0-9a-fA-F]{1,4})):){3})(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):(?:(?:[0-9a-fA-F]{1,4})))|(?:(?:(?:(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9]))\\.){3}(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9])))))))|(?:(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):){0,2}(?:(?:[0-9a-fA-F]{1,4})))?::(?:(?:(?:[0-9a-fA-F]{1,4})):){2})(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):(?:(?:[0-9a-fA-F]{1,4})))|(?:(?:(?:(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9]))\\.){3}(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9])))))))|(?:(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):){0,3}(?:(?:[0-9a-fA-F]{1,4})))?::(?:(?:[0-9a-fA-F]{1,4})):)(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):(?:(?:[0-9a-fA-F]{1,4})))|(?:(?:(?:(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9]))\\.){3}(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9])))))))|(?:(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):){0,4}(?:(?:[0-9a-fA-F]{1,4})))?::)(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):(?:(?:[0-9a-fA-F]{1,4})))|(?:(?:(?:(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9]))\\.){3}(?:(?:25[0-5]|(?:[1-9]|1[0-9]|2[0-4])?[0-9])))))))|(?:(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):){0,5}(?:(?:[0-9a-fA-F]{1,4})))?::)(?:(?:[0-9a-fA-F]{1,4})))|(?:(?:(?:(?:(?:(?:[0-9a-fA-F]{1,4})):){0,6}(?:(?:[0-9a-fA-F]{1,4})))?::))))$", fmt.Sprintf("%v", value))
	if !match {
		return  validError{ Success: false, Message: messages["VALUE_OF"]+filed+messages["INCORRECT"],Filed: filed}
	} else {
		return validError{ Success: true, Message: "",Filed: filed}
	}
}

func parseAttribute(attr string) attribute {
	i := strings.Index(attr, "(")
	j := strings.Index(attr, ")")
	if i != -1 && j != -1 {
		name := attr[0:i]
		param := attr[i+1:j]
		name = strings.Replace(name,",","",1)
		return attribute{Name: name,Param:param}
	}
	return  attribute{Name: attr,Param: ""}
}

func ValidPlease(b interface{},locate string) []validError {
	if locate == "en" || locate == "" {
		messages = messagesEn
	} else if locate == "fa" {
		messages = messagesFa
	}

	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)

		jsonTag := ""
		if jsonTag = t.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
			var endOfAttr = 0
			if endOfAttr = strings.Index(jsonTag, ","); endOfAttr < 0 {
				endOfAttr = len(jsonTag)
			}
		}

		fieldName := t.Name
		if jsonTag != "" {
			fieldName = jsonTag
		}
		filedValue := val.Field(i)

		if vpTag := t.Tag.Get("vp"); vpTag != "" && vpTag != "-" {
			var endOfAttr = 0
			if endOfAttr = strings.Index(vpTag, "\""); endOfAttr < 0 {
				endOfAttr = len(vpTag)
			}

			attrs := strings.Split(vpTag[:endOfAttr], ",")
			for _, s := range attrs {
				parsed := parseAttribute(s)
				ve := call(parsed.Name, fieldName, filedValue, parsed.Param)
				if ve.Success == false {
					return []validError{ve}
				}
			}
		}
	}
	return []validError{}
}

func call(funcName string, filed string, value interface{}, param string) validError {
	f, found := funcMap[funcName]
	if found == true {
		err := f.(func(string,interface{},string) validError)(filed,value,param)
		return err
	}
	return validError{ Success: false, Message: messages["ROLE_NOT_FOUND"],Filed: funcName}
}