package a

func f() {
	for i, v := range []int{1, 2, 3} {
		i := i  // want `The loop variable "i" should not be copied \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		_v := v // want `The loop variable "v" should not be copied \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		_ = i
		_ = _v
	}

	for i := 1; i <= 3; i++ {
		i := i // want `The loop variable "i" should not be copied \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		_ = i
	}

	for i, j := 1, 1; i+j <= 3; i++ {
		i := i       // want `The loop variable "i" should not be copied \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		j, _ := j, 1 // want `The loop variable "j" should not be copied \(GO_VERSION >= 1.22 or GOEXPERIMENT=loopvar\)`
		_, _ = i, j
	}
}
