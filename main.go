// From http://blog.davidsingleton.org/parsing-huge-xml-files-with-go/
package main

import (
	"encoding/xml"
	"os"
	"log"
	"fmt"
)

type ProteinDatabase struct {
	Keyword                                  []string     `xml:"ProteinEntry>keywords>keyword"`
	Author                                   [][]string   `xml:"ProteinEntry>reference>refinfo>authors>author"`
	Year                                     []string     `xml:"ProteinEntry>reference>refinfo>year"`
	Db                                       []string     `xml:"ProteinEntry>reference>refinfo>xrefs>xref>db"`
	SeqSpecAccinfoReferenceProteinEntry      []string     `xml:"ProteinEntry>reference>accinfo>seq-spec"`
	MolType                                  []string     `xml:"ProteinEntry>reference>accinfo>mol-type"`
	Note                                     []string     `xml:"ProteinEntry>reference>note"`
	UidHeaderProteinEntry                    string       `xml:"ProteinEntry>header>uid"`
	Release                                  string       `xml:"release,attr"`
	Database                                 string       `xml:"Database"`
	Feature                                  []Feature    `xml:"ProteinEntry>feature"`
	Status                                   []string     `xml:"ProteinEntry>feature>status"`
	ProteinEntry                             ProteinEntry `xml:"ProteinEntry"`
	Name                                     string       `xml:"ProteinEntry>protein>name"`
	Formal                                   string       `xml:"ProteinEntry>organism>formal"`
	Pages                                    []string     `xml:"ProteinEntry>reference>refinfo>pages"`
	Accinfo                                  []Accinfo    `xml:"ProteinEntry>reference>accinfo"`
	TxtRev_date                              string       `xml:"ProteinEntry>header>txt-rev_date"`
	SeqRev_date                              string       `xml:"ProteinEntry>header>seq-rev_date"`
	FeatureType                              []string     `xml:"ProteinEntry>feature>feature-type"`
	Refinfo                                  []Refinfo    `xml:"ProteinEntry>reference>refinfo"`
	Title                                    []string     `xml:"ProteinEntry>reference>refinfo>title"`
	StatusAccinfoReferenceProteinEntry       []string     `xml:"ProteinEntry>reference>accinfo>status"`
	Sequence                                 string       `xml:"ProteinEntry>sequence"`
	Type                                     string       `xml:"ProteinEntry>summary>type"`
	Common                                   string       `xml:"ProteinEntry>organism>common"`
	Contents                                 []string     `xml:"ProteinEntry>reference>contents"`
	Introns                                  string       `xml:"ProteinEntry>genetics>introns"`
	Description                              []string     `xml:"ProteinEntry>feature>description"`
	Source                                   string       `xml:"ProteinEntry>organism>source"`
	DbXrefXrefsAccinfoReferenceProteinEntry  [][]string   `xml:"ProteinEntry>reference>accinfo>xrefs>xref>db"`
	Date                                     string       `xml:"date,attr"`
	SeqSpec                                  []string     `xml:"ProteinEntry>feature>seq-spec"`
	UidXrefXrefsAccinfoReferenceProteinEntry [][]string   `xml:"ProteinEntry>reference>accinfo>xrefs>xref>uid"`
	Accession                                []string     `xml:"ProteinEntry>reference>accinfo>accession"`
	Superfamily                              []string     `xml:"ProteinEntry>classification>superfamily"`
	Created_date                             string       `xml:"ProteinEntry>header>created_date"`
	Id                                       string       `xml:"id,attr"`
	Citation                                 []string     `xml:"ProteinEntry>reference>refinfo>citation"`
	Volume                                   []string     `xml:"ProteinEntry>reference>refinfo>volume"`
	Uid                                      []string     `xml:"ProteinEntry>reference>refinfo>xrefs>xref>uid"`
	AccessionHeaderProteinEntry              []string     `xml:"ProteinEntry>header>accession"`
	Length                                   string       `xml:"ProteinEntry>summary>length"`
}

type Accinfo struct {
	Label string `xml:"label,attr"`
}
type Feature struct {
	Label string `xml:"label,attr"`
}
type ProteinEntry struct {
	Id string `xml:"id,attr"`
}
type Refinfo struct {
	Refid string `xml:"refid,attr"`
}

func main() {
	done1 := make(chan int)
	done2 := make(chan int)
	go parseXml("/home/nidhind/Downloads/psd7003.xml", done1)
	go parseXml("/home/nidhind/Downloads/psd7003 (copy).xml", done2)
	<-done1
	log.Println("Done 1")
	<-done2
	log.Println("Done 2")
}

func parseXml(path string, done chan int) {
	xmlFile, err := os.Open(path);
	log.Println(err);
	decoder := xml.NewDecoder(xmlFile)

	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		fmt.Print(t)
	}

	done <- 1
}
