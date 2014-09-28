package gosimsimple

import (
	"fmt"
	"testing"
)

func TestSimilaritySimple(t *testing.T) {
	x := Vector{0.789, 0.515, 0.335, 0.0}
	y := Vector{0.832, 0.555, 0.0, 0.0}

	sim, _ := CosineSimilarity(x, y)
	expected_sim := 0.9421693704700748

	if sim != expected_sim {
		t.Errorf("got %v\nwant %v", sim, expected_sim)
	}
}

func TestPersonSimilaritySimple(t *testing.T) {
	x := Vector{1.0, 2.0, 3.0}
	y := Vector{1.0, 3.0, 2.0}

	sim, _ := PearsonSimilarity(x, y)
	expected_sim := 0.5

	if sim != expected_sim {
		t.Errorf("got %v\nwant %v", sim, expected_sim)
	}
}

func ExampleNotMatchSimilaritySimple() {
	x := Vector{1.0, 2.0, 3.0}
	y := Vector{1.0, 2.0, 3.0, 4.0}

	_, err := CosineSimilarity(x, y)
	if err != nil {
		fmt.Println(err)
	}
	// Output: Incompatible dimension for X and Y vectors: X == 3 while Y == 4
}
