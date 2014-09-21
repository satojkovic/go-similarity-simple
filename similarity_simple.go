package similarity_simple

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

func dotProduct(v1 Vector, v2 Vector) (dp float64, err error) {
	dp = 0.0

	if len(v1) != len(v2) {
		return dp, errors.New("Length of two vectors don't match")
	}

	vsize := len(v1)
	for i := 0; i < vsize; i++ {
		dp += (v1[i] * v2[i])
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

func CosineSimilarity(v1 Vector, v2 Vector) (sim float64, err error) {
	sim = 0.0
	if len(v1) == 0 || len(v2) == 0 {
		return sim, errors.New("Length of vector is 0")
	} else if isZeroVector(v1) && isZeroVector(v2) {
		return sim, errors.New("Both are zero vector")
	}

	dp, err := dotProduct(v1, v2)
	if err != nil {
		return sim, err
	}

	sim = dp / (normalize(v1) * normalize(v2))

	return sim, nil
}
