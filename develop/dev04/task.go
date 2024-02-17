package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	// ErrLetter - ошибка некорректного ввода буквы
	ErrLetter = errors.New("anagram: incorrect letter entry")
)

// sortSlice - сортирует срезы строк в map
func sortSlice(anagramSets map[string][]string) {
	for _, words := range anagramSets {
		slices.Sort(words)
	}
}

// generatesAnagram - генерирует уникальный ключ для анаграмм
func generatesAnagram(word string) ([33]uint8, error) {
	// Создаем массив из 33 элементов, у каждой буквы свой индекс
	anagram := [33]uint8{}
	for _, letter := range word {
		// Русский алфавит начинается с 1072 руны - буква 'а'
		// Определяем позицию буквы в массиве
		index := int(letter - 'а')
		// Проверка на русские буквы нижнего регистра
		if index < 0 || index > 32 {
			return anagram, ErrLetter
		}
		anagram[index]++
	}
	return anagram, nil
}

func main() {
	arr := []string{"пятак", "тяпка", "пятка", "пудра", "пятка", "листок", "мох", "ГоРа","рОга", "слиток", "столик"}
	res, err := searchForAnagramSets(&arr)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(res)
}

// searchForAnagramSets - выполняет поиск всех множеств анаграмм по словарю.
// Функция возвращает отсортированное множество анаграмм  map[string][]string.
// Если в функцию передан словарь не с русскими буквами, то функция возвращает ошибку ErrLetter.
func searchForAnagramSets(words *[]string) (*map[string][]string, error) {
	// Объявляем map result для хранения отсортированных множеств анаграмм
	result := make(map[string][]string)
	// Объявляем map для хранения анаграмм
	anagrams := make(map[[33]uint8][]string)
	// Объявляем set для уникальных слов
	set := make(map[string]struct{})

	for _, word := range *words {
		// Переводим слово в нижний регистр
		word = strings.ToLower(word)

		if _, ok := set[word]; ok {
			continue
		}
		set[word] = struct{}{}

		// Генерируем уникальный массив для хранения анаграмм
		anagramWord, err := generatesAnagram(word)
		// Возвращаем ошибку, если буква не соответствует требованиям
		if err != nil {
			return nil, err
		}

		// Добавляем слово в массив анаграмм
		anagrams[anagramWord] = append(anagrams[anagramWord], word)

		// Если в массиве анаграмм больше одного слова, то добавляем его во множество анаграмм
		if len(anagrams[anagramWord]) > 1 {

			// Присваиваем переменной firstWord первое слово из массива анаграмм
			firstWord := anagrams[anagramWord][0]

			// Если во множестве анаграмм еще нет конкретной анаграммы, то добавляем первое слово
			if _, ok := result[firstWord]; !ok {
				// Добавляем первое слово во множество анаграмм
				result[firstWord] = append(result[firstWord], firstWord)
			}

			// Присваиваем переменной lastWord последнее слово из массива анаграмм
			lastWord := anagrams[anagramWord][len(anagrams[anagramWord])-1]

			// Добавляем последнее слово во множество анаграмм
			result[firstWord] = append(result[firstWord], lastWord)
		}
	}
	sortSlice(result)

	return &result, nil
}
