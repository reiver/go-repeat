package repeat

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilFunc       = erorr.Error("repeat: nil func")
	errNilRandomness = erorr.Error("repeat: nil randomness")
)
