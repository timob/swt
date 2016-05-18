package swt

import "github.com/timob/javabind"

type CustomTextChangeListenerInterface interface {

	// public abstract void textChanging(org.eclipse.swt.custom.TextChangingEvent)
	TextChanging(a CustomTextChangingEventInterface) 

	// public abstract void textChanged(org.eclipse.swt.custom.TextChangedEvent)
	TextChanged(a CustomTextChangedEventInterface) 

	// public abstract void textSet(org.eclipse.swt.custom.TextChangedEvent)
	TextSet(a CustomTextChangedEventInterface) 
}

type CustomTextChangeListener struct {
	*javabind.Callable
}

// public abstract void textChanging(org.eclipse.swt.custom.TextChangingEvent)
func (jbobject *CustomTextChangeListener) TextChanging(a CustomTextChangingEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "textChanging", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/TextChangingEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void textChanged(org.eclipse.swt.custom.TextChangedEvent)
func (jbobject *CustomTextChangeListener) TextChanged(a CustomTextChangedEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "textChanged", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/TextChangedEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void textSet(org.eclipse.swt.custom.TextChangedEvent)
func (jbobject *CustomTextChangeListener) TextSet(a CustomTextChangedEventInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "textSet", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/custom/TextChangedEvent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


