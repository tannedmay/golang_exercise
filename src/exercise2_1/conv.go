package tempconv

// FtoC convert temperature from Fahrenheit to Celsius.
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CtoF convert temperature from Celsius to Fahrenheit.
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// KtoC convert temperature from Kelvin to Celsius.
func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CtoK convert temperature from Celsius to Kelvin.
func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// KtoF convert temperature from Kelvin to Fahrenheit.
func KtoF(k Kelvin) Fahrenheit {
	return CtoF(KtoC(k))
}

// FtoK convert temperature from Fahrenheit to Kelvin.
func FtoK(f Fahrenheit) Kelvin {
	return CtoK(FtoC(f))
}
