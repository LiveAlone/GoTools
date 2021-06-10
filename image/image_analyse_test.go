package image

import (
	"fmt"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"log"
	"os"
	"testing"
)

func TestImageAnalyse(t *testing.T) {
	fname := "test.jpg"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)

	lat, long, _ := x.LatLong()
	fmt.Println("lat, long: ", lat, ", ", long)
}

func TestPosition(t *testing.T) {
	geocoder := openstreetmap.Geocoder()
	lat, lon := 31.210254669166666, 121.65915679916667
	addr, err := geocoder.ReverseGeocode(lat, lon)
	fmt.Println(addr, err)
	// addr &{齐爱路, 唐镇, 浦东新区, 201210, 中国 齐爱路   201210     中国 CN 浦东新区}
}
