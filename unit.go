package monitoringplugin

import (
	"fmt"
	"time"
)

type UnitSpecification interface {
	String() string
}

type Unit interface {
	HumanReadable() string
	Value() float64
}

type SimpleUnitSpecification string

func (unitSpec SimpleUnitSpecification) String() string {
	return string(unitSpec)
}

type DurationUnit time.Duration

const DurationUnitSpecification = SimpleUnitSpecification("s")

func (unit DurationUnit) HumanReadable() string {
	return time.Duration(unit).String()
}

func (unit DurationUnit) Value() float64 {
	return time.Duration(unit).Seconds()
}

type CounterUnit uint64

const CounterUnitSpecification = SimpleUnitSpecification("c")

func (unit CounterUnit) HumanReadable() string {
	return fmt.Sprintf("%dc", uint64(unit))
}

func (unit CounterUnit) Value() float64 {
	return float64(unit)
}

type NumberUnit float64

const NumberUnitSpecification = SimpleUnitSpecification("")

func (unit NumberUnit) HumanReadable() string {
	return fmt.Sprintf("%.2f", unit.Value())
}

func (unit NumberUnit) Value() float64 {
	return float64(unit)
}

type PercentageUnit struct {
	Base      float64
	Quantitiy float64
}

const PercentageUnitSpecification = SimpleUnitSpecification("%")

func (unit PercentageUnit) HumanReadable() string {
	return fmt.Sprintf("%.2f%%", unit.Value())
}

func (unit PercentageUnit) Value() float64 {
	return unit.Quantitiy / (unit.Base / 100)
}

type BytesUnit int64

const BytesUnitSpecification = SimpleUnitSpecification("B")

func (unit BytesUnit) HumanReadable() string {
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

func (unit BytesUnit) Value() float64 {
	return float64(unit)
}
