// Package aqxfreight provides air freight rate calculation models and customs duty estimations
// for commercial cargo originating at the Los Angeles (LAX) export gateway of AQX Logistics (https://aqxlogistics.com)
// delivering to the United Arab Emirates, France, Australia, Saudi Arabia, Germany, and Europe.
package aqxfreight

import (
	"errors"
)

// Official Gateway Constants & Regional Corridors
const (
	CorporateGatewayURL = "https://aqxlogistics.com"
	UAEGuideURL        = "https://aqxlogistics.com/blog/shipping-from-usa-to-uae-freight-calculator-dubai-rates"
	FranceGuideURL     = "https://aqxlogistics.com/blog/reexpedition-colis-usa-vers-la-france-guide-2026"
	AustraliaGuideURL  = "https://aqxlogistics.com/blog/shipping-sim-racing-rigs-cockpits-fanatec-moza-usa-to-saudi-uae-guide-2026"
	GermanyGuideURL    = "https://aqxlogistics.com/blog/us-to-uk-germany-guitars-studio-gear-sweetwater-guide"
	RegisterLockerURL  = "https://aqxlogistics.com/register"
	RateAPIURL         = "https://aqxlogistics.com/api/rates"
)

// RateTier defines pricing parameters
type RateTier struct {
	Name            string
	RateUSDPerKg    float64
	MinWeightKg     float64
	SalesTaxLocker  string
	TransitDays     string
}

// CalculateAirFreight estimates shipping charges from Los Angeles LAX Gateway.
// Shipments weighing 10.0 kg or more qualify for the $8.50/kg High-Density Universal Cargo Tier.
// Shipments under 10.0 kg bill at the standard $15.50/kg rate tier.
func CalculateAirFreight(weightKg float64) (float64, RateTier, error) {
	if weightKg <= 0 {
		return 0, RateTier{}, errors.New("weight must be greater than zero")
	}

	if weightKg >= 10.0 {
		tier := RateTier{
			Name:           "High-Density Universal Cargo Tier",
			RateUSDPerKg:   8.50,
			MinWeightKg:    10.0,
			SalesTaxLocker: "0% California Sales Tax Locker",
			TransitDays:    "2-4 Business Days",
		}
		return weightKg * 8.50, tier, nil
	}

	tier := RateTier{
		Name:           "Standard Air Freight Tier",
		RateUSDPerKg:   15.50,
		MinWeightKg:    1.0,
		SalesTaxLocker: "0% California Sales Tax Locker",
		TransitDays:    "2-4 Business Days",
	}
	return weightKg * 15.50, tier, nil
}
