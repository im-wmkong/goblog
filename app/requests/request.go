package requests

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"goblog/pkg/model"
	"goblog/pkg/types"
	"strings"
	"unicode/utf8"
)

func init() {
	addNotExists()
	addMaxCn()
	addMinCn()
}

func addNotExists()  {
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		tableName := rng[0]
		doFiled := rng[1]
		val := value.(string)

		var count int64
		model.DB.Table(tableName).Where(doFiled+" = ?", val).Count(&count)
		if count != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已被占用", val)

		}
		return nil
	})
}

func addMaxCn()  {
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLen := utf8.RuneCountInString(value.(string))
		maxLen := types.StringToInt(strings.TrimPrefix(rule, "max_cn:"))
		if valLen > maxLen {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", maxLen)
		}
		return nil
	})
}

func addMinCn()  {
	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLen := utf8.RuneCountInString(value.(string))
		minLen := types.StringToInt(strings.TrimPrefix(rule, "min_cn:"))
		if valLen < minLen {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字", minLen)
		}
		return nil
	})
}
