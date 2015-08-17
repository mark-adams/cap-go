# cap-go
[![Build Status](https://travis-ci.org/mark-adams/cap-go.svg?branch=master)](https://travis-ci.org/mark-adams/cap-go)
[![Coverage Status](https://coveralls.io/repos/mark-adams/cap-go/badge.svg?branch=master&service=github)](https://coveralls.io/github/mark-adams/cap-go?branch=master)

A Go library for interacting with Common Alerting Protocol messages

The Common Alerting Protocol (CAP) is a simple but general format for exchanging all-hazard emergency alerts and public warnings over all kinds of networks.

The CAP specification can be found on the [OASIS website](http://docs.oasis-open.org/emergency/cap/v1.2/CAP-v1.2-os.html).

## Examples

### Retrieving the CAP feed from NWS

```go
package main

import (
    "fmt"
    "github.com/mark-adams/cap-go/cap"
)

func main(){
    feed, err := GetNWSAtomFeed()

    if err != nil {
        fmt.Errorf(err.Error())
        os.Exit(1)
    }

    fmt.Println(feed.ID)
}
```

### Parsing a CAP alert

```go
package main

import (
    "fmt"
    "github.com/mark-adams/cap-go/cap"
)

func main(){
    xmlData := []byte(`<?xml version='1.0' encoding='UTF-8' standalone='yes'?>
        <alert xmlns='urn:oasis:names:tc:emergency:cap:1.2'>
            <identifier>NOAA-NWS-ALERTS-AR1253BA3B00A4.FloodWarning.1253BA3D4A94AR.LZKFLSLZK.342064b5a5aafb8265dfc3707d6a3b09</identifier>
            <sender>w-nws.webmaster@noaa.gov</sender>
            <sent>2015-08-15T20:45:00-05:00</sent>
            <status>Actual</status>
            <msgType>Alert</msgType>
            <scope>Public</scope>
            <info>
                <category>Met</category>
                <event>Flood Warning</event>
                <urgency>Expected</urgency>
                <severity>Moderate</severity>
                <certainty>Likely</certainty>
            </info>
        </alert>`)

    alert, err := ParseAlert(xmlData)

    if err != nil {
        fmt.Errorf(err.Error())
        os.Exit(1)
    }

    fmt.Println(alert.MessageID)
    fmt.Println(alert.Infos[0].EventType)
}

```
