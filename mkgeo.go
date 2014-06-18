package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(42)

	stypes := []string{"Fisk", "Vann", "Jord", "Plante", "Luft", "Skalldyr"}

	var data bytes.Buffer

	data.WriteString("db = db.getSiblingDB('geo')\nd = db.data\nd.drop()\n\n")
	data.WriteString("u = db.users\nu.drop()\n\nu.insert({'username':'drb', 'password':'passw', 'email':'dag.robole@gmail.com'})\n")
	data.WriteString("u.insert({'username':'ola', 'password':'olsen', 'email':'ola.olsen@gmail.com'})\n\n")

	for i := 0; i < 3000; i++ {
		act := rand.Float32() * 100
		unc := rand.Float32() * 0.3
		refdate := time.Date(2000+rand.Intn(14), time.Month(1+rand.Intn(12)), 1+rand.Intn(31), rand.Intn(24), 0, 0, 0, time.UTC)
		stype := stypes[rand.Intn(6)]
		lat := 55.0 + rand.Float32()*25.0
		lon := -3.0 + rand.Float32()*43.0
		str := fmt.Sprintf("d.insert({ activity: %f, uncertainty: %f, sigma: 2, refdate: new Date('%s'), sample_type: '%s', location: { coordinates: [%f, %f] }})\n", act, unc, refdate.String(), stype, lon, lat)
		data.WriteString(str)
	}

	data.WriteString("\nd.ensureIndex({\"location.coordinates\": \"2d\"})\n")
	data.WriteString("\nd.ensureIndex({\"sample_type\": 1})\n")

	err := ioutil.WriteFile("./geo.js", data.Bytes(), 0644)
	if err != nil {
		panic("Unable to write file")
	}
}
