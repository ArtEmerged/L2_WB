package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Цепочка обязанностей — это поведенческий паттерн, позволяющий передавать запрос по цепочке потенциальных обработчиков,
	пока один из них не обработает запрос.

	Шаблон "цепочка обязанностей" полезен:
		1. Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно,
		какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
		2. Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
		3. Когда набор объектов, способных обработать запрос, должен задаваться динамически.

	+ Уменьшает зависимость между клиентом и обработчиками.
	+ Реализует принцип единственной обязанности.
	+ Реализует принцип открытость/закрытость.

	- Запрос может остаться никем не обработанным.
*/

import "fmt"

type DATA_TYPE int

const (
	DATA DATA_TYPE = iota
	JAVASCRIPT
	CSS
)

type Data struct {
	dataType DATA_TYPE
	key      string
	value    string
}

func NewData(dataType DATA_TYPE, key, value string) *Data {
	return &Data{
		dataType: dataType,
		value:    value,
		key:      key,
	}
}

func (data *Data) GetKey() string {
	return data.key
}

func (data *Data) GetValue() string {
	return data.value
}

func (data *Data) GetDataType() DATA_TYPE {
	return data.dataType
}

type CacheHandler interface {
	HandleRequest(data Data)
}

type CdnCacheHandler struct {
	nextCacheHandler CacheHandler
}

func NewCdnCacheHandler(nextCacheHandler CacheHandler) *CdnCacheHandler {
	return &CdnCacheHandler{nextCacheHandler: nextCacheHandler}
}

func (cdnCacheHandler *CdnCacheHandler) HandleRequest(data Data) {
	if data.GetDataType() == CSS || data.GetDataType() == JAVASCRIPT {
		fmt.Printf("Caching file '%v' in CDN\n", data.GetKey())
		return
	}
	if cdnCacheHandler.nextCacheHandler != nil {
		cdnCacheHandler.nextCacheHandler.HandleRequest(data)
	}
}

type RedisCacheHandler struct {
	nextCacheHandler CacheHandler
}

func NewRedisCacheHandler(nextCacheHandler CacheHandler) *RedisCacheHandler {
	return &RedisCacheHandler{nextCacheHandler: nextCacheHandler}
}

func (redisCacheHandler RedisCacheHandler) HandleRequest(data Data) {
	if data.GetDataType() == DATA && len(data.GetValue()) <= 1024 {
		fmt.Printf("Caching data '%v' in Redis\n", data.GetKey())
		return
	}
	if redisCacheHandler.nextCacheHandler != nil {
		redisCacheHandler.nextCacheHandler.HandleRequest(data)
	}
}

type DiskCacheHandler struct {
	nextCacheHandler CacheHandler
}

func NewDiskCacheHandler(nextCacheHandler CacheHandler) *DiskCacheHandler {
	return &DiskCacheHandler{nextCacheHandler: nextCacheHandler}
}

func (diskCacheHandler *DiskCacheHandler) HandleRequest(data Data) {
	if data.GetDataType() == DATA && len(data.GetValue()) > 1024 {
		fmt.Printf("Caching data '%v' in Disk\n", data.GetKey())
		return
	}
	if diskCacheHandler.nextCacheHandler != nil {
		diskCacheHandler.nextCacheHandler.HandleRequest(data)
	}
}

func main() {
	cacheHandler := NewDiskCacheHandler(NewRedisCacheHandler(NewCdnCacheHandler(nil)))

	data := NewData(DATA, "key1", "ABC320489un3429rn29urn29r82n9jfdn2")

	cacheHandler.HandleRequest(*data)

	data = NewData(CSS, "key2", ".some-class{border: 1px solid red; margin: 10px}")

	cacheHandler.HandleRequest(*data)
}
