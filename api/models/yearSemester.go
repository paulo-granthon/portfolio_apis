package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type YearSemester struct {
	Year     uint16 `json:"year"`
	Semester uint8  `json:"semester"`
}

func NewYearSemester(
	year uint16,
	semester uint8,
) YearSemester {
	return YearSemester{
		Year:     year,
		Semester: ((semester - 1) % 2) + 1, // ensure to 1 or 2
	}
}

func (y YearSemester) Value() (driver.Value, error) {
	return json.Marshal(y)
}

func (y *YearSemester) Scan(value interface{}) error {
	yearSemesterBytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("yearSemester must be a byte array")
	}
	return json.Unmarshal(yearSemesterBytes, y)
}
