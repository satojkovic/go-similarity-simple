package gosimsimple

import (
	"errors"
	"fmt"
	"math"
)

type Vector []interface{}
type normVector []float64

type PairwiseErr struct {
	msg string
}

func newPairwiseErr(dimx, dimy int) *PairwiseErr {
	err := new(PairwiseErr)
	msgBody := "Incompatible dimension for X and Y vectors: "
	msgDetail := fmt.Sprintf("X == %d while Y == %d", dimx, dimy)
	err.msg = msgBody + msgDetail
	return err
}

func (err *PairwiseErr) Error() string {
	return err.msg
}

func checkPairwiseVector(x, y normVector) error {
	if len(x) != len(y) {
		return newPairwiseErr(len(x), len(y))
	} else {
		return nil
	}
}

func ToFloat64(vec []interface{}) (fvec normVector, err error) {
	fvec = make(normVector, len(vec))
	for i, val := range vec {
		switch val.(type) {
		case int:
			fvec[i] = float64(val.(int))
		case float64:
			fvec[i] = val.(float64)
		default:
			return fvec, errors.New("Invalid value type of Vector")
		}
	}

	return fvec, nil
}

func CosineSimilarity(x, y Vector) (sim float64, err error) {
	X, err := ToFloat64(x)
	if err != nil {
		return sim, err
	}

	Y, err := ToFloat64(y)
	if err != nil {
		return sim, err
	}

	if err := checkPairwiseVector(X, Y); err != nil {
		return sim, err
	}

	dp, err := dotProduct(X, Y)
	if err != nil {
		return sim, err
	}

	sim = dp / (normalize(X) * normalize(Y))

	return sim, nil
}

func PearsonSimilarity(x, y Vector) (sim float64, err error) {
	X, err := ToFloat64(x)
	if err != nil {
		return sim, err
	}

	Y, err := ToFloat64(y)
	if err != nil {
		return sim, err
	}

	if err := checkPairwiseVector(X, Y); err != nil {
		return sim, nil
	}

	dp, err := dotProduct(X, Y)
	if err != nil {
		return sim, err
	}

	sumx := vSum(X)
	sumy := vSum(Y)

	sumxSq := vSumSq(X)
	sumySq := vSumSq(Y)

	n := float64(len(X))
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

func dotProduct(x, y normVector) (dp float64, err error) {
	if len(x) != len(y) {
		return dp, errors.New("Length of two vectors don't match")
	}

	vsize := len(x)
	for i := 0; i < vsize; i++ {
		dp += (x[i] * y[i])
	}

	return dp, nil
}

func normalize(v normVector) (nv float64) {
	vsize := len(v)
	for i := 0; i < vsize; i++ {
		nv += (v[i] * v[i])
	}

	return math.Sqrt(nv)
}

func vSum(v normVector) float64 {
	sum := 0.0
	vsize := len(v)
	for i := 0; i < vsize; i++ {
		sum += v[i]
	}
	return sum
}

func vSumSq(v normVector) float64 {
	sumSq := 0.0
	vsize := len(v)
	for i := 0; i < vsize; i++ {
		sumSq += (v[i] * v[i])
	}
	return sumSq
}
