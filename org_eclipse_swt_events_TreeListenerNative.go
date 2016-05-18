package swt

import "github.com/timob/javabind"
import "unsafe"


/*
#include<stdint.h>

extern void go_callback_EventsTreeListenerNative_TreeCollapsed(void *env, uintptr_t obj, uintptr_t arg_0 );
extern void go_callback_EventsTreeListenerNative_TreeExpanded(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_EventsTreeListenerNative_TreeCollapsed
func go_callback_EventsTreeListenerNative_TreeCollapsed(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/TreeListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsTreeListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/TreeEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsTreeEvent{}
	arg_a.Callable = dst_a
i.TreeCollapsed(arg_a)
}

//export go_callback_EventsTreeListenerNative_TreeExpanded
func go_callback_EventsTreeListenerNative_TreeExpanded(env unsafe.Pointer, obj uintptr, arg_0 uintptr) {
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/eclipse/swt/events/TreeListenerNative", false)}
    hash, err := rObj.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := EventsTreeListenerNativeMap[hash.(int)]
        	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "org/eclipse/swt/events/TreeEvent", false)); err != nil {
		panic(err)
	}
	arg_a := &EventsTreeEvent{}
	arg_a.Callable = dst_a
i.TreeExpanded(arg_a)
}

var EventsTreeListenerNativeMap = make(map[int]EventsTreeListenerInterface)

type EventsTreeListenerNative struct {
	*javabind.Callable
	EventsTreeListenerInterface
}

func NewEventsTreeListenerNative(implementation EventsTreeListenerInterface) *EventsTreeListenerNative {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/events/TreeListenerNative")
	if err != nil {
		panic(err)
	}
	
	x := &EventsTreeListenerNative{}
	x.Callable = &javabind.Callable{obj}

    hash, err := x.Callable.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
    if err != nil {
        panic(err)
    }
    EventsTreeListenerNativeMap[hash.(int)] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/eclipse/swt/events/TreeListenerNative", "treeCollapsed", javabind.Void, []interface{}{"org/eclipse/swt/events/TreeEvent"}, C.go_callback_EventsTreeListenerNative_TreeCollapsed)
javabind.GetEnv().RegisterNative("org/eclipse/swt/events/TreeListenerNative", "treeExpanded", javabind.Void, []interface{}{"org/eclipse/swt/events/TreeEvent"}, C.go_callback_EventsTreeListenerNative_TreeExpanded)

        })
    }
    