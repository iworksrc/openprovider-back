package tests

import (
	users "openprovider-back/go"
	"testing"
)

//var tribonacciEtalonList = []uint64{
//	0, 0, 1, 1, 2, 4, 7, 13, 24, 44, 81, 149, 274, 504, 927, 1705, 3136,
//	5768, 10609, 19513, 35890, 66012, 121415, 223317, 410744, 755476,
//	1389537, 2555757, 4700770, 8646064, 15902591, 29249425, 53798080,
//	98950096, 181997601, 334745777, 615693474, 1132436852, 2082876103,
//	3831006429,
//}

var tribonacciEtalonList = []string{
	"0", "0", "1", "1", "2", "4", "7", "13", "24", "44", "81", "149",
	"274", "504", "927", "1705", "3136", "5768", "10609", "19513", "35890",
	"66012", "121415", "223317", "410744", "755476", "1389537", "2555757",
	"4700770", "8646064", "15902591", "29249425", "53798080", "98950096",
	"181997601", "334745777", "615693474", "1132436852", "2082876103",
	"3831006429",
}

func TestTribonacciValues(t *testing.T) {

	for i := 0; i < len(tribonacciEtalonList); i++ {
		result:= users.TribonacciThroughCache(i, 100)

		if result != tribonacciEtalonList[i] {
			t.Error(
				"For", i,
				"expected", tribonacciEtalonList[i],
				"got", result,
			)
		}
	}
}
