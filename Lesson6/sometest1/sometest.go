//1 Дополните код из раздела «Тестирование» функцией подсчета суммы переданных элементов и тестом для этой функции.

package sometest

//Тут, чтоб не переписывать, просто отправляем два значения
//Среднее и сумму

//AverageSumm function
func AverageSumm(xs []float64) (float64, float64) {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs)), total
}
