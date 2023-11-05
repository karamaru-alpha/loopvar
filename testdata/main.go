package testdata

func main() {
	for i := range []int{1, 2, 3} {
		i := i // want `overwriting loop iterator "i" is unnecessary \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		_ = i
	}

	for _, v := range []int{1, 2, 3} {
		v := v // want `overwriting loop iterator "v" is unnecessary \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		_ = v
	}

	for i := 1; i <= 3; i++ {
		i := i // want `overwriting loop iterator "i" is unnecessary \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		_ = i
	}

	for i, j := 1, 1; i+j <= 3; i++ {
		i := i // want `overwriting loop iterator "i" is unnecessary \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		j := j // want `overwriting loop iterator "j" is unnecessary \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		_, _ = i, j
	}
}
