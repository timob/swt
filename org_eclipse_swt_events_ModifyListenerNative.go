package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsModifyListenerNative_ModifyText(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsModifyListenerNative_ModifyText
func go_callback_EventsModifyListenerNative_ModifyText(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ModifyListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsModifyListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ModifyEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsModifyEvent{}
	arg_a.Callable = dst_a
i.ModifyText(arg_a)
}

var EventsModifyListenerNativeMap = make(map[int]EventsModifyListenerInterface)

type EventsModifyListenerNative struct {
	*javabind.Callable
	EventsModifyListenerInterface
}

func NewEventsModifyListenerNative(implementation EventsModifyListenerInterface) *EventsModifyListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ModifyListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsModifyListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsModifyListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ModifyListenerNative", "modifyText", javabind.Void, []interface{}{"org/eclipse/swt/events/ModifyEvent"}, C.go_callback_EventsModifyListenerNative_ModifyText)

        })
    }
    