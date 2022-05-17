package datetime

import "time"

type Datetime struct {
	num     int
	now     time.Time
	dateStr string
}

func Now() *time.Time {
	now := time.Now().UTC()
	return &now
}

func IsFuture(d1 *time.Time) bool {
	now := Now()
	return d1.UTC().After(*now)
}

func ToString(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(time.RFC3339)
}

func Parse(s string) *Datetime {
	return &Datetime{
		dateStr: s,
	}
}

func (d *Datetime) ToISO() *time.Time {
	dt, _ := time.Parse(time.RFC3339, d.dateStr)
	return &dt
}

func Get(num int) *Datetime {
	now := time.Now().UTC()
	return &Datetime{
		num: num,
		now: now,
	}
}

func (d *Datetime) HourAgo() *time.Time {
	return d.HoursAgo()
}

func (d *Datetime) Hour() *time.Time {
	return d.Hours()
}

func (d *Datetime) HoursAgo() *time.Time {
	dt := d.now.Add(-time.Hour * time.Duration(d.num))
	return &dt
}

func (d *Datetime) Hours() *time.Time {
	dt := d.now.Add(time.Hour * time.Duration(d.num))
	return &dt
}
