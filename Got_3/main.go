package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	//http.HandleFunc("/surface", handlerSurface)
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
	StringConverToNumber()
}

func handlerSurface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	start := time.Now()
	surface(w)
	fmt.Printf("process time %f \n", time.Since(start).Seconds())
}

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, b1 := corner(i+1, j)
			bx, by, b2 := corner(i, j)
			cx, cy, b3 := corner(i, j+1)
			dx, dy, b4 := corner(i+1, j+1)
			if b1 && b2 && b3 && b4 {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#ff0000;stroke:#000000;stroke-width:1'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			} else if !b1 && !b2 && !b3 && !b4 {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#0000ff;stroke:#000000;stroke-width:1'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#cccccc;stroke:#000000;stroke-width:1'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}

		}
	}
	fmt.Fprintf(out, "</svg>\n")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z > 0
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
