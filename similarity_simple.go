package gosimsimple

import (
	"errors"
	"math"
)

type Vector []float64

func isZeroVector(v Vector) bool {
	vsize := len(v)
	for i := 0; i < vsize; i++ {
		if v[i] != 0.0 {
			return false
		}
	}
	return true
}

func dotProduct(x Vector, y Vector) (dp float64, err error) {
	dp = 0.0

	if len(x) != len(y) {
		return dp, errors.New("Length of two vectors don't match")
	}

	vsize := len(x)
	for i := 0; i < vsize; i++ {
		dp += (x[i] * y[i])
	}

	return dp, nil
}

func normalize(v Vector) (nv float64) {
	nv = 0.0

	vsize := len(v)
	for i := 0; i < vsize; i++ {
		nv += (v[i] * v[i])
	}

	return math.Sqrt(nv)
}

func checkVector(x Vector, y Vector) (err error) {
	if len(x) == 0 || len(y) == 0 {
		return errors.New("Length of vector is 0")
	} else if isZeroVector(x) && isZeroVector(y) {
		return errors.New("Both are zero vector")
	} else {
		return nil
	}
}

func vSum(v Vector) float64 {
	sum := 0.0
	vsize := len(v)
	for i := 0; i < vsize; i++ {
		sum += v[i]
	}
	return sum
}

func vSumSq(v Vector) float64 {
	sumSq := 0.0
	vsize := len(v)
	for i := 0; i < vsize; i++ {
		sumSq += (v[i] * v[i])
	}
	return sumSq
}

func CosineSimilarity(x Vector, y Vector) (sim float64, err error) {
	sim = 0.0
	if err := checkVector(x, y); err != nil {
		return sim, err
	}

	dp, err := dotProduct(x, y)
	if err != nil {
		return sim, err
	}

	sim = dp / (normalize(x) * normalize(y))

	return sim, nil
}

func PearsonSimilarity(x Vector, y Vector) (sim float64, err error) {
	sim = 0.0
	if err := checkVector(x, y); err != nil {
		return sim, nil
	}

	dp, err := dotProduct(x, y)
	if err != nil {
		return sim, err
	}

	sumx := vSum(x)
	sumy := vSum(y)

	sumxSq := vSumSq(x)
	sumySq := vSumSq(y)

	n := float64(len(x))
	num := dp - (sumx * sumy / n)
	den := (sumxSq - math.Pow(sumx, 2)/n) *
		(sumySq - math.Pow(sumy, 2)/n)
	den = math.Sqrt(den)

	if den == 0.0 {
		return sim, errors.New("Pearson similarity could not calculated because of division by zero")
	} else {
		sim = num / den
	}

	return sim, nil
}
