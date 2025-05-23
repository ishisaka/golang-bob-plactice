// Code generated by BobGen mysql v0.34.2. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

type Factory struct {
	baseCityMods            CityModSlice
	baseCountryMods         CountryModSlice
	baseCountrylanguageMods CountrylanguageModSlice
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) NewCity(mods ...CityMod) *CityTemplate {
	o := &CityTemplate{f: f}

	if f != nil {
		f.baseCityMods.Apply(o)
	}

	CityModSlice(mods).Apply(o)

	return o
}

func (f *Factory) NewCountry(mods ...CountryMod) *CountryTemplate {
	o := &CountryTemplate{f: f}

	if f != nil {
		f.baseCountryMods.Apply(o)
	}

	CountryModSlice(mods).Apply(o)

	return o
}

func (f *Factory) NewCountrylanguage(mods ...CountrylanguageMod) *CountrylanguageTemplate {
	o := &CountrylanguageTemplate{f: f}

	if f != nil {
		f.baseCountrylanguageMods.Apply(o)
	}

	CountrylanguageModSlice(mods).Apply(o)

	return o
}

func (f *Factory) ClearBaseCityMods() {
	f.baseCityMods = nil
}

func (f *Factory) AddBaseCityMod(mods ...CityMod) {
	f.baseCityMods = append(f.baseCityMods, mods...)
}

func (f *Factory) ClearBaseCountryMods() {
	f.baseCountryMods = nil
}

func (f *Factory) AddBaseCountryMod(mods ...CountryMod) {
	f.baseCountryMods = append(f.baseCountryMods, mods...)
}

func (f *Factory) ClearBaseCountrylanguageMods() {
	f.baseCountrylanguageMods = nil
}

func (f *Factory) AddBaseCountrylanguageMod(mods ...CountrylanguageMod) {
	f.baseCountrylanguageMods = append(f.baseCountrylanguageMods, mods...)
}
