package main

//Date : Struct depciting the types
type Date struct {
	year  int16
	month int8
	day   int8
}

//Adding setters
func (d *Date) SetYear(year int16) {
	d.year = year
}

//Adding setters
func (d *Date) SetDay(value int8) {
	d.day = value
}

//Adding setters
func (d *Date) SetMonth(value int8) {
	d.month = value
}

//Adding setters
func (d *Date) Year() int16 {
	return d.year
}

//Adding setters
func (d *Date) Day() int8 {
	return d.day
}

//Adding setters
func (d *Date) Month() int8 {
	return d.month
}
