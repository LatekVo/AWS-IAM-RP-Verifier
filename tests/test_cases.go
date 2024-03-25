package tests

type Testcase struct {
	In   string
	Want bool
}

func GetTestCases() []Testcase {
	return []Testcase{
		{In: "valid_payload_clean.json", Want: true},
		{In: "valid_misordered.json", Want: true},
		{In: "valid_missing_extension", Want: true},          // make sure we don't care about externalities
		{In: "valid_nonstring_field_types.json", Want: true}, // fields are automatically converted to str
		{In: "invalid_payload_clean.json", Want: false},
		{In: "invalid_wrong_format.json", Want: false},
		{In: "bad_filename.json", Want: false},
	}
}
