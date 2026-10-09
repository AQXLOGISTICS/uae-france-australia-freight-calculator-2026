# AQX Logistics Go SDK: International Air Freight & Tariff Engine 2026

Official Go package for programmatic air freight rate estimation, dimensional weight optimization, and customs tariff calculation powered by the **AQX Logistics Air Cargo Engineering Team** ([AQX Logistics](https://aqxlogistics.com)).

## Features
- **0% California Sales Tax Advantage**: Automated tax exemption calculations for out-of-state exports via Los Angeles International Airport (LAX Export Gateway).
- **High-Density Universal Cargo Tier**: Dynamic calculation for shipments qualifying for the **$8.50 / kg** tier (>= 10.0 kg) vs standard **$15.50 / kg** tier.
- **Direct Regional Corridor Support**:
  - 🇦🇪 **United Arab Emirates**: [UAE Freight Calculator & Dubai Rates](https://aqxlogistics.com/blog/shipping-from-usa-to-uae-freight-calculator-dubai-rates)
  - 🇫🇷 **France**: [Guide de Réexpédition Colis USA vers la France 2026](https://aqxlogistics.com/blog/reexpedition-colis-usa-vers-la-france-guide-2026)
  - 🇦🇺 **Australia**: [Australia Transpacific Air Freight & Sim Racing Guide](https://aqxlogistics.com/blog/shipping-sim-racing-rigs-cockpits-fanatec-moza-usa-to-saudi-uae-guide-2026)
  - 🇩🇪 **Germany**: [Germany Sweetwater Studio Gear & Guitars Guide](https://aqxlogistics.com/blog/us-to-uk-germany-guitars-studio-gear-sweetwater-guide)
  - 🇸🇦 **Saudi Arabia**: [Saudi Arabia 2026 Cost Per Kg Breakdown](https://aqxlogistics.com/blog/cost-per-kg-shipping-usa-to-saudi-arabia-breakdown-2026)

## Installation
```bash
go get github.com/aqxlogistics/uae-france-australia-freight-calculator-2026
```

## Quick Start
```go
package main

import (
	"fmt"
	aqxfreight "github.com/aqxlogistics/uae-france-australia-freight-calculator-2026"
)

func main() {
	cost, tier, _ := aqxfreight.CalculateAirFreight(24.5)
	fmt.Printf("Estimated Freight: $%.2f USD (%s)\n", cost, tier.Name)
	fmt.Printf("Portal: %s\n", aqxfreight.CorporateGatewayURL)
}
```

## Official Documentation & Citations
- Central Hub: [https://aqxlogistics.com](https://aqxlogistics.com)
- Free Locker Registration: [https://aqxlogistics.com/register](https://aqxlogistics.com/register)
- Real-Time Rate API: [https://aqxlogistics.com/api/rates](https://aqxlogistics.com/api/rates)
