package xmap

import (
	"strings"
)

// GetMapValue 根据给定的路径（如"x.y.z"）递归地从map[string]interface{}中读取值。
// 如果值存在且可以转换为目标类型，则返回该值和true；否则，返回默认值和false。
// func GetMapValue(data map[string]interface{}, path string) Result {
// 	segments := strings.Split(path, ".")
// 	for i, segment := range segments {
// 		if val, exists := data[segment]; exists {
// 			if i == len(segments)-1 {
// 				return Result{Raw: val}
// 			}
// 			if nextMap, ok := val.(map[string]interface{}); ok {
// 				data = nextMap
// 			} else {
// 				return Result{}
// 			}
// 		} else {
// 			return Result{}
// 		}
// 	}
// 	return Result{}
// }

// getValue 根据给定的路径（如"x.y.z"）递归地从map[string]interface{}中读取值。
// 如果值存在且可以转换为目标类型，则返回该值和true；否则，返回默认值和false。
func GetMapValue(data map[string]any, path string) Result {
	// 尝试直接使用整个路径作为键进行匹配
	if val, exists := data[path]; exists {
		return Result{raw: val}
	}

	segments := strings.Split(path, ".")
	var currentPath string
	// fmt.Println("s", segments)
	for i, segment := range segments {
		// 生成从当前段到末尾的路径
		if i > 0 {
			currentPath += "."
		}
		currentPath += segment
		// fmt.Println("currentPath", currentPath)
		// 尝试直接匹配生成的路径作为键
		if val, exists := data[currentPath]; exists {
			if i == len(segments)-1 {
				return Result{raw: val}
			}
			// 如果不是最后一个段，并且值可以是map[string]interface{}，则对剩余的路径进行递归搜索
			if nextMap, ok := val.(map[string]interface{}); ok {
				restPath := strings.Join(segments[i+1:], ".")
				return GetMapValue(nextMap, restPath)
			}
			return Result{}
		}
		// fmt.Println("currentPath2", currentPath)
		// 如果当前路径段作为键未找到，但不是最后一段，则尝试按传统路径深入查找
		// if i < len(segments)-1 {
		// 	if nextMap, ok := data[segment].(map[string]interface{}); ok {
		// 		data = nextMap
		// 	} else {

		// 		return nil, false
		// 	}
		// }
	}

	// 处理传统路径最后一段的情况
	// if val, exists := data[segments[len(segments)-1]]; exists {
	// 	return val, true
	// }

	return Result{}
}

// func GetMapString(data map[string]interface{}, path string, defaultValue string) string {
// 	if val, ok := GetMapValue(data, path); ok {
// 		if stringVal, ok := val.(string); ok {
// 			return stringVal
// 		}
// 	}
// 	return defaultValue
// }

// func GetMapFloat64(data map[string]interface{}, path string, defaultValue float64) float64 {
// 	if val, ok := GetMapValue(data, path); ok {
// 		if stringVal, ok := val.(float64); ok {
// 			return stringVal
// 		}
// 	}
// 	return defaultValue
// }

// func GetMapInt(data map[string]interface{}, path string, defaultValue int) int {
// 	if val, ok := GetMapValue(data, path); ok {
// 		if stringVal, ok := val.(int); ok {
// 			return int(stringVal)
// 		}
// 	}
// 	return defaultValue
// }

// func updateResult(result *map[string]any, file *fileInfo) {

// 	current := *result

// 	for i, field := range file.Field {
// 		if i == len(file.Field)-1 { // Last
// 			// if current == nil {
// 			// 	// tmp := make(map[string]any)
// 			// 	// *result = tmp
// 			// 	current = make(map[string]any)
// 			// }
// 			if file.Merge == MergeTypeOver {
// 				if file.Value == nil {
// 					delete(current, field)
// 				} else {
// 					current[field] = file.Value
// 				}
// 			} else if file.Merge == MergeTypeUpsert {
// 				if _, ok := current[field]; !ok {
// 					current[field] = make(map[string]any)
// 				}
// 				for k, v := range file.Value.(map[string]any) {
// 					if v == nil {
// 						delete(current[field].(map[string]any), k)
// 					} else {
// 						current[field].(map[string]any)[k] = v
// 					}
// 				}
// 			}
// 		} else {
// 			if _, ok := current[field]; !ok {
// 				current[field] = make(map[string]any)
// 			}
// 			current = current[field].(map[string]any)
// 		}
// 	}
// 	if len(file.Field) == 0 { // Root directory operation
// 		if file.Merge == MergeTypeOver {
// 			if file.Value == nil {
// 				*result = make(map[string]any)
// 			} else {
// 				*result = file.Value.(map[string]any)
// 			}
// 		} else if file.Merge == MergeTypeUpsert {
// 			// if _, ok := (*result).(map[string]any); !ok {
// 			for k, v := range file.Value.(map[string]any) {
// 				if v == nil {
// 					delete((*result), k)
// 				} else {
// 					(*result)[k] = v
// 				}
// 			}
// 			// }
// 		} else {
// 			panic("unknow merge")
// 		}
// 	}
// }

// func MergaMap(result *map[string]any, fileds []string, value any, mergeType MergeType) error {
// 	current := *result
// 	if mergeType == MergeTypeUpsert {
// 		_, found := value.(map[string]any)
// 		if !found {
// 			return fmt.Errorf("value is not a map, but mergeType is Upsert")
// 		}
// 	}
// 	for i, field := range fileds {
// 		if i == len(fileds)-1 { // Last
// 			if mergeType == MergeTypeOver {
// 				if value == nil {
// 					delete(current, field)
// 				} else {
// 					current[field] = value
// 				}
// 			} else if mergeType == MergeTypeUpsert {
// 				if _, ok := current[field]; !ok {
// 					current[field] = make(map[string]any)
// 				}
// 				for k, v := range value.(map[string]any) {
// 					if v == nil {
// 						delete(current[field].(map[string]any), k)
// 					} else {
// 						current[field].(map[string]any)[k] = v
// 					}
// 				}
// 			}
// 		} else {
// 			if _, ok := current[field]; !ok {
// 				current[field] = make(map[string]any)
// 			}
// 			current = current[field].(map[string]any)
// 		}
// 	}
// 	if len(fileds) == 0 { // Root directory operation
// 		if mergeType == MergeTypeOver {
// 			if value == nil {
// 				*result = make(map[string]any)
// 			} else {
// 				*result = value.(map[string]any)
// 			}
// 		} else if mergeType == MergeTypeUpsert {
// 			for k, v := range value.(map[string]any) {
// 				if v == nil {
// 					delete((*result), k)
// 				} else {
// 					(*result)[k] = v
// 				}
// 			}
// 		} else {
// 			return fmt.Errorf("unknow mergeType: %v", mergeType)
// 		}
// 	}
// 	return nil
// }
