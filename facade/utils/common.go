package utils

import "regexp"

// StringInSlice 判断字段串是否在slice中
func StringInSlice(target string, list []string) bool {
	for _, elem := range list {
		if elem == target {
			return true
		}
	}
	return false
}

// RemoveEmptyStringInSlice 从slice中滤掉空字符串
func RemoveEmptyStringInSlice(list []string) []string {
	var resultList []string
	for _, elem := range list {
		if elem != "" {
			resultList = append(resultList, elem)
		}
	}
	return resultList
}

// GetSafeFirst 从slice中获取第一个字符串，没有则返回空串
func GetSafeFirst(s []string) string {
	if len(s) > 0 {
		return s[0]
	}
	return ""
}

// MatchOneOf match one of the patterns
func MatchOneOf(text string, patterns ...string) []string {
	var (
		re    *regexp.Regexp
		value []string
	)
	for _, pattern := range patterns {
		// (?flags): set flags within current group; non-capturing
		// s: let . match \n (default false)
		// https://github.com/google/re2/wiki/Syntax
		re = regexp.MustCompile(pattern)
		value = re.FindStringSubmatch(text)
		if len(value) > 0 {
			return value
		}
	}
	return nil
}

// Domain get the domain of given URL
// Example: www.baidu.com -> baidu
func Domain(url string) string {
	domainPattern := `([a-z0-9][-a-z0-9]{0,62})\.` +
		`(com\.cn|com\.hk|` +
		`cn|com|net|edu|gov|biz|org|info|pro|name|xxx|xyz|be|` +
		`me|top|cc|tv|tt)`
	domain := MatchOneOf(url, domainPattern)
	if domain != nil {
		return domain[1]
	}
	return "Universal"
}

//// 从结构体根据名称获取slice字段的值
//func (e Extracted) GetFieldStringSlice(field string) []string {
//	r := reflect.ValueOf(e)
//	f := reflect.Indirect(r).FieldByName(field)
//	return f.Interface().([]string)
//}
//
//// 从结构体根据名称设置slice字段的值
//func (e *Extracted) SetFieldStringSlice(field string, setValue []string) {
//	value := reflect.ValueOf(setValue)
//
//	s := reflect.ValueOf(e).Elem()
//	if s.Kind() == reflect.Struct {
//		// exported field
//		f := s.FieldByName(field)
//		if f.IsValid() {
//			// A Value can be changed only if it is
//			// addressable and was not obtained by
//			// the use of unexported struct fields.
//			if f.CanSet() {
//				f.Set(value)
//			}
//		}
//	}
//}
