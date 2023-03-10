package main

import "fmt"

func summArr(arr []int) int {
	lenght := len(arr)
	if lenght == 0 {
		return 0 //arr[lenght]
	} else {
		return arr[0] + summArr(arr[1:lenght])
	}
}

func main() {
	arra := []int{1, 5, 4, 6}
	fmt.Println(summArr(arra))

}


//описание
/*
Данный код представляет функцию с именем summArr, которая принимает на вход массив целых чисел arr и возвращает сумму всех элементов в этом массиве.

Как работает функция?

    Сначала функция определяет длину массива с помощью функции len(arr) и сохраняет ее в переменную lenght.

    Затем функция проверяет, равна ли длина массива нулю. Если это так, то функция возвращает 0, так как массив пустой. В противном случае, функция переходит к следующему шагу.

    Если длина массива не равна нулю, функция берет первый элемент массива (arr[0]) и складывает его со значением, возвращаемым функцией summArr(arr[1:lenght]).

    Что происходит в выражении summArr(arr[1:lenght])? Здесь мы передаем функции summArr подмассив массива arr, начиная со второго элемента (arr[1]) и заканчивая последним элементом (arr[lenght-1]).

    Функция summArr будет вызвана рекурсивно, пока не останется элементов в переданном ей подмассиве, и затем вернет сумму всех элементов этого подмассива.

    Этот процесс продолжится до тех пор, пока все элементы массива не будут просуммированы, и сумма будет возвращена в качестве результата работы функции.

Таким образом, мы можем увидеть, что функция summArr использует рекурсивный подход для вычисления суммы элементов в массиве, и это происходит путем деления массива на подмассивы и последующего их суммирования.
*/
