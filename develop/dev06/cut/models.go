package cut

import "errors"

const (
	//EmptyStart - default value для FromStartTo
	EmptyStart = 0
	//EmptyEnd - default value для FromToEnd
	EmptyEnd = 999999
)

var (
	// ErrNoSuchFile - не удалось прочитать файл
	ErrNoSuchFile = errors.New("cut: no such file or directory")
	// ErrInvalidRange - в флаг -f передан диапазон i-j где j < i
	ErrInvalidRange = errors.New("cut: invalid decreasing range")
	// ErrNumberLessOne - в флаг -f передано значение меньше 1
	ErrNumberLessOne = errors.New("cut: fields are numbered from 1")
	// ErrNotNumber - в флаг -f переданы не цифры
	ErrNotNumber = errors.New("cut: invalid field value")
	// ErrNotCharacter - в флаг -d передан не один символ
	ErrNotCharacter = errors.New("cut: the delimiter must be a single character")
)

// Fields - параметры для флага -f (fields)
type Fields struct {
	FromStartTo int
	FromToEnd   int
	FieldsList  []int
}
