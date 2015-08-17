package cap

import (
	"io/ioutil"
	"testing"
)

func getCAPAlertExample() (*Alert11, error) {
	xmlData, err := ioutil.ReadFile("../examples/nws_alert.xml")

	if err != nil {
		return nil, err
	}

	return ParseAlert11(xmlData)
}

func TestUnmarshalAlertHasProperValues(t *testing.T) {
	alert, err := getCAPAlertExample()

	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t,
		alert.MessageID,
		"NOAA-NWS-ALERTS-AR1253BA3B00A4.FloodWarning.1253BA3D4A94AR.LZKFLSLZK.342064b5a5aafb8265dfc3707d6a3b09",
		"MessageID does not match!")

	assertEqual(t,
		alert.SenderID,
		"w-nws.webmaster@noaa.gov",
		"SenderID does not match!")

	assertEqual(t,
		alert.SentDate,
		"2015-08-15T20:45:00-05:00",
		"SenderDate does not match!")

	assertEqual(t,
		alert.MessageStatus,
		"Actual",
		"MessageStatus does not match!")

	assertEqual(t,
		alert.MessageType,
		"Alert",
		"MessageType does not match!")

	assertEqual(t,
		alert.Scope,
		"Public",
		"Scope does not match!")

	assertEqual(t,
		alert.Note,
		"Alert for Jackson; Woodruff (Arkansas) Issued by the National Weather Service",
		"Note does not match!")

	assertEqual(t,
		len(alert.Infos),
		1,
		"One <info> should be present")
}

func TestUnmarshalAlertInfoHasProperValues(t *testing.T) {
	alert, err := getCAPAlertExample()

	if err != nil {
		t.Fatal(err)
	}

	var info = alert.Infos[0]
	assertEqual(t,
		info.EventCategory,
		"Met",
		"EventCategory does not match!")

	assertEqual(t,
		info.EventType,
		"Flood Warning",
		"EventType does not match!")

	assertEqual(t,
		info.Urgency,
		"Expected",
		"Urgency does not match!")

	assertEqual(t,
		info.Certainty,
		"Likely",
		"Certainty does not match!")

	assertEqual(t,
		info.EventCode[0].ValueName,
		"SAME",
		"EventCode-ValueName does not match!")

	assertEqual(t,
		info.EventCode[0].Value,
		"",
		"EventCode-Value does not match!")

	assertEqual(t,
		info.EffectiveDate,
		"2015-08-15T20:45:00-05:00",
		"EffectiveDate does not match!")

	assertEqual(t,
		info.ExpiresDate,
		"2015-08-16T11:45:00-05:00",
		"ExpiresDate does not match")

	assertEqual(t,
		info.SenderName,
		"NWS Little Rock (Arkansas)",
		"SenderName does not match")

	assertEqual(t,
		info.Headline,
		"Flood Warning issued August 15 at 8:45PM CDT until further notice by NWS Little Rock",
		"Headline does not match!")

	assertStartsWith(t,
		info.EventDescription,
		"...From the National Weather Service in Little Rock",
		"EventDescription does not match!")

	assertStartsWith(t,
		info.Instruction,
		"Safety message...",
		"Instruction does not match!")

	assertEqual(t,
		len(info.Parameters),
		4,
		"Number of Parameters does not match!")

	assertEqual(t,
		len(info.Areas),
		1,
		"Number of Areas does not match!")
}

func TestUnmarshalAlertInfoParameterHasProperValues(t *testing.T) {
	alert, err := getCAPAlertExample()

	if err != nil {
		t.Fatal(err)
	}

	var info = alert.Infos[0]

	assertEqual(t,
		info.Parameter("WMOHEADER"),
		"",
		"WMOHEADER does not match!")

	assertEqual(t,
		info.Parameter("UGC"),
		"ARC067-147",
		"UGC does not match!")

	assertStartsWith(t,
		info.Parameter("VTEC"),
		"/O.CON.KLZK.FL.W.0108.000000T0000Z-000000T0000Z/\n",
		"VTEC does not match!")

	assertEqual(t,
		info.Parameter("TIME...MOT...LOC"),
		"",
		"TIME...MOT...LOC does not match!")

	assertEqual(t,
		info.Parameter("TIME...MOT...LOC"),
		"",
		"TIME...MOT...LOC does not match!")

	assertEqual(t,
		info.Parameter("TIME...MOT...LOC"),
		"",
		"TIME...MOT...LOC does not match!")
}

func TestUnmarshalAlertInfoAreaHasProperValues(t *testing.T) {
	alert, err := getCAPAlertExample()

	if err != nil {
		t.Fatal(err)
	}

	var info = alert.Infos[0]
	var area = info.Areas[0]

	assertEqual(t,
		area.Description,
		"Jackson; Woodruff",
		"Description does not match!")

	assertEqual(t,
		area.Polygon,
		"35.1,-91.33 35.22,-91.28 35.39,-91.23 35.38,-91.13 35.21,-91.17 35.08,-91.22 35.1,-91.33",
		"Polygon does not match!")

	assertEqual(t,
		len(area.Geocodes),
		4,
		"Area does not have the proper number of Geocode elements")

	assertIn(t,
		"005067",
		area.GeocodeAll("FIPS6"),
		"Value not found in Geocode[FIPS6]!")

	assertIn(t,
		"005147",
		area.GeocodeAll("FIPS6"),
		"Value not found in Geocode[FIPS6]!")

	assertIn(t,
		"ARC067",
		area.GeocodeAll("UGC"),
		"Value not found in Geocode[UGC]!")

	assertIn(t,
		"ARC147",
		area.GeocodeAll("UGC"),
		"Value not found in Geocode[UGC]!")
}

func TestAddParameterToInfoSetsProperValue(t *testing.T) {
	parameterName := "testcode"
	parameterValue := "1234"
	var info Info

	assertEqual(t, len(info.Parameters), 0, "info.Parameters should be empty")

	info.AddParameter(parameterName, parameterValue)

	assertEqual(t, len(info.Parameters), 1, "info.Parameters should have len = 1")

	parameter := info.Parameters[0]
	assertEqual(t, parameter.ValueName, parameterName, "info.Parameters[0] does not have the correct name")
	assertEqual(t, parameter.Value, parameterValue, "info.Parameters[0] does not have the correct value")
}

func TestAddGeocodeToAreaSetsProperValue(t *testing.T) {
	geocodeName := "testcode"
	geocodeValue := "1234"
	var area Area

	assertEqual(t, len(area.Geocodes), 0, "area.Geocodes should be empty")

	area.AddGeocode(geocodeName, geocodeValue)

	assertEqual(t, len(area.Geocodes), 1, "area.Geocodes should have len = 1")

	geocode := area.Geocodes[0]
	assertEqual(t, geocode.ValueName, geocodeName, "area.Geocodes[0] does not have the correct name")
	assertEqual(t, geocode.Value, geocodeValue, "area.Geocodes[0] does not have the correct value")
}

func TestAreaGecodeReturnsFirstValue(t *testing.T) {
	geocode1 := NamedValue{"test-name", "1234"}
	geocode2 := NamedValue{"test-name", "5678"}

	var area Area

	area.AddGeocode(geocode1.ValueName, geocode1.Value)
	area.AddGeocode(geocode2.ValueName, geocode2.Value)

	assertEqual(t, len(area.Geocodes), 2, "area.Geocodes should have len = 2")

	geocodeValue := area.Geocode("test-name")
	assertEqual(t, geocodeValue, geocode1.Value, "Geocode does not have the correct name")
}

func TestAreaGecodeReturnsEmptyStringIfNotFound(t *testing.T) {
	geocode := NamedValue{"test-name", "1234"}

	var area Area

	area.AddGeocode(geocode.ValueName, geocode.Value)

	geocodeValue := area.Geocode("not-a-real-key")
	assertEqual(t, geocodeValue, "", "Geocode did not return an empty string")
}

func TestParseCAPDateReturnsCorrectValue(t *testing.T) {
	alert, err := getCAPAlertExample()

	if err != nil {
		t.Fatal(err)
	}

	info := alert.Infos[0]
	dt, _ := ParseCAPDate(info.EffectiveDate)
	_, zoneOffset := dt.Zone()
	zoneOffsetHours := zoneOffset / 3600

	// 2015-08-15T20:45:00-05:00
	assertEqual(t, 2015, int(dt.Year()), "Year does not equal expected value")
	assertEqual(t, 8, int(dt.Month()), "Month does not equal expected value")
	assertEqual(t, 15, int(dt.Day()), "Day does not equal expected value")
	assertEqual(t, 20, int(dt.Hour()), "Hour does not equal expected value")
	assertEqual(t, 45, int(dt.Minute()), "Minute does not equal expected value")
	assertEqual(t, 0, int(dt.Second()), "Second does not equal expected value")
	assertEqual(t, -5, zoneOffsetHours, "TZ offset does not equal expected value")
}

func TestParseAlertReturnsErrForInvalidXml(t *testing.T) {
	_, err := ParseAlert([]byte("invalid xml"))
	assertEqual(t, "EOF", err.Error(), "Unexpected or missing error message")
}

func TestParseAlert11ReturnsErrForInvalidXml(t *testing.T) {
	_, err := ParseAlert11([]byte("invalid xml"))
	assertEqual(t, "EOF", err.Error(), "Unexpected or missing error message")
}
