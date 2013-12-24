package ngram

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	english := NGram{"english", 3, make(map[string]int)}
	english.ParseText("hello my name is polly")
	spanish := NGram{"spanish", 3, make(map[string]int)}
	spanish.ParseText("hola mi amo polly")
	data := NGram{"data", 3, make(map[string]int)}
	data.ParseText("my name is stan")
	best, _ := data.BestMatch([]*NGram{&english, &spanish})
	fmt.Println(best)

	// Output:
	// english
}
