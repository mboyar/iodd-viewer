package main

type Column struct {
	TextId string
	Value  string
}

type Table struct {
	Rows []Column
}
