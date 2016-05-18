package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsMouseTrackListenerNative_MouseEnter(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsMouseTrackListenerNative_MouseExit(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsMouseTrackListenerNative_MouseHover(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsMouseTrackListenerNative_MouseEnter
func go_callback_EventsMouseTrackListenerNative_MouseEnter(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MouseTrackListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMouseTrackListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MouseEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMouseEvent{}
	arg_a.Callable = dst_a
i.MouseEnter(arg_a)
}

//export go_callback_EventsMouseTrackListenerNative_MouseExit
func go_callback_EventsMouseTrackListenerNative_MouseExit(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MouseTrackListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMouseTrackListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MouseEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMouseEvent{}
	arg_a.Callable = dst_a
i.MouseExit(arg_a)
}

//export go_callback_EventsMouseTrackListenerNative_MouseHover
func go_callback_EventsMouseTrackListenerNative_MouseHover(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MouseTrackListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMouseTrackListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MouseEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMouseEvent{}
	arg_a.Callable = dst_a
i.MouseHover(arg_a)
}

var EventsMouseTrackListenerNativeMap = make(map[int]EventsMouseTrackListenerInterface)

type EventsMouseTrackListenerNative struct {
	*javabind.Callable
	EventsMouseTrackListenerInterface
}

func NewEventsMouseTrackListenerNative(implementation EventsMouseTrackListenerInterface) *EventsMouseTrackListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MouseTrackListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsMouseTrackListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsMouseTrackListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MouseTrackListenerNative", "mouseEnter", javabind.Void, []interface{}{"org/eclipse/swt/events/MouseEvent"}, C.go_callback_EventsMouseTrackListenerNative_MouseEnter)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MouseTrackListenerNative", "mouseExit", javabind.Void, []interface{}{"org/eclipse/swt/events/MouseEvent"}, C.go_callback_EventsMouseTrackListenerNative_MouseExit)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MouseTrackListenerNative", "mouseHover", javabind.Void, []interface{}{"org/eclipse/swt/events/MouseEvent"}, C.go_callback_EventsMouseTrackListenerNative_MouseHover)

        })
    }
    