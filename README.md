go-similarity-simple
=====================
A simple library of calculating similarity in Go. Inspired by scikit-learn. 

Example
=====================

    package main

    import "fmt"
    import "github.com/satojkovic/go-similarity-simple"

    func main() {
	    v1 := Vector{0.789, 0.515, 0.335, 0.0}
	    v2 := Vector{0.832, 0.555, 0.0, 0.0}

        sim, _ := CosineSimilarity(v1, v2)
        fmt.Printf("Similarity = %f\n", sim)
    }


Features
=====================

* Cosine similarity
* Pearson similarity

License
=====================

This software is released under the MIT License, see LICENSE.txt
