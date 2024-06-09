package validators

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func InitValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("IsDateFormat", IsDate)
	v.RegisterValidation("IsDateTimeZFormat", IsDateTimeZ)
	v.RegisterValidation("GtDateAddDay", GtDate)
	v.RegisterValidation("json", IsJson)
	v.RegisterValidation("IsInt", IsInt)

	return v
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func IsJson(fl validator.FieldLevel) bool {
	s := fl.Field().Interface()
	json, err := json.Marshal(s)
	if err != nil {
		return false
	}
	return string(json) != "{}"
}

func IsDateTimeZ(fl validator.FieldLevel) bool {
	if _, err := time.Parse("2006-01-02 15:04:05 -07:00", fl.Field().String()); err != nil {
		return false
	}

	return true
}
func IsInt(fl validator.FieldLevel) bool {
	strQtyOrder := strconv.Itoa(int(fl.Field().Int()))
	return strings.HasPrefix(strQtyOrder, "0")
}

func IsDate(fl validator.FieldLevel) bool {
	if _, err := time.Parse("2006-01-02", fl.Field().String()); err != nil {
		return false
	}

	return true
}

func GtDate(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	param := strings.Split(fl.Param(), `:`)
	paramField := param[0]
	paramValue := param[1]

	if paramField == `` {
		return true
	}

	var paramFieldValue reflect.Value

	if fl.Parent().Kind() == reflect.Ptr {
		paramFieldValue = fl.Parent().Elem().FieldByName(paramField)
	} else {
		paramFieldValue = fl.Parent().FieldByName(paramField)
	}

	date1, err := time.Parse("2006-01-02", value)
	if err != nil {
		return false
	}

	date2, err := time.Parse("2006-01-02", paramFieldValue.String())
	if err != nil {
		return false
	}

	addDate, err := strconv.Atoi(paramValue)
	if err != nil {
		addDate = 0
	}
	date2 = date2.AddDate(0, 0, addDate)

	return date1.After(date2)
}

func OnlyAcceptNumber(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	_, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return true
}

func LengthBetween(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	param := strings.Split(fl.Param(), `:`)
	paramMin := param[0]
	paramMax := param[1]

	if paramMin == `` && paramMax == `` {
		return true
	}

	min, err := strconv.Atoi(paramMin)
	if err != nil {
		min = 0
	}

	max, err := strconv.Atoi(paramMax)
	if err != nil {
		max = 0
	}

	return len(value) >= min && len(value) <= max
}
