package swt

import "github.com/timob/javabind"

type GraphicsPathDataInterface interface {
	JavaLangObjectInterface
}

type GraphicsPathData struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.PathData()
func NewGraphicsPathData() (*GraphicsPathData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/PathData")
	if err != nil {
		panic(err)
	}
	x := &GraphicsPathData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *GraphicsPathData) Points() []float32 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "points", javabind.Float | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]float32)
}

func (jbobject *GraphicsPathData) SetFieldPoints(val []float32) {
	err := jbobject.SetField(javabind.GetEnv(), "points", val)
	if err != nil {
		panic(err)
	}

}


