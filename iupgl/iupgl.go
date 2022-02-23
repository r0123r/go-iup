/*
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of go-iup.

	go-iup is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	go-iup is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with go-iup.  If not, see <http://www.gnu.org/licenses/>.
*/

package iupgl

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liup -liup_plot
#cgo CFLAGS : -I../../libs/cd/include
#cgo LDFLAGS: -L../../libs/cd -liupcd -lcd
#cgo LDFLAGS: -liupgl -liupglcontrols -liupcontrols
#cgo windows LDFLAGS: -lopengl32 -lglu32
#cgo LDFLAGS: -lfreeglut

#include <stdlib.h>
#include "iup.h"
#include "iupgl.h"
#include "iupglcontrols.h"
#include <GL/glut.h>
*/
import "C"
import (
	"unsafe"

	"github.com/r0123r/go-iup/iup"
)

const (
	GLCANVAS      = "glcanvas"
	GLCANVASBOX   = "glcanvasbox"
	GLSUBCANVAS   = "glsubcanvas"
	GLFRAME       = "glframe"
	GLLABEL       = "gllabel"
	GLSEPARATOR   = "glseparator"
	GLBUTTON      = "glbutton"
	GLTOGGLE      = "gltoggle"
	GLLINK        = "gllink"
	GLPROGRESSBAR = "glprogressbar"
	GLVAL         = "glval"
	GLEXPANDER    = "glexpander"
	GLSCROLLBOX   = "glscrollbox"
	GLSIZEBOX     = "glsizebox"
	GLTEXT        = "gltext"
)

func Canvas(a ...interface{}) *iup.Handle {
	return iup.New(GLCANVAS, a...)
}
func Text(a ...interface{}) *iup.Handle {
	return iup.New(GLTEXT, a...)
}
func CanvasOpen() *iup.Error {
	C.IupGLCanvasOpen()
	return nil
}
func ControlsOpen() {
	C.IupGLControlsOpen()
}
func toNative(h *iup.Handle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(h.Native()))
}

func MakeCurrent(ih *iup.Handle) {
	C.IupGLMakeCurrent(toNative(ih))
}

func IsCurrent(ih *iup.Handle) int {
	return int(C.IupGLIsCurrent(toNative(ih)))
}

func SwapBuffers(ih *iup.Handle) {
	C.IupGLSwapBuffers(toNative(ih))
}

func Palette(ih *iup.Handle, index int, r, g, b float64) {
	C.IupGLPalette(toNative(ih), C.int(index), C.float(r), C.float(g), C.float(b))
}

func UseFont(ih *iup.Handle, first, count, list_base int) {
	C.IupGLUseFont(toNative(ih), C.int(first), C.int(count), C.int(list_base))
}

func Wait(gl int) {
	C.IupGLWait(C.int(gl))
}

func init() {
	iup.RegisterClass(GLCANVAS, iup.NewClassInfo(GLCANVAS, iup.Canvas_SetCallback))
	iup.RegisterClass(GLTEXT, iup.NewClassInfo(GLTEXT, iup.Canvas_SetCallback))
}
func Perspective(fovy, aspect, zNear, zFar float64) {
	C.gluPerspective(C.GLdouble(fovy), C.GLdouble(aspect), C.GLdouble(zNear), C.GLdouble(zFar))

}
func LookAt(eyex, eyey, eyez, centerx, centery, centerz, upx, upy, upz float64) {
	C.gluLookAt(C.GLdouble(eyex), C.GLdouble(eyey), C.GLdouble(eyez),
		C.GLdouble(centerx), C.GLdouble(centery), C.GLdouble(centerz),
		C.GLdouble(upx), C.GLdouble(upy), C.GLdouble(upz))
}
