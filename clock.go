package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var offset, r_s, x_s, y_s, r_m, x_m, y_m, r_l, x_l, y_l, index float64

	var secondCommand, hourCommand, minutesCommand, ticCommand string

	index = 0

	offset = 95

	fmt.Printf("r=2000\n")
	fmt.Printf("e=1\n")

	x_s = 0
	y_s = 0
	x_m = 0
	y_m = 0
	x_l = 0
	y_l = 0

	r_s = 500
	r_m = 550
	r_l = 750

	// tics
	for n := 0; n < 360; n++ {
		if math.Mod(float64(n), 30) == 0 {

			tick := n - int(offset)

			x_m = (math.Sin(float64(tick)*(math.Pi/180)) * (r_m)) + 2000
			y_m = (math.Cos(float64(tick)*(math.Pi/180)) * (r_m)) + 2000

			x_l = (math.Sin(float64(tick)*(math.Pi/180)) * (r_l)) + 2000
			y_l = (math.Cos(float64(tick)*(math.Pi/180)) * (r_l)) + 2000

			ticCommand = ticCommand + fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", int(x_m), int(y_m), "0", "0", "0", "0")
			ticCommand = ticCommand + fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", int(x_l), int(y_l), "4095", "4095", "1", "1")
			ticCommand = ticCommand + fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", int(x_l), int(y_l), "0", "0", "0", "0")
		}
	}

	for true {
		x_l = (math.Sin(index*(math.Pi/180)) * r_l) + 2000
		y_l = (math.Cos(index*(math.Pi/180)) * r_l) + 2000

		x_m = (math.Sin(index*(math.Pi/180)) * r_m) + 2000
		y_m = (math.Cos(index*(math.Pi/180)) * r_m) + 2000

		x_s = (math.Sin(index*(math.Pi/180)) * r_s) + 2000
		y_s = (math.Cos(index*(math.Pi/180)) * r_s) + 2000

		_ = x_m
		_ = y_m
		_ = x_l
		_ = y_l

		t := time.Now()
		h := t.Hour()
		m := t.Minute()
		s := t.Second()

		// hands
		seconds := float64((360 / 60) * s)
		x_l = (math.Sin(seconds*(math.Pi/180)) * r_l) + 2000
		y_l = (math.Cos(seconds*(math.Pi/180)) * r_l) + 2000
		secondCommand = fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", 2000, 2000, "0", "0", "0", "0")
		secondCommand = secondCommand + fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", int(x_l), int(y_l), "4095", "4095", "1", "1")
		secondCommand = secondCommand + fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", int(x_l), int(y_l), "0", "0", "0", "0")
		fmt.Printf(secondCommand)

		minutes := float64((360 / 60) * m)
		minutes = minutes - offset
		x_l = (math.Sin(minutes*(math.Pi/180)) * r_l) + 2000
		y_l = (math.Cos(minutes*(math.Pi/180)) * r_l) + 2000
		minutesCommand = fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", 2000, 2000, "0", "0", "0", "0")
		minutesCommand = minutesCommand + fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", int(x_l), int(y_l), "4095", "4095", "1", "1")
		minutesCommand = minutesCommand + fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", int(x_l), int(y_l), "0", "0", "0", "0")
		fmt.Printf(minutesCommand)

		hours := float64((360 / 12) * h)
		if h > 12 {
			hours = float64((360 / 12) * (h - 12))
		}
		hours = hours - offset
		x_s = (math.Sin(hours*(math.Pi/180)) * r_s) + 2000
		y_s = (math.Cos(hours*(math.Pi/180)) * r_s) + 2000

		hourCommand = fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", 2000, 2000, "0", "0", "0", "0")
		hourCommand = hourCommand + fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", int(x_s), int(y_s), "4095", "4095", "1", "1")
		hourCommand = hourCommand + fmt.Sprintf("s=%v,%v,%v,%v,%v,%v\n", int(x_s), int(y_s), "0", "0", "0", "0")

		fmt.Printf(hourCommand)
		fmt.Printf(ticCommand)
	}
}
