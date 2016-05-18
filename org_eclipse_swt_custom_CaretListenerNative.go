package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomCaretListenerNative_CaretMoved(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomCaretListenerNative_CaretMoved
func go_callback_CustomCaretListenerNative_CaretMoved(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/CaretListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomCaretListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/CaretEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomCaretEvent{}
	arg_a.Callable = dst_a
i.CaretMoved(arg_a)
}

var CustomCaretListenerNativeMap = make(map[int]CustomCaretListenerInterface)

type CustomCaretListenerNative struct {
	*javabind.Callable
	CustomCaretListenerInterface
}

func NewCustomCaretListenerNative(implementation CustomCaretListenerInterface) *CustomCaretListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CaretListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomCaretListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomCaretListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/CaretListenerNative", "caretMoved", javabind.Void, []interface{}{"org/eclipse/swt/custom/CaretEvent"}, C.go_callback_CustomCaretListenerNative_CaretMoved)

        })
    }
    