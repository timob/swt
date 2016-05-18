package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_WidgetsListenerNative_HandleEvent(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_WidgetsListenerNative_HandleEvent
func go_callback_WidgetsListenerNative_HandleEvent(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/widgets/ListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := WidgetsListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/widgets/Event", false)); err != nil {
		panic(err)
	}
	arg_a := &WidgetsEvent{}
	arg_a.Callable = dst_a
i.HandleEvent(arg_a)
}

var WidgetsListenerNativeMap = make(map[int]WidgetsListenerInterface)

type WidgetsListenerNative struct {
	*javabind.Callable
	WidgetsListenerInterface
}

func NewWidgetsListenerNative(implementation WidgetsListenerInterface) *WidgetsListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/widgets/ListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &WidgetsListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    WidgetsListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/widgets/ListenerNative", "handleEvent", javabind.Void, []interface{}{"org/eclipse/swt/widgets/Event"}, C.go_callback_WidgetsListenerNative_HandleEvent)

        })
    }
    