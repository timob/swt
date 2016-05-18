package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomTextChangeListenerNative_TextChanging(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_CustomTextChangeListenerNative_TextChanged(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_CustomTextChangeListenerNative_TextSet(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomTextChangeListenerNative_TextChanging
func go_callback_CustomTextChangeListenerNative_TextChanging(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/TextChangeListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomTextChangeListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/TextChangingEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomTextChangingEvent{}
	arg_a.Callable = dst_a
i.TextChanging(arg_a)
}

//export go_callback_CustomTextChangeListenerNative_TextChanged
func go_callback_CustomTextChangeListenerNative_TextChanged(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/TextChangeListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomTextChangeListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/TextChangedEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomTextChangedEvent{}
	arg_a.Callable = dst_a
i.TextChanged(arg_a)
}

//export go_callback_CustomTextChangeListenerNative_TextSet
func go_callback_CustomTextChangeListenerNative_TextSet(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/TextChangeListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomTextChangeListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/TextChangedEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomTextChangedEvent{}
	arg_a.Callable = dst_a
i.TextSet(arg_a)
}

var CustomTextChangeListenerNativeMap = make(map[int]CustomTextChangeListenerInterface)

type CustomTextChangeListenerNative struct {
	*javabind.Callable
	CustomTextChangeListenerInterface
}

func NewCustomTextChangeListenerNative(implementation CustomTextChangeListenerInterface) *CustomTextChangeListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TextChangeListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomTextChangeListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomTextChangeListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/TextChangeListenerNative", "textChanging", javabind.Void, []interface{}{"org/eclipse/swt/custom/TextChangingEvent"}, C.go_callback_CustomTextChangeListenerNative_TextChanging)
javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/TextChangeListenerNative", "textChanged", javabind.Void, []interface{}{"org/eclipse/swt/custom/TextChangedEvent"}, C.go_callback_CustomTextChangeListenerNative_TextChanged)
javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/TextChangeListenerNative", "textSet", javabind.Void, []interface{}{"org/eclipse/swt/custom/TextChangedEvent"}, C.go_callback_CustomTextChangeListenerNative_TextSet)

        })
    }
    