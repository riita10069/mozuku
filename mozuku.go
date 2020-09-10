package mozuku

import (
	"github.com/riita10069/mozuku/adapter"
	"github.com/riita10069/mozuku/domain"
)

func GetSimilarWords(str string) *domain.Dictionary {
	fileReader := adapter.NewFileAdapter()
	fileReader.ReadByLine()
	dictionaryDomain := domain.NewDictionary().MakeBySlice(fileReader.Lines)
	searchedDictionary := dictionaryDomain.BinarySearch(str)
	return searchedDictionary
}
