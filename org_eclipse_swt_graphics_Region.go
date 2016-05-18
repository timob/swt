package swt

import "github.com/timob/javabind"

type GraphicsRegionInterface interface {
	GraphicsResourceInterface

	// public void add(int[])
	Add(a []int) 

	// public void add(org.eclipse.swt.graphics.Rectangle)
	Add3(a GraphicsRectangleInterface) 

	// public void add(int, int, int, int)
	Add2(a int, b int, c int, d int) 

	// public void add(org.eclipse.swt.graphics.Region)
	Add4(a GraphicsRegionInterface) 

	// public boolean contains(int, int)
	Contains(a int, b int) bool

	// public boolean contains(org.eclipse.swt.graphics.Point)
	Contains2(a GraphicsPointInterface) bool

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public static org.eclipse.swt.graphics.Region gtk_new(org.eclipse.swt.graphics.Device, long)
	Gtk_new(a GraphicsDeviceInterface, b int64) *GraphicsRegion

	// public void intersect(org.eclipse.swt.graphics.Rectangle)
	Intersect2(a GraphicsRectangleInterface) 

	// public void intersect(int, int, int, int)
	Intersect(a int, b int, c int, d int) 

	// public void intersect(org.eclipse.swt.graphics.Region)
	Intersect3(a GraphicsRegionInterface) 

	// public boolean intersects(int, int, int, int)
	Intersects(a int, b int, c int, d int) bool

	// public boolean intersects(org.eclipse.swt.graphics.Rectangle)
	Intersects2(a GraphicsRectangleInterface) bool

	// public boolean isEmpty()
	IsEmpty() bool

	// public void subtract(int[])
	Subtract(a []int) 

	// public void subtract(org.eclipse.swt.graphics.Rectangle)
	Subtract3(a GraphicsRectangleInterface) 

	// public void subtract(int, int, int, int)
	Subtract2(a int, b int, c int, d int) 

	// public void subtract(org.eclipse.swt.graphics.Region)
	Subtract4(a GraphicsRegionInterface) 

	// public void translate(int, int)
	Translate(a int, b int) 

	// public void translate(org.eclipse.swt.graphics.Point)
	Translate2(a GraphicsPointInterface) 
}

type GraphicsRegion struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.Region()
func NewGraphicsRegion() (*GraphicsRegion) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Region")
	if err != nil {
		panic(err)
	}
	x := &GraphicsRegion{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Region(org.eclipse.swt.graphics.Device)
func NewGraphicsRegion2(a GraphicsDeviceInterface) (*GraphicsRegion) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Region", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsRegion{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void add(int[])
func (jbobject *GraphicsRegion) Add(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "add", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void add(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsRegion) Add3(a GraphicsRectangleInterface)  {
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

// public void add(int, int, int, int)
func (jbobject *GraphicsRegion) Add2(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "add", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void add(org.eclipse.swt.graphics.Region)
func (jbobject *GraphicsRegion) Add4(a GraphicsRegionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "add", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Region"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean contains(int, int)
func (jbobject *GraphicsRegion) Contains(a int, b int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean contains(org.eclipse.swt.graphics.Point)
func (jbobject *GraphicsRegion) Contains2(a GraphicsPointInterface) bool {
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
func (jbobject *GraphicsRegion) Equals(a interface{}) bool {
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

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *GraphicsRegion) GetBounds() *GraphicsRectangle {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBounds", "org/eclipse/swt/graphics/Rectangle")
	if err != nil {
		panic(err)
	}
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

// public static org.eclipse.swt.graphics.Region gtk_new(org.eclipse.swt.graphics.Device, long)
func (jbobject *GraphicsRegion) Gtk_new(a GraphicsDeviceInterface, b int64) *GraphicsRegion {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/Region", "gtk_new", "org/eclipse/swt/graphics/Region", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b)
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
	unique_x := &GraphicsRegion{}
	unique_x.Callable = dst
	return unique_x
}

// public int hashCode()
func (jbobject *GraphicsRegion) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void intersect(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsRegion) Intersect2(a GraphicsRectangleInterface)  {
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

// public void intersect(int, int, int, int)
func (jbobject *GraphicsRegion) Intersect(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "intersect", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void intersect(org.eclipse.swt.graphics.Region)
func (jbobject *GraphicsRegion) Intersect3(a GraphicsRegionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "intersect", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Region"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean intersects(int, int, int, int)
func (jbobject *GraphicsRegion) Intersects(a int, b int, c int, d int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intersects", javabind.Boolean, a, b, c, d)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean intersects(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsRegion) Intersects2(a GraphicsRectangleInterface) bool {
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

// public boolean isDisposed()
func (jbobject *GraphicsRegion) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isEmpty()
func (jbobject *GraphicsRegion) IsEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void subtract(int[])
func (jbobject *GraphicsRegion) Subtract(a []int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void subtract(org.eclipse.swt.graphics.Rectangle)
func (jbobject *GraphicsRegion) Subtract3(a GraphicsRectangleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void subtract(int, int, int, int)
func (jbobject *GraphicsRegion) Subtract2(a int, b int, c int, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void subtract(org.eclipse.swt.graphics.Region)
func (jbobject *GraphicsRegion) Subtract4(a GraphicsRegionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "subtract", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Region"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void translate(int, int)
func (jbobject *GraphicsRegion) Translate(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "translate", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void translate(org.eclipse.swt.graphics.Point)
func (jbobject *GraphicsRegion) Translate2(a GraphicsPointInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "translate", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Point"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String toString()
func (jbobject *GraphicsRegion) ToString() string {
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

func (jbobject *GraphicsRegion) Handle() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "handle", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsRegion) SetFieldHandle(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "handle", val)
	if err != nil {
		panic(err)
	}

}


