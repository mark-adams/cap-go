package cap

import "encoding/xml"

// Alert provides basic information about the current message: its purpose, its source and its status
type Alert struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:emergency:cap:1.2 alert"`

	MessageID     string   `xml:"identifier"`
	SenderID      string   `xml:"sender"`
	SentDate      string   `xml:"sent"`
	MessageStatus string   `xml:"status"`
	MessageType   string   `xml:"msgType"`
	Source        string   `xml:"source,omitempty"`
	Scope         string   `xml:"scope"`
	Restriction   string   `xml:"restriction,omitempty"`
	Addresses     string   `xml:"addresses,omitempty"`
	HandlingCode  string   `xml:"code,omitempty"`
	Note          string   `xml:"note,omitempty"`
	ReferenceIDs  []string `xml:"references,omitempty"`
	IncidentIDs   []string `xml:"incidents,omitempty"`
	Infos         []Info   `xml:"info,omitempty"`
}

// Alert11 is the same as Alert but using the CAP 1.1 namespace
type Alert11 struct {
	Alert
	XMLName xml.Name `xml:"urn:oasis:names:tc:emergency:cap:1.1 alert"`
}

// Info describes an anticipated or actual event
type Info struct {
	XMLName xml.Name `xml:"info"`

	Language         string       `xml:"language,omitempty"`
	EventCategory    string       `xml:"category"`
	EventType        string       `xml:"event"`
	ResponseType     string       `xml:"responseType,omitempty"`
	Urgency          string       `xml:"urgency"`
	Severity         string       `xml:"severity"`
	Certainty        string       `xml:"certainty"`
	Audience         string       `xml:"audience,omitempty"`
	EventCode        []NamedValue `xml:"eventCode,omitempty"`
	EffectiveDate    string       `xml:"effective,omitempty"`
	ExpiresDate      string       `xml:"expires,omitempty"`
	OnsetDate        string       `xml:"onset,omitempty"`
	SenderName       string       `xml:"senderName,omitempty"`
	Headline         string       `xml:"headline,omitempty"`
	EventDescription string       `xml:"description,omitempty"`
	Instruction      string       `xml:"instruction,omitempty"`
	InformationURL   string       `xml:"web,omitempty"`
	ContactInfo      string       `xml:"contact,omitempty"`
	Parameters       []NamedValue `xml:"parameter,omitempty"`
	Areas            []Area       `xml:"area,omitempty"`
	Resources        []Resource   `xml:"resource,omitempty"`
}

// Resource provides an optional reference to additional information related to Info
type Resource struct {
	XMLName xml.Name `xml:"resource"`

	Description      string `xml:"resourceDesc"`
	MIMEType         string `xml:"mimeType,omitempty"`
	FileSize         string `xml:"size,omitempty"`
	URI              string `xml:"uri,omitempty"`
	DeereferencedURI string `xml:"derefUri,omitempty"`
	Digest           string `xml:"digest,omitempty"`
}

// Area describes a geographic area to which the Info segment applies
type Area struct {
	XMLName xml.Name `xml:"area"`

	Description string       `xml:"areaDesc"`
	Polygon     string       `xml:"polygon,omitempty"`
	Circle      string       `xml:"circle,omitempty"`
	Geocodes    []NamedValue `xml:"geocode,omitempty"`
	Altitude    string       `xml:"altitude,omitempty"`
	Ceiling     string       `xml:"ceiling,omitempty"`
}

// NamedValue contains a name and a value associated with that name
type NamedValue struct {
	ValueName string `xml:"valueName"`
	Value     string `xml:"value"`
}

// search checks a slice of NamedValues for the first value with a specific name
func search(nva *[]NamedValue, name string) string {
	for _, element := range *nva {
		if element.ValueName == name {
			return element.Value
		}
	}

	return ""
}

// searchAll checks a slice of NamedValues for all values with the specified name
func searchAll(nva *[]NamedValue, name string) []string {
	var found = make([]string, 0, len(*nva))

	for _, element := range *nva {
		if element.ValueName == name {
			found = append(found, element.Value)
		}
	}

	return found
}

// Parameter returns back the value for the first parameter with the specified name
func (info *Info) Parameter(name string) string {
	return search(&info.Parameters, name)
}

// AddParameter adds a Parameter with the specified name and value
func (info *Info) AddParameter(name string, value string) {
	param := NamedValue{ValueName: name, Value: value}
	info.Parameters = append(info.Parameters, param)
}

// Geocode returns back the value for the first Geocode value with the specified name
func (a *Area) Geocode(name string) string {
	return search(&a.Geocodes, name)
}

// GeocodeAll returns back the Geocode values with the specified name
func (a *Area) GeocodeAll(name string) []string {
	return searchAll(&a.Geocodes, name)
}

// AddGeocode adds a Geocode with the specified name and value
func (a *Area) AddGeocode(name string, value string) {
	geocode := NamedValue{ValueName: name, Value: value}
	a.Geocodes = append(a.Geocodes, geocode)
}
