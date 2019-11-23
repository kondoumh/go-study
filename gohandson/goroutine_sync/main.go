package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
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

	var wg sync.WaitGroup

	var hotWater HotWater
	var hwmu sync.Mutex
	for water > 0 {
		water -= 600 * MilliLitterWater
		go func() {
			hw := boil(ctx, 600*MilliLitterWater)
			wg.Add(1)
			hwmu.Lock()
			defer hwmu.Unlock()
			hotWater += hw
		}()
	}

	var groundBeans GroundBean
	var gbmu sync.Mutex
	for beans > 0 {
		beans -= 20 * GramBeans
		wg.Add(1)
		go func() {
			defer wg.Done()
			gb := grind(ctx, 20*GramBeans)
			gbmu.Lock()
			defer gbmu.Unlock()
			groundBeans += gb
		}()
	}

	wg.Wait()
	fmt.Println(hotWater)
	fmt.Println(groundBeans)

	var wg2 sync.WaitGroup
	var coffee Coffee
	var cfmu sync.Mutex
	cups := 4 * CupsCoffee
	for hotWater >= cups.HotWater() && groundBeans >= cups.GroundBeans() {
		hotWater -= cups.HotWater()
		groundBeans -= cups.GroundBeans()
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			cf := brew(ctx, cups.HotWater(), cups.GroundBeans())
			cfmu.Lock()
			defer cfmu.Unlock()
			coffee += cf
		}()
	}

	wg2.Wait()
	fmt.Println(coffee)
}

func boil(ctx context.Context, water Water) HotWater {
	defer trace.StartRegion(ctx, "boil").End()
	time.Sleep(400 * time.Millisecond)
	return HotWater(water)
}

func grind(ctx context.Context, beans Bean)  GroundBean {
	defer trace.StartRegion(ctx, "grind").End()
	time.Sleep(200 * time.Millisecond)
	return GroundBean(beans)
}

func brew(ctx context.Context, hotWater HotWater, groundBeans GroundBean) Coffee {
	defer trace.StartRegion(ctx, "brew").End()
	time.Sleep(1 * time.Second)
	cups1 := Coffee(hotWater / (1 * CupsCoffee).HotWater())
	cups2 := Coffee(groundBeans / (1 * CupsCoffee).GroundBeans())
	if cups1 < cups2 {
		return cups1
	}
	return cups2
}
