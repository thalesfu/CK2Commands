package warzazat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WarzazatCounty interface {
    feud.County
    BChtaouna什塔乌纳() 	feud.Barony
    BEl_jehifa杰希费() 	feud.Barony
    BQalat_sless斯莱斯堡() 	feud.Barony
    BShwan施旺() 	feud.Barony
    BTasagdah泰塞格达() 	feud.Barony
    BTimhata蒂姆哈塔() 	feud.Barony
    BWarzazat瓦尔扎扎特() 	feud.Barony
}

type 瓦尔扎扎特WarzazatCounty struct {
	feud.BaseCounty
	什塔乌纳Chtaouna 	feud.Barony
	杰希费El_jehifa 	feud.Barony
	斯莱斯堡Qalat_sless 	feud.Barony
	施旺Shwan 	feud.Barony
	泰塞格达Tasagdah 	feud.Barony
	蒂姆哈塔Timhata 	feud.Barony
	瓦尔扎扎特Warzazat 	feud.Barony
}

func (c *瓦尔扎扎特WarzazatCounty) BChtaouna什塔乌纳() feud.Barony {
	return c.什塔乌纳Chtaouna
}
    
func (c *瓦尔扎扎特WarzazatCounty) BEl_jehifa杰希费() feud.Barony {
	return c.杰希费El_jehifa
}
    
func (c *瓦尔扎扎特WarzazatCounty) BQalat_sless斯莱斯堡() feud.Barony {
	return c.斯莱斯堡Qalat_sless
}
    
func (c *瓦尔扎扎特WarzazatCounty) BShwan施旺() feud.Barony {
	return c.施旺Shwan
}
    
func (c *瓦尔扎扎特WarzazatCounty) BTasagdah泰塞格达() feud.Barony {
	return c.泰塞格达Tasagdah
}
    
func (c *瓦尔扎扎特WarzazatCounty) BTimhata蒂姆哈塔() feud.Barony {
	return c.蒂姆哈塔Timhata
}
    
func (c *瓦尔扎扎特WarzazatCounty) BWarzazat瓦尔扎扎特() feud.Barony {
	return c.瓦尔扎扎特Warzazat
}
    
var CWarzazat瓦尔扎扎特 WarzazatCounty = &瓦尔扎扎特WarzazatCounty{}

func init() {
	f := CWarzazat瓦尔扎扎特.(*瓦尔扎扎特WarzazatCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1776",
		Title:     "warzazat",
		TitleName: "瓦尔扎扎特",
		TitleCode: "c_warzazat",
		Baronies:  map[string]feud.Barony{},
	}

	f.什塔乌纳Chtaouna = BChtaouna什塔乌纳
	f.什塔乌纳Chtaouna.SetParent(f)
	
	f.杰希费El_jehifa = BEl_jehifa杰希费
	f.杰希费El_jehifa.SetParent(f)
	
	f.斯莱斯堡Qalat_sless = BQalat_sless斯莱斯堡
	f.斯莱斯堡Qalat_sless.SetParent(f)
	
	f.施旺Shwan = BShwan施旺
	f.施旺Shwan.SetParent(f)
	
	f.泰塞格达Tasagdah = BTasagdah泰塞格达
	f.泰塞格达Tasagdah.SetParent(f)
	
	f.蒂姆哈塔Timhata = BTimhata蒂姆哈塔
	f.蒂姆哈塔Timhata.SetParent(f)
	
	f.瓦尔扎扎特Warzazat = BWarzazat瓦尔扎扎特
	f.瓦尔扎扎特Warzazat.SetParent(f)
	
}
