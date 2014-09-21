package similarity_simple

import (
	"fmt"
	"testing"
)

func TestSimilaritySimple(t *testing.T) {
	v1 := Vector{0.789, 0.515, 0.335, 0.0}
	v2 := Vector{0.832, 0.555, 0.0, 0.0}

	sim, _ := CosineSimilarity(v1, v2)
	expected_sim := 0.9421693704700748

	if sim != expected_sim {
		t.Errorf("got %v\nwant %v", sim, expected_sim)
	}
}

func ExampleEmptySimilaritySimple() {
	v1 := Vector{}
	v2 := Vector{}

	_, err := CosineSimilarity(v1, v2)
	if err != nil {
		fmt.Println(err)
	}
	// Output: Length of vector is 0
}

func ExampleZeroSimilaritySimple() {
	v1 := Vector{0.0, 0.0, 0.0}
	v2 := Vector{0.0, 0.0, 0.0}

	_, err := CosineSimilarity(v1, v2)
	if err != nil {
		fmt.Println(err)
	}
	// Output: Both are zero vector
}

func ExampleNotMatchSimilaritySimple() {
	v1 := Vector{1.0, 2.0, 3.0}
	v2 := Vector{1.0, 2.0, 3.0, 4.0}

	_, err := CosineSimilarity(v1, v2)
	if err != nil {
		fmt.Println(err)
	}
	// Output: Length of two vectors don't match
}
