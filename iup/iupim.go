// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the COPYRIGHT file.

package iup

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS : -L../../libs/iup -liup -liupim
#cgo linux CFLAGS : -I../../libs/im/include
#cgo linux LDFLAGS : -L../../libs/im -lim
#include <iup.h>
#include <iupim.h>
*/
import "C"

func LoadImage(filename string) *Handle {
	cname := NewCS(filename)
	defer FreeCS(cname)
	return NewHandle(C.IupLoadImage((*C.char)(cname)))
}
