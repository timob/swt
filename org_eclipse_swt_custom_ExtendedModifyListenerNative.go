package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomExtendedModifyListenerNative_ModifyText(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomExtendedModifyListenerNative_ModifyText
func go_callback_CustomExtendedModifyListenerNative_ModifyText(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/ExtendedModifyListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomExtendedModifyListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/ExtendedModifyEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomExtendedModifyEvent{}
	arg_a.Callable = dst_a
i.ModifyText(arg_a)
}

var CustomExtendedModifyListenerNativeMap = make(map[int]CustomExtendedModifyListenerInterface)

type CustomExtendedModifyListenerNative struct {
	*javabind.Callable
	CustomExtendedModifyListenerInterface
}

func NewCustomExtendedModifyListenerNative(implementation CustomExtendedModifyListenerInterface) *CustomExtendedModifyListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/ExtendedModifyListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomExtendedModifyListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomExtendedModifyListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/ExtendedModifyListenerNative", "modifyText", javabind.Void, []interface{}{"org/eclipse/swt/custom/ExtendedModifyEvent"}, C.go_callback_CustomExtendedModifyListenerNative_ModifyText)

        })
    }
    