package draa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DraaCounty interface {
    feud.County
    BAlogoum艾卢古姆() 	feud.Barony
    BBeni_djenade贝尼杰纳德() 	feud.Barony
    BJenane_el_menzeh杰纳奈曼泽() 	feud.Barony
    BKasbaeljoua朱阿堡() 	feud.Barony
    BTamchennennt坦尚南特() 	feud.Barony
    BTazagourt塔扎古尔特() 	feud.Barony
    BTizgui提兹古() 	feud.Barony
}

type 德拉DraaCounty struct {
	feud.BaseCounty
	艾卢古姆Alogoum 	feud.Barony
	贝尼杰纳德Beni_djenade 	feud.Barony
	杰纳奈曼泽Jenane_el_menzeh 	feud.Barony
	朱阿堡Kasbaeljoua 	feud.Barony
	坦尚南特Tamchennennt 	feud.Barony
	塔扎古尔特Tazagourt 	feud.Barony
	提兹古Tizgui 	feud.Barony
}

func (c *德拉DraaCounty) BAlogoum艾卢古姆() feud.Barony {
	return c.艾卢古姆Alogoum
}
    
func (c *德拉DraaCounty) BBeni_djenade贝尼杰纳德() feud.Barony {
	return c.贝尼杰纳德Beni_djenade
}
    
func (c *德拉DraaCounty) BJenane_el_menzeh杰纳奈曼泽() feud.Barony {
	return c.杰纳奈曼泽Jenane_el_menzeh
}
    
func (c *德拉DraaCounty) BKasbaeljoua朱阿堡() feud.Barony {
	return c.朱阿堡Kasbaeljoua
}
    
func (c *德拉DraaCounty) BTamchennennt坦尚南特() feud.Barony {
	return c.坦尚南特Tamchennennt
}
    
func (c *德拉DraaCounty) BTazagourt塔扎古尔特() feud.Barony {
	return c.塔扎古尔特Tazagourt
}
    
func (c *德拉DraaCounty) BTizgui提兹古() feud.Barony {
	return c.提兹古Tizgui
}
    
var CDraa德拉 DraaCounty = &德拉DraaCounty{}

func init() {
	f := CDraa德拉.(*德拉DraaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1779",
		Title:     "draa",
		TitleName: "德拉",
		TitleCode: "c_draa",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾卢古姆Alogoum = BAlogoum艾卢古姆
	f.艾卢古姆Alogoum.SetParent(f)
	
	f.贝尼杰纳德Beni_djenade = BBeni_djenade贝尼杰纳德
	f.贝尼杰纳德Beni_djenade.SetParent(f)
	
	f.杰纳奈曼泽Jenane_el_menzeh = BJenane_el_menzeh杰纳奈曼泽
	f.杰纳奈曼泽Jenane_el_menzeh.SetParent(f)
	
	f.朱阿堡Kasbaeljoua = BKasbaeljoua朱阿堡
	f.朱阿堡Kasbaeljoua.SetParent(f)
	
	f.坦尚南特Tamchennennt = BTamchennennt坦尚南特
	f.坦尚南特Tamchennennt.SetParent(f)
	
	f.塔扎古尔特Tazagourt = BTazagourt塔扎古尔特
	f.塔扎古尔特Tazagourt.SetParent(f)
	
	f.提兹古Tizgui = BTizgui提兹古
	f.提兹古Tizgui.SetParent(f)
	
}
