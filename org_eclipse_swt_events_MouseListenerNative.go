package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsMouseListenerNative_MouseDoubleClick(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsMouseListenerNative_MouseDown(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsMouseListenerNative_MouseUp(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsMouseListenerNative_MouseDoubleClick
func go_callback_EventsMouseListenerNative_MouseDoubleClick(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MouseListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMouseListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MouseEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMouseEvent{}
	arg_a.Callable = dst_a
i.MouseDoubleClick(arg_a)
}

//export go_callback_EventsMouseListenerNative_MouseDown
func go_callback_EventsMouseListenerNative_MouseDown(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MouseListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMouseListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MouseEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMouseEvent{}
	arg_a.Callable = dst_a
i.MouseDown(arg_a)
}

//export go_callback_EventsMouseListenerNative_MouseUp
func go_callback_EventsMouseListenerNative_MouseUp(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MouseListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMouseListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MouseEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMouseEvent{}
	arg_a.Callable = dst_a
i.MouseUp(arg_a)
}

var EventsMouseListenerNativeMap = make(map[int]EventsMouseListenerInterface)

type EventsMouseListenerNative struct {
	*javabind.Callable
	EventsMouseListenerInterface
}

func NewEventsMouseListenerNative(implementation EventsMouseListenerInterface) *EventsMouseListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MouseListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsMouseListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsMouseListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MouseListenerNative", "mouseDoubleClick", javabind.Void, []interface{}{"org/eclipse/swt/events/MouseEvent"}, C.go_callback_EventsMouseListenerNative_MouseDoubleClick)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MouseListenerNative", "mouseDown", javabind.Void, []interface{}{"org/eclipse/swt/events/MouseEvent"}, C.go_callback_EventsMouseListenerNative_MouseDown)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MouseListenerNative", "mouseUp", javabind.Void, []interface{}{"org/eclipse/swt/events/MouseEvent"}, C.go_callback_EventsMouseListenerNative_MouseUp)

        })
    }
    