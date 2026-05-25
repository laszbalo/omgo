package omgo

// CurrentMetric represents a metric that can be requested for current conditions.
type CurrentMetric string

// Current weather metrics available from the Open-Meteo API.
const (

	// NOTE: From Open Meteo docs: "Every weather variable available in hourly data, is available as current condition as well."
	// It seems that Hourly and Current both using the same set of metrics

	// Basic weather
	CurrentTemperature2m       CurrentMetric = "temperature_2m"
	CurrentRelativeHumidity2m  CurrentMetric = "relative_humidity_2m"
	CurrentDewPoint2m          CurrentMetric = "dew_point_2m"
	CurrentApparentTemperature CurrentMetric = "apparent_temperature"

	// Pressure
	CurrentPressureMSL     CurrentMetric = "pressure_msl"
	CurrentSurfacePressure CurrentMetric = "surface_pressure"

	// Cloud cover
	CurrentCloudCover     CurrentMetric = "cloud_cover"
	CurrentCloudCoverLow  CurrentMetric = "cloud_cover_low"
	CurrentCloudCoverMid  CurrentMetric = "cloud_cover_mid"
	CurrentCloudCoverHigh CurrentMetric = "cloud_cover_high"

	// Wind at different heights
	CurrentWindSpeed10m      CurrentMetric = "wind_speed_10m"
	CurrentWindSpeed80m      CurrentMetric = "wind_speed_80m"
	CurrentWindSpeed120m     CurrentMetric = "wind_speed_120m"
	CurrentWindSpeed180m     CurrentMetric = "wind_speed_180m"
	CurrentWindDirection10m  CurrentMetric = "wind_direction_10m"
	CurrentWindDirection80m  CurrentMetric = "wind_direction_80m"
	CurrentWindDirection120m CurrentMetric = "wind_direction_120m"
	CurrentWindDirection180m CurrentMetric = "wind_direction_180m"
	CurrentWindGusts10m      CurrentMetric = "wind_gusts_10m"

	// Solar radiation
	CurrentShortwaveRadiation     CurrentMetric = "shortwave_radiation"
	CurrentDirectRadiation        CurrentMetric = "direct_radiation"
	CurrentDirectNormalIrradiance CurrentMetric = "direct_normal_irradiance"
	CurrentDiffuseRadiation       CurrentMetric = "diffuse_radiation"
	CurrentGlobalTiltedIrradiance CurrentMetric = "global_tilted_irradiance"

	// Precipitation
	CurrentPrecipitation            CurrentMetric = "precipitation"
	CurrentSnowfall                 CurrentMetric = "snowfall"
	CurrentPrecipitationProbability CurrentMetric = "precipitation_probability"
	CurrentRain                     CurrentMetric = "rain"
	CurrentShowers                  CurrentMetric = "showers"

	// Other
	CurrentWeatherCode              CurrentMetric = "weather_code"
	CurrentSnowDepth                CurrentMetric = "snow_depth"
	CurrentFreezingLevelHeight      CurrentMetric = "freezing_level_height"
	CurrentVisibility               CurrentMetric = "visibility"
	CurrentEvapotranspiration       CurrentMetric = "evapotranspiration"
	CurrentET0FAOEvapotranspiration CurrentMetric = "et0_fao_evapotranspiration"
	CurrentVapourPressureDeficit    CurrentMetric = "vapour_pressure_deficit"
	CurrentCape                     CurrentMetric = "cape"
	CurrentSunshineDuration         CurrentMetric = "sunshine_duration"
	CurrentIsDay                    CurrentMetric = "is_day"

	// Soil temperature
	CurrentSoilTemperature0cm  CurrentMetric = "soil_temperature_0cm"
	CurrentSoilTemperature6cm  CurrentMetric = "soil_temperature_6cm"
	CurrentSoilTemperature18cm CurrentMetric = "soil_temperature_18cm"
	CurrentSoilTemperature54cm CurrentMetric = "soil_temperature_54cm"

	// Soil moisture
	CurrentSoilMoisture0to1cm   CurrentMetric = "soil_moisture_0_to_1cm"
	CurrentSoilMoisture1to3cm   CurrentMetric = "soil_moisture_1_to_3cm"
	CurrentSoilMoisture3to9cm   CurrentMetric = "soil_moisture_3_to_9cm"
	CurrentSoilMoisture9to27cm  CurrentMetric = "soil_moisture_9_to_27cm"
	CurrentSoilMoisture27to81cm CurrentMetric = "soil_moisture_27_to_81cm"

	// Pressure level: Temperature
	CurrentTemperature1000hPa CurrentMetric = "temperature_1000hPa"
	CurrentTemperature975hPa  CurrentMetric = "temperature_975hPa"
	CurrentTemperature950hPa  CurrentMetric = "temperature_950hPa"
	CurrentTemperature925hPa  CurrentMetric = "temperature_925hPa"
	CurrentTemperature900hPa  CurrentMetric = "temperature_900hPa"
	CurrentTemperature850hPa  CurrentMetric = "temperature_850hPa"
	CurrentTemperature800hPa  CurrentMetric = "temperature_800hPa"
	CurrentTemperature700hPa  CurrentMetric = "temperature_700hPa"
	CurrentTemperature600hPa  CurrentMetric = "temperature_600hPa"
	CurrentTemperature500hPa  CurrentMetric = "temperature_500hPa"
	CurrentTemperature400hPa  CurrentMetric = "temperature_400hPa"
	CurrentTemperature300hPa  CurrentMetric = "temperature_300hPa"
	CurrentTemperature250hPa  CurrentMetric = "temperature_250hPa"
	CurrentTemperature200hPa  CurrentMetric = "temperature_200hPa"
	CurrentTemperature150hPa  CurrentMetric = "temperature_150hPa"
	CurrentTemperature100hPa  CurrentMetric = "temperature_100hPa"
	CurrentTemperature70hPa   CurrentMetric = "temperature_70hPa"
	CurrentTemperature50hPa   CurrentMetric = "temperature_50hPa"
	CurrentTemperature30hPa   CurrentMetric = "temperature_30hPa"

	// Pressure level: Relative Humidity
	CurrentRelativeHumidity1000hPa CurrentMetric = "relative_humidity_1000hPa"
	CurrentRelativeHumidity975hPa  CurrentMetric = "relative_humidity_975hPa"
	CurrentRelativeHumidity950hPa  CurrentMetric = "relative_humidity_950hPa"
	CurrentRelativeHumidity925hPa  CurrentMetric = "relative_humidity_925hPa"
	CurrentRelativeHumidity900hPa  CurrentMetric = "relative_humidity_900hPa"
	CurrentRelativeHumidity850hPa  CurrentMetric = "relative_humidity_850hPa"
	CurrentRelativeHumidity800hPa  CurrentMetric = "relative_humidity_800hPa"
	CurrentRelativeHumidity700hPa  CurrentMetric = "relative_humidity_700hPa"
	CurrentRelativeHumidity600hPa  CurrentMetric = "relative_humidity_600hPa"
	CurrentRelativeHumidity500hPa  CurrentMetric = "relative_humidity_500hPa"
	CurrentRelativeHumidity400hPa  CurrentMetric = "relative_humidity_400hPa"
	CurrentRelativeHumidity300hPa  CurrentMetric = "relative_humidity_300hPa"
	CurrentRelativeHumidity250hPa  CurrentMetric = "relative_humidity_250hPa"
	CurrentRelativeHumidity200hPa  CurrentMetric = "relative_humidity_200hPa"
	CurrentRelativeHumidity150hPa  CurrentMetric = "relative_humidity_150hPa"
	CurrentRelativeHumidity100hPa  CurrentMetric = "relative_humidity_100hPa"
	CurrentRelativeHumidity70hPa   CurrentMetric = "relative_humidity_70hPa"
	CurrentRelativeHumidity50hPa   CurrentMetric = "relative_humidity_50hPa"
	CurrentRelativeHumidity30hPa   CurrentMetric = "relative_humidity_30hPa"

	// Pressure level: Dew Point
	CurrentDewPoint1000hPa CurrentMetric = "dew_point_1000hPa"
	CurrentDewPoint975hPa  CurrentMetric = "dew_point_975hPa"
	CurrentDewPoint950hPa  CurrentMetric = "dew_point_950hPa"
	CurrentDewPoint925hPa  CurrentMetric = "dew_point_925hPa"
	CurrentDewPoint900hPa  CurrentMetric = "dew_point_900hPa"
	CurrentDewPoint850hPa  CurrentMetric = "dew_point_850hPa"
	CurrentDewPoint800hPa  CurrentMetric = "dew_point_800hPa"
	CurrentDewPoint700hPa  CurrentMetric = "dew_point_700hPa"
	CurrentDewPoint600hPa  CurrentMetric = "dew_point_600hPa"
	CurrentDewPoint500hPa  CurrentMetric = "dew_point_500hPa"
	CurrentDewPoint400hPa  CurrentMetric = "dew_point_400hPa"
	CurrentDewPoint300hPa  CurrentMetric = "dew_point_300hPa"
	CurrentDewPoint250hPa  CurrentMetric = "dew_point_250hPa"
	CurrentDewPoint200hPa  CurrentMetric = "dew_point_200hPa"
	CurrentDewPoint150hPa  CurrentMetric = "dew_point_150hPa"
	CurrentDewPoint100hPa  CurrentMetric = "dew_point_100hPa"
	CurrentDewPoint70hPa   CurrentMetric = "dew_point_70hPa"
	CurrentDewPoint50hPa   CurrentMetric = "dew_point_50hPa"
	CurrentDewPoint30hPa   CurrentMetric = "dew_point_30hPa"

	// Pressure level: Cloud Cover
	CurrentCloudCover1000hPa CurrentMetric = "cloud_cover_1000hPa"
	CurrentCloudCover975hPa  CurrentMetric = "cloud_cover_975hPa"
	CurrentCloudCover950hPa  CurrentMetric = "cloud_cover_950hPa"
	CurrentCloudCover925hPa  CurrentMetric = "cloud_cover_925hPa"
	CurrentCloudCover900hPa  CurrentMetric = "cloud_cover_900hPa"
	CurrentCloudCover850hPa  CurrentMetric = "cloud_cover_850hPa"
	CurrentCloudCover800hPa  CurrentMetric = "cloud_cover_800hPa"
	CurrentCloudCover700hPa  CurrentMetric = "cloud_cover_700hPa"
	CurrentCloudCover600hPa  CurrentMetric = "cloud_cover_600hPa"
	CurrentCloudCover500hPa  CurrentMetric = "cloud_cover_500hPa"
	CurrentCloudCover400hPa  CurrentMetric = "cloud_cover_400hPa"
	CurrentCloudCover300hPa  CurrentMetric = "cloud_cover_300hPa"
	CurrentCloudCover250hPa  CurrentMetric = "cloud_cover_250hPa"
	CurrentCloudCover200hPa  CurrentMetric = "cloud_cover_200hPa"
	CurrentCloudCover150hPa  CurrentMetric = "cloud_cover_150hPa"
	CurrentCloudCover100hPa  CurrentMetric = "cloud_cover_100hPa"
	CurrentCloudCover70hPa   CurrentMetric = "cloud_cover_70hPa"
	CurrentCloudCover50hPa   CurrentMetric = "cloud_cover_50hPa"
	CurrentCloudCover30hPa   CurrentMetric = "cloud_cover_30hPa"

	// Pressure level: Wind Speed
	CurrentWindSpeed1000hPa CurrentMetric = "wind_speed_1000hPa"
	CurrentWindSpeed975hPa  CurrentMetric = "wind_speed_975hPa"
	CurrentWindSpeed950hPa  CurrentMetric = "wind_speed_950hPa"
	CurrentWindSpeed925hPa  CurrentMetric = "wind_speed_925hPa"
	CurrentWindSpeed900hPa  CurrentMetric = "wind_speed_900hPa"
	CurrentWindSpeed850hPa  CurrentMetric = "wind_speed_850hPa"
	CurrentWindSpeed800hPa  CurrentMetric = "wind_speed_800hPa"
	CurrentWindSpeed700hPa  CurrentMetric = "wind_speed_700hPa"
	CurrentWindSpeed600hPa  CurrentMetric = "wind_speed_600hPa"
	CurrentWindSpeed500hPa  CurrentMetric = "wind_speed_500hPa"
	CurrentWindSpeed400hPa  CurrentMetric = "wind_speed_400hPa"
	CurrentWindSpeed300hPa  CurrentMetric = "wind_speed_300hPa"
	CurrentWindSpeed250hPa  CurrentMetric = "wind_speed_250hPa"
	CurrentWindSpeed200hPa  CurrentMetric = "wind_speed_200hPa"
	CurrentWindSpeed150hPa  CurrentMetric = "wind_speed_150hPa"
	CurrentWindSpeed100hPa  CurrentMetric = "wind_speed_100hPa"
	CurrentWindSpeed70hPa   CurrentMetric = "wind_speed_70hPa"
	CurrentWindSpeed50hPa   CurrentMetric = "wind_speed_50hPa"
	CurrentWindSpeed30hPa   CurrentMetric = "wind_speed_30hPa"

	// Pressure level: Wind Direction
	CurrentWindDirection1000hPa CurrentMetric = "wind_direction_1000hPa"
	CurrentWindDirection975hPa  CurrentMetric = "wind_direction_975hPa"
	CurrentWindDirection950hPa  CurrentMetric = "wind_direction_950hPa"
	CurrentWindDirection925hPa  CurrentMetric = "wind_direction_925hPa"
	CurrentWindDirection900hPa  CurrentMetric = "wind_direction_900hPa"
	CurrentWindDirection850hPa  CurrentMetric = "wind_direction_850hPa"
	CurrentWindDirection800hPa  CurrentMetric = "wind_direction_800hPa"
	CurrentWindDirection700hPa  CurrentMetric = "wind_direction_700hPa"
	CurrentWindDirection600hPa  CurrentMetric = "wind_direction_600hPa"
	CurrentWindDirection500hPa  CurrentMetric = "wind_direction_500hPa"
	CurrentWindDirection400hPa  CurrentMetric = "wind_direction_400hPa"
	CurrentWindDirection300hPa  CurrentMetric = "wind_direction_300hPa"
	CurrentWindDirection250hPa  CurrentMetric = "wind_direction_250hPa"
	CurrentWindDirection200hPa  CurrentMetric = "wind_direction_200hPa"
	CurrentWindDirection150hPa  CurrentMetric = "wind_direction_150hPa"
	CurrentWindDirection100hPa  CurrentMetric = "wind_direction_100hPa"
	CurrentWindDirection70hPa   CurrentMetric = "wind_direction_70hPa"
	CurrentWindDirection50hPa   CurrentMetric = "wind_direction_50hPa"
	CurrentWindDirection30hPa   CurrentMetric = "wind_direction_30hPa"

	// Pressure level: Geopotential Height
	CurrentGeopotentialHeight1000hPa CurrentMetric = "geopotential_height_1000hPa"
	CurrentGeopotentialHeight975hPa  CurrentMetric = "geopotential_height_975hPa"
	CurrentGeopotentialHeight950hPa  CurrentMetric = "geopotential_height_950hPa"
	CurrentGeopotentialHeight925hPa  CurrentMetric = "geopotential_height_925hPa"
	CurrentGeopotentialHeight900hPa  CurrentMetric = "geopotential_height_900hPa"
	CurrentGeopotentialHeight850hPa  CurrentMetric = "geopotential_height_850hPa"
	CurrentGeopotentialHeight800hPa  CurrentMetric = "geopotential_height_800hPa"
	CurrentGeopotentialHeight700hPa  CurrentMetric = "geopotential_height_700hPa"
	CurrentGeopotentialHeight600hPa  CurrentMetric = "geopotential_height_600hPa"
	CurrentGeopotentialHeight500hPa  CurrentMetric = "geopotential_height_500hPa"
	CurrentGeopotentialHeight400hPa  CurrentMetric = "geopotential_height_400hPa"
	CurrentGeopotentialHeight300hPa  CurrentMetric = "geopotential_height_300hPa"
	CurrentGeopotentialHeight250hPa  CurrentMetric = "geopotential_height_250hPa"
	CurrentGeopotentialHeight200hPa  CurrentMetric = "geopotential_height_200hPa"
	CurrentGeopotentialHeight150hPa  CurrentMetric = "geopotential_height_150hPa"
	CurrentGeopotentialHeight100hPa  CurrentMetric = "geopotential_height_100hPa"
	CurrentGeopotentialHeight70hPa   CurrentMetric = "geopotential_height_70hPa"
	CurrentGeopotentialHeight50hPa   CurrentMetric = "geopotential_height_50hPa"
	CurrentGeopotentialHeight30hPa   CurrentMetric = "geopotential_height_30hPa"

	// UVIndex
	CurrentUVIndex                   CurrentMetric = "uv_index"
	CurrentUVIndexClearSky           CurrentMetric = "uv_index_clear_sky"

)

// String returns the API parameter string for the metric.
func (m CurrentMetric) String() string {
	return string(m)
}
