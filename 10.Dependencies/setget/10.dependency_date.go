package setget

import(
	"errors"
)


type Date struct{
	year int
	month int
	day int
}


// Setter methods for fields of Date type
// As Event type has anonymous field of Date type, all the below methods 
// on Date type are exported to Event type as well
func (d *Date) SetYear(year int) error{
	if year < 1 || year > 2020{
		return errors.New("Invalid Year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error{
	if month < 1 || month > 12{
		return errors.New("Invalid Month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error{
	if day < 1 || day > 31{
		return errors.New("Invalid Day")
	}
	d.day = day
	return nil
}


// Getter Methods for fields of Date type
func (d *Date) Day() int{
	return d.day
}

func (d *Date) Month() int{
	return d.month
}

func (d *Date) Year() int{
	return d.year
}


