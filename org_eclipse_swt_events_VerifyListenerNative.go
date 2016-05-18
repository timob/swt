package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsVerifyListenerNative_VerifyText(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsVerifyListenerNative_VerifyText
func go_callback_EventsVerifyListenerNative_VerifyText(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/VerifyListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsVerifyListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/VerifyEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsVerifyEvent{}
	arg_a.Callable = dst_a
i.VerifyText(arg_a)
}

var EventsVerifyListenerNativeMap = make(map[int]EventsVerifyListenerInterface)

type EventsVerifyListenerNative struct {
	*javabind.Callable
	EventsVerifyListenerInterface
}

func NewEventsVerifyListenerNative(implementation EventsVerifyListenerInterface) *EventsVerifyListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/VerifyListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsVerifyListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsVerifyListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/VerifyListenerNative", "verifyText", javabind.Void, []interface{}{"org/eclipse/swt/events/VerifyEvent"}, C.go_callback_EventsVerifyListenerNative_VerifyText)

        })
    }
    