package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsSelectionListenerNative_WidgetSelected(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsSelectionListenerNative_WidgetDefaultSelected(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsSelectionListenerNative_WidgetSelected
func go_callback_EventsSelectionListenerNative_WidgetSelected(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/SelectionListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsSelectionListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/SelectionEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsSelectionEvent{}
	arg_a.Callable = dst_a
i.WidgetSelected(arg_a)
}

//export go_callback_EventsSelectionListenerNative_WidgetDefaultSelected
func go_callback_EventsSelectionListenerNative_WidgetDefaultSelected(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/SelectionListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsSelectionListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/SelectionEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsSelectionEvent{}
	arg_a.Callable = dst_a
i.WidgetDefaultSelected(arg_a)
}

var EventsSelectionListenerNativeMap = make(map[int]EventsSelectionListenerInterface)

type EventsSelectionListenerNative struct {
	*javabind.Callable
	EventsSelectionListenerInterface
}

func NewEventsSelectionListenerNative(implementation EventsSelectionListenerInterface) *EventsSelectionListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/SelectionListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsSelectionListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsSelectionListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/SelectionListenerNative", "widgetSelected", javabind.Void, []interface{}{"org/eclipse/swt/events/SelectionEvent"}, C.go_callback_EventsSelectionListenerNative_WidgetSelected)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/SelectionListenerNative", "widgetDefaultSelected", javabind.Void, []interface{}{"org/eclipse/swt/events/SelectionEvent"}, C.go_callback_EventsSelectionListenerNative_WidgetDefaultSelected)

        })
    }
    