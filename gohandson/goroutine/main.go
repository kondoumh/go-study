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
	GramBeans            Bean       = 1
	GramGroundBeans      GroundBean = 1
	MilliLitterWater     Water      = 1
	MilliLiterHotWater   HotWater   = 1
	CupsCofee            Coffee     = 1
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