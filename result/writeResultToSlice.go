package result

import (
	"fmt"
	"strings"
)

func writeResultToSlice(dict map[string]float64, resultSlice *resultStruct) {
	for operator, runTime := range dict {
		if strings.HasPrefix(operator, "select") {
			// 说明是select操作
			resultSlice.SelectOptionSlice = append(resultSlice.SelectOptionSlice, operator)
			resultSlice.SelectRunTimeSlice = append(resultSlice.SelectRunTimeSlice, runTime)
		} else if strings.HasPrefix(operator, "insert") {
			// 说明是insert操作
			resultSlice.InsertOptionSlice = append(resultSlice.InsertOptionSlice, operator)
			resultSlice.InsertRunTimeSlice = append(resultSlice.InsertRunTimeSlice, runTime)
		} else if strings.HasPrefix(operator, "update") {
			// 说明是update操作
			resultSlice.UpdateOptionSlice = append(resultSlice.UpdateOptionSlice, operator)
			resultSlice.UpdateRunTimeSlice = append(resultSlice.UpdateRunTimeSlice, runTime)
		} else {
			// 说明不符合压测类型
			fmt.Println("操作不符合测试类型")
		}
	}

}
