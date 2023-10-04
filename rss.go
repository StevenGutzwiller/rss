package rss

import (
	"encoding/xml"
	"io"
)

type Page struct {
	RSS RSS `xml:"rss"`
	Channel Channel `xml:"channel"`

}

type RSS struct {
	Version string `xml:"id,version"`
}

type Channel struct {
	Title string `xml:"title"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	Copyright string `xml:"copyright"`
	PubDate string `xml:"pubDate"`
	TTL string `xml:"ttl"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	GUID GUID `xml:"guid"`
	PubDate string `xml:"pubDate"`
	Category string `xml:"category"`

}

type GUID struct {
	IsPermaLink bool `xml:"isPermaLink,attr"`
	GUID string `xml:",chardata"`
}

func Decode[T any](r io.Reader) (T, error) {
	var val T
	dec := xml.NewDecoder(r)
	err := dec.Decode(&val)
	return val, err
}

func Encode[T any](w io.Writer, val T) error {
	enc := xml.NewEncoder(w)
	defer enc.Close()
	return enc.Encode(val)
}