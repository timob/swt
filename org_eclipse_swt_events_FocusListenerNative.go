package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsFocusListenerNative_FocusGained(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsFocusListenerNative_FocusLost(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsFocusListenerNative_FocusGained
func go_callback_EventsFocusListenerNative_FocusGained(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/FocusListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsFocusListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/FocusEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsFocusEvent{}
	arg_a.Callable = dst_a
i.FocusGained(arg_a)
}

//export go_callback_EventsFocusListenerNative_FocusLost
func go_callback_EventsFocusListenerNative_FocusLost(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/FocusListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsFocusListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/FocusEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsFocusEvent{}
	arg_a.Callable = dst_a
i.FocusLost(arg_a)
}

var EventsFocusListenerNativeMap = make(map[int]EventsFocusListenerInterface)

type EventsFocusListenerNative struct {
	*javabind.Callable
	EventsFocusListenerInterface
}

func NewEventsFocusListenerNative(implementation EventsFocusListenerInterface) *EventsFocusListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/FocusListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsFocusListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsFocusListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/FocusListenerNative", "focusGained", javabind.Void, []interface{}{"org/eclipse/swt/events/FocusEvent"}, C.go_callback_EventsFocusListenerNative_FocusGained)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/FocusListenerNative", "focusLost", javabind.Void, []interface{}{"org/eclipse/swt/events/FocusEvent"}, C.go_callback_EventsFocusListenerNative_FocusLost)

        })
    }
    