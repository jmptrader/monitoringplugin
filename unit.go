package monitoringplugin

import (
	"fmt"
	"time"
)

type Unit interface {
	String() string
	Value() float64
	UnitString() string
}

type DurationUnit time.Duration

func (unit DurationUnit) String() string {
	return time.Duration(unit).String()
}

func (unit DurationUnit) UnitString() string {
	return "s"
}

func (unit DurationUnit) Value() float64 {
	return time.Duration(unit).Seconds()
}

type CounterUnit uint64

func (unit CounterUnit) String() string {
	return fmt.Sprintf("%dc", uint64(unit))
}

func (unit CounterUnit) UnitString() string {
	return "c"
}

func (unit CounterUnit) Value() float64 {
	return float64(unit)
}

type NumberUnit float64

func (unit NumberUnit) String() string {
	return fmt.Sprintf("%.2f", unit.Value())
}

func (unit NumberUnit) UnitString() string {
	return ""
}

func (unit NumberUnit) Value() float64 {
	return float64(unit)
}

type PercentageUnit struct {
	Base      float64
	Quantitiy float64
}

func (unit PercentageUnit) String() string {
	return fmt.Sprintf("%.2f%%", unit.Value())
}

func (unit PercentageUnit) UnitString() string {
	return "%"
}

func (unit PercentageUnit) Value() float64 {
	return unit.Quantitiy / (unit.Base / 100)
}

type BytesUnit int64

func (unit BytesUnit) String() string {
	result := ""
	kibi := float64(unit) / 1024
	if kibi > 1 {
		mibi := kibi / 1024
		if mibi > 1 {
			gibi := mibi / 1024
			if gibi > 1 {
				tibi := gibi / 1024
				if tibi > 1 {
					result = fmt.Sprintf("%.2fTiB", tibi)
				} else {
					result = fmt.Sprintf("%.2fGiB", gibi)
				}
			} else {
				result = fmt.Sprintf("%.2fMiB", mibi)
			}
		} else {
			result = fmt.Sprintf("%.2fKiB", kibi)
		}
	} else {
		result = fmt.Sprintf("%dB", unit)
	}
	return result
}

func (unit BytesUnit) UnitString() string {
	return "B"
}

func (unit BytesUnit) Value() float64 {
	return float64(unit)
}
