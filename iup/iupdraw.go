// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the COPYRIGHT file.

package iup

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liup
#cgo CFLAGS : -I../../libs/cd/include
#cgo LDFLAGS: -L../../libs/cd -liupcd -lcd
#include <stdlib.h>
#include <iup.h>
#include <iupdraw.h>
#include <cd.h>
*/
import "C"
import (
	"unsafe"
)

type Draw struct {
	*Handle
}

func NewDraw(h *Handle) *Draw {
	return &Draw{h}
}
func (h *Draw) Begin() {
	C.IupDrawBegin(toNative(h))
}

func (h *Draw) End() {
	C.IupDrawEnd(toNative(h))
}

/* all functions can be called only between calls to Begin and End */

func (h *Draw) SetClipRect(x1, y1, x2, y2 int) {
	C.IupDrawSetClipRect(toNative(h), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}
func (h *Draw) GetClipRect(x1, y1, x2, y2 *int) {
	C.IupDrawGetClipRect(toNative(h), (*C.int)(unsafe.Pointer(x1)), (*C.int)(unsafe.Pointer(y1)), (*C.int)(unsafe.Pointer(x2)), (*C.int)(unsafe.Pointer(y2)))
}
func (h *Draw) ResetClip() {
	C.IupDrawResetClip(toNative(h))
}

/* color controlled by the attribute DRAWCOLOR */
/* line style or fill controlled by the attribute DRAWSTYLE */

func (h *Draw) ParentBackground() {
	C.IupDrawParentBackground(toNative(h))
}
func (h *Draw) Line(x1, y1, x2, y2 int) {
	C.IupDrawLine(toNative(h), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}
func (h *Draw) Rectangle(x1, y1, x2, y2 int) {
	C.IupDrawRectangle(toNative(h), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}
func (h *Draw) Arc(x1, y1, x2, y2 int, a1, a2 float64) {
	C.IupDrawArc(toNative(h), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.double(a1), C.double(a2))
}
func (h *Draw) Polygon(points []int) {
	C.IupDrawPolygon(toNative(h), (*C.int)(unsafe.Pointer(&points[0])), C.int(len(points)))
}

func (h *Draw) SelectRect(x1, y1, x2, y2 int) {
	C.IupDrawSelectRect(toNative(h), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}
func (h *Draw) FocusRect(x1, y1, x2, y2 int) {
	C.IupDrawFocusRect(toNative(h), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

func (hi *Draw) GetSize(w, h *int) {
	C.IupDrawGetSize(toNative(hi), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

func (hi *Draw) GetTextSize(text string, w, h *int) {
	cx := NewCS(text)
	FreeCS(cx)

	C.IupDrawGetTextSize(toNative(hi), (*C.char)(unsafe.Pointer(cx)), C.int(len(text)), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}
func (hi *Draw) Text(text string, x, y, w, h int) {
	cx := NewCS(text)
	FreeCS(cx)
	C.IupDrawText(toNative(hi), (*C.char)(unsafe.Pointer(cx)), C.int(len(text)), C.int(x), C.int(y), C.int(w), C.int(h))
}

// func (h *Draw) IupDrawGetImageInfo(const char* name, int *w, int *h, int *bpp);
func (hi *Draw) Image(name string, x, y, w, h int) {
	cx := NewCS(name)
	FreeCS(cx)

	C.IupDrawImage(toNative(hi), (*C.char)(unsafe.Pointer(cx)), C.int(x), C.int(y), C.int(w), C.int(h))
}

type CD struct {
	Canvas unsafe.Pointer
}

func (cd *CD) Background() int {
	return int(C.cdCanvasBackground((*C.cdCanvas)(cd.Canvas), C.CD_QUERY))
}
func (cd *CD) Foreground(color int) {
	C.cdCanvasForeground((*C.cdCanvas)(cd.Canvas), C.long(color))
}
func (cd *CD) Box(xmin, xmax, ymin, ymax int) {
	C.cdCanvasBox((*C.cdCanvas)(cd.Canvas), C.int(xmin), C.int(xmax), C.int(ymin), C.int(ymax))
}
func (cd *CD) EncodeColor(red, green, blue byte) int {
	return int(C.cdEncodeColor(C.uchar(red), C.uchar(green), C.uchar(blue)))
}
func (cd *CD) DecodeColor(color int) (r byte, g byte, b byte) {
	C.cdDecodeColor(C.long(color), (*C.uchar)(unsafe.Pointer(&r)), (*C.uchar)(unsafe.Pointer(&g)), (*C.uchar)(unsafe.Pointer(&b)))
	return
}
func (cd *CD) TextAlignment(alignment string) {
	var a C.int
	switch alignment {
	case "CD_NORTH":
		a = C.CD_NORTH
	case "CD_SOUTH":
		a = C.CD_SOUTH
	case "CD_EAST":
		a = C.CD_EAST
	case "CD_WEST":
		a = C.CD_WEST
	case "CD_NORTH_EAST":
		a = C.CD_NORTH_EAST
	case "CD_NORTH_WEST":
		a = C.CD_NORTH_WEST
	case "CD_SOUTH_EAST":
		a = C.CD_SOUTH_EAST
	case "CD_SOUTH_WEST":
		a = C.CD_SOUTH_WEST
	case "CD_CENTER":
		a = C.CD_CENTER
	case "CD_BASE_LEFT":
		a = C.CD_BASE_LEFT
	case "CD_BASE_CENTER":
		a = C.CD_BASE_CENTER
	case "CD_BASE_RIGHT":
		a = C.CD_BASE_RIGHT
	}

	C.cdCanvasTextAlignment((*C.cdCanvas)(cd.Canvas), C.int(a))
}
func (cd *CD) Text(x, y int, s string) {
	cx := NewCS(s)
	FreeCS(cx)
	C.cdCanvasText((*C.cdCanvas)(cd.Canvas), C.int(x), C.int(y), (*C.char)(unsafe.Pointer(cx)))
}

// "System" (which depends on the driver and platform),
//"Courier" (mono spaced with serif),
//"Times" (proportional with serif) and
//"Helvetica" (proportional without serif).
func (cd *CD) Font(typeface, style string, size int) {
	cx := NewCS(typeface)
	FreeCS(cx)
	var a C.int
	switch style {
	case "CD_PLAIN", "":
		a = C.CD_PLAIN
	case "CD_BOLD", "b":
		a = C.CD_BOLD
	case "CD_ITALIC", "i":
		a = C.CD_ITALIC
	case "CD_UNDERLINE", "u":
		a = C.CD_UNDERLINE
	case "CD_STRIKEOUT", "s":
		a = C.CD_STRIKEOUT
	}
	C.cdCanvasFont((*C.cdCanvas)(cd.Canvas), (*C.char)(unsafe.Pointer(cx)), a, C.int(size))
}
