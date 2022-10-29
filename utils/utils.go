package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

// ConvertStrSlice2Map 将字符串 slice 转为 map[string]struct{}。
func ConvertStrSlice2Map(sl []interface{}) map[string]struct{} {
	set := make(map[string]struct{}, len(sl))
	for _, v := range sl {
		set[v.(string)] = struct{}{}
	}
	return set
}

// 将字符串 slice 转为 map[string]struct{}。
func ConvertStrSlice1Map(sl []string) map[string]struct{} {
	set := make(map[string]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

// InMap 判断字符串是否在 map 中。
func InMap(m map[string]struct{}, s string) bool {
	_, ok := m[s]
	return ok
}

// 判断是否是子切片-string类型
func Subset(first, second []string) bool {
	set := make(map[string]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}
func SubsetLInt(first, second []int64) bool {
	set := make(map[int64]int64)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}
	return true
}

// DeleteStrSliceElms 删除切片指定元素（不许改原切片）。
func DeleteStrSliceElms(sl []string, elms ...string) []string {
	// 先将元素转为 set。
	m := make(map[string]struct{})
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 过滤掉指定元素。
	res := make([]string, 0, len(sl))
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}

// 数组切片去重

func RemoveRepeatedElement(s []interface{}) []string {
	result := make([]string, 0)
	m := make(map[string]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v.(string)]; !ok {
			result = append(result, v.(string))
			m[v.(string)] = true
		}
	}
	return result
}

func DisableEscapeHtml(data interface{}) (string, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return "", err
	}
	return bf.String(), nil
}

func DoubleRoundFive(val float64) float64 {
	formattedVal := fmt.Sprintf("%.2f", val)
	roundedVal, _ := strconv.ParseFloat(formattedVal, 64)
	return roundedVal
}

func TestEq(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func RemoveDuplicatesInPlace(userIDs []int64) []int64 {
	// 如果有0或1个元素，则返回切片本身。
	if len(userIDs) < 2 {
		return userIDs
	}

	//  使切片升序排序
	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })

	uniqPointer := 0

	for i := 1; i < len(userIDs); i++ {
		// 比较当前元素和唯一指针指向的元素
		//  如果它们不相同，则将项写入唯一指针的右侧。
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}

// 结构体切片排序

type Persons struct {
	Age    int
	Height int
}
