// Code generated by BobGen mysql v0.34.2. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"database/sql/driver"

	"github.com/shopspring/decimal"
	"github.com/stephenafamo/bob"
)

// Make sure the type City runs hooks after queries
var _ bob.HookableType = &City{}

// Make sure the type Country runs hooks after queries
var _ bob.HookableType = &Country{}

// Make sure the type Countrylanguage runs hooks after queries
var _ bob.HookableType = &Countrylanguage{}

// Make sure the type CountryContinent satisfies database/sql.Scanner
var _ sql.Scanner = (*CountryContinent)(nil)

// Make sure the type CountryContinent satisfies database/sql/driver.Valuer
var _ driver.Valuer = *new(CountryContinent)

// Make sure the type decimal.Decimal satisfies database/sql.Scanner
var _ sql.Scanner = (*decimal.Decimal)(nil)

// Make sure the type decimal.Decimal satisfies database/sql/driver.Valuer
var _ driver.Valuer = *new(decimal.Decimal)

// Make sure the type CountrylanguageIsOfficial satisfies database/sql.Scanner
var _ sql.Scanner = (*CountrylanguageIsOfficial)(nil)

// Make sure the type CountrylanguageIsOfficial satisfies database/sql/driver.Valuer
var _ driver.Valuer = *new(CountrylanguageIsOfficial)
