package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsMenuListenerNative_MenuHidden(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsMenuListenerNative_MenuShown(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsMenuListenerNative_MenuHidden
func go_callback_EventsMenuListenerNative_MenuHidden(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MenuListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMenuListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MenuEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMenuEvent{}
	arg_a.Callable = dst_a
i.MenuHidden(arg_a)
}

//export go_callback_EventsMenuListenerNative_MenuShown
func go_callback_EventsMenuListenerNative_MenuShown(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/MenuListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsMenuListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/MenuEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsMenuEvent{}
	arg_a.Callable = dst_a
i.MenuShown(arg_a)
}

var EventsMenuListenerNativeMap = make(map[int]EventsMenuListenerInterface)

type EventsMenuListenerNative struct {
	*javabind.Callable
	EventsMenuListenerInterface
}

func NewEventsMenuListenerNative(implementation EventsMenuListenerInterface) *EventsMenuListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/MenuListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsMenuListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsMenuListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MenuListenerNative", "menuHidden", javabind.Void, []interface{}{"org/eclipse/swt/events/MenuEvent"}, C.go_callback_EventsMenuListenerNative_MenuHidden)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/MenuListenerNative", "menuShown", javabind.Void, []interface{}{"org/eclipse/swt/events/MenuEvent"}, C.go_callback_EventsMenuListenerNative_MenuShown)

        })
    }
    