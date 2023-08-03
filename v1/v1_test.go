package meteo

import (
	"testing"
)

var c = NewClient(nil)

func TestGetTemperature2m(t *testing.T) {
	temp, err := c.GetTemperature2m("21.02", "105.84")
	if err != nil {
		t.FailNow()
	}

	if temp.Latitude != 21.0 ||
		temp.Longitude != 105.875 ||
		temp.Timezone != "GMT" ||
		len(temp.Hourly.Time) < 1 || len(temp.HourlyUnits.Temperature2M) < 1 {
		t.FailNow()
	}
}
