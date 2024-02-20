package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(4*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	// Объявляем single канал
	single := make(chan interface{})
	var wg sync.WaitGroup

	// Добавляем в WaitGroup 1
	wg.Add(1)

	// Циклом проходимся по всем каналам
	for _, ch := range channels {
		// Вызываем анонимную функцию в отдельной горутине и передаем в неё канал
		go func(ch <-chan interface{}) {
			// После закрытия одного из каналов, у нас получится прочитать из канала zero value
			// и перед выходом из функции выполнится отложенный вызов, который декрементирует счетчик wg на 1
			defer wg.Done()
			// Важно понимать, что WaitGroup нам нужен для синхронизации, мы не можем закрывать в этом месте канал single
			// Если мы выполним здесь close(single), то вызывающая функция, в нашем случае main(), может начать закрывать остальные каналы, 
			// наши анонимные горутины разблокируются и выполнят close(single). В результате мы получим panic 
			<-ch
		}(ch)
	}

	// Вызываем анонимную функцию в отдельной горутине и ожидаем закрытия одного из каналов,
	// Wait блокирует выполнение горутины, пока wg не равен нулю
	// перед выходом из функции выполнится отложенный вызов, который закрывает канал single
	go func() {
		defer close(single)
		wg.Wait()
	}()

	return single
}