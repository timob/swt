package swt

import "github.com/timob/javabind"

type CustomCTabFolderRendererInterface interface {
	JavaLangObjectInterface
}

type CustomCTabFolderRenderer struct {
	JavaLangObject
}

func (jbobject *CustomCTabFolderRenderer) PART_BODY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_BODY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolderRenderer) SetFieldPART_BODY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_BODY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolderRenderer) PART_HEADER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_HEADER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolderRenderer) SetFieldPART_HEADER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_HEADER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolderRenderer) PART_BORDER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_BORDER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolderRenderer) SetFieldPART_BORDER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_BORDER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolderRenderer) PART_BACKGROUND() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_BACKGROUND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolderRenderer) SetFieldPART_BACKGROUND(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_BACKGROUND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolderRenderer) PART_MAX_BUTTON() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_MAX_BUTTON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolderRenderer) SetFieldPART_MAX_BUTTON(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_MAX_BUTTON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolderRenderer) PART_MIN_BUTTON() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_MIN_BUTTON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolderRenderer) SetFieldPART_MIN_BUTTON(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_MIN_BUTTON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolderRenderer) PART_CHEVRON_BUTTON() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_CHEVRON_BUTTON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolderRenderer) SetFieldPART_CHEVRON_BUTTON(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_CHEVRON_BUTTON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolderRenderer) PART_CLOSE_BUTTON() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_CLOSE_BUTTON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolderRenderer) SetFieldPART_CLOSE_BUTTON(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "PART_CLOSE_BUTTON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomCTabFolderRenderer) MINIMUM_SIZE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "MINIMUM_SIZE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomCTabFolderRenderer) SetFieldMINIMUM_SIZE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/CTabFolderRenderer", "MINIMUM_SIZE", val)
	if err != nil {
		panic(err)
	}

}


