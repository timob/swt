package swt

import "github.com/timob/javabind"

type GraphicsRectangleInterface interface {
	JavaLangObjectInterface

	// public void add(org.eclipse.swt.graphics.Rectangle)
	Add(a GraphicsRectangleInterface) 

	// public boolean contains(int, int)
	Contains(a int, b int) bool

	// public boolean contains(org.eclipse.swt.graphics.Point)
	Contains2(a GraphicsPointInterface) bool

	// public void intersect(org.eclipse.swt.graphics.Rectangle)
	Intersect(a GraphicsRectangleInterface) 

	// public org.eclipse.swt.graphics.Rectangle intersection(org.eclipse.swt.graphics.Rectangle)
	Intersection(a GraphicsRectangleInterface) *GraphicsRectangle

	// public boolean intersects(int, int, int, int)
	Intersects(a int, b int, c int, d int) bool

	// public boolean intersects(org.eclipse.swt.graphics.Rectangle)
	Intersects2(a GraphicsRectangleInterface) bool

	// public boolean isEmpty()
	IsEmpty() bool

	// public org.eclipse.swt.graphics.Rectangle union(org.eclipse.swt.graphics.Rectangle)
	Union(a GraphicsRectangleInterface) *GraphicsRectangle
}

type GraphicsRectangle struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.Rectangle(int, int, int, int)
func NewGraphicsRectangle(a int, b int, c int, d int) (*GraphicsRectangle) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Rectangle", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &GraphicsRectangle{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void add(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsRectangle) Add(a GraphicsRectangleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "add", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean contains(int, int)
func (jbobject *GraphicsRectangle) Contains(a int, b int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean contains(org.eclipse.swt.graphics.Point)
func (jbobject *GraphicsRectangle) Contains2(a GraphicsPointInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsRectangle) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *GraphicsRectangle) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void intersect(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsRectangle) Intersect(a GraphicsRectangleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "intersect", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.eclipse.swt.graphics.Rectangle intersection(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsRectangle) Intersection(a GraphicsRectangleInterface) *GraphicsRectangle {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intersection", "org/eclipse/swt/graphics/Rectangle", conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean intersects(int, int, int, int)
func (jbobject *GraphicsRectangle) Intersects(a int, b int, c int, d int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intersects", javabind.Boolean, a, b, c, d)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean intersects(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsRectangle) Intersects2(a GraphicsRectangleInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intersects", javabind.Boolean, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean isEmpty()
func (jbobject *GraphicsRectangle) IsEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *GraphicsRectangle) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.graphics.Rectangle union(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsRectangle) Union(a GraphicsRectangleInterface) *GraphicsRectangle {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/eclipse/swt/graphics/Rectangle", conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &GraphicsRectangle{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *GraphicsRectangle) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsRectangle) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsRectangle) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsRectangle) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsRectangle) Width() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsRectangle) SetFieldWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsRectangle) Height() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "height", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsRectangle) SetFieldHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "height", val)
	if err != nil {
		panic(err)
	}

}


