package result

import "fmt"

func GetResult(ch chan map[string]float64, runTime int64)  {
	resultSlice := resultStruct{}
	for input := range ch {
		writeResultToSlice(input, &resultSlice)
	}

	selectCount := int64(len(resultSlice.SelectOptionSlice))
	insertCount := int64(len(resultSlice.InsertOptionSlice))
	updateCount := int64(len(resultSlice.UpdateOptionSlice))
	ops := (selectCount + insertCount + updateCount) * 1000 /runTime
	fmt.Println("ops is", ops)
}
