package main

import (
	"encoding/xml"
	"os"
	"strings"
)

type annotation struct {
	AppInfo string `xml:"appinfo"`
}

type enumeration struct {
	Value      string     `xml:"value,attr"`
	Annotation annotation `xml:"annotation"`
}

type restriction struct {
	Base        string        `xml:"base,attr"`
	Enumeration []enumeration `xml:"enumeration"`
}

type simpleType struct {
	Name        string      `xml:"name,attr"`
	Restriction restriction `xml:"restriction"`
}

type gpaisXSD struct {
	Version            string       `xml:"version,attr"`
	TargetNamespace    string       `xml:"targetNamespace,attr"`
	ElementFormDefault string       `xml:"elementFormDefault,attr"`
	SimpleType         []simpleType `xml:"simpleType"`
}

func main() {
	xsd, err := os.ReadFile("./gpais-klasifikatoriai.xsd")
	if err != nil {
		panic(err)
	}

	m := gpaisXSD{}

	err = xml.Unmarshal(xsd, &m)
	if err != nil {
		panic(err)
	}

	for _, st := range m.SimpleType {
		f, err := os.Create(st.Name + ".csv")
		if err != nil {
			panic(err)
		}

		for _, e := range st.Restriction.Enumeration {
			var ws []string

			va := strings.Split(e.Value, ":")
			ai := strings.Split(e.Annotation.AppInfo, ":")

			ws = append(ws, va...)
			ws = append(ws, ai...)

			f.WriteString(strings.Join(ws, ",") + "\n")
		}

		f.Close()
	}

}
