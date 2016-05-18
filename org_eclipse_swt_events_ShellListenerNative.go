package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsShellListenerNative_ShellActivated(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsShellListenerNative_ShellClosed(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsShellListenerNative_ShellDeactivated(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsShellListenerNative_ShellDeiconified(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsShellListenerNative_ShellIconified(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsShellListenerNative_ShellActivated
func go_callback_EventsShellListenerNative_ShellActivated(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ShellListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsShellListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ShellEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsShellEvent{}
	arg_a.Callable = dst_a
i.ShellActivated(arg_a)
}

//export go_callback_EventsShellListenerNative_ShellClosed
func go_callback_EventsShellListenerNative_ShellClosed(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ShellListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsShellListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ShellEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsShellEvent{}
	arg_a.Callable = dst_a
i.ShellClosed(arg_a)
}

//export go_callback_EventsShellListenerNative_ShellDeactivated
func go_callback_EventsShellListenerNative_ShellDeactivated(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ShellListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsShellListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ShellEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsShellEvent{}
	arg_a.Callable = dst_a
i.ShellDeactivated(arg_a)
}

//export go_callback_EventsShellListenerNative_ShellDeiconified
func go_callback_EventsShellListenerNative_ShellDeiconified(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ShellListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsShellListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ShellEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsShellEvent{}
	arg_a.Callable = dst_a
i.ShellDeiconified(arg_a)
}

//export go_callback_EventsShellListenerNative_ShellIconified
func go_callback_EventsShellListenerNative_ShellIconified(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/ShellListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsShellListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/ShellEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsShellEvent{}
	arg_a.Callable = dst_a
i.ShellIconified(arg_a)
}

var EventsShellListenerNativeMap = make(map[int]EventsShellListenerInterface)

type EventsShellListenerNative struct {
	*javabind.Callable
	EventsShellListenerInterface
}

func NewEventsShellListenerNative(implementation EventsShellListenerInterface) *EventsShellListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/ShellListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsShellListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsShellListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ShellListenerNative", "shellActivated", javabind.Void, []interface{}{"org/eclipse/swt/events/ShellEvent"}, C.go_callback_EventsShellListenerNative_ShellActivated)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ShellListenerNative", "shellClosed", javabind.Void, []interface{}{"org/eclipse/swt/events/ShellEvent"}, C.go_callback_EventsShellListenerNative_ShellClosed)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ShellListenerNative", "shellDeactivated", javabind.Void, []interface{}{"org/eclipse/swt/events/ShellEvent"}, C.go_callback_EventsShellListenerNative_ShellDeactivated)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ShellListenerNative", "shellDeiconified", javabind.Void, []interface{}{"org/eclipse/swt/events/ShellEvent"}, C.go_callback_EventsShellListenerNative_ShellDeiconified)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/ShellListenerNative", "shellIconified", javabind.Void, []interface{}{"org/eclipse/swt/events/ShellEvent"}, C.go_callback_EventsShellListenerNative_ShellIconified)

        })
    }
    