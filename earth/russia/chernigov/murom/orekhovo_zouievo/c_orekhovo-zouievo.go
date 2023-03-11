package orekhovo_zouievo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Orekhovo_zouievoCounty interface {
    feud.County
    BDegtyanoye杰格佳诺耶() 	feud.Barony
    BGorki戈尔基() 	feud.Barony
    BMalakhovo马拉霍沃() 	feud.Barony
    BMurmino穆尔米诺() 	feud.Barony
    BOrekhovo_zouievo奥列霍沃_祖耶沃() 	feud.Barony
    BTuma图马() 	feud.Barony
    BYegoryevsk叶戈里耶夫斯克() 	feud.Barony
}

type 奥列霍沃_祖耶沃Orekhovo_zouievoCounty struct {
	feud.BaseCounty
	杰格佳诺耶Degtyanoye 	feud.Barony
	戈尔基Gorki 	feud.Barony
	马拉霍沃Malakhovo 	feud.Barony
	穆尔米诺Murmino 	feud.Barony
	奥列霍沃_祖耶沃Orekhovo_zouievo 	feud.Barony
	图马Tuma 	feud.Barony
	叶戈里耶夫斯克Yegoryevsk 	feud.Barony
}

func (c *奥列霍沃_祖耶沃Orekhovo_zouievoCounty) BDegtyanoye杰格佳诺耶() feud.Barony {
	return c.杰格佳诺耶Degtyanoye
}
    
func (c *奥列霍沃_祖耶沃Orekhovo_zouievoCounty) BGorki戈尔基() feud.Barony {
	return c.戈尔基Gorki
}
    
func (c *奥列霍沃_祖耶沃Orekhovo_zouievoCounty) BMalakhovo马拉霍沃() feud.Barony {
	return c.马拉霍沃Malakhovo
}
    
func (c *奥列霍沃_祖耶沃Orekhovo_zouievoCounty) BMurmino穆尔米诺() feud.Barony {
	return c.穆尔米诺Murmino
}
    
func (c *奥列霍沃_祖耶沃Orekhovo_zouievoCounty) BOrekhovo_zouievo奥列霍沃_祖耶沃() feud.Barony {
	return c.奥列霍沃_祖耶沃Orekhovo_zouievo
}
    
func (c *奥列霍沃_祖耶沃Orekhovo_zouievoCounty) BTuma图马() feud.Barony {
	return c.图马Tuma
}
    
func (c *奥列霍沃_祖耶沃Orekhovo_zouievoCounty) BYegoryevsk叶戈里耶夫斯克() feud.Barony {
	return c.叶戈里耶夫斯克Yegoryevsk
}
    
var COrekhovo_zouievo奥列霍沃_祖耶沃 Orekhovo_zouievoCounty = &奥列霍沃_祖耶沃Orekhovo_zouievoCounty{}

func init() {
	f := COrekhovo_zouievo奥列霍沃_祖耶沃.(*奥列霍沃_祖耶沃Orekhovo_zouievoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1678",
		Title:     "orekhovo_zouievo",
		TitleName: "奥列霍沃_祖耶沃",
		TitleCode: "c_orekhovo-zouievo",
		Baronies:  map[string]feud.Barony{},
	}

	f.杰格佳诺耶Degtyanoye = BDegtyanoye杰格佳诺耶
	f.杰格佳诺耶Degtyanoye.SetParent(f)
	
	f.戈尔基Gorki = BGorki戈尔基
	f.戈尔基Gorki.SetParent(f)
	
	f.马拉霍沃Malakhovo = BMalakhovo马拉霍沃
	f.马拉霍沃Malakhovo.SetParent(f)
	
	f.穆尔米诺Murmino = BMurmino穆尔米诺
	f.穆尔米诺Murmino.SetParent(f)
	
	f.奥列霍沃_祖耶沃Orekhovo_zouievo = BOrekhovo_zouievo奥列霍沃_祖耶沃
	f.奥列霍沃_祖耶沃Orekhovo_zouievo.SetParent(f)
	
	f.图马Tuma = BTuma图马
	f.图马Tuma.SetParent(f)
	
	f.叶戈里耶夫斯克Yegoryevsk = BYegoryevsk叶戈里耶夫斯克
	f.叶戈里耶夫斯克Yegoryevsk.SetParent(f)
	
}
