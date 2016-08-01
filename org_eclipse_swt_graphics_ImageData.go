package swt

import "github.com/timob/javabind"

type GraphicsImageDataInterface interface {
	JavaLangObjectInterface

	// public static org.eclipse.swt.graphics.ImageData internal_new(int, int, int, org.eclipse.swt.graphics.PaletteData, int, byte[], int, byte[], byte[], int, int, int, int, int, int, int)
	Internal_new(a int, b int, c int, d GraphicsPaletteDataInterface, e int, f []byte, g int, h []byte, i []byte, j int, k int, l int, m int, n int, o int, p int) *GraphicsImageData

	// public java.lang.Object clone()
	Clone() *JavaLangObject

	// public int getAlpha(int, int)
	GetAlpha(a int, b int) int

	// public void getAlphas(int, int, int, byte[], int)
	GetAlphas(a int, b int, c int, d []byte, e int) 

	// public int getPixel(int, int)
	GetPixel(a int, b int) int

	// public void getPixels(int, int, int, byte[], int)
	GetPixels(a int, b int, c int, d []byte, e int) 

	// public void getPixels(int, int, int, int[], int)
	GetPixels2(a int, b int, c int, d []int, e int) 

	// public org.eclipse.swt.graphics.RGB[] getRGBs()
	GetRGBs() []*GraphicsRGB

	// public org.eclipse.swt.graphics.ImageData getTransparencyMask()
	GetTransparencyMask() *GraphicsImageData

	// public int getTransparencyType()
	GetTransparencyType() int

	// public org.eclipse.swt.graphics.ImageData scaledTo(int, int)
	ScaledTo(a int, b int) *GraphicsImageData

	// public void setAlpha(int, int, int)
	SetAlpha(a int, b int, c int) 

	// public void setAlphas(int, int, int, byte[], int)
	SetAlphas(a int, b int, c int, d []byte, e int) 

	// public void setPixel(int, int, int)
	SetPixel(a int, b int, c int) 

	// public void setPixels(int, int, int, byte[], int)
	SetPixels(a int, b int, c int, d []byte, e int) 

	// public void setPixels(int, int, int, int[], int)
	SetPixels2(a int, b int, c int, d []int, e int) 
}

type GraphicsImageData struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.ImageData(int, int, int, org.eclipse.swt.graphics.PaletteData)
func NewGraphicsImageData2(a int, b int, c int, d GraphicsPaletteDataInterface) (*GraphicsImageData) {
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/ImageData", a, b, c, conv_d.Value().Cast("org/eclipse/swt/graphics/PaletteData"))
	if err != nil {
		panic(err)
	}
	conv_d.CleanUp()
	x := &GraphicsImageData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.ImageData(int, int, int, org.eclipse.swt.graphics.PaletteData, int, byte[])
func NewGraphicsImageData3(a int, b int, c int, d GraphicsPaletteDataInterface, e int, f []byte) (*GraphicsImageData) {
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/ImageData", a, b, c, conv_d.Value().Cast("org/eclipse/swt/graphics/PaletteData"), e, f)
	if err != nil {
		panic(err)
	}
	conv_d.CleanUp()
	x := &GraphicsImageData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.ImageData(java.lang.String)
func NewGraphicsImageData(a string) (*GraphicsImageData) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/ImageData", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsImageData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.eclipse.swt.graphics.ImageData internal_new(int, int, int, org.eclipse.swt.graphics.PaletteData, int, byte[], int, byte[], byte[], int, int, int, int, int, int, int)
func (jbobject *GraphicsImageData) Internal_new(a int, b int, c int, d GraphicsPaletteDataInterface, e int, f []byte, g int, h []byte, i []byte, j int, k int, l int, m int, n int, o int, p int) *GraphicsImageData {
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/eclipse/swt/graphics/ImageData", "internal_new", "org/eclipse/swt/graphics/ImageData", a, b, c, conv_d.Value().Cast("org/eclipse/swt/graphics/PaletteData"), e, f, g, h, i, j, k, l, m, n, o, p)
	if err != nil {
		panic(err)
	}
	conv_d.CleanUp()
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

// public java.lang.Object clone()
func (jbobject *GraphicsImageData) Clone() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public int getAlpha(int, int)
func (jbobject *GraphicsImageData) GetAlpha(a int, b int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAlpha", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void getAlphas(int, int, int, byte[], int)
func (jbobject *GraphicsImageData) GetAlphas(a int, b int, c int, d []byte, e int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getAlphas", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

// public int getPixel(int, int)
func (jbobject *GraphicsImageData) GetPixel(a int, b int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPixel", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void getPixels(int, int, int, byte[], int)
func (jbobject *GraphicsImageData) GetPixels(a int, b int, c int, d []byte, e int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getPixels", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

// public void getPixels(int, int, int, int[], int)
func (jbobject *GraphicsImageData) GetPixels2(a int, b int, c int, d []int, e int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getPixels", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

// public org.eclipse.swt.graphics.RGB[] getRGBs()
func (jbobject *GraphicsImageData) GetRGBs() []*GraphicsRGB {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRGBs", javabind.ObjectArrayType("org/eclipse/swt/graphics/RGB"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/eclipse/swt/graphics/RGB")
	dst := new([]*GraphicsRGB)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.eclipse.swt.graphics.ImageData getTransparencyMask()
func (jbobject *GraphicsImageData) GetTransparencyMask() *GraphicsImageData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTransparencyMask", "org/eclipse/swt/graphics/ImageData")
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

// public int getTransparencyType()
func (jbobject *GraphicsImageData) GetTransparencyType() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTransparencyType", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.eclipse.swt.graphics.ImageData scaledTo(int, int)
func (jbobject *GraphicsImageData) ScaledTo(a int, b int) *GraphicsImageData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "scaledTo", "org/eclipse/swt/graphics/ImageData", a, b)
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

// public void setAlpha(int, int, int)
func (jbobject *GraphicsImageData) SetAlpha(a int, b int, c int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlpha", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void setAlphas(int, int, int, byte[], int)
func (jbobject *GraphicsImageData) SetAlphas(a int, b int, c int, d []byte, e int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAlphas", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

// public void setPixel(int, int, int)
func (jbobject *GraphicsImageData) SetPixel(a int, b int, c int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPixel", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void setPixels(int, int, int, byte[], int)
func (jbobject *GraphicsImageData) SetPixels(a int, b int, c int, d []byte, e int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPixels", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

// public void setPixels(int, int, int, int[], int)
func (jbobject *GraphicsImageData) SetPixels2(a int, b int, c int, d []int, e int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPixels", javabind.Void, a, b, c, d, e)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) Width() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "width", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldWidth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "width", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) Height() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "height", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldHeight(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "height", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) Depth() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "depth", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldDepth(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "depth", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) ScanlinePad() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "scanlinePad", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldScanlinePad(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "scanlinePad", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) BytesPerLine() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "bytesPerLine", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldBytesPerLine(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "bytesPerLine", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) Data() []byte {
	jret, err := jbobject.GetField(javabind.GetEnv(), "data", javabind.Byte | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]byte)
}

func (jbobject *GraphicsImageData) SetFieldData(val []byte) {
	err := jbobject.SetField(javabind.GetEnv(), "data", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) Palette() *GraphicsPaletteData {
	jret, err := jbobject.GetField(javabind.GetEnv(), "palette", "org/eclipse/swt/graphics/PaletteData")
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
	unique_x := &GraphicsPaletteData{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *GraphicsImageData) SetFieldPalette(val GraphicsPaletteDataInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "palette", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsImageData) TransparentPixel() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "transparentPixel", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldTransparentPixel(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "transparentPixel", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) MaskData() []byte {
	jret, err := jbobject.GetField(javabind.GetEnv(), "maskData", javabind.Byte | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]byte)
}

func (jbobject *GraphicsImageData) SetFieldMaskData(val []byte) {
	err := jbobject.SetField(javabind.GetEnv(), "maskData", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) MaskPad() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "maskPad", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldMaskPad(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "maskPad", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) AlphaData() []byte {
	jret, err := jbobject.GetField(javabind.GetEnv(), "alphaData", javabind.Byte | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]byte)
}

func (jbobject *GraphicsImageData) SetFieldAlphaData(val []byte) {
	err := jbobject.SetField(javabind.GetEnv(), "alphaData", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) Alpha() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "alpha", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldAlpha(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "alpha", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) Type() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "type", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldType(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "type", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) X() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "x", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldX(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "x", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) Y() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "y", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldY(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "y", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) DisposalMethod() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "disposalMethod", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldDisposalMethod(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "disposalMethod", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsImageData) DelayTime() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "delayTime", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsImageData) SetFieldDelayTime(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "delayTime", val)
	if err != nil {
		panic(err)
	}

}


