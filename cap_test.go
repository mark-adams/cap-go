package cap

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/kr/pretty"
)

func assertEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Error(message)
		for _, desc := range pretty.Diff(expected, actual) {
			t.Error(desc)
		}
	}
}

func TestUnmarshalAlertHasProperValues(t *testing.T) {
	xmlData, err := ioutil.ReadFile("examples/nws_alert.xml")

	if err != nil {
		t.Error(err)
	}

	var testAlert Alert
	xml.Unmarshal(xmlData, &testAlert)

	assertEqual(t,
		testAlert.MessageID,
		"NOAA-NWS-ALERTS-AR1253BA3B00A4.FloodWarning.1253BA3D4A94AR.LZKFLSLZK.342064b5a5aafb8265dfc3707d6a3b09",
		"MessageID does not match!")

	assertEqual(t,
		testAlert.SenderID,
		"w-nws.webmaster@noaa.gov",
		"SenderID does not match!")

	assertEqual(t,
		testAlert.SentDate,
		"2015-08-15T20:45:00-05:00",
		"SenderDate does not match!")

	assertEqual(t,
		testAlert.MessageStatus,
		"Actual",
		"MessageStatus does not match!")

	assertEqual(t,
		testAlert.MessageType,
		"Alert",
		"MessageType does not match!")

	assertEqual(t,
		testAlert.Scope,
		"Public",
		"Scope does not match!")

	assertEqual(t,
		testAlert.Note,
		"Alert for Jackson; Woodruff (Arkansas) Issued by the National Weather Service",
		"Note does not match!")

	assertEqual(t,
		len(testAlert.Infos),
		1,
		"One <info> should be present")
}
