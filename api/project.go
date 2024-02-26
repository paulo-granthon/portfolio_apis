package main

type Project struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Semester uint8  `json:"semester"`
	Company  string `json:"company"`
}
}
