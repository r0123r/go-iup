// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the COPYRIGHT file.

package iuppplot

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liup -liup_plot
#cgo CFLAGS : -I../../libs/cd/include
#cgo LDFLAGS: -L../../libs/cd -liupcd -lcd
#include <stdlib.h>
#include <iup.h>
#include <iup_plot.h>
#include <cd.h>
*/
import "C"
import (
	"unsafe"

	"github.com/r0123r/go-iup/iup"
)

func toNative(h iup.IHandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(h.Native()))
}

func Open() *iup.Error {
	C.IupPlotOpen()
	return nil
}

type IupPlot struct {
	*iup.Handle
}

func PPlot(a ...interface{}) *IupPlot {
	return &IupPlot{iup.PPlot(a...)}
}

func AttachPPloat(h *iup.Handle) *IupPlot {
	return &IupPlot{h}
}

func (h *IupPlot) Begin(strXdata int) {
	C.IupPlotBegin(toNative(h), C.int(strXdata))
}

func (h *IupPlot) End() {
	C.IupPlotEnd(toNative(h))
}

func (h *IupPlot) Add(x, y float64) {
	C.IupPlotAdd(toNative(h), C.double(x), C.double(y))
}

func (h *IupPlot) AddStr(x string, y float64) {
	cx := iup.NewCS(x)
	iup.FreeCS(cx)
	C.IupPlotAddStr(toNative(h), (*C.char)(cx), C.double(y))
}

func (h *IupPlot) Insert(index, sample_index int, x, y float64) {
	C.IupPlotInsert(toNative(h), C.int(index), C.int(sample_index), C.double(x), C.double(y))
}

func (h *IupPlot) InsertStr(index, sample_index int, x string, y float64) {
	cx := iup.NewCS(x)
	iup.FreeCS(cx)
	C.IupPlotInsertStr(toNative(h), C.int(index), C.int(sample_index), (*C.char)(cx), C.double(y))
}

func (h *IupPlot) Transform(x, y float64, ix, iy *float64) {
	C.IupPlotTransform(toNative(h), C.double(x), C.double(y), (*C.double)(unsafe.Pointer(ix)), (*C.double)(unsafe.Pointer(iy)))
}

func (h *IupPlot) PaintTo(cnv uintptr) {
	C.IupPlotPaintTo(toNative(h), (*C.cdCanvas)(unsafe.Pointer(cnv)))
}
