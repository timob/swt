# SWT
Go bindings for Eclipse SWT widget toolkit library. https://www.eclipse.org/swt/

The binding uses JNI to use SWT. The classpath of the JVM must be setup to include your local SWT install (swt.jar) and the native
classes used by this binding (classpath directory in the source). This can be done with javabind.SetupJVM().

## Notes
### Thread model
SWT requires that functions that modify the UI must be run on the UI thread. In this binding when javabind.SetupJVM is called 
the Go routine is locked to the thread. Call swt.NewWidgetsDisplay() from this Go routine and run the event loop here. The util package
can be used to run functions from other Go routines on the UI Go routine. Place ExecUiFuncs() just before display.Sleep() in the 
event loop, so functions from other Go routines execute here.

Note listener callback functions always run on the UI Go routine, so the util package is not needed for these.

## Example
Demonstrates adding controls to a window, tab control, tree and setting up a button selection listener. (from examples dir)

```` go
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
````
