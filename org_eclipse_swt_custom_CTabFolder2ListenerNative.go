package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomCTabFolder2ListenerNative_Close(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_CustomCTabFolder2ListenerNative_Minimize(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_CustomCTabFolder2ListenerNative_Maximize(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_CustomCTabFolder2ListenerNative_Restore(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_CustomCTabFolder2ListenerNative_ShowList(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomCTabFolder2ListenerNative_Close
func go_callback_CustomCTabFolder2ListenerNative_Close(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/CTabFolder2ListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomCTabFolder2ListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/CTabFolderEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomCTabFolderEvent{}
	arg_a.Callable = dst_a
i.Close(arg_a)
}

//export go_callback_CustomCTabFolder2ListenerNative_Minimize
func go_callback_CustomCTabFolder2ListenerNative_Minimize(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/CTabFolder2ListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomCTabFolder2ListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/CTabFolderEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomCTabFolderEvent{}
	arg_a.Callable = dst_a
i.Minimize(arg_a)
}

//export go_callback_CustomCTabFolder2ListenerNative_Maximize
func go_callback_CustomCTabFolder2ListenerNative_Maximize(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/CTabFolder2ListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomCTabFolder2ListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/CTabFolderEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomCTabFolderEvent{}
	arg_a.Callable = dst_a
i.Maximize(arg_a)
}

//export go_callback_CustomCTabFolder2ListenerNative_Restore
func go_callback_CustomCTabFolder2ListenerNative_Restore(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/CTabFolder2ListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomCTabFolder2ListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/CTabFolderEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomCTabFolderEvent{}
	arg_a.Callable = dst_a
i.Restore(arg_a)
}

//export go_callback_CustomCTabFolder2ListenerNative_ShowList
func go_callback_CustomCTabFolder2ListenerNative_ShowList(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/CTabFolder2ListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomCTabFolder2ListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/CTabFolderEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomCTabFolderEvent{}
	arg_a.Callable = dst_a
i.ShowList(arg_a)
}

var CustomCTabFolder2ListenerNativeMap = make(map[int]CustomCTabFolder2ListenerInterface)

type CustomCTabFolder2ListenerNative struct {
	*javabind.Callable
	CustomCTabFolder2ListenerInterface
}

func NewCustomCTabFolder2ListenerNative(implementation CustomCTabFolder2ListenerInterface) *CustomCTabFolder2ListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CTabFolder2ListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomCTabFolder2ListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomCTabFolder2ListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/CTabFolder2ListenerNative", "close", javabind.Void, []interface{}{"org/eclipse/swt/custom/CTabFolderEvent"}, C.go_callback_CustomCTabFolder2ListenerNative_Close)
javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/CTabFolder2ListenerNative", "minimize", javabind.Void, []interface{}{"org/eclipse/swt/custom/CTabFolderEvent"}, C.go_callback_CustomCTabFolder2ListenerNative_Minimize)
javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/CTabFolder2ListenerNative", "maximize", javabind.Void, []interface{}{"org/eclipse/swt/custom/CTabFolderEvent"}, C.go_callback_CustomCTabFolder2ListenerNative_Maximize)
javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/CTabFolder2ListenerNative", "restore", javabind.Void, []interface{}{"org/eclipse/swt/custom/CTabFolderEvent"}, C.go_callback_CustomCTabFolder2ListenerNative_Restore)
javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/CTabFolder2ListenerNative", "showList", javabind.Void, []interface{}{"org/eclipse/swt/custom/CTabFolderEvent"}, C.go_callback_CustomCTabFolder2ListenerNative_ShowList)

        })
    }
    