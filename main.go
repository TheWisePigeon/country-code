package countrycode

type CountryCode struct {
	ISO  string
	Name string
}

var countryCodes = map[string]CountryCode{
	"93": {
		ISO:  "AF/AFG",
		Name: "Afghanistan",
	},
	"355": {
		ISO:  "AL/ALB",
		Name: "Albania",
	},
	"213": {
		ISO:  "DZ/DZA",
		Name: "Algeria",
	},
	"1-684": {
		ISO:  "AS/ASM",
		Name: "American Samoa",
	},
	"376": {
		ISO:  "AD/AND",
		Name: "Andorra",
	},
	"244": {
		ISO:  "AO/AGO",
		Name: "Angola",
	},
	"1-264": {
		ISO:  "AI/AIA",
		Name: "Anguilla",
	},
	"672": {
		ISO:  "AQ/ATA",
		Name: "Antarctica",
	},
	"1-268": {
		ISO:  "AG/ATG",
		Name: "Antigua and Barbuda",
	},
	"54": {
		ISO:  "AR/ARG",
		Name: "Argentina",
	},
	"374": {
		ISO:  "AM/ARM",
		Name: "Armenia",
	},
	"297": {
		ISO:  "AW/ABW",
		Name: "Aruba",
	},
	"61": {
		ISO:  "AU/AUS",
		Name: "Australia",
	},
	"43": {
		ISO:  "AT/AUT",
		Name: "Austria",
	},
	"994": {
		ISO:  "AZ/AZE",
		Name: "Azerbaijan",
	},
}
