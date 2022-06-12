package unitconv

type Celsius float64
type Fahrenheit float64
type Feet float64
type Pound float64
type Meter float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	FtToMRatio    float64 = 0.3048
	KgToLbRatio   float64 = 0.45359237
)

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func MToFts(f Feet) Meter {
	return Meter(f / Feet(FtToMRatio))
}

func FtToM(m Meter) Feet {
	return Feet(m * Meter(FtToMRatio))
}

func KgToLbs(k Kilogram) Pound {
	return Pound(k / Kilogram(KgToLbRatio))
}

func LbsToKg(lbs Pound) Pound {
	return Pound(lbs * Pound(KgToLbRatio))
}
