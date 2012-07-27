package main

import (
	"fmt"
)

func (a AffectData) swap(i, j int, b *AffectData) {
	copy(*b, a)
	//~ println(i,j)
	(*b)[i], (*b)[j] = a[j], a[i]
}

func (a AffectData) invertRange(i, j int, b *AffectData) {
	if j < i {
		a.invertRange(j, i, b)
		return
	}
	copy(*b, a)
	for k := 0; k <= j-i; k++ {
		(*b)[i+k] = a[j-k]
	}
}

func (a AffectData) cutAt(i int, b *AffectData) {
	n := len(a)

	for k := 0; k < i; k++ {
		(*b)[n-i+k] = a[k]
	}
	for k := 0; k < n-i; k++ {
		(*b)[k] = a[i+k]
	}
}

func (a Affectation) swap(i, j int, b *Affectation) {
	a.data.swap(i, j, &(*b).data)
	(*b).name = fmt.Sprintf("swap %d %d", i, j)
}

func (a Affectation) invertRange(i, j int, b *Affectation) {
	a.data.invertRange(i, j, &(*b).data)
	(*b).name = fmt.Sprintf("invert range %d %d", i, j)
}

func (a Affectation) cutAt(i int, b *Affectation) {
	//~ println("in cut at", i)
	a.data.cutAt(i, &(*b).data)
	(*b).name = fmt.Sprintf("cut at %d", i)
}
