//1 Дополните код из раздела «Тестирование» функцией подсчета суммы переданных элементов и тестом для этой функции.
package sometest

import "testing"

type testpair struct {
	values  []float64
	average float64
	sumdig  float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 6},
	{[]float64{-1, 1}, 0, 0},
}

//TestAverageSummSet function
func TestAverageSummSet(t *testing.T) {
	for _, pair := range tests {
		v, s := AverageSumm(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
		if s != pair.sumdig {
			t.Error(
				"For", pair.values,
				"expected", pair.sumdig,
				"got", s,
			)
		}
	}
}

// go test
// PASS
// ok      _/mnt/c/Geekbrains/Golang/HW/Lesson6/sometest   0.047s
