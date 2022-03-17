package nestedmap

type nestedMap[A, B comparable, C any] map[A]map[B]C

// NewNestedMap create nestedMap.
func NewNestedMap[A, B comparable, C any]() nestedMap[A, B, C] {
	nm := nestedMap[A, B, C](make(map[A]map[B]C))
	return nm
}

// Set sets value to nestedMap.
func (nm nestedMap[A, B, C]) Set(k1 A, k2 B, v C) {
	nm.nilMapCheck(k1)
	nm[k1][k2] = v
}

// GetOuter gets map of nestedMap.
func (nm nestedMap[A, B, C]) GetOuter(k1 A) (map[B]C, bool) {
	if ret, ok := nm[k1]; ok {
		return ret, ok
	}
	return nil, false
}

// GetInner gets value of nestedMap.
func (nm nestedMap[A, B, C]) GetInner(k1 A, k2 B) (C, bool) {
	var zero C
	nm.nilMapCheck(k1)
	outer, ok := nm[k1]
	if !ok {
		return zero, false
	}
	if inner, ok := outer[k2]; ok {
		return inner, true
	}
	return zero, false
}

func (nm nestedMap[A, B, C]) nilMapCheck(k1 A) {
	if nm[k1] == nil {
		nm[k1] = make(map[B]C)
	}
}
