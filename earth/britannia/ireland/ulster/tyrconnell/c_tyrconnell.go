package tyrconnell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TyrconnellCounty interface {
    feud.County
    BBallymacswiney巴利马克斯万尼() 	feud.Barony
    BBallyshannon巴利香农() 	feud.Barony
    BDonegal多尼戈尔() 	feud.Barony
    BFahan法罕() 	feud.Barony
    BGartan加坦() 	feud.Barony
    BKilmacrenan基尔麦克雷南() 	feud.Barony
    BMoville莫维尔() 	feud.Barony
    BRaphoe拉福() 	feud.Barony
}

type 蒂尔康奈TyrconnellCounty struct {
	feud.BaseCounty
	巴利马克斯万尼Ballymacswiney 	feud.Barony
	巴利香农Ballyshannon 	feud.Barony
	多尼戈尔Donegal 	feud.Barony
	法罕Fahan 	feud.Barony
	加坦Gartan 	feud.Barony
	基尔麦克雷南Kilmacrenan 	feud.Barony
	莫维尔Moville 	feud.Barony
	拉福Raphoe 	feud.Barony
}

func (c *蒂尔康奈TyrconnellCounty) BBallymacswiney巴利马克斯万尼() feud.Barony {
	return c.巴利马克斯万尼Ballymacswiney
}
    
func (c *蒂尔康奈TyrconnellCounty) BBallyshannon巴利香农() feud.Barony {
	return c.巴利香农Ballyshannon
}
    
func (c *蒂尔康奈TyrconnellCounty) BDonegal多尼戈尔() feud.Barony {
	return c.多尼戈尔Donegal
}
    
func (c *蒂尔康奈TyrconnellCounty) BFahan法罕() feud.Barony {
	return c.法罕Fahan
}
    
func (c *蒂尔康奈TyrconnellCounty) BGartan加坦() feud.Barony {
	return c.加坦Gartan
}
    
func (c *蒂尔康奈TyrconnellCounty) BKilmacrenan基尔麦克雷南() feud.Barony {
	return c.基尔麦克雷南Kilmacrenan
}
    
func (c *蒂尔康奈TyrconnellCounty) BMoville莫维尔() feud.Barony {
	return c.莫维尔Moville
}
    
func (c *蒂尔康奈TyrconnellCounty) BRaphoe拉福() feud.Barony {
	return c.拉福Raphoe
}
    
var CTyrconnell蒂尔康奈 TyrconnellCounty = &蒂尔康奈TyrconnellCounty{}

func init() {
	f := CTyrconnell蒂尔康奈.(*蒂尔康奈TyrconnellCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "3",
		Title:     "tyrconnell",
		TitleName: "蒂尔康奈",
		TitleCode: "c_tyrconnell",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴利马克斯万尼Ballymacswiney = BBallymacswiney巴利马克斯万尼
	f.巴利马克斯万尼Ballymacswiney.SetParent(f)
	
	f.巴利香农Ballyshannon = BBallyshannon巴利香农
	f.巴利香农Ballyshannon.SetParent(f)
	
	f.多尼戈尔Donegal = BDonegal多尼戈尔
	f.多尼戈尔Donegal.SetParent(f)
	
	f.法罕Fahan = BFahan法罕
	f.法罕Fahan.SetParent(f)
	
	f.加坦Gartan = BGartan加坦
	f.加坦Gartan.SetParent(f)
	
	f.基尔麦克雷南Kilmacrenan = BKilmacrenan基尔麦克雷南
	f.基尔麦克雷南Kilmacrenan.SetParent(f)
	
	f.莫维尔Moville = BMoville莫维尔
	f.莫维尔Moville.SetParent(f)
	
	f.拉福Raphoe = BRaphoe拉福
	f.拉福Raphoe.SetParent(f)
	
}
