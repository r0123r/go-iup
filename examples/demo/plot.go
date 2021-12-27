package main

import (
	"math"

	"github.com/r0123r/go-iup/iup"
	"github.com/r0123r/go-iup/iup_plot"
)

func vbox6(name string) *iup.Handle {
	pplot := iup_plot.PPlot(
		iup.Attr("TITLE", "A simple XY Plot"),
		iup.Attrs(
			"MARGINBOTTOM", "35",
			"MARGINLEFT", "35",
			//"AXS_XLABEL", "X",
			//"AXS_YLABEL", "Y",
			"SHOWCROSSHAIR", "HORIZONTAL",
			"LEGENDSHOW", "YES",
			"LEGENDPOS", "BOTTOMCENTER",
			"LEGENDBOX", "NO",
			"GRID", "YES",
			"MENUITEMPROPERTIES", "YES",
		),
	)
	pplot.Begin(0)
	for x := -3.14; x <= 3.14; x += .1 {
		pplot.Add(x, math.Cos(x))
	}
	pplot.End()
	pplot.SetAttribute("DS_NAME", "P0")
	pplot.SetAttribute("DS_LEGEND", "Cos(x)")
	pplot.Begin(0)
	for x := -3.14; x <= 3.14; x += .1 {
		pplot.Add(x, math.Cosh(x)/10)
	}
	pplot.End()
	pplot.SetAttribute("DS_NAME", "P1")
	pplot.SetAttribute("DS_LEGEND", "Cosh(x)/10")
	pplot1 := iup_plot.PPlot(
		//iup.Attr("TITLE", "A simple XY Plot"),
		iup.Attrs(
			"MARGINBOTTOM", "35",
			"MARGINLEFT", "35",
			//"AXS_XLABEL", "X",
			//"AXS_YLABEL", "Y",
			"SHOWCROSSHAIR", "HORIZONTAL",
			"LEGENDSHOW", "YES",
			"LEGENDPOS", "BOTTOMCENTER",
			"LEGENDBOX", "NO",
			"GRID", "YES",
			"MENUITEMPROPERTIES", "YES",
		),
	)
	pplot1.Begin(0)
	for x := -3.14; x <= 3.14; x += .01 {
		pplot1.Add(x, math.Sin(x))
	}
	pplot1.End()
	res := iup.Vbox(
		"TABTITLE="+name,
		pplot,
		pplot1,
	)
	return res
}
