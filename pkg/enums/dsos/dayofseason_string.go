// Code generated by "stringer -type=DayOfSeason"; DO NOT EDIT.

package dsos

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
	_ = x[D32-32]
	_ = x[D33-33]
	_ = x[D34-34]
	_ = x[D35-35]
	_ = x[D36-36]
	_ = x[D37-37]
	_ = x[D38-38]
	_ = x[D39-39]
	_ = x[D40-40]
	_ = x[D41-41]
	_ = x[D42-42]
	_ = x[D43-43]
	_ = x[D44-44]
	_ = x[D45-45]
	_ = x[D46-46]
	_ = x[D47-47]
	_ = x[D48-48]
	_ = x[D49-49]
	_ = x[D50-50]
	_ = x[D51-51]
	_ = x[D52-52]
	_ = x[D53-53]
	_ = x[D54-54]
	_ = x[D55-55]
	_ = x[D56-56]
	_ = x[D57-57]
	_ = x[D58-58]
	_ = x[D59-59]
	_ = x[D60-60]
	_ = x[D61-61]
	_ = x[D62-62]
	_ = x[D63-63]
	_ = x[D64-64]
	_ = x[D65-65]
	_ = x[D66-66]
	_ = x[D67-67]
	_ = x[D68-68]
	_ = x[D69-69]
	_ = x[D70-70]
}

const _DayOfSeason_name = "D01D02D03D04D05D06D07D08D09D10D11D12D13D14D15D16D17D18D19D20D21D22D23D24D25D26D27D28D29D30D31D32D33D34D35D36D37D38D39D40D41D42D43D44D45D46D47D48D49D50D51D52D53D54D55D56D57D58D59D60D61D62D63D64D65D66D67D68D69D70"

var _DayOfSeason_index = [...]uint8{0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99, 102, 105, 108, 111, 114, 117, 120, 123, 126, 129, 132, 135, 138, 141, 144, 147, 150, 153, 156, 159, 162, 165, 168, 171, 174, 177, 180, 183, 186, 189, 192, 195, 198, 201, 204, 207, 210}

func (i DayOfSeason) String() string {
	i -= 1
	if i < 0 || i >= DayOfSeason(len(_DayOfSeason_index)-1) {
		return "DayOfSeason(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _DayOfSeason_name[_DayOfSeason_index[i]:_DayOfSeason_index[i+1]]
}