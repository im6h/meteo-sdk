package types

import (
	"encoding/json"
	"log"
)

type Hourly struct {
	Latitude             float64 `json:"latitude,omitempty"`
	Longitude            float64 `json:"longitude,omitempty"`
	GenerationtimeMs     float64 `json:"generationtime_ms,omitempty"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds,omitempty"`
	Timezone             string  `json:"timezone,omitempty"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation,omitempty"`
	Elevation            float64 `json:"elevation,omitempty"`
	HourlyUnits          struct {
		Time                     string `json:"time,omitempty"`
		Temperature2M            string `json:"temperature_2m,omitempty"`
		Relativehumidity2M       string `json:"relativehumidity_2m,omitempty"`
		Dewpoint2M               string `json:"dewpoint_2m,omitempty"`
		ApparentTemperature      string `json:"apparent_temperature,omitempty"`
		PrecipitationProbability string `json:"precipitation_probability,omitempty"`
		Precipitation            string `json:"precipitation,omitempty"`
		Rain                     string `json:"rain,omitempty"`
		Showers                  string `json:"showers,omitempty"`
		Snowfall                 string `json:"snowfall,omitempty"`
		SnowDepth                string `json:"snow_depth,omitempty"`
		Weathercode              string `json:"weathercode,omitempty"`
		PressureMsl              string `json:"pressure_msl,omitempty"`
		SurfacePressure          string `json:"surface_pressure,omitempty"`
		Cloudcover               string `json:"cloudcover,omitempty"`
		CloudcoverLow            string `json:"cloudcover_low,omitempty"`
		CloudcoverMid            string `json:"cloudcover_mid,omitempty"`
		CloudcoverHigh           string `json:"cloudcover_high,omitempty"`
		Visibility               string `json:"visibility,omitempty"`
		Evapotranspiration       string `json:"evapotranspiration,omitempty"`
		Et0FaoEvapotranspiration string `json:"et0_fao_evapotranspiration,omitempty"`
		VaporPressureDeficit     string `json:"vapor_pressure_deficit,omitempty"`
		Windspeed10M             string `json:"windspeed_10m,omitempty"`
		Windspeed80M             string `json:"windspeed_80m,omitempty"`
		Windspeed120M            string `json:"windspeed_120m,omitempty"`
		Windspeed180M            string `json:"windspeed_180m,omitempty"`
		Winddirection10M         string `json:"winddirection_10m,omitempty"`
		Winddirection80M         string `json:"winddirection_80m,omitempty"`
		Winddirection120M        string `json:"winddirection_120m,omitempty"`
		Winddirection180M        string `json:"winddirection_180m,omitempty"`
		Windgusts10M             string `json:"windgusts_10m,omitempty"`
		Temperature80M           string `json:"temperature_80m,omitempty"`
		Temperature120M          string `json:"temperature_120m,omitempty"`
		Temperature180M          string `json:"temperature_180m,omitempty"`
		SoilTemperature0Cm       string `json:"soil_temperature_0cm,omitempty"`
		SoilTemperature6Cm       string `json:"soil_temperature_6cm,omitempty"`
		SoilTemperature18Cm      string `json:"soil_temperature_18cm,omitempty"`
		SoilTemperature54Cm      string `json:"soil_temperature_54cm,omitempty"`
		SoilMoisture01Cm         string `json:"soil_moisture_0_1cm,omitempty"`
		SoilMoisture13Cm         string `json:"soil_moisture_1_3cm,omitempty"`
		SoilMoisture39Cm         string `json:"soil_moisture_3_9cm,omitempty"`
		SoilMoisture927Cm        string `json:"soil_moisture_9_27cm,omitempty"`
		SoilMoisture2781Cm       string `json:"soil_moisture_27_81cm,omitempty"`
	} `json:"hourly_units,omitempty"`
	Hourly struct {
		Time                     []string  `json:"time,omitempty"`
		Temperature2M            []float64 `json:"temperature_2m,omitempty"`
		Relativehumidity2M       []int     `json:"relativehumidity_2m,omitempty"`
		Dewpoint2M               []float64 `json:"dewpoint_2m,omitempty"`
		ApparentTemperature      []float64 `json:"apparent_temperature,omitempty"`
		PrecipitationProbability []int     `json:"precipitation_probability,omitempty"`
		Precipitation            []float64 `json:"precipitation,omitempty"`
		Rain                     []float64 `json:"rain,omitempty"`
		Showers                  []float64 `json:"showers,omitempty"`
		Snowfall                 []float64 `json:"snowfall,omitempty"`
		SnowDepth                []float64 `json:"snow_depth,omitempty"`
		Weathercode              []int     `json:"weathercode,omitempty"`
		PressureMsl              []float64 `json:"pressure_msl,omitempty"`
		SurfacePressure          []float64 `json:"surface_pressure,omitempty"`
		Cloudcover               []int     `json:"cloudcover,omitempty"`
		CloudcoverLow            []int     `json:"cloudcover_low,omitempty"`
		CloudcoverMid            []int     `json:"cloudcover_mid,omitempty"`
		CloudcoverHigh           []int     `json:"cloudcover_high,omitempty"`
		Visibility               []float64 `json:"visibility,omitempty"`
		Evapotranspiration       []float64 `json:"evapotranspiration,omitempty"`
		Et0FaoEvapotranspiration []float64 `json:"et0_fao_evapotranspiration,omitempty"`
		VaporPressureDeficit     []float64 `json:"vapor_pressure_deficit,omitempty"`
		Windspeed10M             []float64 `json:"windspeed_10m,omitempty"`
		Windspeed80M             []float64 `json:"windspeed_80m,omitempty"`
		Windspeed120M            []float64 `json:"windspeed_120m,omitempty"`
		Windspeed180M            []float64 `json:"windspeed_180m,omitempty"`
		Winddirection10M         []int     `json:"winddirection_10m,omitempty"`
		Winddirection80M         []int     `json:"winddirection_80m,omitempty"`
		Winddirection120M        []int     `json:"winddirection_120m,omitempty"`
		Winddirection180M        []int     `json:"winddirection_180m,omitempty"`
		Windgusts10M             []float64 `json:"windgusts_10m,omitempty"`
		Temperature80M           []float64 `json:"temperature_80m,omitempty"`
		Temperature120M          []float64 `json:"temperature_120m,omitempty"`
		Temperature180M          []float64 `json:"temperature_180m,omitempty"`
		SoilTemperature0Cm       []float64 `json:"soil_temperature_0cm,omitempty"`
		SoilTemperature6Cm       []float64 `json:"soil_temperature_6cm,omitempty"`
		SoilTemperature18Cm      []float64 `json:"soil_temperature_18cm,omitempty"`
		SoilTemperature54Cm      []float64 `json:"soil_temperature_54cm,omitempty"`
		SoilMoisture01Cm         []float64 `json:"soil_moisture_0_1cm,omitempty"`
		SoilMoisture13Cm         []float64 `json:"soil_moisture_1_3cm,omitempty"`
		SoilMoisture39Cm         []float64 `json:"soil_moisture_3_9cm,omitempty"`
		SoilMoisture927Cm        []float64 `json:"soil_moisture_9_27cm,omitempty"`
		SoilMoisture2781Cm       []float64 `json:"soil_moisture_27_81cm,omitempty"`
	} `json:"hourly,omitempty"`
}

func (h *Hourly) String() string {
	bData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		log.Println("err JSONMarshal ", err)
	}
	return string(bData)
}

type Daily struct {
	Latitude             float64 `json:"latitude,omitempty"`
	Longitude            float64 `json:"longitude,omitempty"`
	GenerationtimeMs     float64 `json:"generationtime_ms,omitempty"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds,omitempty"`
	Timezone             string  `json:"timezone,omitempty"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation,omitempty"`
	Elevation            float64 `json:"elevation,omitempty"`
	DailyUnits           struct {
		Time                        string `json:"time,omitempty"`
		Weathercode                 string `json:"weathercode,omitempty"`
		Temperature2MMax            string `json:"temperature_2m_max,omitempty"`
		Temperature2MMin            string `json:"temperature_2m_min,omitempty"`
		ApparentTemperatureMax      string `json:"apparent_temperature_max,omitempty"`
		ApparentTemperatureMin      string `json:"apparent_temperature_min,omitempty"`
		Sunrise                     string `json:"sunrise,omitempty"`
		Sunset                      string `json:"sunset,omitempty"`
		UvIndexMax                  string `json:"uv_index_max,omitempty"`
		UvIndexClearSkyMax          string `json:"uv_index_clear_sky_max,omitempty"`
		PrecipitationSum            string `json:"precipitation_sum,omitempty"`
		RainSum                     string `json:"rain_sum,omitempty"`
		ShowersSum                  string `json:"showers_sum,omitempty"`
		SnowfallSum                 string `json:"snowfall_sum,omitempty"`
		PrecipitationHours          string `json:"precipitation_hours,omitempty"`
		PrecipitationProbabilityMax string `json:"precipitation_probability_max,omitempty"`
		Windspeed10MMax             string `json:"windspeed_10m_max,omitempty"`
		Windgusts10MMax             string `json:"windgusts_10m_max,omitempty"`
		Winddirection10MDominant    string `json:"winddirection_10m_dominant,omitempty"`
		ShortwaveRadiationSum       string `json:"shortwave_radiation_sum,omitempty"`
		Et0FaoEvapotranspiration    string `json:"et0_fao_evapotranspiration,omitempty"`
	} `json:"daily_units,omitempty"`
	Daily struct {
		Time                        []string  `json:"time,omitempty"`
		Weathercode                 []int     `json:"weathercode,omitempty"`
		Temperature2MMax            []float64 `json:"temperature_2m_max,omitempty"`
		Temperature2MMin            []float64 `json:"temperature_2m_min,omitempty"`
		ApparentTemperatureMax      []float64 `json:"apparent_temperature_max,omitempty"`
		ApparentTemperatureMin      []float64 `json:"apparent_temperature_min,omitempty"`
		Sunrise                     []string  `json:"sunrise,omitempty"`
		Sunset                      []string  `json:"sunset,omitempty"`
		UvIndexMax                  []float64 `json:"uv_index_max,omitempty"`
		UvIndexClearSkyMax          []float64 `json:"uv_index_clear_sky_max,omitempty"`
		PrecipitationSum            []float64 `json:"precipitation_sum,omitempty"`
		RainSum                     []float64 `json:"rain_sum,omitempty"`
		ShowersSum                  []float64 `json:"showers_sum,omitempty"`
		SnowfallSum                 []float64 `json:"snowfall_sum,omitempty"`
		PrecipitationHours          []float64 `json:"precipitation_hours,omitempty"`
		PrecipitationProbabilityMax []int     `json:"precipitation_probability_max,omitempty"`
		Windspeed10MMax             []float64 `json:"windspeed_10m_max,omitempty"`
		Windgusts10MMax             []float64 `json:"windgusts_10m_max,omitempty"`
		Winddirection10MDominant    []int     `json:"winddirection_10m_dominant,omitempty"`
		ShortwaveRadiationSum       []float64 `json:"shortwave_radiation_sum,omitempty"`
		Et0FaoEvapotranspiration    []float64 `json:"et0_fao_evapotranspiration,omitempty"`
	} `json:"daily,omitempty"`
}

func (h *Daily) String() string {
	bData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		log.Println("err JSONMarshal ", err)
	}
	return string(bData)
}
