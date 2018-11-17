package exercise9

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestRandomArrays(t *testing.T) {
	length := 1000
	maxint := 100
	arr1 := RandomArray(length, maxint)
	assert.Equal(t, length, len(arr1))
	for _, v := range arr1 {
		assert.True(t, 0 <= v, "contains a negative integer")
		assert.True(t, v < maxint, "contains an integer >=maxint")
	}

	// check that different calls return different results
	arr2 := RandomArray(length, maxint)
	assert.False(t, reflect.DeepEqual(arr1, arr2))
}

func TestRandomArrays2(t *testing.T) {
	length := 10
	maxint := 1000
	arr1 := RandomArray(length, maxint)
	assert.Equal(t, length, len(arr1))
	for _, v := range arr1 {
		assert.True(t, 0 <= v, "contains a negative integer")
		assert.True(t, v < maxint, "contains an integer >=maxint")
	}

	// check that different calls return different results
	arr2 := RandomArray(length, maxint)
	assert.False(t, reflect.DeepEqual(arr1, arr2))
}

func TestArrayStatParallel(t *testing.T) {
	length := 30000000
	maxint := 10000
	arr2 := RandomArray(length, maxint)

	// call MeanStddev single-threaded
	start := time.Now()
	mean2, stddev2 := MeanStddev(arr2, 1)
	end := time.Now()
	dur1 := end.Sub(start)

	// now turn on cuncurrency and make sure we get the same results, but faster
	start = time.Now()
	mean3, stddev3 := MeanStddev(arr2, 3)
	end = time.Now()
	dur2 := end.Sub(start)

	speedup := float64(dur1) / float64(dur2)
	assert.True(t, speedup > 1.25, "Running MeanStddev with concurrency didn't speed up as expected. Sped up by %g.", speedup)
	assert.Equal(t, float32(mean2), float32(mean3)) // compare as float32 to avoid rounding differences
	assert.Equal(t, float32(stddev2), float32(stddev3))
}

func TestArrayStats(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	mean1, stddev1 := MeanStddev(arr1, 3)
	assert.Equal(t, 3.5, mean1)
	assert.Equal(t, float32(1.70782512766), float32(stddev1)) // compare as float32 to avoid rounding differences

	length := 50000000
	maxint := 10000
	arr2 := RandomArray(length, maxint)
	mean2, stddev2 := MeanStddev(arr2, 1)

	// Check that mean and stddev are reasonably close to what the true values should be. There is a *small* probability
	// that the values will be outside this range just by chance.
	meanErr := mean2 - float64(maxint-1)/2.0
	stddevErr := stddev2 - float64(maxint)/math.Sqrt(12.0) // https://www.dummies.com/education/math/business-statistics/how-to-calculate-the-variance-and-standard-deviation-in-the-uniform-distribution/
	assert.True(t, math.Abs(meanErr) < 2.0*float64(length)/float64(maxint), "mean %v from the truth", meanErr)
	assert.True(t, math.Abs(stddevErr) < 2.0*float64(length)/float64(maxint), "stddev %v from the truth", stddevErr)
}

func TestPointScaling(t *testing.T) {
	p := Point{1, 3}
	p.Scale(3)
	assert.Equal(t, 3.0, p.x)
	assert.Equal(t, 9.0, p.y)
}

func TestPointScaling2(t *testing.T) {
	p := Point{1, 3}
	p.Scale(2)
	assert.Equal(t, 2.0, p.x)
	assert.Equal(t, 6.0, p.y)
}

func TestPointRotate(t *testing.T) {
	p := Point{1, 0}
	// compare as float32 to avoid rounding differences
	p.Rotate(math.Pi / 4)
	assert.Equal(t, float32(math.Sqrt(2)/2), float32(p.x))
	assert.Equal(t, float32(math.Sqrt(2)/2), float32(p.y))
	p.Rotate(math.Pi / 4)
	p.Rotate(math.Pi / 4)
	assert.Equal(t, float32(-math.Sqrt(2)/2), float32(p.x))
	assert.Equal(t, float32(math.Sqrt(2)/2), float32(p.y))
}

func TestSortingSmall(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr := []float64{5, 7.2, 3.1, 4, 9.1, 1, 7}
	QuickSort(arr)
	if !reflect.DeepEqual(arr, []float64{1, 3.1, 4, 5, 7, 7.2, 9.1}) {
		t.Error("Array not actually sorted", arr)
	}
}

// copied from https://golang.org/src/math/rand/rand.go?s=7456:7506#L225 for Go <1.10 compatibility
func shuffle(n int, swap func(i, j int)) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	if n < 0 {
		panic("invalid argument to shuffle")
	}

	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(r.Int31n(int32(i + 1)))
		swap(i, j)
	}
}

func TestSortingLarge(t *testing.T) {
	const length = 1000
	arr := make([]float64, length)
	for i := 0; i < length; i++ {
		arr[i] = float64(i)
	}
	shuffle(length, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	QuickSort(arr)
	for i := 0; i < length; i++ {
		if arr[i] != float64(i) {
			t.Error("Value not in correct sorted location", arr[i])
		}
	}
}

func benchmarkSorting(b *testing.B, sort func(arr []float64)) {
	const length = 1000
	arr := make([]float64, length)
	for i := 0; i < length; i++ {
		arr[i] = float64(i)
	}

	// run the benchmark
	for n := 0; n < b.N; n++ {
		rand.Shuffle(length, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
		sort(arr)
	}
}

func BenchmarkInsertionSort(b *testing.B) { benchmarkSorting(b, InsertionSort) }
func BenchmarkQuickSort(b *testing.B)     { benchmarkSorting(b, QuickSort) }
func BenchmarkFloat64s(b *testing.B)      { benchmarkSorting(b, sort.Float64s) }
