package mozuku

import (
	"fmt"
	"github.com/riita10069/mozuku/adapter"
	"github.com/riita10069/mozuku/domain"
)

func GetSimilarWords(str string, filename string) *domain.Dictionary {
	fileReader := adapter.NewFileAdapter(filename)
	fileReader.ReadByLine()
	dictionaryDomain := domain.NewDictionary().MakeBySlice(fileReader.Lines)
	searchedDictionary := dictionaryDomain.BinarySearch(str)
	fmt.Println(searchedDictionary)
	return searchedDictionary
}
