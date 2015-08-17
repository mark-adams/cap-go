package cap

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func getNwsAtomFeedExample() (*NWSAtomFeed, error) {
	xmlData, err := ioutil.ReadFile("../examples/nws_atom.xml")

	if err != nil {
		return nil, err
	}

	var feed NWSAtomFeed
	err = xml.Unmarshal(xmlData, &feed)

	if err != nil {
		return nil, err
	}

	return &feed, nil
}

func TestUnmarshalNWSAtomFeedHasProperValues(t *testing.T) {
	feed, err := getNwsAtomFeedExample()

	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t,
		feed.ID,
		"http://alerts.weather.gov/cap/us.atom",
		"Feed ID does not match!")

	assertEqual(t,
		feed.Logo,
		"http://alerts.weather.gov/images/xml_logo.gif",
		"Feed logo does not match!")

	assertEqual(t,
		feed.Generator,
		"NWS CAP Server",
		"Feed generator does not match!")

	assertEqual(t,
		feed.UpdatedDate,
		"2015-08-15T18:06:00-06:00",
		"Feed updated date does not match!")

	assertEqual(t,
		feed.Author.Name,
		"w-nws.webmaster@noaa.gov",
		"Feed author does not match!")

	assertEqual(t,
		feed.Title,
		"Current Watches, Warnings and Advisories for the United States Issued by the National Weather Service",
		"Feed title does not match!")

	assertEqual(t,
		feed.Link[0].Href,
		"http://alerts.weather.gov/cap/us.atom",
		"Feed link does not match!")

	assertEqual(t,
		len(feed.Entries),
		5,
		"Number of feed entries is not correct")
}

func TestUnmarshalNWSAtomFeedEntryHasProperValues(t *testing.T) {
	feed, err := getNwsAtomFeedExample()

	if err != nil {
		t.Fatal(err)
	}

	var entry = feed.Entries[0]

	assertEqual(t,
		entry.ID,
		"http://alerts.weather.gov/cap/wwacapget.php?x=AR1253BA2D9194.FloodWarning.1253BA3B7444AR.LZKFLSLZK.342064b5a5aafb8265dfc3707d6a3b09",
		"Entry ID does not match!")

	assertEqual(t,
		entry.UpdatedDate,
		"2015-08-15T08:41:00-05:00",
		"Entry update date does not match!")

	assertEqual(t,
		entry.PublishedDate,
		"2015-08-15T08:41:00-05:00",
		"Entry published date does not match!")

	assertEqual(t,
		entry.Author.Name,
		"w-nws.webmaster@noaa.gov",
		"Entry author name does not match!")

	assertEqual(t,
		entry.Title,
		"Flood Warning issued August 15 at 8:41AM CDT until further notice by NWS",
		"Entry title does not match!")

	assertEqual(t,
		entry.Link[0].Href,
		"http://alerts.weather.gov/cap/wwacapget.php?x=AR1253BA2D9194.FloodWarning.1253BA3B7444AR.LZKFLSLZK.342064b5a5aafb8265dfc3707d6a3b09",
		"Entry link does not match")

	assertEqual(t,
		entry.Summary,
		"...From the National Weather Service in Little Rock...the Flood Warning continues for the following river in Arkansas... White River At Augusta affecting White and Woodruff Counties Cache River Near Patterson affecting Jackson and Woodruff Counties River forecasts are based on current conditions and rainfall forecasted to occur over the next 12 hours. During periods of",
		"Entry summary does not match!")

	assertEqual(t,
		entry.EventType,
		"Flood Warning",
		"Entry event type does not match!")

	assertEqual(t,
		entry.EffectiveDate,
		"2015-08-15T08:41:00-05:00",
		"Entry effective date does not match!")

	assertEqual(t,
		entry.ExpiresDate,
		"2015-08-15T23:41:00-05:00",
		"Entry expires date does not match!")

	assertEqual(t,
		entry.MessageStatus,
		"Actual",
		"Entry message status does not match!")

	assertEqual(t,
		entry.MessageType,
		"Alert",
		"Entry message type does not match!")

	assertEqual(t,
		entry.EventCategory,
		"Met",
		"Entry event category does not match!")

	assertEqual(t,
		entry.Urgency,
		"Expected",
		"Entry urgencfy does not match!")

	assertEqual(t,
		entry.Severity,
		"Moderate",
		"Entry severity does not match!")

	assertEqual(t,
		entry.Certainty,
		"Likely",
		"Entry certainty does not match!")

	assertEqual(t,
		entry.AreaDescription,
		"Jackson; Woodruff",
		"Entry area description does not match!")

	assertEqual(t,
		entry.Polygon,
		"35.1,-91.33 35.22,-91.28 35.39,-91.23 35.38,-91.13 35.21,-91.17 35.08,-91.22 35.1,-91.33",
		"Entry polygon does not match!")

	assertEqual(t,
		len(entry.Geocode.Names),
		2,
		"Number of geocode names in entry does not match!")

	assertEqual(t,
		len(entry.Geocode.Values),
		2,
		"Number of geocode values in entry does not match!")
}

func TestUnmarshalNWSAtomFeedEntryGeocodeHasProperValues(t *testing.T) {
	feed, err := getNwsAtomFeedExample()

	if err != nil {
		t.Fatal(err)
	}

	var entry = feed.Entries[0]
	var geocode = entry.Geocode

	assertIn(t,
		"005067",
		geocode.GetValues("FIPS6"),
		"Value not found in Geocode[FIPS6]!")

	assertIn(t,
		"005147",
		geocode.GetValues("FIPS6"),
		"Value not found in Geocode[FIPS6]!")

	assertIn(t,
		"ARC067",
		geocode.GetValues("UGC"),
		"Value not found in Geocode[UGC]!")

	assertIn(t,
		"ARC147",
		geocode.GetValues("UGC"),
		"Value not found in Geocode[UGC]!")
}

func TestUnmarshalNWSAtomFeedEntryParameterHasProperValues(t *testing.T) {
	feed, err := getNwsAtomFeedExample()

	if err != nil {
		t.Fatal(err)
	}

	var entry = feed.Entries[0]

	assertStartsWith(t,
		entry.Parameter("VTEC"),
		"/O.CON.KLZK.FL.W.0108.000000T0000Z-000000T0000Z/\n",
		"VTEC parameter does not match!")
}

func TestNWSAtomGeocodeGetValuesReturnsEmptyArrIfNotFound(t *testing.T) {
	var geocode NWSAtomGeocode

	found := geocode.GetValues("not-a-real-key")

	assertEqual(t, len(found), 0, "No items should be found")
}

func TestLinkFollowAlertReturnsErrorForInvalidURL(t *testing.T) {
	link := Link{Href: "abcdef"}

	_, err := link.FollowAlert()

	assertEqual(t,
		"Get abcdef: unsupported protocol scheme \"\"",
		err.Error(),
		"Incorrect error was returned")
}

func TestHandleHttpResponseReturnsStartingErr(t *testing.T) {
	existingError := errors.New("Prexisting error!")

	_, err := handleHTTPResponse(nil, existingError)

	assertEqual(t, existingError, err, "The returned error was not the expected error")
}

func TestHandleHttpResponseReturnsErrOnNon200StatusCode(t *testing.T) {
	var response http.Response

	response.StatusCode = 400
	_, err := handleHTTPResponse(&response, nil)

	assertEqual(t, "Non-200 status code received from server: 400", err.Error(), "The returned error was not the expected error")
}

func TestHandleHttpResponseReturnsErrOnZeroContentLength(t *testing.T) {
	var response http.Response

	response.StatusCode = 200
	response.ContentLength = 0

	_, err := handleHTTPResponse(&response, nil)

	assertEqual(t, "No content was returned", err.Error(), "The returned error was not the expected error")
}

func TestHandleHttpResponseReturnsErrOnContentLengthTooLarge(t *testing.T) {
	var response http.Response

	response.StatusCode = 200
	response.ContentLength = MaxFeedSize + 1

	_, err := handleHTTPResponse(&response, nil)

	assertStartsWith(t,
		err.Error(),
		"Feed exceeds maximum size of",
		"The returned error was not the expected error")
}

// Integration test! Requires Internet access!
func TestGetNwsAtomFeed(t *testing.T) {
	feed, err := GetNWSAtomFeed()

	if err != nil {
		t.Fatal(err)
	}

	if feed.ID != "http://alerts.weather.gov/cap/us.atom" {
		t.Fatal("Feed did not download / parse correctly")
	}
}

// Integration test! Requires Internet access!
func TestGetNwsAtomFeedEntryFromLink(t *testing.T) {
	feed, err := GetNWSAtomFeed()

	if err != nil {
		t.Fatal(err)
	}

	if len(feed.Entries) == 0 {
		t.Skip("No alert entries in the Atom feed. Skipping...")
	}

	entry := feed.Entries[0]
	alert, err := entry.Link[0].FollowAlert()

	if err != nil {
		t.Fatal(err)
	}

	entryID, err := url.Parse(entry.ID)

	assertEqual(t,
		"NOAA-NWS-ALERTS-"+entryID.Query().Get("x"),
		alert.MessageID,
		"Entry ID and Message ID do not match!")

}
