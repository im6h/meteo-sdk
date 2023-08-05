package types

import (
	"encoding/json"
	"log"
)

type Hourly struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	HourlyUnits          struct {
		Time                     string `json:"time"`
		Temperature2M            string `json:"temperature_2m"`
		Relativehumidity2M       string `json:"relativehumidity_2m"`
		Dewpoint2M               string `json:"dewpoint_2m"`
		ApparentTemperature      string `json:"apparent_temperature"`
		PrecipitationProbability string `json:"precipitation_probability"`
		Precipitation            string `json:"precipitation"`
		Rain                     string `json:"rain"`
		Showers                  string `json:"showers"`
		Snowfall                 string `json:"snowfall"`
		SnowDepth                string `json:"snow_depth"`
		Weathercode              string `json:"weathercode"`
		PressureMsl              string `json:"pressure_msl"`
		SurfacePressure          string `json:"surface_pressure"`
		Cloudcover               string `json:"cloudcover"`
		CloudcoverLow            string `json:"cloudcover_low"`
		CloudcoverMid            string `json:"cloudcover_mid"`
		CloudcoverHigh           string `json:"cloudcover_high"`
		Visibility               string `json:"visibility"`
		Evapotranspiration       string `json:"evapotranspiration"`
		Et0FaoEvapotranspiration string `json:"et0_fao_evapotranspiration"`
		VaporPressureDeficit     string `json:"vapor_pressure_deficit"`
		Windspeed10M             string `json:"windspeed_10m"`
		Windspeed80M             string `json:"windspeed_80m"`
		Windspeed120M            string `json:"windspeed_120m"`
		Windspeed180M            string `json:"windspeed_180m"`
		Winddirection10M         string `json:"winddirection_10m"`
		Winddirection80M         string `json:"winddirection_80m"`
		Winddirection120M        string `json:"winddirection_120m"`
		Winddirection180M        string `json:"winddirection_180m"`
		Windgusts10M             string `json:"windgusts_10m"`
		Temperature80M           string `json:"temperature_80m"`
		Temperature120M          string `json:"temperature_120m"`
		Temperature180M          string `json:"temperature_180m"`
		SoilTemperature0Cm       string `json:"soil_temperature_0cm"`
		SoilTemperature6Cm       string `json:"soil_temperature_6cm"`
		SoilTemperature18Cm      string `json:"soil_temperature_18cm"`
		SoilTemperature54Cm      string `json:"soil_temperature_54cm"`
		SoilMoisture01Cm         string `json:"soil_moisture_0_1cm"`
		SoilMoisture13Cm         string `json:"soil_moisture_1_3cm"`
		SoilMoisture39Cm         string `json:"soil_moisture_3_9cm"`
		SoilMoisture927Cm        string `json:"soil_moisture_9_27cm"`
		SoilMoisture2781Cm       string `json:"soil_moisture_27_81cm"`
	} `json:"hourly_units"`
	Hourly struct {
		Time                     []string  `json:"time"`
		Temperature2M            []float64 `json:"temperature_2m"`
		Relativehumidity2M       []int     `json:"relativehumidity_2m"`
		Dewpoint2M               []float64 `json:"dewpoint_2m"`
		ApparentTemperature      []float64 `json:"apparent_temperature"`
		PrecipitationProbability []int     `json:"precipitation_probability"`
		Precipitation            []float64 `json:"precipitation"`
		Rain                     []float64 `json:"rain"`
		Showers                  []float64 `json:"showers"`
		Snowfall                 []float64 `json:"snowfall"`
		SnowDepth                []float64 `json:"snow_depth"`
		Weathercode              []int     `json:"weathercode"`
		PressureMsl              []float64 `json:"pressure_msl"`
		SurfacePressure          []float64 `json:"surface_pressure"`
		Cloudcover               []int     `json:"cloudcover"`
		CloudcoverLow            []int     `json:"cloudcover_low"`
		CloudcoverMid            []int     `json:"cloudcover_mid"`
		CloudcoverHigh           []int     `json:"cloudcover_high"`
		Visibility               []float64 `json:"visibility"`
		Evapotranspiration       []float64 `json:"evapotranspiration"`
		Et0FaoEvapotranspiration []float64 `json:"et0_fao_evapotranspiration"`
		VaporPressureDeficit     []float64 `json:"vapor_pressure_deficit"`
		Windspeed10M             []float64 `json:"windspeed_10m"`
		Windspeed80M             []float64 `json:"windspeed_80m"`
		Windspeed120M            []float64 `json:"windspeed_120m"`
		Windspeed180M            []float64 `json:"windspeed_180m"`
		Winddirection10M         []int     `json:"winddirection_10m"`
		Winddirection80M         []int     `json:"winddirection_80m"`
		Winddirection120M        []int     `json:"winddirection_120m"`
		Winddirection180M        []int     `json:"winddirection_180m"`
		Windgusts10M             []float64 `json:"windgusts_10m"`
		Temperature80M           []float64 `json:"temperature_80m"`
		Temperature120M          []float64 `json:"temperature_120m"`
		Temperature180M          []float64 `json:"temperature_180m"`
		SoilTemperature0Cm       []float64 `json:"soil_temperature_0cm"`
		SoilTemperature6Cm       []float64 `json:"soil_temperature_6cm"`
		SoilTemperature18Cm      []float64 `json:"soil_temperature_18cm"`
		SoilTemperature54Cm      []float64 `json:"soil_temperature_54cm"`
		SoilMoisture01Cm         []float64 `json:"soil_moisture_0_1cm"`
		SoilMoisture13Cm         []float64 `json:"soil_moisture_1_3cm"`
		SoilMoisture39Cm         []float64 `json:"soil_moisture_3_9cm"`
		SoilMoisture927Cm        []float64 `json:"soil_moisture_9_27cm"`
		SoilMoisture2781Cm       []float64 `json:"soil_moisture_27_81cm"`
	} `json:"hourly"`
}

func (h *Hourly) String() string {
	bData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		log.Println("err JSONMarshal ", err)
	}
	return string(bData)
}

type Daily struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	DailyUnits           struct {
		Time                        string `json:"time"`
		Weathercode                 string `json:"weathercode"`
		Temperature2MMax            string `json:"temperature_2m_max"`
		Temperature2MMin            string `json:"temperature_2m_min"`
		ApparentTemperatureMax      string `json:"apparent_temperature_max"`
		ApparentTemperatureMin      string `json:"apparent_temperature_min"`
		Sunrise                     string `json:"sunrise"`
		Sunset                      string `json:"sunset"`
		UvIndexMax                  string `json:"uv_index_max"`
		UvIndexClearSkyMax          string `json:"uv_index_clear_sky_max"`
		PrecipitationSum            string `json:"precipitation_sum"`
		RainSum                     string `json:"rain_sum"`
		ShowersSum                  string `json:"showers_sum"`
		SnowfallSum                 string `json:"snowfall_sum"`
		PrecipitationHours          string `json:"precipitation_hours"`
		PrecipitationProbabilityMax string `json:"precipitation_probability_max"`
		Windspeed10MMax             string `json:"windspeed_10m_max"`
		Windgusts10MMax             string `json:"windgusts_10m_max"`
		Winddirection10MDominant    string `json:"winddirection_10m_dominant"`
		ShortwaveRadiationSum       string `json:"shortwave_radiation_sum"`
		Et0FaoEvapotranspiration    string `json:"et0_fao_evapotranspiration"`
	} `json:"daily_units"`
	Daily struct {
		Time                        []string  `json:"time"`
		Weathercode                 []int     `json:"weathercode"`
		Temperature2MMax            []float64 `json:"temperature_2m_max"`
		Temperature2MMin            []float64 `json:"temperature_2m_min"`
		ApparentTemperatureMax      []float64 `json:"apparent_temperature_max"`
		ApparentTemperatureMin      []float64 `json:"apparent_temperature_min"`
		Sunrise                     []string  `json:"sunrise"`
		Sunset                      []string  `json:"sunset"`
		UvIndexMax                  []float64 `json:"uv_index_max"`
		UvIndexClearSkyMax          []float64 `json:"uv_index_clear_sky_max"`
		PrecipitationSum            []float64 `json:"precipitation_sum"`
		RainSum                     []float64 `json:"rain_sum"`
		ShowersSum                  []float64 `json:"showers_sum"`
		SnowfallSum                 []float64 `json:"snowfall_sum"`
		PrecipitationHours          []float64 `json:"precipitation_hours"`
		PrecipitationProbabilityMax []int     `json:"precipitation_probability_max"`
		Windspeed10MMax             []float64 `json:"windspeed_10m_max"`
		Windgusts10MMax             []float64 `json:"windgusts_10m_max"`
		Winddirection10MDominant    []int     `json:"winddirection_10m_dominant"`
		ShortwaveRadiationSum       []float64 `json:"shortwave_radiation_sum"`
		Et0FaoEvapotranspiration    []float64 `json:"et0_fao_evapotranspiration"`
	} `json:"daily"`
}

func (h *Daily) String() string {
	bData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		log.Println("err JSONMarshal ", err)
	}
	return string(bData)
}
