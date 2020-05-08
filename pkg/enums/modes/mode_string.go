// Code generated by "stringer -type=Mode"; DO NOT EDIT.

package modes

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[M1-1]
	_ = x[M2-2]
	_ = x[M3-3]
	_ = x[M4-4]
	_ = x[M5-5]
	_ = x[M6-6]
	_ = x[M7-7]
	_ = x[M8-8]
}

const _Mode_name = "M1M2M3M4M5M6M7M8"

var _Mode_index = [...]uint8{0, 2, 4, 6, 8, 10, 12, 14, 16}

func (i Mode) String() string {
	i -= 1
	if i < 0 || i >= Mode(len(_Mode_index)-1) {
		return "Mode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Mode_name[_Mode_index[i]:_Mode_index[i+1]]
}