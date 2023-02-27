package kozelsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KozelskCounty interface {
    feud.County
    BBelev别列夫() 	feud.Barony
    BKozelsk科泽利斯克() 	feud.Barony
    BMezetsk梅泽茨克() 	feud.Barony
    BMosalsk莫萨利斯克() 	feud.Barony
    BSerpeisk谢尔佩伊斯克() 	feud.Barony
    BSosensky索先斯基() 	feud.Barony
    BSukhinichi苏希尼奇() 	feud.Barony
}

type 科泽利斯克KozelskCounty struct {
	feud.BaseCounty
	别列夫Belev 	feud.Barony
	科泽利斯克Kozelsk 	feud.Barony
	梅泽茨克Mezetsk 	feud.Barony
	莫萨利斯克Mosalsk 	feud.Barony
	谢尔佩伊斯克Serpeisk 	feud.Barony
	索先斯基Sosensky 	feud.Barony
	苏希尼奇Sukhinichi 	feud.Barony
}

func (c *科泽利斯克KozelskCounty) BBelev别列夫() feud.Barony {
	return c.别列夫Belev
}
    
func (c *科泽利斯克KozelskCounty) BKozelsk科泽利斯克() feud.Barony {
	return c.科泽利斯克Kozelsk
}
    
func (c *科泽利斯克KozelskCounty) BMezetsk梅泽茨克() feud.Barony {
	return c.梅泽茨克Mezetsk
}
    
func (c *科泽利斯克KozelskCounty) BMosalsk莫萨利斯克() feud.Barony {
	return c.莫萨利斯克Mosalsk
}
    
func (c *科泽利斯克KozelskCounty) BSerpeisk谢尔佩伊斯克() feud.Barony {
	return c.谢尔佩伊斯克Serpeisk
}
    
func (c *科泽利斯克KozelskCounty) BSosensky索先斯基() feud.Barony {
	return c.索先斯基Sosensky
}
    
func (c *科泽利斯克KozelskCounty) BSukhinichi苏希尼奇() feud.Barony {
	return c.苏希尼奇Sukhinichi
}
    
var CKozelsk科泽利斯克 KozelskCounty = &科泽利斯克KozelskCounty{}

func init() {
	f := CKozelsk科泽利斯克.(*科泽利斯克KozelskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1682",
		Title:     "kozelsk",
		TitleName: "科泽利斯克",
		TitleCode: "c_kozelsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.别列夫Belev = BBelev别列夫
	f.别列夫Belev.SetParent(f)
	
	f.科泽利斯克Kozelsk = BKozelsk科泽利斯克
	f.科泽利斯克Kozelsk.SetParent(f)
	
	f.梅泽茨克Mezetsk = BMezetsk梅泽茨克
	f.梅泽茨克Mezetsk.SetParent(f)
	
	f.莫萨利斯克Mosalsk = BMosalsk莫萨利斯克
	f.莫萨利斯克Mosalsk.SetParent(f)
	
	f.谢尔佩伊斯克Serpeisk = BSerpeisk谢尔佩伊斯克
	f.谢尔佩伊斯克Serpeisk.SetParent(f)
	
	f.索先斯基Sosensky = BSosensky索先斯基
	f.索先斯基Sosensky.SetParent(f)
	
	f.苏希尼奇Sukhinichi = BSukhinichi苏希尼奇
	f.苏希尼奇Sukhinichi.SetParent(f)
	
}
