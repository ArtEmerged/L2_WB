package main

import (
	"develop/dev5/grep"
	"fmt"
	"os"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	options, err := grep.ParseArgs()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	err = grep.Run(os.Stdin, os.Stdout, options)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

}
