package main

func (a Affectation) swap( i, j int) Affectation {
	b := Affectation(make([]int, len(a)))
	copy(b,a)
	b[i], b[j] = a[j], a[i]
	return b
}

func (a Affectation) invertRange(i, j int) Affectation {
	if j<i {
	  return a.invertRange(j, i)
	}
	b := Affectation(make([]int, len(a)))
	copy(b,a)
	for k:=0; k<=j-i; k++{
	  b[i+k] = a[j-k]
	}
	return b
}

func (a Affectation) cutAt(i int) Affectation{
	n := len(a)
	b:= Affectation(make([]int, n))
	
	for k:=0; k<i; k++ {
		b[n-i+k] = a[k]
	}
	for k := 0; k<n-i; k++ {
		b[k] = a[i+k]
	}
	return b
	}
