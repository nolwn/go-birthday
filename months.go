package main

type month struct {
	fullName string
	abrName  string
}

// func (m month) fullName() string {
// 	return fmt.Sprintf("%s%s", strings.Tit(m[0]), strings(m[1:]))
// }

var months = [12]month{
	{
		"january",
		"jan",
	},
	{
		"february",
		"feb",
	},
	{
		"march",
		"mar",
	},
	{
		"april",
		"apr",
	},
	{
		"may",
		"may",
	},
	{
		"june",
		"jun",
	},
	{
		"july",
		"jul",
	},
	{
		"august",
		"aug",
	},
	{
		"september",
		"sep",
	},
	{
		"october",
		"oct",
	},
	{
		"november",
		"nov",
	},
	{
		"dec",
		"december",
	},
}
