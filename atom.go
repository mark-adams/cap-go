package cap

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// NwsNationalAtomFeedURL is the URL for the NWS National Atom feed
const NwsNationalAtomFeedURL string = "https://alerts.weather.gov/cap/us.php?x=1"

// MaxFeedSize is the maximum size for a downloaded feed
const MaxFeedSize int64 = 1024 * 1024 * 5

// GetNWSAtomFeed retrieves the main National Weather Service CAP v1.1 ATOM feed
func GetNWSAtomFeed() (feed *NWSAtomFeed, err error) {
	response, err := http.Get(NwsNationalAtomFeedURL)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Non-200 status code received from server: %d", 200)
	}

	if response.ContentLength == 0 {
		return nil, fmt.Errorf("No content was returned")
	}

	if response.ContentLength > MaxFeedSize {
		return nil, fmt.Errorf("Feed exceeds maximum size of %d bytes", MaxFeedSize)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var downloadedFeed NWSAtomFeed
	err = xml.Unmarshal(body, &downloadedFeed)

	if err != nil {
		return nil, err
	}

	return &downloadedFeed, nil
}

// NWSAtomFeed represents a AtomFeed of CAP alerts from the National Weather Service
type NWSAtomFeed struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`

	ID          string         `xml:"id"`
	Logo        string         `xml:"logo"`
	Generator   string         `xml:"generator"`
	UpdatedDate string         `xml:"updated"`
	Author      Author         `xml:"author"`
	Title       string         `xml:"title"`
	Link        []Link         `xml:"link"`
	Entries     []NWSAtomEntry `xml:"entry"`
}

// Link represents a link related to the parent entity
type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr"`
}

// NWSAtomEntry represents an entry on a NWSAtomFeed
type NWSAtomEntry struct {
	XMLName xml.Name `xml:"entry"`

	ID              string         `xml:"id"`
	UpdatedDate     string         `xml:"updated"`
	PublishedDate   string         `xml:"published"`
	Author          Author         `xml:"author"`
	Title           string         `xml:"title"`
	Link            []Link         `xml:"link"`
	Summary         string         `xml:"summary"`
	EventType       string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 event"`
	EffectiveDate   string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 effective,omitempty"`
	ExpiresDate     string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 expires,omitempty"`
	MessageStatus   string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 status"`
	MessageType     string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 msgType"`
	EventCategory   string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 category"`
	Urgency         string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 urgency"`
	Severity        string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 severity"`
	Certainty       string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 certainty"`
	AreaDescription string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 areaDesc"`
	Polygon         string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 polygon,omitempty"`
	Circle          string         `xml:"urn:oasis:names:tc:emergency:cap:1.1 circle,omitempty"`
	Geocode         NWSAtomGeocode `xml:"urn:oasis:names:tc:emergency:cap:1.1 geocode,omitempty"`
	Parameters      []NamedValue   `xml:"urn:oasis:names:tc:emergency:cap:1.1 parameter,omitempty"`
}

// NWSAtomGeocode describes the special version of Geocode elements used by the NWS Atom feed
//
// See note in GetValues for more information
type NWSAtomGeocode struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:emergency:cap:1.1 geocode,omitempty"`

	Names  []string `xml:"valueName,omitempty"`
	Values []string `xml:"value,omitempty"`
}

// Author represents the author of an NWSAtomFeed
type Author struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name"`
}

// GetValues returns back an array of values for the Geocode element with the same name
//
// Unfortunately, the NWS Atom format deviates from the normal CAP Geocode element
// and does not create a new Geocode element for each name / value pair.
// Instead, multiple name / value pairs are listed in order inside a single
// geocode tag. As a result, we store both a list of names and list of values
// that have been unmarshalled from the XML and retrieve them by searching
// the Names list and returning the corresponding item from the Values list
// at the same index
func (g *NWSAtomGeocode) GetValues(name string) []string {
	for index, value := range g.Names {
		if value == name {
			return strings.Split(g.Values[index], " ")
		}
	}

	return []string{}
}

// Parameter returns back the value for the first parameter with the specified name
func (ae *NWSAtomEntry) Parameter(name string) string {
	return search(&ae.Parameters, name)
}

// Follow retrieves the resource that the link's href attribute points to
func (l *Link) Follow() (*http.Response, error) {
	return http.Get(l.Href)
}

// FollowAlert retrieves the Alert that the link's href attribute points to
func (l *Link) FollowAlert() (*Alert11, error) {
	response, err := http.Get(l.Href)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Non-200 status code received from server: %d", 200)
	}

	if response.ContentLength == 0 {
		return nil, fmt.Errorf("No content was returned")
	}

	if response.ContentLength > MaxFeedSize {
		return nil, fmt.Errorf("Feed exceeds maximum size of %d bytes", MaxFeedSize)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var alert Alert11
	err = xml.Unmarshal(body, &alert)

	if err != nil {
		return nil, err
	}

	return &alert, nil
}
