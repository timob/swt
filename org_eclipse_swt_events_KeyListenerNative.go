package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsKeyListenerNative_KeyPressed(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsKeyListenerNative_KeyReleased(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsKeyListenerNative_KeyPressed
func go_callback_EventsKeyListenerNative_KeyPressed(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/KeyListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsKeyListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/KeyEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsKeyEvent{}
	arg_a.Callable = dst_a
i.KeyPressed(arg_a)
}

//export go_callback_EventsKeyListenerNative_KeyReleased
func go_callback_EventsKeyListenerNative_KeyReleased(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/KeyListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsKeyListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/KeyEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsKeyEvent{}
	arg_a.Callable = dst_a
i.KeyReleased(arg_a)
}

var EventsKeyListenerNativeMap = make(map[int]EventsKeyListenerInterface)

type EventsKeyListenerNative struct {
	*javabind.Callable
	EventsKeyListenerInterface
}

func NewEventsKeyListenerNative(implementation EventsKeyListenerInterface) *EventsKeyListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/KeyListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsKeyListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsKeyListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/KeyListenerNative", "keyPressed", javabind.Void, []interface{}{"org/eclipse/swt/events/KeyEvent"}, C.go_callback_EventsKeyListenerNative_KeyPressed)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/KeyListenerNative", "keyReleased", javabind.Void, []interface{}{"org/eclipse/swt/events/KeyEvent"}, C.go_callback_EventsKeyListenerNative_KeyReleased)

        })
    }
    