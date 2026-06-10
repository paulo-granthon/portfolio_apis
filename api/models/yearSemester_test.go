package models

import "testing"

func TestNewYearSemester_NormalizesSemester(t *testing.T) {
	cases := []struct {
		in   uint8
		want uint8
	}{
		{1, 1},
		{2, 2},
		{3, 1},
		{4, 2},
	}
	for _, c := range cases {
		if got := NewYearSemester(2022, c.in).Semester; got != c.want {
			t.Errorf("NewYearSemester semester %d: want %d, got %d", c.in, c.want, got)
		}
	}
}

func TestYearSemester_ValueScanRoundTrip(t *testing.T) {
	original := YearSemester{Year: 2023, Semester: 2}

	value, err := original.Value()
	if err != nil {
		t.Fatalf("Value() error: %v", err)
	}

	bytes, ok := value.([]byte)
	if !ok {
		t.Fatalf("Value() should return []byte, got %T", value)
	}

	var scanned YearSemester
	if err := scanned.Scan(bytes); err != nil {
		t.Fatalf("Scan() error: %v", err)
	}

	if scanned != original {
		t.Errorf("round trip mismatch: want %+v, got %+v", original, scanned)
	}
}
