package brycheiniog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrycheiniogCounty interface {
    feud.County
    BBrecon布雷肯() 	feud.Barony
    BCradoc克拉多克() 	feud.Barony
    BCrickhowell克里克豪威尔() 	feud.Barony
    BLlanddew兰迪尤() 	feud.Barony
    BMerthyr_tydfil梅瑟蒂德菲尔() 	feud.Barony
    BTalgarth塔尔加斯() 	feud.Barony
    BTretower特雷陶尔() 	feud.Barony
}

type 布里黑尼奥格BrycheiniogCounty struct {
	feud.BaseCounty
	布雷肯Brecon 	feud.Barony
	克拉多克Cradoc 	feud.Barony
	克里克豪威尔Crickhowell 	feud.Barony
	兰迪尤Llanddew 	feud.Barony
	梅瑟蒂德菲尔Merthyr_tydfil 	feud.Barony
	塔尔加斯Talgarth 	feud.Barony
	特雷陶尔Tretower 	feud.Barony
}

func (c *布里黑尼奥格BrycheiniogCounty) BBrecon布雷肯() feud.Barony {
	return c.布雷肯Brecon
}
    
func (c *布里黑尼奥格BrycheiniogCounty) BCradoc克拉多克() feud.Barony {
	return c.克拉多克Cradoc
}
    
func (c *布里黑尼奥格BrycheiniogCounty) BCrickhowell克里克豪威尔() feud.Barony {
	return c.克里克豪威尔Crickhowell
}
    
func (c *布里黑尼奥格BrycheiniogCounty) BLlanddew兰迪尤() feud.Barony {
	return c.兰迪尤Llanddew
}
    
func (c *布里黑尼奥格BrycheiniogCounty) BMerthyr_tydfil梅瑟蒂德菲尔() feud.Barony {
	return c.梅瑟蒂德菲尔Merthyr_tydfil
}
    
func (c *布里黑尼奥格BrycheiniogCounty) BTalgarth塔尔加斯() feud.Barony {
	return c.塔尔加斯Talgarth
}
    
func (c *布里黑尼奥格BrycheiniogCounty) BTretower特雷陶尔() feud.Barony {
	return c.特雷陶尔Tretower
}
    
var CBrycheiniog布里黑尼奥格 BrycheiniogCounty = &布里黑尼奥格BrycheiniogCounty{}

func init() {
	f := CBrycheiniog布里黑尼奥格.(*布里黑尼奥格BrycheiniogCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1945",
		Title:     "brycheiniog",
		TitleName: "布里黑尼奥格",
		TitleCode: "c_brycheiniog",
		Baronies:  map[string]feud.Barony{},
	}

	f.布雷肯Brecon = BBrecon布雷肯
	f.布雷肯Brecon.SetParent(f)
	
	f.克拉多克Cradoc = BCradoc克拉多克
	f.克拉多克Cradoc.SetParent(f)
	
	f.克里克豪威尔Crickhowell = BCrickhowell克里克豪威尔
	f.克里克豪威尔Crickhowell.SetParent(f)
	
	f.兰迪尤Llanddew = BLlanddew兰迪尤
	f.兰迪尤Llanddew.SetParent(f)
	
	f.梅瑟蒂德菲尔Merthyr_tydfil = BMerthyr_tydfil梅瑟蒂德菲尔
	f.梅瑟蒂德菲尔Merthyr_tydfil.SetParent(f)
	
	f.塔尔加斯Talgarth = BTalgarth塔尔加斯
	f.塔尔加斯Talgarth.SetParent(f)
	
	f.特雷陶尔Tretower = BTretower特雷陶尔
	f.特雷陶尔Tretower.SetParent(f)
	
}
