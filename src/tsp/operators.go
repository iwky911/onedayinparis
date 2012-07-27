package main

import(
	"fmt"
)

func (a AffectData) swap(i, j int) AffectData {
	b := AffectData(make([]int, len(a)))
	copy(b, a)
	b[i], b[j] = a[j], a[i]
	return b
}

func (a AffectData) invertRange(i, j int) AffectData {
	if j < i {
		return a.invertRange(j, i)
	}
	b := AffectData(make([]int, len(a)))
	copy(b, a)
	for k := 0; k <= j-i; k++ {
		b[i+k] = a[j-k]
	}
	return b
}

func (a AffectData) cutAt(i int) AffectData {
	n := len(a)
	b := AffectData(make([]int, n))

	for k := 0; k < i; k++ {
		b[n-i+k] = a[k]
	}
	for k := 0; k < n-i; k++ {
		b[k] = a[i+k]
	}
	return b
}

func (a Affectation) swap(i, j int) Affectation {
	return Affectation{a.data.swap(i,j), fmt.Sprintf("swap %d %d", i,j)}
}

func (a Affectation) invertRange(i, j int) Affectation {
	return Affectation{a.data.invertRange(i,j), fmt.Sprintf("invert range %d %d", i,j)}
}

func (a Affectation) cutAt(i int) Affectation {
	//~ println("in cut at", i)
	return Affectation{a.data.cutAt(i), fmt.Sprintf("cut at %d", i)}
}
