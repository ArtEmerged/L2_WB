Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```

Интерфейс считается `nil`, только если его **тип** и **значение** оба равны `nil`. Переменная `err` имеет тип `error` (интерфейс). Интерфейс хранит информацию о типе данных (в данном случае о `*customError`). Значение переменной `err` не равна `nil`

Функция `test()` возвращает `nil`, что удовлетворяет интерфейсу `error`. В данном случае, тип возвращаемого значения функции `test()` - `*customError`, который является указателем на структуру `customError`, реализующую интерфейс `error`. Даже если внутренний указатель равен `nil`, переменная `err` не будет равна `nil`, потому что она содержит значение интерфейса `error` (в данном случае, указатель на `customError`). В результате условие `if err != nil` будет истинным, и программа выведет "error".