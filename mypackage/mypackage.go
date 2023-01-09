package mypackage

import "errors"

type Date struct {
	day, month, year int
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) SetDay(day int) error {
	if err := d.checkDay(day); err != nil {
		return err
	}
	d.day = day
	return nil
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) SetMonth(month int) {
	d.month = month
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) SetYear(year int) {
	d.year = year
}

func (d *Date) checkDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("incorrect day range")
	}
	return nil
}
