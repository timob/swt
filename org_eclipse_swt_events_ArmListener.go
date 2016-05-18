package swt

import "github.com/timob/javabind"

type EventsArmListenerInterface interface {

	// public abstract void widgetArmed(org.eclipse.swt.events.ArmEvent)
	WidgetArmed(a EventsArmEventInterface) 
}

type EventsArmListener struct {
	*javabind.Callable
}

// public abstract void widgetArmed(org.eclipse.swt.events.ArmEvent)
func (jbobject *EventsArmListener) WidgetArmed(a EventsArmEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "widgetArmed", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/events/ArmEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


