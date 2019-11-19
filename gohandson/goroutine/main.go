package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

// type aliases
type (
	Bean       int
	GroundBean int
	Water      int
	HotWater   int
	Coffee     int
)

// const values
const (
	GramBeans          Bean       = 1
	GramGroundBeans    GroundBean = 1
	MilliLitterWater   Water      = 1
	MilliLiterHotWater HotWater   = 1
	CupsCoffee         Coffee     = 1
)

func (w Water) String() string {
	return fmt.Sprintf("%d[ml] water", int(w))
}

func (hw HotWater) String() string {
	return fmt.Sprintf("%d[ml] hot water", int(hw))
}

func (b Bean) String() string {
	return fmt.Sprintf("%d[g] beans", int(b))
}

func (gb GroundBean) String() string {
	return fmt.Sprintf("%d[g] ground beans", int(gb))
}

func (cups Coffee) String() string {
	return fmt.Sprintf("%d cup(s) coffee", int(cups))
}

// Water needed to make cups of coffee
func (cups Coffee) Water() Water {
	return Water(180*cups) / MilliLitterWater
}

// HotWater needed to make cups of coffee
func (cups Coffee) HotWater() HotWater {
	return HotWater(180*cups) / MilliLiterHotWater
}

// Beans needed to make cups of coffee
func (cups Coffee) Beans() Bean {
	return Bean(20*cups) / GramBeans
}

// GroundBeans needed to make cups of coffee
func (cups Coffee) GroundBeans() GroundBean {
	return GroundBean(20*cups) / GramGroundBeans
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()
	_main()
}

func _main() {
	const amountCoffee = 20 * CupsCoffee

	ctx, task := trace.NewTask(context.Background(), "make coffee")
	defer task.End()

	water := amountCoffee.Water()
	beans := amountCoffee.Beans()

	fmt.Println(water)
	fmt.Println(beans)

	hwch := make(chan HotWater)
	gbch := make(chan GroundBean)
	cfch := make(chan Coffee)

	var hwCount int
	for water > 0 {
		hwCount++
		water -= 600 * MilliLitterWater
		go boil(ctx, hwch, 600 * MilliLitterWater)
	}

	var gbCount int
	for beans > 0 {
		beans -= 20 * GramBeans
		gbCount++
		go grind(ctx, gbch, 20 * GramBeans)
	}

	var hotWater HotWater
	for i := 0; i < hwCount; i++ {
		hotWater += <-hwch
	}
	fmt.Println(hotWater)

	var groundBeans GroundBean
	for i := 0; i < gbCount; i++ {
		groundBeans += <-gbch
	}
	fmt.Println(groundBeans)

	var cfCount int
	cups := 4 * CupsCoffee
	for hotWater >= cups.HotWater() && groundBeans >= cups.GroundBeans() {
		hotWater -= cups.HotWater()
		groundBeans -= cups.GroundBeans()
		cfCount++
		go brew(ctx, cfch, cups.HotWater(), cups.GroundBeans())
	}

	var coffee Coffee
	for i := 0; i < cfCount; i++ {
		coffee += <-cfch
	}
	fmt.Println(coffee)
}

func boil(ctx context.Context, ch chan<- HotWater, water Water) {
	defer trace.StartRegion(ctx, "boil").End()
	time.Sleep(400 * time.Millisecond)
	ch <- HotWater(water)
}

func grind(ctx context.Context, ch chan<- GroundBean, beans Bean) {
	defer trace.StartRegion(ctx, "grind").End()
	time.Sleep(200 * time.Millisecond)
	ch <- GroundBean(beans)
}

func brew(ctx context.Context, ch chan<- Coffee, hotWater HotWater, groundBeans GroundBean) {
	defer trace.StartRegion(ctx, "brew").End()
	time.Sleep(1 * time.Second)
	cups1 := Coffee(hotWater / (1 * CupsCoffee).HotWater())
	cups2 := Coffee(groundBeans / (1 * CupsCoffee).GroundBeans())
	if cups1 < cups2 {
		ch <- cups1
	} else {
		ch <- cups2
	}
}
