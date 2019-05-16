package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

func main() {
	outp := flag.String("output", "console", "Output location for game. Must be console or apa102")
	config := flag.String("config", "config.json", "Location of config file")
	flag.Parse()

	cfg, err := parseConfig(*config)
	if err != nil {
		fmt.Printf("Cannot parse config file %s: %v", *config, err)
		os.Exit(1)
	}

	_, err = host.Init()
	if err != nil {
		fmt.Printf("failed to initialize periph: %v", err)
		os.Exit(1)
	}

	p, err := spireg.Open(cfg.APA102.Port)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	c, err := p.Connect(physic.Frequency(cfg.APA102.Mhz*float32(physic.MegaHertz)), spi.Mode0, 8)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	_ = c

	_ = outp
	_ = cfg
}

// parseConfig reads and validates the config file
func parseConfig(fileName string) (*Config, error) {
	// read config
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fb, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(fb, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
