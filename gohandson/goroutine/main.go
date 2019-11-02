package main

import (
	"fmt"
	"time"
)

type (
	Bean       int
	GroundBean int
	Water      int
	HotWater   int
	Coffee     int
)

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

func (cups Coffee) Water() Water {
	return Water(180*cups) / MilliLitterWater
}

func (cups Coffee) HotWater() HotWater {
	return HotWater(180*cups) / MilliLiterHotWater
}

func (cups Coffee) GroundBeans() GroundBean {
	return GroundBean(20*cups) / GramGroundBeans
}

func boil(water Water) HotWater {
	time.Sleep(400 * time.Millisecond)
	return HotWater(water)
}

func grind(beans Bean) GroundBean {
	time.Sleep(200 * time.Microsecond)
	return GroundBean(beans)
}

func brew(hotWater HotWater, groundBeans GroundBean) Coffee {
	time.Sleep(1 * time.Second)
	cups1 := Coffee(hotWater / (1 * CupsCoffee).HotWater())
	cups2 := Coffee(groundBeans / (1 * CupsCoffee).GroundBeans())
	if cups1 < cups2 {
		return cups1
	}
	return cups2
}

func main() {
	const amountCoffee = 20 * CupsCoffee

	water := amountCoffee.Water()
	beans := amountCoffee.Beans()

	fmt.Println(water)
	fmt.Println(beans)

	var hotWater HotWater
	for water > 0 {
		water -= 600 * MilliLiterHotWater
		hotWater += boil(500 * MilliLitterWater)
	}
	fmt.Println(hotWater)

	var groundBeans GroundBean
	for beans > 0 {
		beans -= 20 * GramBeans
		groundBeans += grind(20 * GramBeans)
	}
	fmt.Println(groundBeans)

	var coffee Coffee
	cups := 4 * CupsCoffee
	for hotWater >= cups.HotWater() && groundBeans >= cups.GroundBeans() {
		hotWater -= cups.HotWater()
		groundBeans -= cups.GroundBeans()
		coffee += brew(cups.HotWater(), cups.GroundBeans())
	}
	fmt.Println(coffee)
}
