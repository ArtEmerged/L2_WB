Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
[3 2 3]
```

#### Слайс в Go представляет собой структуру с 3 полями:

-  Указатель на первый элемент массива.
- Длина слайса: количество элементов в слайсе, которые доступны для использования.
- Емкость: количество элементов в базовом массиве.


При передаче слайса в качестве аргумента функции в Go происходит передача копии самого слайса (структуры).

Если принимающая функция выполняет операции с элементами слайса, которые не изменяют поля слайса, то изменения затронут оригинальный слайс. Однако, если происходят изменения в полях структуры, то они не отразятся на оригинальном слайсе.

В функции `modifySlice(i []string)` происходят следующие операции:

1. `i[0] = "3"` - присваивание нового значения под индексом 0 (отразится на оригинальном слайсе).
2. `i = append(i, "4")` - добавление нового элемента в слайс через функцию `append()`. Длина и емкость слайса были равны  После добавления нового элемента функция `append()` создает новый базовый массив для копии нашего слайса, его длина будет равна 4, а емкость - 6. (новый слайс уже не повлияет на оригинальный).
3. `i[1] = "5"` - присваивание нового значения под индексом 1 в новом слайсе. (не отразится на оригинальном слайсе).
4. `i = append(i, "6")` - добавление нового элемента в новый слайс. Его длина будет равна 5, а емкость - 6. (не отразится на оригинальном слайсе).