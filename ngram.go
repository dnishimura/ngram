package ngram

import (
	"math"
	"strings"
)

type NGramError struct {
	err string
}

func (e *NGramError) Error() string {
	return e.err
}

type NGram struct {
	text  string
	n     int
	table map[string]int
}

func (ng *NGram) ParseText(text string) error {
	chars := "   "
	for _, c := range text {
		chars = strings.Join([]string{chars[1:], string(c)}, "")
		ng.table[chars] += 1
	}
	return nil
}

func (ng *NGram) CalcLength() float64 {
	length := 0
	for _, v := range ng.table {
		length += v * v
	}
	return math.Sqrt(float64(length))
}

func (ng *NGram) VectorDist(other *NGram) (float64, error) {
	if ng.n != other.n {
		return 0, &NGramError{"Must use same size NGrams"}
	}
	total := 0
	for k, v := range ng.table {
		total += v * other.table[k]
	}
	return 1.0 - float64(total)/(ng.CalcLength()*other.CalcLength()), nil
}

type LangDist struct {
	lang     string
	distance float64
}

func (ng *NGram) BestMatch(langs []*NGram) (string, error) {
	min := math.MaxFloat64
	count := len(langs)
	messages := make(chan LangDist, count)

	for _, lang := range langs {
		go func(ngram *NGram) {
			dist, err := ng.VectorDist(ngram)
			if err != nil {
				messages <- LangDist{ngram.text, dist}
			} else {
				messages <- LangDist{ngram.text, math.MaxFloat64}
			}
		}(lang)
	}

	best := ""
	for i := 0; i < count; i++ {
		message, ok := <-messages
		if ok {
			if message.distance < min {
				min = message.distance
				best = message.lang
			}
		} else {
			break
		}
	}

	return best, nil
}
