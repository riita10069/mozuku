package domain

import (
	"strings"
	)

type (
	Dictionary struct {
		Headword string
		Items    []string
	}

	Dictionaries []Dictionary
)

func NewDictionary() *Dictionary {
	return &Dictionary{}
}

func NewDictionaries() *Dictionaries {
	return &Dictionaries{}
}

func (this *Dictionary) MakeBySlice(lists []string) *Dictionaries {
	dictionaries := NewDictionaries()
	dictionary := NewDictionary()
	for _, v := range lists {
		dictionary.Headword = strings.Split(v, ",")[0]
		dictionary.Items = strings.Split(v, ",")[1:]
		*dictionaries = append(*dictionaries, *dictionary)
	}
	return dictionaries
}

func (these *Dictionaries) BinarySearch(str string) *Dictionary {
	left := 0; right := len(*these) - 1
	
	for left <= right{
		median := (left + right) / 2

		if (*these)[median].Headword == str {
			return &(*these)[median]
		}
		if (*these)[median].Headword < str {
			left = median + 1
		}else{
			right = median - 1
		}
	}

	if left == len(*these) || (*these)[left].Headword != str {
		return nil
	}

	return nil
}
