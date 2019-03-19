package main
 
import (
	"encoding/xml"
	"io/ioutil"
)
 
type notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}
 
func main() {
	note := &notes{To: "Nicky",
		From:    "Rock",
		Heading: "Meeting",
		Body:    "Meeting at 5pm!",
	}
 
	file, _ := xml.MarshalIndent(note, "", " ")
 
	_ = ioutil.WriteFile("notes1.xml", file, 0644)
 
}
