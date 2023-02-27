package utva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UtvaCounty interface {
    feud.County
    BAktau阿克套() 	feud.Barony
    BAraltobe阿拉尔托别() 	feud.Barony
    BAshchesay阿谢_赛() 	feud.Barony
    BErsary叶尔萨雷() 	feud.Barony
    BKirovo基洛沃() 	feud.Barony
    BSaryumir萨雷乌米尔() 	feud.Barony
    BUtva乌特瓦() 	feud.Barony
}

type 乌特瓦UtvaCounty struct {
	feud.BaseCounty
	阿克套Aktau 	feud.Barony
	阿拉尔托别Araltobe 	feud.Barony
	阿谢_赛Ashchesay 	feud.Barony
	叶尔萨雷Ersary 	feud.Barony
	基洛沃Kirovo 	feud.Barony
	萨雷乌米尔Saryumir 	feud.Barony
	乌特瓦Utva 	feud.Barony
}

func (c *乌特瓦UtvaCounty) BAktau阿克套() feud.Barony {
	return c.阿克套Aktau
}
    
func (c *乌特瓦UtvaCounty) BAraltobe阿拉尔托别() feud.Barony {
	return c.阿拉尔托别Araltobe
}
    
func (c *乌特瓦UtvaCounty) BAshchesay阿谢_赛() feud.Barony {
	return c.阿谢_赛Ashchesay
}
    
func (c *乌特瓦UtvaCounty) BErsary叶尔萨雷() feud.Barony {
	return c.叶尔萨雷Ersary
}
    
func (c *乌特瓦UtvaCounty) BKirovo基洛沃() feud.Barony {
	return c.基洛沃Kirovo
}
    
func (c *乌特瓦UtvaCounty) BSaryumir萨雷乌米尔() feud.Barony {
	return c.萨雷乌米尔Saryumir
}
    
func (c *乌特瓦UtvaCounty) BUtva乌特瓦() feud.Barony {
	return c.乌特瓦Utva
}
    
var CUtva乌特瓦 UtvaCounty = &乌特瓦UtvaCounty{}

func init() {
	f := CUtva乌特瓦.(*乌特瓦UtvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1852",
		Title:     "utva",
		TitleName: "乌特瓦",
		TitleCode: "c_utva",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克套Aktau = BAktau阿克套
	f.阿克套Aktau.SetParent(f)
	
	f.阿拉尔托别Araltobe = BAraltobe阿拉尔托别
	f.阿拉尔托别Araltobe.SetParent(f)
	
	f.阿谢_赛Ashchesay = BAshchesay阿谢_赛
	f.阿谢_赛Ashchesay.SetParent(f)
	
	f.叶尔萨雷Ersary = BErsary叶尔萨雷
	f.叶尔萨雷Ersary.SetParent(f)
	
	f.基洛沃Kirovo = BKirovo基洛沃
	f.基洛沃Kirovo.SetParent(f)
	
	f.萨雷乌米尔Saryumir = BSaryumir萨雷乌米尔
	f.萨雷乌米尔Saryumir.SetParent(f)
	
	f.乌特瓦Utva = BUtva乌特瓦
	f.乌特瓦Utva.SetParent(f)
	
}
