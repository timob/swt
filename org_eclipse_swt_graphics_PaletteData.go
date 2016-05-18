package swt

import "github.com/timob/javabind"

type GraphicsPaletteDataInterface interface {
	JavaLangObjectInterface

	// public int getPixel(org.eclipse.swt.graphics.RGB)
	GetPixel(a GraphicsRGBInterface) int

	// public org.eclipse.swt.graphics.RGB getRGB(int)
	GetRGB(a int) *GraphicsRGB

	// public org.eclipse.swt.graphics.RGB[] getRGBs()
	GetRGBs() []*GraphicsRGB
}

type GraphicsPaletteData struct {
	JavaLangObject
}

// public org.eclipse.swt.graphics.PaletteData(org.eclipse.swt.graphics.RGB[])
func NewGraphicsPaletteData2(a []*GraphicsRGB) (*GraphicsPaletteData) {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/RGB")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/PaletteData", conv_a.Value().Cast("org/eclipse/swt/graphics/RGB"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &GraphicsPaletteData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.eclipse.swt.graphics.PaletteData(int, int, int)
func NewGraphicsPaletteData(a int, b int, c int) (*GraphicsPaletteData) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/graphics/PaletteData", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &GraphicsPaletteData{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getPixel(org.eclipse.swt.graphics.RGB)
func (jbobject *GraphicsPaletteData) GetPixel(a GraphicsRGBInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPixel", javabind.Int, conv_a.Value().Cast("org/eclipse/swt/graphics/RGB"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public org.eclipse.swt.graphics.RGB getRGB(int)
func (jbobject *GraphicsPaletteData) GetRGB(a int) *GraphicsRGB {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRGB", "org/eclipse/swt/graphics/RGB", a)
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
	unique_x := &GraphicsRGB{}
	unique_x.Callable = dst
	return unique_x
}

// public org.eclipse.swt.graphics.RGB[] getRGBs()
func (jbobject *GraphicsPaletteData) GetRGBs() []*GraphicsRGB {
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

func (jbobject *GraphicsPaletteData) IsDirect() bool {
	jret, err := jbobject.GetField(javabind.GetEnv(), "isDirect", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *GraphicsPaletteData) SetFieldIsDirect(val bool) {
	err := jbobject.SetField(javabind.GetEnv(), "isDirect", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsPaletteData) Colors() []*GraphicsRGB {
	jret, err := jbobject.GetField(javabind.GetEnv(), "colors", javabind.ObjectArrayType("org/eclipse/swt/graphics/RGB"))
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

func (jbobject *GraphicsPaletteData) SetFieldColors(val []*GraphicsRGB) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/eclipse/swt/graphics/RGB")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "colors", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *GraphicsPaletteData) RedMask() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "redMask", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsPaletteData) SetFieldRedMask(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "redMask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsPaletteData) GreenMask() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "greenMask", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsPaletteData) SetFieldGreenMask(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "greenMask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsPaletteData) BlueMask() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "blueMask", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsPaletteData) SetFieldBlueMask(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "blueMask", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsPaletteData) RedShift() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "redShift", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsPaletteData) SetFieldRedShift(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "redShift", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsPaletteData) GreenShift() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "greenShift", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsPaletteData) SetFieldGreenShift(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "greenShift", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *GraphicsPaletteData) BlueShift() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "blueShift", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *GraphicsPaletteData) SetFieldBlueShift(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "blueShift", val)
	if err != nil {
		panic(err)
	}

}


