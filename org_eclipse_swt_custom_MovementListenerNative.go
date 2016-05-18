package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomMovementListenerNative_GetNextOffset(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_CustomMovementListenerNative_GetPreviousOffset(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomMovementListenerNative_GetNextOffset
func go_callback_CustomMovementListenerNative_GetNextOffset(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/MovementListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomMovementListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/MovementEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomMovementEvent{}
	arg_a.Callable = dst_a
i.GetNextOffset(arg_a)
}

//export go_callback_CustomMovementListenerNative_GetPreviousOffset
func go_callback_CustomMovementListenerNative_GetPreviousOffset(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/MovementListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomMovementListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/MovementEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomMovementEvent{}
	arg_a.Callable = dst_a
i.GetPreviousOffset(arg_a)
}

var CustomMovementListenerNativeMap = make(map[int]CustomMovementListenerInterface)

type CustomMovementListenerNative struct {
	*javabind.Callable
	CustomMovementListenerInterface
}

func NewCustomMovementListenerNative(implementation CustomMovementListenerInterface) *CustomMovementListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/MovementListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomMovementListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomMovementListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/MovementListenerNative", "getNextOffset", javabind.Void, []interface{}{"org/eclipse/swt/custom/MovementEvent"}, C.go_callback_CustomMovementListenerNative_GetNextOffset)
javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/MovementListenerNative", "getPreviousOffset", javabind.Void, []interface{}{"org/eclipse/swt/custom/MovementEvent"}, C.go_callback_CustomMovementListenerNative_GetPreviousOffset)

        })
    }
    