package main

import (
	"flag"
	// "fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/gousb"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {

		ctx := gousb.NewContext()
		defer ctx.Close()

		dev, err := ctx.OpenDeviceWithVIDPID(0x16c0, 0x5dc)
		if err != nil {
			log.Fatal("Could not open device: %v", err)
		}
		defer dev.Close()

		var buf = []byte{0, 0}

		for {
			ret, err := dev.Control(gousb.ControlIn|gousb.ControlVendor, 4, 0, 0, buf)
			if err != nil {
				log.Fatal("dev.Control returned %s", err)
			}
			if ret == 0 {
				log.Fatal("Control returned 0 bytes: %d", ret, err)
			}
			/*
			log.Printf("%#08b %#08b, \n", buf[1], buf[0])
			log.Printf("7 %01b MAX   0=off 1=on\n", (buf[1] >> 7 & 1))
			log.Printf("6 %01b speed 0=fast 1=slow\n", (buf[1] >> 6 & 1))
			log.Printf("5 %01b db    0=A 1=C\n", (buf[1] >> 5 & 1))
			log.Printf("4 %01b\n", (buf[1] >> 4 & 1))
			log.Printf("3 %01b\n", (buf[1] >> 3 & 1))
			log.Printf("2 %01b\n", (buf[1] >> 2 & 1))
			log.Printf("r %03b %d range\n", (buf[1] >> 2 & 7), (buf[1] >> 2 & 7))
			*/
			/*
				buf[1] 	msar rruu
				               ^^ = upper two bits of decibel value*10 - 30.
					   ^ ^^   = range 0: 30-80 1: 40-90 2: 50-100 3: 60-110 4: 70-120 5: 80-130 7: 30-130
					  ^       = ac 0=A 1=C
					 ^        = speed 0=fast 1=slow
					^         = max 0=off 1=on

				buf[0]
			*/

			db := (float64(buf[0])+float64(buf[1]&3)*256)*0.1 + 30

			decibelGauge.Set(db)
			time.Sleep(1 * time.Second)

		}
	}()
}

var (
	decibelGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "ws1361",
		Name:      "decibel",
		Help:      "sound pressure",
	})
)

func main() {
	var listen_string = flag.String("listen", ":1361", "Address:Port to listen on")
	flag.Parse()
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Listening on http://%v/metrics\n", *listen_string)
	http.ListenAndServe(*listen_string, nil)
}
