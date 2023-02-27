package dariya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DariyaCounty interface {
    feud.County
    BAl_ajfur阿杰富尔() 	feud.Barony
    BAl_qaryatan盖尔亚廷() 	feud.Barony
    BAn_nibaj尼拜杰() 	feud.Barony
    BAt_talabiya塔拉比亚() 	feud.Barony
    BDariya德拉伊耶() 	feud.Barony
    BFayd费德() 	feud.Barony
    BZubala祖巴莱() 	feud.Barony
}

type 德拉伊耶DariyaCounty struct {
	feud.BaseCounty
	阿杰富尔Al_ajfur 	feud.Barony
	盖尔亚廷Al_qaryatan 	feud.Barony
	尼拜杰An_nibaj 	feud.Barony
	塔拉比亚At_talabiya 	feud.Barony
	德拉伊耶Dariya 	feud.Barony
	费德Fayd 	feud.Barony
	祖巴莱Zubala 	feud.Barony
}

func (c *德拉伊耶DariyaCounty) BAl_ajfur阿杰富尔() feud.Barony {
	return c.阿杰富尔Al_ajfur
}
    
func (c *德拉伊耶DariyaCounty) BAl_qaryatan盖尔亚廷() feud.Barony {
	return c.盖尔亚廷Al_qaryatan
}
    
func (c *德拉伊耶DariyaCounty) BAn_nibaj尼拜杰() feud.Barony {
	return c.尼拜杰An_nibaj
}
    
func (c *德拉伊耶DariyaCounty) BAt_talabiya塔拉比亚() feud.Barony {
	return c.塔拉比亚At_talabiya
}
    
func (c *德拉伊耶DariyaCounty) BDariya德拉伊耶() feud.Barony {
	return c.德拉伊耶Dariya
}
    
func (c *德拉伊耶DariyaCounty) BFayd费德() feud.Barony {
	return c.费德Fayd
}
    
func (c *德拉伊耶DariyaCounty) BZubala祖巴莱() feud.Barony {
	return c.祖巴莱Zubala
}
    
var CDariya德拉伊耶 DariyaCounty = &德拉伊耶DariyaCounty{}

func init() {
	f := CDariya德拉伊耶.(*德拉伊耶DariyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1530",
		Title:     "dariya",
		TitleName: "德拉伊耶",
		TitleCode: "c_dariya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿杰富尔Al_ajfur = BAl_ajfur阿杰富尔
	f.阿杰富尔Al_ajfur.SetParent(f)
	
	f.盖尔亚廷Al_qaryatan = BAl_qaryatan盖尔亚廷
	f.盖尔亚廷Al_qaryatan.SetParent(f)
	
	f.尼拜杰An_nibaj = BAn_nibaj尼拜杰
	f.尼拜杰An_nibaj.SetParent(f)
	
	f.塔拉比亚At_talabiya = BAt_talabiya塔拉比亚
	f.塔拉比亚At_talabiya.SetParent(f)
	
	f.德拉伊耶Dariya = BDariya德拉伊耶
	f.德拉伊耶Dariya.SetParent(f)
	
	f.费德Fayd = BFayd费德
	f.费德Fayd.SetParent(f)
	
	f.祖巴莱Zubala = BZubala祖巴莱
	f.祖巴莱Zubala.SetParent(f)
	
}
