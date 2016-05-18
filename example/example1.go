// Copyright 2015 Tim O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/timob/javabind"
	"github.com/timob/swt"
	"log"
	"os"
)

// Example as shown at https://www.eclipse.org/swt/ . But with an example of a listener too.

type selectionListener struct {
	button swt.WidgetsButtonInterface
	*swt.EventsSelectionListenerNative
}

func (m *selectionListener) WidgetSelected(a swt.EventsSelectionEventInterface) {
	if m.button.GetSelection() {
		m.button.SetSelection(false)
	} else {
		m.button.SetSelection(true)
	}
}

func (m *selectionListener) WidgetDefaultSelected(a swt.EventsSelectionEventInterface) {}

func NewSelectionListener(button swt.WidgetsButtonInterface) *selectionListener {
	s := &selectionListener{button, nil}
	s.EventsSelectionListenerNative = swt.NewEventsSelectionListenerNative(s)
	return s
}

func main() {
	err := javabind.SetupJVM(os.Getenv("CLASSPATH"))
	if err != nil {
		log.Fatal(err)
	}

	var SWT *swt.SWT

	display := swt.NewWidgetsDisplay()
	shell := swt.NewWidgetsShell3(display)
	shell.SetLayout(swt.NewLayoutGridLayout2(2, true))

	tabFolder := swt.NewWidgetsTabFolder(shell, SWT.NONE())
	tabFolder.SetLayoutData(swt.NewLayoutGridData5(SWT.FILL(), SWT.FILL(), true, true, 2, 1))
	item := swt.NewWidgetsTabItem(tabFolder, SWT.NONE())
	item.SetText("Widget")

	composite := swt.NewWidgetsComposite(tabFolder, SWT.NONE())
	composite.SetLayout(swt.NewLayoutGridLayout())

	tree := swt.NewWidgetsTree(composite, SWT.BORDER())
	item.SetControl(composite)
	tree.SetHeaderVisible(true)
	tree.SetLayoutData(swt.NewLayoutGridData4(SWT.FILL(), SWT.FILL(), true, true))
	column1 := swt.NewWidgetsTreeColumn(tree, SWT.NONE())
	column1.SetText("Standard")
	column2 := swt.NewWidgetsTreeColumn(tree, SWT.NONE())
	column2.SetText("Widget")
	branch := swt.NewWidgetsTreeItem(tree, SWT.NONE())
	branch.SetText2([]string{"Efficient", "Portable"})
	leaf := swt.NewWidgetsTreeItem2(branch, SWT.NONE())
	leaf.SetText2([]string{"Cross", "Platform"})
	branch.SetExpanded(true)
	branch = swt.NewWidgetsTreeItem(tree, SWT.NONE())
	branch.SetText2([]string{"Native", "Controls"})
	leaf = swt.NewWidgetsTreeItem2(branch, SWT.NONE())
	leaf.SetText2([]string{"Cross", "Platform"})
	branch = swt.NewWidgetsTreeItem(tree, SWT.NONE())
	branch.SetText2([]string{"Cross", "Platform"})
	column1.Pack()
	column2.Pack()

	item = swt.NewWidgetsTabItem(tabFolder, SWT.NONE())
	item.SetText("Toolkit")

	button := swt.NewWidgetsButton(shell, SWT.CHECK())
	button.SetText("Totally")
	button.SetSelection(true)
	button.SetLayoutData(swt.NewLayoutGridData4(SWT.CENTER(), SWT.CENTER(), false, false))
	listener := NewSelectionListener(button)

	button = swt.NewWidgetsButton(shell, SWT.PUSH())
	button.SetText("Awesome")
	button.SetLayoutData(swt.NewLayoutGridData4(SWT.CENTER(), SWT.CENTER(), false, false))
	button.AddSelectionListener(listener)

	shell.Pack()
	shell.Open()

	for !shell.IsDisposed() {
		if !display.ReadAndDispatch() {
			display.Sleep()
		}
	}
}
