package cap

import "encoding/xml"

// Alert provides basic information about the current message: its purpose, its source and its status
type Alert struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:emergency:cap:1.1 alert"`

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

// Info describes an anticipated or actual event
type Info struct {
	XMLName          xml.Name     `xml:"info"`
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
	ExpirationDate   string       `xml:"onset,omitempty"`
	OnsetDate        string       `xml:"expires,omitempty"`
	SenderName       string       `xml:"senderName,omitempty"`
	Headline         string       `xml:"headline,omitempty"`
	EventDescription string       `xml:"description,omitempty"`
	Instructions     string       `xml:"instruction,omitempty"`
	InformationURL   string       `xml:"web,omitempty"`
	ContactInfo      string       `xml:"contact,omitempty"`
	Parameter        []NamedValue `xml:"parameter,omitempty"`
	Areas            []Area       `xml:"area,omitempty"`
	Resources        []Resource   `xml:"resource,omitempty"`
}

// Resource provides an optional reference to additional information related to Info
type Resource struct {
	Description      string `xml:"resourceDesc"`
	MIMEType         string `xml:"mimeType,omitempty"`
	FileSize         string `xml:"size,omitempty"`
	URI              string `xml:"uri,omitempty"`
	DeereferencedURI string `xml:"derefUri,omitempty"`
	Digest           string `xml:"digest,omitempty"`
}

// Area describes a geographic area to which the Info segment applies
type Area struct {
	Description string `xml:"areaDesc"`
	Polygon     string `xml:"polygon,omitempty"`
	Circle      string `xml:"circle,omitempty"`
	Geocode     string `xml:"geocode,omitempty"`
	Altitude    string `xml:"altitude,omitempty"`
	Ceiling     string `xml:"ceiling,omitempty"`
}

// NamedValue contains a name and a value associated with that name
type NamedValue struct {
	ValueName string `xml:"valueName"`
	Value     string `xml:"value"`
}
