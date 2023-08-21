package srv

import (
	"bytes"
	"fmt"
	"sort"
)

// Comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

// Commav2 inserts commas in a non-negative decimal integer string.
func Commav2(in string) string {
	byteA := []byte(in)
	n := len(byteA)
	if n <= 3 {
		return in
	}
	// bytes.Index()
	count := len(byteA)

	byB := new(bytes.Buffer)
	for i := 0; i < count; i++ {
		err := byB.WriteByte(byteA[i])
		if err != nil {
			panic(err)
		}
		// newBytes = append(newBytes, byteA[i])
		if (count-i)%3 == 1 && i != count-1 {
			err := byB.WriteByte(',')
			if err != nil {
				panic(err)
			}
		}
	}
	return byB.String()
}

// CompareStrHasSameWord compare two string, if they have same word, return true
func CompareStrHasSameWord(in1, in2 string) bool {
	if len(in1) == 0 || len(in2) == 0 {
		return true
	}
	str1 := []byte(in1)
	str2 := []byte(in2)
	sort.Slice(str1, func(i, j int) bool {
		if bytes.Compare([]byte{str1[i]}, []byte{str1[j]}) == -1 {
			return true
		}
		return false
	})
	sort.Slice(str2, func(i, j int) bool {
		if bytes.Compare([]byte{str2[i]}, []byte{str2[j]}) == -1 {
			return true
		}
		return false
	})

	if bytes.Equal(str1, str2) {
		return true
	}
	return false
}

// reverse reverses a slice of ints in place.
func reversev1(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

// Point represents a point in the Euclidean plane.
type Point struct {
	X, Y int
}

// ScaleBy scales p by the factor f.
func (p *Point) ScaleBy(factor int) {
	p.X *= factor
	p.Y *= factor
}

// Distance returns the distance between points p and q.
func (p Point) Distance(q Point) int {
	return 1
}
func test1() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance
	fmt.Println(distance(p, q)) // "5"
}

func De1() {
	ch := make(chan int, 1)
	for i := 1; i < 11; i++ {
		fmt.Println(i, "---", len(ch))
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
