// Copyright 2015 Tim O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"github.com/timob/javabind"
	"github.com/timob/swt"
)

type UIRunner struct {
	uiFuncChan chan func()
	uiDone     chan byte
	cbFuncChan chan func()
	jniThread  *javabind.AttachedThread
	display    *swt.WidgetsDisplay
	awake		bool
	check		*javabind.CheckEnv
}

func NewUIRunner(d *swt.WidgetsDisplay) *UIRunner {
	return &UIRunner{make(chan func(), 10), make(chan byte), make(chan func(), 1000), javabind.NewAttachedThread(), d, false, nil}
}

func (u *UIRunner) Run(f func()) {
	if u.check != nil && u.check.SameEnv() {
		f()
	} else {
		u.uiFuncChan <- f
		u.jniThread.Run(u.display.Wake)
		<-u.uiDone
	}
}

func (u *UIRunner) ExecUiFuncs() {
	if u.check == nil {
		u.check = javabind.NewCheckEnv()
	}
	select {
	case f := <-u.uiFuncChan:
		f()
		u.uiDone <- 1
	case f := <-u.cbFuncChan:
		f()
	A:
		for {
			select {
			case f := <-u.cbFuncChan:
				f()
			default:
				break A
			}
		}
	default:
	}
}

func (u *UIRunner) CallBackRun(f func()) {
	u.cbFuncChan <- f
}
