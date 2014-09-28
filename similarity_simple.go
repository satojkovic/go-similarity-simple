package gosimsimple

import (
	"errors"
	"fmt"
	"math"
)

type Vector []float64

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

func checkPairwiseVector(x, y Vector) error {
	if len(x) != len(y) {
		return newPairwiseErr(len(x), len(y))
	} else {
		return nil
	}
}

func CosineSimilarity(x, y Vector) (sim float64, err error) {
	if err := checkPairwiseVector(x, y); err != nil {
		return sim, err
	}

	dp, err := dotProduct(x, y)
	if err != nil {
		return sim, err
	}

	sim = dp / (normalize(x) * normalize(y))

	return sim, nil
}

func PearsonSimilarity(x, y Vector) (sim float64, err error) {
	if err := checkPairwiseVector(x, y); err != nil {
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

func dotProduct(x, y Vector) (dp float64, err error) {
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
	vsize := len(v)
	for i := 0; i < vsize; i++ {
		nv += (v[i] * v[i])
	}

	return math.Sqrt(nv)
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
