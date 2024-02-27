Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
3
2
4
5
6
8
7
0
0
0
0
...
```
#
В фунуциии `merge(a, b <-chan int) <-chan int`, мы будем читать данные из двух каналов `a` и `b` в блоке `select`. `select` будет выбирать тот `case`, котрый в данный момент не блокиет операцию чтения из канала. Если чтение из обоих канлаво являются неблокирующими, то мы будем заходить в случайный `case`.

После тогдо, как отправляющая горутина передаст все данные, она закроет канал. В функции `merge` мы будет принимать zero value (для `int` - 0) и отправлять их в канал `c`. Затем мы читаем данные из канала `c` до тех пор, пока канал `c` незакрыт. 

Этот процесс будет продолжаться до завершения выполнения программы.


#### Один из вариантов решения проблемы.
```go
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	closeA, closeB := true, true
	go func() {
		// закрываем канал `c` перед выходом из функции
		defer close(c)
		// Цикл будет выполняться пока один из каналов открыт
		for closeA || closeB {
			select {
			case v, ok := <-a:
				// передавим данные в канал `c` пока канал `a` открыт
				if ok {
					c <- v
				}
				// сигнализируем закрытие канала `a`
				closeA = ok
			case v, ok := <-b:
				// передавим данные в канал `c` пока канал `b` открыт
				if ok {
					c <- v
				}
				// сигнализируем закрытие канала `b`
				closeB = ok
			}
		}
	}()
	return c
}

```