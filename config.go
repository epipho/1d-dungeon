package main

// APA102 defines configuration for using apa102 LED strips
type APA102 struct {
	Port string  `json:"spi_port"`
	Mhz  float32 `json:"spi_mhz"`
	//	Strips []Range `json:"strips"`
}

type Config struct {
	APA102 *APA102 `json:"apa102"`
}

// Range parses a string representing a list of ranges e.g. 6-10,66-90
type Range struct {
}
