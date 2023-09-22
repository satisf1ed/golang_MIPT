package main

import "fmt"

const hashConstant int = 31
const hashMod int = 1000033

// алгоритм хеширования строки
func hashString(s string) int {
	var hash int = 0
	var power int = 1
	for i := 0; i < len(s); i++ {
		hash += int(s[i]) * power % hashMod
		power *= hashConstant % hashMod
	}

	return hash
}

type library interface {
	generateId(Book) int
	putBook(Book)
	getBookByName(string) interface{}
}
type storage interface {
	getBookById(int) interface{}
	putBook(int, Book) string
}

type Book struct {
	author  string
	name    string
	year    int
	isTaken bool
}

func (book *Book) identifyAuthor() string {
	return book.author
}
func (book *Book) identifyName() string {
	return book.name
}
func (book *Book) identifyYear() int {
	return book.year
}
func (book *Book) isBookTaken() bool {
	return book.isTaken
}

type Library struct {
	bookStorage storage
}

func (lib *Library) generateId(book Book) int {
	return hashString(book.identifyName())
}
func (lib *Library) putBook(book Book) {
	hash := lib.generateId(book)
	lib.bookStorage.putBook(hash, book)
}
func (lib *Library) getBookByName(name string) interface{} {
	hash := hashString(name)
	return lib.bookStorage.getBookById(hash)
}

type StorageMap struct {
	books map[int]Book
}

func (storage *StorageMap) getBookById(id int) interface{} {
	if value, ok := storage.books[id]; ok {
		return value
	}

	return "No such book in storage" // Если запрошенной книги нет в базе, сообщаем об этом
}
func (storage *StorageMap) putBook(hash int, book Book) string {
	if value, ok := storage.books[hash]; ok {
		value.isTaken = false // Если книга присутствовала в базе, мы просто помечаем, что ее вернули
	} else {
		storage.books[hash] = book // Иначе добавляем в базу новую книгу
	}
	return "ok"
}

type StorageSlice struct {
	books []Book
	id    map[int]int
	index int
}

func (storage *StorageSlice) getBookById(hash int) interface{} {
	if value, ok := storage.id[hash]; ok {
		return storage.books[value]
	}

	return "No such book in storage"
}
func (storage *StorageSlice) putBook(hash int, book Book) string {
	if value, ok := storage.id[hash]; ok {
		storage.books[value].isTaken = false
	} else {
		storage.id[hash] = storage.index
		storage.books[storage.index] = book
		storage.index += 1
	}
	return "ok"
}

func main() {
	var storMap storage = &StorageMap{make(map[int]Book)} // инициализация storage на мапке
	var libMap library = &Library{storMap}                // инициализация библиотеки на storageMap

	book1 := Book{"Privet", "Omlet", 2020, true} // создадим книгу
	libMap.putBook(book1)                        // добавляем книгу в библу

	book2 := Book{"aaaa", "Маленький принц", 0, false}
	book3 := Book{"HoMM III", "Fortress", 1999, false}
	book4 := Book{"meeh", "meeh", 2004, true}
	book5 := Book{"Наум Яковлевич Виленкин", "Комбинаторика", 2018, true}
	libMap.putBook(book2)
	libMap.putBook(book3)
	libMap.putBook(book4)
	libMap.putBook(book5)

	fmt.Println(libMap.getBookByName("Omlet"))                         // {Privet Omlet 2020 true}
	fmt.Println(libMap.getBookByName("выаооывраоыаываорыворлаорлывф")) // No such book in storage
	fmt.Println(libMap.getBookByName("Fortress"))                      // {HoMM III Fortress 1999 false}

	fmt.Println("\n")

	var storSlice storage = &StorageSlice{make([]Book, 5), make(map[int]int), 0} // инициализация storage на slice
	var libSlice library = &Library{storSlice}                                   // инициализация библиотеки на storageSlice

	libSlice.putBook(book1)
	libSlice.putBook(book2)
	libSlice.putBook(book3)
	libSlice.putBook(book4)
	libSlice.putBook(book5)

	fmt.Println(libMap.getBookByName("Omlet"))                         // {Privet Omlet 2020 true}
	fmt.Println(libMap.getBookByName("выаооывраоыаываорыворлаорлывф")) // No such book in storage
	fmt.Println(libMap.getBookByName("Fortress"))                      // {HoMM III Fortress 1999 false}
}
