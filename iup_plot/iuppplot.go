// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the COPYRIGHT file.

package iup_plot

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

func (h *IupPlot) End() int {
	return int(C.IupPlotEnd(toNative(h)))
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

func (h *IupPlot) PlotAddSegment(x, y float64) {
	C.IupPlotAddSegment(toNative(h), C.double(x), C.double(y))
}
func (h *IupPlot) PlotLoadData(filename string, strXdata int) int {
	cx := iup.NewCS(filename)
	iup.FreeCS(cx)
	i := C.IupPlotLoadData(toNative(h), (*C.char)(cx), C.int(strXdata))
	return int(i)
}
func (h *IupPlot) InsertSegment(index, ds_index, sample_index int, x, y float64) {
	C.IupPlotInsertSegment(toNative(h), C.int(ds_index), C.int(sample_index), C.double(x), C.double(y))
}
func (h *IupPlot) PlotInsertStrSamples(ds_index, sample_index int, x []string, y *float64) {
	count := len(x)
	cx := iup.NewCSA(x)
	defer iup.FreeCSA(cx)

	C.IupPlotInsertStrSamples(toNative(h), C.int(ds_index), C.int(sample_index), (**C.char)(unsafe.Pointer(&cx[0])), (*C.double)(unsafe.Pointer(y)), C.int(count))
}
func (h *IupPlot) PlotInsertSamples(ds_index, sample_index int, x, y *float64, count int) {
	C.IupPlotInsertSamples(toNative(h), C.int(ds_index), C.int(sample_index), (*C.double)(unsafe.Pointer(x)), (*C.double)(unsafe.Pointer(y)), C.int(count))
}

func (h *IupPlot) PlotAddSamples(ds_index int, x, y *float64, count int) {
	C.IupPlotAddSamples(toNative(h), C.int(ds_index), (*C.double)(unsafe.Pointer(x)), (*C.double)(unsafe.Pointer(y)), C.int(count))
}
func (h *IupPlot) PlotAddStrSamples(ds_index int, x []string, y *float64) {
	count := len(x)
	cx := iup.NewCSA(x)
	defer iup.FreeCSA(cx)
	C.IupPlotAddStrSamples(toNative(h), C.int(ds_index), (**C.char)(unsafe.Pointer(&cx[0])), (*C.double)(unsafe.Pointer(y)), C.int(count))
}

func (h *IupPlot) PlotGetSample(ds_index, sample_index int, x, y *float64) {
	C.IupPlotGetSample(toNative(h), C.int(ds_index), C.int(sample_index), (*C.double)(unsafe.Pointer(x)), (*C.double)(unsafe.Pointer(y)))
}
func (h *IupPlot) PlotGetSampleStr(ds_index, sample_index int) (x string, y float64) {
	cx := iup.NewCSN(x, 2000)
	defer iup.FreeCS(cx)

	C.IupPlotGetSampleStr(toNative(h), C.int(ds_index), C.int(sample_index), (**C.char)(unsafe.Pointer(&cx)), (*C.double)(unsafe.Pointer(&y)))
	return x, y
}
func (h *IupPlot) PlotGetSampleSelection(ds_index, sample_index int, x, y float64) int {
	i := C.IupPlotGetSampleSelection(toNative(h), C.int(ds_index), C.int(sample_index))
	return int(i)
}
func (h *IupPlot) PlotGetSampleExtra(ds_index, sample_index int) float64 {
	d := C.IupPlotGetSampleExtra(toNative(h), C.int(ds_index), C.int(sample_index))
	return float64(d)
}

func (h *IupPlot) PlotSetSample(ds_index, sample_index int, x, y float64) {
	C.IupPlotSetSample(toNative(h), C.int(ds_index), C.int(sample_index), C.double(x), C.double(y))
}
func (h *IupPlot) PlotSetSampleStr(ds_index, sample_index int, x string, y float64) {
	cx := iup.NewCS(x)
	iup.FreeCS(cx)

	C.IupPlotSetSampleStr(toNative(h), C.int(ds_index), C.int(sample_index), (*C.char)(cx), C.double(y))
}
func (h *IupPlot) PlotSetSampleSelection(ds_index, sample_index, selected int) {
	C.IupPlotSetSampleSelection(toNative(h), C.int(ds_index), C.int(sample_index), C.int(selected))
}
func (h *IupPlot) PlotSetSampleExtra(ds_index, sample_index int, extra float64) {
	C.IupPlotSetSampleExtra(toNative(h), C.int(ds_index), C.int(sample_index), C.double(extra))
}

func (h *IupPlot) PlotTransformTo(cnv_x, cnv_y float64, x, y *float64) {
	C.IupPlotTransformTo(toNative(h), C.double(cnv_x), C.double(cnv_y), (*C.double)(unsafe.Pointer(x)), (*C.double)(unsafe.Pointer(y)))
}
func (h *IupPlot) PlotFindSample(cnv_x, cnv_y float64, ds_index, sample_index *int32) int {
	return int(C.IupPlotFindSample(toNative(h), C.double(cnv_x), C.double(cnv_y), (*C.int)(unsafe.Pointer(ds_index)), (*C.int)(unsafe.Pointer(sample_index))))
}
func (h *IupPlot) PlotFindSegment(cnv_x, cnv_y float64, ds_index, sample_index1, sample_index2 *int32) int {
	return int(C.IupPlotFindSegment(toNative(h), C.double(cnv_x), C.double(cnv_y), (*C.int)(unsafe.Pointer(ds_index)), (*C.int)(unsafe.Pointer(sample_index1)), (*C.int)(unsafe.Pointer(sample_index2))))
}
