# mozuku
Mozuku returns a thesaurus of the received words. 

Mozuku is an OSS for natural language processing. 
It is useful when you want to compensate for similar words and notational distortions.

## Getting Started
### How to install

```
go get github.com/riita10069/mozuku
```

### Example Usage
```go
package main

import "fmt"
import "github.com/riita10069/mozuku"

func main() {
  d := mozuku.GetSimilarWords("海")
  fmt.Printf("-----%v", d)
}
```

Let's run above script!

```
ryotanoMacBook:~ ryota$ go run mozukunotest.go 
go: finding github.com/riita10069/mozuku latest
go: downloading github.com/riita10069/mozuku v0.0.0-20200910125341-3db416c94e6f
go: extracting github.com/riita10069/mozuku v0.0.0-20200910125341-3db416c94e6f

-----&{海 [Sea うみ 古海洋 海中 海洋 み わた]}
```

## GetSimilarWords
signature is here
`GetSimilarWords(str string) *domain.Dictionary`

str means word you want to know thesaurus.\
Dictionary is struct.
```go
Dictionary struct {
		Headword string
		Items    []string
}
```

Headword is str.
Items is thesaurus.
