package cap

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func getCAPAlertExample() (*Alert11, error) {
	xmlData, err := ioutil.ReadFile("examples/nws_alert.xml")

	if err != nil {
		return nil, err
	}

	var alert Alert11

	err = xml.Unmarshal(xmlData, &alert)

	if err != nil {
		return nil, err
	}

	return &alert, nil
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
