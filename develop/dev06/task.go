package main

import (
	"develop/dev6/cut"
	"fmt"
	"os"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель вместо TAB
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	options := cut.ParseArgs()
	err := cut.Run(os.Stdin, os.Stdout, options)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

}
