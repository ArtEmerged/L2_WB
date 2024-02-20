package cut

import "flag"

// Options - структура для хранения опций
type Options struct {
	FilesName []string // имена файлов
	Fields    string   // -f
	Delimiter string   // -d
	Separated bool     // -s
}

// NewOptions - конструктор для Options
func NewOptions(filesName []string, fields, delimiter string, separated bool) *Options {
	return &Options{
		FilesName: filesName,
		Fields:    fields,
		Delimiter: delimiter,
		Separated: separated,
	}
}

// ParseArgs - функция для парсинга флагов и аргументов командной строки
func ParseArgs() *Options {
	// Определение флагов
	fields := flag.String("f", "", "select only these fields; also print any line that contains no delimiter character")
	delimiter := flag.String("d", "\t", "use DELIM instead of TAB for field delimiter")
	separated := flag.Bool("s", false, "do not print lines not containing delimiters")

	// Парсинг флагов
	flag.Parse()

	// Получение аргументов (имен файлов)
	filesName := flag.Args()

	// Создание и возврат структуры Options
	return NewOptions(filesName, *fields, *delimiter, *separated)
}
