package edessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EdessaCounty interface {
    feud.County
    BBile比勒() 	feud.Barony
    BEdessa埃德萨() 	feud.Barony
    BEdesurfa乌尔法() 	feud.Barony
    BHarran哈兰() 	feud.Barony
    BKaisun凯松() 	feud.Barony
    BSamsat萨姆萨特() 	feud.Barony
    BSruk苏尔克() 	feud.Barony
    BTulupe图鲁匹() 	feud.Barony
}

type 埃德萨EdessaCounty struct {
	feud.BaseCounty
	比勒Bile 	feud.Barony
	埃德萨Edessa 	feud.Barony
	乌尔法Edesurfa 	feud.Barony
	哈兰Harran 	feud.Barony
	凯松Kaisun 	feud.Barony
	萨姆萨特Samsat 	feud.Barony
	苏尔克Sruk 	feud.Barony
	图鲁匹Tulupe 	feud.Barony
}

func (c *埃德萨EdessaCounty) BBile比勒() feud.Barony {
	return c.比勒Bile
}
    
func (c *埃德萨EdessaCounty) BEdessa埃德萨() feud.Barony {
	return c.埃德萨Edessa
}
    
func (c *埃德萨EdessaCounty) BEdesurfa乌尔法() feud.Barony {
	return c.乌尔法Edesurfa
}
    
func (c *埃德萨EdessaCounty) BHarran哈兰() feud.Barony {
	return c.哈兰Harran
}
    
func (c *埃德萨EdessaCounty) BKaisun凯松() feud.Barony {
	return c.凯松Kaisun
}
    
func (c *埃德萨EdessaCounty) BSamsat萨姆萨特() feud.Barony {
	return c.萨姆萨特Samsat
}
    
func (c *埃德萨EdessaCounty) BSruk苏尔克() feud.Barony {
	return c.苏尔克Sruk
}
    
func (c *埃德萨EdessaCounty) BTulupe图鲁匹() feud.Barony {
	return c.图鲁匹Tulupe
}
    
var CEdessa埃德萨 EdessaCounty = &埃德萨EdessaCounty{}

func init() {
	f := CEdessa埃德萨.(*埃德萨EdessaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "699",
		Title:     "edessa",
		TitleName: "埃德萨",
		TitleCode: "c_edessa",
		Baronies:  map[string]feud.Barony{},
	}

	f.比勒Bile = BBile比勒
	f.比勒Bile.SetParent(f)
	
	f.埃德萨Edessa = BEdessa埃德萨
	f.埃德萨Edessa.SetParent(f)
	
	f.乌尔法Edesurfa = BEdesurfa乌尔法
	f.乌尔法Edesurfa.SetParent(f)
	
	f.哈兰Harran = BHarran哈兰
	f.哈兰Harran.SetParent(f)
	
	f.凯松Kaisun = BKaisun凯松
	f.凯松Kaisun.SetParent(f)
	
	f.萨姆萨特Samsat = BSamsat萨姆萨特
	f.萨姆萨特Samsat.SetParent(f)
	
	f.苏尔克Sruk = BSruk苏尔克
	f.苏尔克Sruk.SetParent(f)
	
	f.图鲁匹Tulupe = BTulupe图鲁匹
	f.图鲁匹Tulupe.SetParent(f)
	
}
