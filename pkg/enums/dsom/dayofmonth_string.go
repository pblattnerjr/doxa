// Code generated by "stringer -type=DayOfMonth"; DO NOT EDIT.

package dsom

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[D01-1]
	_ = x[D02-2]
	_ = x[D03-3]
	_ = x[D04-4]
	_ = x[D05-5]
	_ = x[D06-6]
	_ = x[D07-7]
	_ = x[D08-8]
	_ = x[D09-9]
	_ = x[D10-10]
	_ = x[D11-11]
	_ = x[D12-12]
	_ = x[D13-13]
	_ = x[D14-14]
	_ = x[D15-15]
	_ = x[D16-16]
	_ = x[D17-17]
	_ = x[D18-18]
	_ = x[D19-19]
	_ = x[D20-20]
	_ = x[D21-21]
	_ = x[D22-22]
	_ = x[D23-23]
	_ = x[D24-24]
	_ = x[D25-25]
	_ = x[D26-26]
	_ = x[D27-27]
	_ = x[D28-28]
	_ = x[D29-29]
	_ = x[D30-30]
	_ = x[D31-31]
}

const _DayOfMonth_name = "D01D02D03D04D05D06D07D08D09D10D11D12D13D14D15D16D17D18D19D20D21D22D23D24D25D26D27D28D29D30D31"

var _DayOfMonth_index = [...]uint8{0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93}

func (i DayOfMonth) String() string {
	i -= 1
	if i < 0 || i >= DayOfMonth(len(_DayOfMonth_index)-1) {
		return "DayOfMonth(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _DayOfMonth_name[_DayOfMonth_index[i]:_DayOfMonth_index[i+1]]
}