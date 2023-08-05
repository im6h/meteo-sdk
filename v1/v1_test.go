package meteo

import (
	"fmt"
	"testing"
)

var c = NewClient(nil)

func TestGetHourly(t *testing.T) {
	temp, err := c.GetHourlyData("21.02", "105.84")

	if err != nil {
		t.FailNow()
	}

	if temp.Latitude != 21.0 ||
		temp.Longitude != 105.875 ||
		temp.Timezone != "GMT" ||
		len(temp.Hourly.Time) < 1 {
		t.FailNow()
	}

	fmt.Println(temp.String())
}

func TestGetDaily(t *testing.T) {
	temp, err := c.GetDailyData("21.02", "105.84", "GMT")
	if err != nil {
		t.FailNow()
	}

	if temp.Latitude != 21.0 ||
		temp.Longitude != 105.875 ||
		temp.Timezone != "GMT" ||
		len(temp.Daily.Time) < 1 || len(temp.Daily.Weathercode) < 1 {
		t.FailNow()
	}
}
