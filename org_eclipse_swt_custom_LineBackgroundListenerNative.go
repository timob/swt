package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomLineBackgroundListenerNative_LineGetBackground(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomLineBackgroundListenerNative_LineGetBackground
func go_callback_CustomLineBackgroundListenerNative_LineGetBackground(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/LineBackgroundListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomLineBackgroundListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/LineBackgroundEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomLineBackgroundEvent{}
	arg_a.Callable = dst_a
i.LineGetBackground(arg_a)
}

var CustomLineBackgroundListenerNativeMap = make(map[int]CustomLineBackgroundListenerInterface)

type CustomLineBackgroundListenerNative struct {
	*javabind.Callable
	CustomLineBackgroundListenerInterface
}

func NewCustomLineBackgroundListenerNative(implementation CustomLineBackgroundListenerInterface) *CustomLineBackgroundListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/LineBackgroundListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomLineBackgroundListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomLineBackgroundListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/LineBackgroundListenerNative", "lineGetBackground", javabind.Void, []interface{}{"org/eclipse/swt/custom/LineBackgroundEvent"}, C.go_callback_CustomLineBackgroundListenerNative_LineGetBackground)

        })
    }
    