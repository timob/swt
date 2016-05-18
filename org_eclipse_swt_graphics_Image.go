package swt

import "github.com/timob/javabind"

type GraphicsImageInterface interface {
	GraphicsResourceInterface

	// public org.eclipse.swt.graphics.Color getBackground()
	GetBackground() *GraphicsColor

	// public org.eclipse.swt.graphics.Rectangle getBounds()
	GetBounds() *GraphicsRectangle

	// public org.eclipse.swt.graphics.ImageData getImageData()
	GetImageData() *GraphicsImageData

	// public static org.eclipse.swt.graphics.Image gtk_new(org.eclipse.swt.graphics.Device, int, long, long)
	Gtk_new(a GraphicsDeviceInterface, b int, c int64, d int64) *GraphicsImage

	// public static org.eclipse.swt.graphics.Image gtk_new_from_pixbuf(org.eclipse.swt.graphics.Device, int, long)
	Gtk_new_from_pixbuf(a GraphicsDeviceInterface, b int, c int64) *GraphicsImage

	// public long internal_new_GC(org.eclipse.swt.graphics.GCData)
	Internal_new_GC(a GraphicsGCDataInterface) int64

	// public void internal_dispose_GC(long, org.eclipse.swt.graphics.GCData)
	Internal_dispose_GC(a int64, b GraphicsGCDataInterface) 

	// public void setBackground(org.eclipse.swt.graphics.Color)
	SetBackground(a GraphicsColorInterface) 
}

type GraphicsImage struct {
	GraphicsResource
}

// public org.eclipse.swt.graphics.Image(org.eclipse.swt.graphics.Device, int, int)
func NewGraphicsImage(a GraphicsDeviceInterface, b int, c int) (*GraphicsImage) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Image", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsImage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Image(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.Image, int)
func NewGraphicsImage5(a GraphicsDeviceInterface, b GraphicsImageInterface, c int) (*GraphicsImage) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Image", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/Image"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsImage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Image(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.Rectangle)
func NewGraphicsImage4(a GraphicsDeviceInterface, b GraphicsRectangleInterface) (*GraphicsImage) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Image", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/Rectangle"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsImage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Image(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.ImageData)
func NewGraphicsImage3(a GraphicsDeviceInterface, b GraphicsImageDataInterface) (*GraphicsImage) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Image", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/ImageData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsImage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Image(org.eclipse.swt.graphics.Device, org.eclipse.swt.graphics.ImageData, org.eclipse.swt.graphics.ImageData)
func NewGraphicsImage6(a GraphicsDeviceInterface, b GraphicsImageDataInterface, c GraphicsImageDataInterface) (*GraphicsImage) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Image", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("org/eclipse/swt/graphics/ImageData"), conv_c.Value().Cast("org/eclipse/swt/graphics/ImageData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &GraphicsImage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.Image(org.eclipse.swt.graphics.Device, java.lang.String)
func NewGraphicsImage2(a GraphicsDeviceInterface, b string) (*GraphicsImage) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/Image", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &GraphicsImage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean equals(java.lang.Object)
func (jbobject *GraphicsImage) Equals(a interface{}) bool {
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

// public org.eclipse.swt.graphics.Color getBackground()
func (jbobject *GraphicsImage) GetBackground() *GraphicsColor {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBackground", "org/eclipse/swt/graphics/Color")
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
	unique_x := &GraphicsColor{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.Rectangle getBounds()
func (jbobject *GraphicsImage) GetBounds() *GraphicsRectangle {
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

// public org.eclipse.swt.graphics.ImageData getImageData()
func (jbobject *GraphicsImage) GetImageData() *GraphicsImageData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageData", "org/eclipse/swt/graphics/ImageData")
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
	unique_x := &GraphicsImageData{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.eclipse.swt.graphics.Image gtk_new(org.eclipse.swt.graphics.Device, int, long, long)
func (jbobject *GraphicsImage) Gtk_new(a GraphicsDeviceInterface, b int, c int64, d int64) *GraphicsImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/Image", "gtk_new", "org/eclipse/swt/graphics/Image", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b, c, d)
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
	unique_x := &GraphicsImage{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.eclipse.swt.graphics.Image gtk_new_from_pixbuf(org.eclipse.swt.graphics.Device, int, long)
func (jbobject *GraphicsImage) Gtk_new_from_pixbuf(a GraphicsDeviceInterface, b int, c int64) *GraphicsImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/Image", "gtk_new_from_pixbuf", "org/eclipse/swt/graphics/Image", conv_a.Value().Cast("org/eclipse/swt/graphics/Device"), b, c)
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
	unique_x := &GraphicsImage{}
	unique_x.Callable = dst
	return unique_x
}

// public int hashCode()
func (jbobject *GraphicsImage) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public long internal_new_GC(org.eclipse.swt.graphics.GCData)
func (jbobject *GraphicsImage) Internal_new_GC(a GraphicsGCDataInterface) int64 {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "internal_new_GC", javabind.Long, conv_a.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public void internal_dispose_GC(long, org.eclipse.swt.graphics.GCData)
func (jbobject *GraphicsImage) Internal_dispose_GC(a int64, b GraphicsGCDataInterface)  {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "internal_dispose_GC", javabind.Void, a, conv_b.Value().Cast("org/eclipse/swt/graphics/GCData"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public boolean isDisposed()
func (jbobject *GraphicsImage) IsDisposed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisposed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setBackground(org.eclipse.swt.graphics.Color)
func (jbobject *GraphicsImage) SetBackground(a GraphicsColorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBackground", javabind.Void, conv_a.Value().Cast("org/eclipse/swt/graphics/Color"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String toString()
func (jbobject *GraphicsImage) ToString() string {
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

func (jbobject *GraphicsImage) Type() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "type", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImage) SetFieldType(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "type", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImage) Pixmap() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "pixmap", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsImage) SetFieldPixmap(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "pixmap", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImage) Mask() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "mask", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsImage) SetFieldMask(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "mask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImage) Surface() int64 {
	jret, err := jbobject.GetField(javabind.GetEnv(), "surface", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

func (jbobject *GraphicsImage) SetFieldSurface(val int64) {
	err := jbobject.SetField(javabind.GetEnv(), "surface", val)
	if err != nil {
		panic(err)
	}

}


