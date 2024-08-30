package enums

const (
	XYZ_COMMONS_EAN_TYPE_13 = "EAN_13"
	XYZ_COMMONS_EAN_TYPE_8  = "EAN_8"

	XYZ_COMMONS_COUNTRY_AREA_CODE_US = "US"
	XYZ_COMMONS_COUNTRY_AREA_CODE_FR = "FR"
	XYZ_COMMONS_COUNTRY_AREA_CODE_ES = "ES"
	XYZ_COMMONS_COUNTRY_AREA_CODE_IT = "IT"
	XYZ_COMMONS_COUNTRY_AREA_CODE_GB = "GB"
	XYZ_COMMONS_COUNTRY_AREA_CODE_DE = "DE"
	XYZ_COMMONS_COUNTRY_AREA_CODE_AR = "AR"
	XYZ_COMMONS_COUNTRY_AREA_CODE_BR = "BR"
	XYZ_COMMONS_COUNTRY_AREA_CODE_JP = "JP"
	XYZ_COMMONS_COUNTRY_AREA_CODE_CN = "CN"
	XYZ_COMMONS_COUNTRY_AREA_CODE_PT = "PT"

	XYZ_COMMONS_COUNTRY_BR_ACRONYM = "BR"
	XYZ_COMMONS_COUNTRY_US_ACRONYM = "US"
	XYZ_COMMONS_COUNTRY_MX_ACRONYM = "MX"

	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_PNG  = ".png"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_JPG  = ".jpg"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_JPEG = ".jpeg"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_GIF  = ".gif"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_TIF  = ".tif"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_TIFF = ".tiff"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_CSS  = ".css"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_JS   = ".js"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_JSON = ".json"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_PDF  = ".pdf"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_PPT  = ".ppt"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_DOC  = ".doc"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_XLS  = ".xls"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_XLSX = ".xlsx"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_XML  = ".xml"
	XYZ_COMMONS_CONTENT_TYPE_SUFFIX_SVG  = ".svg"

	XYZ_COMMONS_CONTENT_TYPE_PNG  = "image/png"
	XYZ_COMMONS_CONTENT_TYPE_JPG  = "image/jpeg"
	XYZ_COMMONS_CONTENT_TYPE_JPEG = "image/jpeg"
	XYZ_COMMONS_CONTENT_TYPE_GIF  = "image/gif"
	XYZ_COMMONS_CONTENT_TYPE_TIF  = "image/tiff"
	XYZ_COMMONS_CONTENT_TYPE_TIFF = "image/tiff"
	XYZ_COMMONS_CONTENT_TYPE_CSS  = "text/css"
	XYZ_COMMONS_CONTENT_TYPE_JS   = "text/javascript"
	XYZ_COMMONS_CONTENT_TYPE_JSON = "application/json"
	XYZ_COMMONS_CONTENT_TYPE_PDF  = "application/pdf"
	XYZ_COMMONS_CONTENT_TYPE_PPT  = "application/vnd.ms-powerpoint"
	XYZ_COMMONS_CONTENT_TYPE_DOC  = "application/msword"
	XYZ_COMMONS_CONTENT_TYPE_XLS  = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	XYZ_COMMONS_CONTENT_TYPE_XLSX = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	XYZ_COMMONS_CONTENT_TYPE_XML  = "application/xml"
	XYZ_COMMONS_CONTENT_TYPE_SVG  = "image/svg+xml"

	XYZ_COMMONS_RANDOM_STRING_CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	XYZ_COMMONS_REGEX_EMAIL_MATCH_STRING = "^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

	XYZ_COMMONS_ADD_TIMEZONE_IN_TIME_ERROR_MESSAGE = "the timezone is invalid: "

	XYZ_COMMONS_MASK_EMAIL_ERROR_MESSAGE = "the email is invalid"

	XYZ_COMMONS_DOCUMENT_TYPE_CNPJ = "CNPJ"
)

// countries and their states
var (
	BRASIL_ACRONYMS_AND_STATES = map[string]string{
		"AC": "Acre",
		"AL": "Alagoas",
		"AP": "Amapá",
		"AM": "Amazonas",
		"BA": "Bahia",
		"CE": "Ceará",
		"DF": "Distrito Federal",
		"ES": "Espírito Santo",
		"GO": "Goiás",
		"MA": "Maranhão",
		"MG": "Minas Gerais",
		"MT": "Mato Grosso",
		"MS": "Mato Grosso do Sul",
		"PA": "Pará",
		"PB": "Paraíba",
		"PR": "Paraná",
		"PE": "Pernambuco",
		"PI": "Piauí",
		"RJ": "Rio de Janeiro",
		"RN": "Rio Grande do Norte",
		"RO": "Rondônia",
		"RR": "Roraima",
		"RS": "Rio Grande do Sul",
		"SC": "Santa Catarina",
		"SE": "Sergipe",
		"SP": "São Paulo",
		"TO": "Tocantins",
	}

	MEXICO_ACRONYMS_AND_STATES = map[string]string{
		"AG": "Aguascalientes",
		"BN": "Baja California",
		"BS": "Baja California Sur",
		"CP": "Campeche",
		"CS": "Chiapas",
		"CI": "Chihuahua",
		"CH": "Coahuila",
		"CL": "Colima",
		"DF": "Ciudad de México",
		"DG": "Durango",
		"GJ": "Guanajuato",
		"GE": "Guerrero",
		"HD": "Hidalgo",
		"JA": "Jalisco",
		"MX": "México",
		"MC": "Michoacán",
		"MR": "Morelos",
		"NA": "Nayarit",
		"NL": "Nuevo León",
		"OA": "Oaxaca",
		"PU": "Puebla",
		"QE": "Querétaro",
		"QI": "Quintana Roo",
		"SL": "San Luis Potosí",
		"SI": "Sinaloa",
		"SO": "Sonora",
		"TB": "Tabasco",
		"TA": "Tamaulipas",
		"TL": "Tlaxcala",
		"VC": "Veracruz",
		"YU": "Yucatán",
		"ZA": "Zacatecas",
	}

	COUNTRIES_STATES = map[string]map[string]string{
		"BR": BRASIL_ACRONYMS_AND_STATES,
		"MX": MEXICO_ACRONYMS_AND_STATES,
	}
)
