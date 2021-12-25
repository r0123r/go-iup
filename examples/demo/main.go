// Copyright (C) 2011-2012 r0123r. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the COPYRIGHT file.

package main

import (
	"fmt"

	"github.com/r0123r/go-iup/iup"
	"github.com/r0123r/go-iup/iup_plot"
	"github.com/r0123r/go-iup/iupctls"
	"github.com/r0123r/go-iup/iupimglib"
)

func main() {
	e := iup.Open()
	if e != nil {
		fmt.Println(e)
		return
	}
	defer iup.Close()

	iupctls.Open()
	iup_plot.Open()
	iupimglib.Open()

	ui()
}

func ui() {

	tree := iup.Tree(
		"EXPANDALL=YES",
	)
	tree.SetName("tree")
	pr := iup.ProgressBar("MIN=0,MAX=100,VALUE=50,EXPAND=HORIZONTAL")
	vbox1 := iup.Vbox(
		"TABTITLE=Standard,MARGIN=0x0,GAP=5",
		iup.Frame("TITLE=Frame,EXPAND=YES",
			iup.Vbox(
				iup.Label("TITLE=\"List Label\""),
				iup.List("EXPAND=YES,1=list1", "2=list2", "3=list3"),
			),
		),
		iup.Hbox(
			iup.List("DROPDOWN=YES,EDITBOX=YES", "1=list1", "2=list2"),
			iup.SpinBox(
				iup.Text("VALUE=100,SPIN=YES"),
			),
			iup.Toggle("TITLE=Toggle,EXPAND=VERTICAL"),
			iup.Val("MIN=1,MAX=100"),
		),
		iup.Frame(
			iup.Radio(
				iup.Hbox(
					iup.Toggle("TITLE=1"),
					iup.Toggle("TITLE=2"),
					iup.Toggle("TITLE=3"),
					iup.Toggle("TITLE=4"),
				),
			),
		),
		iup.Text("EXPAND=HORIZONTAL,SPIN=YES,SPINVALUE=50", func(arg *iup.TextSpin) {
			pr.SetAttribute("VALUE", fmt.Sprint(arg.Inc))
			arg.Return = iup.DEFAULT
		}),
		pr,
		tree,
	)
	vbox2 := iup.Vbox(
		"TABTITLE=Text",
		iup.Text("EXPAND=YES,MULTILINE=YES"),
	)
	vbox3 := iup.Vbox(
		"TABTITLE=Colorbar",
		iup.Colorbar(
			"ORIENTATION=HORIZONTAL",
			"NUM_PARTS=2",
			"SHOW_SECONDARY=YES",
			func(arg *iup.ColorbarSelect) {
				fmt.Println(arg)
			},
		),
	)
	vbox4 := iup.Vbox(
		"TABTITLE=ColorBrowser",
		iup.ColorBrowser(
			"EXPAND=YES",
			func(arg *iup.ColorBrowserDrag) {
				fmt.Printf("R=%d,G=%d,B=%d\n", arg.R, arg.G, arg.B)
			},
		),
	)
	dlg := iup.Dialog(
		iup.Attr("TITLE", "GO-IUP Demo 1.0"),
		"SIZE=350x200",
		"SHRINK=YES",
		"MARGIN=10x10",
		"FULLSCREEN=YES",
		iup.Tabs(
			vbox1,
			vbox2,
			vbox3,
			vbox4,
			vbox5("Dial"),
			vbox6("Plot"),
			vbox7("Matrix"),
			vbox8("Image"),
			vbox9("Messagebox"),
		),
	)
	dlg.Show()

	tree.SetAttribute("ADDBRANCH", "Item2")
	tree.SetAttribute("ADDLEAF1", "leaf3")
	tree.SetAttribute("ADDLEAF2", "leaf4")
	tree.SetAttribute("ADDBRANCH", "Item1")
	tree.SetAttribute("ADDLEAF1", "leaf1")
	tree.SetAttribute("ADDLEAF2", "leaf2")

	defer dlg.Destroy()
	iup.MainLoop()
}
