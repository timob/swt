package swt

import "github.com/timob/javabind"

type CustomCaretEventInterface interface {
	EventsTypedEventInterface
}

type CustomCaretEvent struct {
	EventsTypedEvent
}

func (jbobject *CustomCaretEvent) CaretOffset() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "caretOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCaretEvent) SetFieldCaretOffset(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "caretOffset", val)
	if err != nil {
		panic(err)
	}

}


