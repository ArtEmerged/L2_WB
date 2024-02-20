package grep

import (
	"errors"
	"flag"
)

var (
	// ErrArgs - аргументы переданы некорректно
	ErrArgs = errors.New("usage: grep [OPTION]... PATTERNS [FILE]... ")
	// ErrNoSuchFile - не удалось прочитать файл
	ErrNoSuchFile = errors.New("grep: no such file or directory")
)

// Options - структура для хранения опций
type Options struct {
	Pattern    string   // pattern для поиска
	FilesName  []string // имена файлов для поиска
	After      int      // -A: Кол-во строк для вывода после найденного совпадения +N
	Before     int      // -B: Кол-во строк для вывода перед найденным совпадением -N
	Context    int      // -C: Кол-во строк контекста для вывода ±N
	Count      bool     // -c: Выводить только кол-во совпадений
	IgnoreCase bool     // -i: Игнорировать регистр при поиске
	Invert     bool     // -v: Инвертировать результат поиска
	Fixed      bool     // -F: Использовать фиксированные строки вместо регулярных выражений
	LineNum    bool     // -n: Добавить номера строк к результатам
}

// NewOptions - конструктор для Options
func NewOptions(pattern string, filesName []string, after, before, context int, count, ignoreCase, invert, fixed, lineNum bool) *Options {
	return &Options{
		Pattern:    pattern,
		FilesName:  filesName,
		After:      after,
		Before:     before,
		Context:    context,
		Count:      count,
		IgnoreCase: ignoreCase,
		Invert:     invert,
		Fixed:      fixed,
		LineNum:    lineNum,
	}
}

// ParseArgs - функция для парсинга флагов и аргументов командной строки
func ParseArgs() (*Options, error) {
	// Определение флагов
	after := flag.Int("A", 0, "-A <NUM>. Print <NUM> lines of trailing context after matching lines.")
	before := flag.Int("B", 0, "-B <NUM>. Print <NUM> lines of trailing context befor matching lines.")
	context := flag.Int("C", 0, "-C <NUM>. Print <NUM> lines of output context.")
	count := flag.Bool("c", false, "Suppress normal output; instead print a count of matching lines for each input file.")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions in patterns and input data, so that characters that differ only in case match each other.")
	invert := flag.Bool("v", false, "Invert the sense of matching, to select non-matching lines.")
	fixed := flag.Bool("F", false, "Interpret PATTERNS as fixed strings, not regular expressions.")
	lineNum := flag.Bool("n", false, "Prefix each line of output with the 1-based line number within its input file.")

	// Парсинг флагов
	flag.Parse()

	args := flag.Args()

	// Возвращаем ошибку, если шаблон не передан
	if len(args) < 1 {
		return nil, ErrArgs
	}
	pattern := args[0]

	filesName := args[1:]

	return NewOptions(pattern, filesName, *after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum), nil
}
