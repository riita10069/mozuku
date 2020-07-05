package main

import (
	"fmt"
	"github.com/Tech-Design-Inc/mozuku/adapter"
	"github.com/Tech-Design-Inc/mozuku/domain"
)

func GetSimilarWords(str string, filename string) *domain.Dictionary {
	fileReader := adapter.NewFileAdapter(filename)
	fileReader.ReadByLine()
	dictionaryDomain := domain.NewDictionary().MakeBySlice(fileReader.Lines)
	searchedDictionary := dictionaryDomain.BinarySearch(str)
	fmt.Println(searchedDictionary)
	return searchedDictionary
}

func main() {
	GetSimilarWords("あいどりんぐ", "dic.csv")
}
