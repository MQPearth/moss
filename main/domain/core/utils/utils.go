package utils

import (
	"moss/domain/core/aggregate"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// SortByIds 根据ids排序
func SortByIds[M aggregate.EntityInterface](items []M, ids []int) (res []M) {
	if len(items) < 2 || len(ids) < 2 {
		return items
	}
	m := SliceToMap(items, "ID")
	for _, id := range ids {
		item, ok := m[id]
		if !ok {
			continue
		}
		res = append(res, item)
	}
	return
}

// SliceToMap 切片转map
func SliceToMap[M aggregate.EntityInterface](s []M, field string) map[any]M {
	var res = make(map[any]M)
	for _, v := range s {
		id := int(reflect.ValueOf(v).FieldByName(field).Int())
		res[id] = v
	}
	return res
}

func TodayBeginTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func ValidateOrderInput(val any) any {
	str, ok := val.(string)
	if !ok || str == "" {
		return val
	}
	str = strings.ToLower(strings.TrimSpace(str))
	if str == "rand()" || str == "random()" {
		return str
	}
	if regexp.MustCompile(`^([a-z0-9_]+(?:\s+(asc|desc))?\s*,\s*)*[a-z0-9_]+(?:\s+(asc|desc))?$`).MatchString(str) {
		return str
	}
	if regexp.MustCompile(`^(count|sum|max|min|avg|length|year)\((\w+)\)(?:\s+(asc|desc))?$`).MatchString(str) {
		return str
	}
	return nil
}
