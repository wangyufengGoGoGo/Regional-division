package model

type Dict struct {
	Adcode    string
	Name      string
	Center    string
	Level     string
	Districts []*Dict
}

type Result struct {
	Info      string
	Districts []*Dict
}
