package swt

import "github.com/timob/javabind"

type CustomMovementListenerInterface interface {

	// public abstract void getNextOffset(org.eclipse.swt.custom.MovementEvent)
	GetNextOffset(a CustomMovementEventInterface) 

	// public abstract void getPreviousOffset(org.eclipse.swt.custom.MovementEvent)
	GetPreviousOffset(a CustomMovementEventInterface) 
}

type CustomMovementListener struct {
	*javabind.Callable
}

// public abstract void getNextOffset(org.eclipse.swt.custom.MovementEvent)
func (jbobject *CustomMovementListener) GetNextOffset(a CustomMovementEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getNextOffset", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/MovementEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void getPreviousOffset(org.eclipse.swt.custom.MovementEvent)
func (jbobject *CustomMovementListener) GetPreviousOffset(a CustomMovementEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getPreviousOffset", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/MovementEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


