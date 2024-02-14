package main

import (
	"errors"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// ErrInvalidString ошибка для невалидного запроса
var ErrInvalidString = errors.New("invalid string")

// UnpackString выполняет примитивную распаковку строки
func UnpackString(s string) (string, error) {
	sRune := []rune(s)

	res := make([]rune, 0, len(s))

	var letterFlag bool

	for i, r := range sRune {
		// Программа работает только с числами от 0 до 9.
		// Если letterFlag != true, значит до этого была цифра, а это > 9.
		// Ещё условие сработает если первая руна в строке цифра.
		if unicode.IsDigit(r) && !letterFlag {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(r) {
			// Переводим цифру из rune в int
			digit := int(r - '0')
			// Если пришел 0, то нужно удалить предедущую руну
			if digit == 0 {
				res = res[:len(res)-1]
			}
			// Добавляем руны через цикл к результату
			for j := 1; j < digit; j++ {
				res = append(res, sRune[i-1])
			}
			// Меняем letterFlag на false, делаем пометку, что сейчас была цифра (нужно для проверки чисел которые больше 9)
			letterFlag = false
			continue
		}
		// Добавляем все символы кроме цифр
		res = append(res, r)
		// Меняем letterFlag на true чтобы не зайти в первый if
		letterFlag = true
	}
	return string(res), nil
}
