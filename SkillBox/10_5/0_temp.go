//Какой целочисленный тип можно сохранить без потери точности в float32 и какой в float64? Для проверки,
//напишите программу сохраняющую во float32 и float64 максимальные значения целочисленных типов со знаком
//и без знака и проверьте, преобразовав назад к целому числу, будет ли совпадать результат преобразования
//с максимумом или нет.
package main

import "fmt"

func main() {

	var a int16 = 25
	var a1 int32 = -290

	b := float32(a)
	fmt.Println(b)
	bf := float64(a)
	fmt.Println(bf)
	//----------------

	b1 := float32(a1)
	fmt.Println(b1)
	b1f := float64(a1)
	fmt.Println(b1f)
	//------------------

	c := int16(b)
	fmt.Println(c)

	c1 := int32(b1)
	fmt.Println(c1)

	cf := int32(b1f)
	fmt.Println(cf)

}
