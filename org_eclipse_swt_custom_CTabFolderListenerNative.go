package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_CustomCTabFolderListenerNative_ItemClosed(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_CustomCTabFolderListenerNative_ItemClosed
func go_callback_CustomCTabFolderListenerNative_ItemClosed(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/custom/CTabFolderListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := CustomCTabFolderListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/custom/CTabFolderEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &CustomCTabFolderEvent{}
	arg_a.Callable = dst_a
i.ItemClosed(arg_a)
}

var CustomCTabFolderListenerNativeMap = make(map[int]CustomCTabFolderListenerInterface)

type CustomCTabFolderListenerNative struct {
	*javabind.Callable
	CustomCTabFolderListenerInterface
}

func NewCustomCTabFolderListenerNative(implementation CustomCTabFolderListenerInterface) *CustomCTabFolderListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/CTabFolderListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &CustomCTabFolderListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    CustomCTabFolderListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/custom/CTabFolderListenerNative", "itemClosed", javabind.Void, []interface{}{"org/eclipse/swt/custom/CTabFolderEvent"}, C.go_callback_CustomCTabFolderListenerNative_ItemClosed)

        })
    }
    