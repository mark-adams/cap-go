package cap

import "encoding/xml"
import "strings"

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
