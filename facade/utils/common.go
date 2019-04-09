package utils

// 判断字段串是否在slice中
func StringInSlice(target string, list []string) bool {
	for _, elem := range list {
		if elem == target {
			return true
		}
	}
	return false
}

// 滤掉空字符串
func RemoveEmptyStringInSlice(list []string) []string {
	var resultList []string
	for _, elem := range list {
		if elem != "" {
			resultList = append(resultList, elem)
		}
	}
	return resultList
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
