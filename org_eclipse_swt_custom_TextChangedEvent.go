package swt

import "github.com/timob/javabind"

type CustomTextChangedEventInterface interface {
	EventsTypedEventInterface
}

type CustomTextChangedEvent struct {
	EventsTypedEvent
}

// public org.eclipse.swt.custom.TextChangedEvent(org.eclipse.swt.custom.StyledTextContent)
func NewCustomTextChangedEvent(a CustomStyledTextContentInterface) (*CustomTextChangedEvent) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/TextChangedEvent", conv_a.Value().Cast("org/eclipse/swt/custom/StyledTextContent"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CustomTextChangedEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}


