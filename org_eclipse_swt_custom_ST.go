package swt

import "github.com/timob/javabind"

type CustomSTInterface interface {
	JavaLangObjectInterface
}

type CustomST struct {
	JavaLangObject
}

// public org.eclipse.swt.custom.ST()
func NewCustomST() (*CustomST) {

	obj, err := javabind.GetEnv().NewObject("org/eclipse/swt/custom/ST")
	if err != nil {
		panic(err)
	}
	x := &CustomST{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *CustomST) LINE_UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "LINE_UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldLINE_UP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "LINE_UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) LINE_DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "LINE_DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldLINE_DOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "LINE_DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) LINE_START() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "LINE_START", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldLINE_START(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "LINE_START", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) LINE_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "LINE_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldLINE_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "LINE_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) COLUMN_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "COLUMN_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldCOLUMN_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "COLUMN_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) COLUMN_NEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "COLUMN_NEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldCOLUMN_NEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "COLUMN_NEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) PAGE_UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "PAGE_UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldPAGE_UP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "PAGE_UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) PAGE_DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "PAGE_DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldPAGE_DOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "PAGE_DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) WORD_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "WORD_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldWORD_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "WORD_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) WORD_NEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "WORD_NEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldWORD_NEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "WORD_NEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) TEXT_START() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "TEXT_START", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldTEXT_START(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "TEXT_START", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) TEXT_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "TEXT_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldTEXT_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "TEXT_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) WINDOW_START() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "WINDOW_START", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldWINDOW_START(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "WINDOW_START", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) WINDOW_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "WINDOW_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldWINDOW_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "WINDOW_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_ALL() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_ALL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_ALL(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_ALL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_LINE_UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_LINE_UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_LINE_UP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_LINE_UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_LINE_DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_LINE_DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_LINE_DOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_LINE_DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_LINE_START() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_LINE_START", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_LINE_START(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_LINE_START", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_LINE_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_LINE_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_LINE_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_LINE_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_COLUMN_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_COLUMN_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_COLUMN_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_COLUMN_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_COLUMN_NEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_COLUMN_NEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_COLUMN_NEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_COLUMN_NEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_PAGE_UP() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_PAGE_UP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_PAGE_UP(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_PAGE_UP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_PAGE_DOWN() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_PAGE_DOWN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_PAGE_DOWN(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_PAGE_DOWN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_WORD_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_WORD_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_WORD_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_WORD_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_WORD_NEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_WORD_NEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_WORD_NEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_WORD_NEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_TEXT_START() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_TEXT_START", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_TEXT_START(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_TEXT_START", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_TEXT_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_TEXT_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_TEXT_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_TEXT_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_WINDOW_START() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_WINDOW_START", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_WINDOW_START(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_WINDOW_START", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) SELECT_WINDOW_END() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "SELECT_WINDOW_END", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldSELECT_WINDOW_END(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "SELECT_WINDOW_END", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) CUT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "CUT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldCUT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "CUT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) COPY() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "COPY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldCOPY(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "COPY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) PASTE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "PASTE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldPASTE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "PASTE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) DELETE_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "DELETE_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldDELETE_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "DELETE_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) DELETE_NEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "DELETE_NEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldDELETE_NEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "DELETE_NEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) DELETE_WORD_PREVIOUS() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "DELETE_WORD_PREVIOUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldDELETE_WORD_PREVIOUS(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "DELETE_WORD_PREVIOUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) DELETE_WORD_NEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "DELETE_WORD_NEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldDELETE_WORD_NEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "DELETE_WORD_NEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) TOGGLE_OVERWRITE() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "TOGGLE_OVERWRITE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldTOGGLE_OVERWRITE(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "TOGGLE_OVERWRITE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) TOGGLE_BLOCKSELECTION() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "TOGGLE_BLOCKSELECTION", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldTOGGLE_BLOCKSELECTION(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "TOGGLE_BLOCKSELECTION", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) BULLET_DOT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "BULLET_DOT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldBULLET_DOT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "BULLET_DOT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) BULLET_NUMBER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "BULLET_NUMBER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldBULLET_NUMBER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "BULLET_NUMBER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) BULLET_LETTER_LOWER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "BULLET_LETTER_LOWER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldBULLET_LETTER_LOWER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "BULLET_LETTER_LOWER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) BULLET_LETTER_UPPER() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "BULLET_LETTER_UPPER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldBULLET_LETTER_UPPER(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "BULLET_LETTER_UPPER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) BULLET_TEXT() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "BULLET_TEXT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldBULLET_TEXT(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "BULLET_TEXT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) BULLET_CUSTOM() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "BULLET_CUSTOM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldBULLET_CUSTOM(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "BULLET_CUSTOM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) ExtendedModify() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "ExtendedModify", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldExtendedModify(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "ExtendedModify", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) LineGetBackground() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "LineGetBackground", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldLineGetBackground(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "LineGetBackground", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) LineGetStyle() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "LineGetStyle", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldLineGetStyle(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "LineGetStyle", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) TextChanging() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "TextChanging", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldTextChanging(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "TextChanging", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) TextSet() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "TextSet", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldTextSet(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "TextSet", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) VerifyKey() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "VerifyKey", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldVerifyKey(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "VerifyKey", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) TextChanged() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "TextChanged", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldTextChanged(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "TextChanged", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) LineGetSegments() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "LineGetSegments", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldLineGetSegments(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "LineGetSegments", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) PaintObject() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "PaintObject", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldPaintObject(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "PaintObject", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) WordNext() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "WordNext", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldWordNext(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "WordNext", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) WordPrevious() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "WordPrevious", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldWordPrevious(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "WordPrevious", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *CustomST) CaretMoved() int {
	jret, err := javabind.GetEnv().GetStaticField("org/eclipse/swt/custom/ST", "CaretMoved", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *CustomST) SetFieldCaretMoved(val int) {
	err := javabind.GetEnv().SetStaticField("org/eclipse/swt/custom/ST", "CaretMoved", val)
	if err != nil {
		panic(err)
	}

}


