package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomVerifyKeyListenerNative_VerifyKey(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomVerifyKeyListenerNative_VerifyKey
func go_callback_CustomVerifyKeyListenerNative_VerifyKey(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/VerifyKeyListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomVerifyKeyListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/VerifyEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsVerifyEvent{}
	arg_a.Callable = dst_a
i.VerifyKey(arg_a)
}

var CustomVerifyKeyListenerNativeMap = make(map[int]CustomVerifyKeyListenerInterface)

type CustomVerifyKeyListenerNative struct {
	*javabind.Callable
	CustomVerifyKeyListenerInterface
}

func NewCustomVerifyKeyListenerNative(implementation CustomVerifyKeyListenerInterface) *CustomVerifyKeyListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/VerifyKeyListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomVerifyKeyListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomVerifyKeyListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/VerifyKeyListenerNative", "verifyKey", javabind.Void, []interface{}{"org/eclipse/swt/events/VerifyEvent"}, C.go_callback_CustomVerifyKeyListenerNative_VerifyKey)

        })
    }
    