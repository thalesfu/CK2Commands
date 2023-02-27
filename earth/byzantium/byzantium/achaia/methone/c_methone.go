package methone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MethoneCounty interface {
    feud.County
    BAndrousa安德鲁萨() 	feud.Barony
    BCoron科龙() 	feud.Barony
    BGritzena格里采纳() 	feud.Barony
    BKalamata卡拉马塔() 	feud.Barony
    BKarytaina卡里泰纳() 	feud.Barony
    BKiparissia基帕里西亚() 	feud.Barony
    BModon莫顿() 	feud.Barony
    BPilos皮洛斯() 	feud.Barony
}

type 墨托涅MethoneCounty struct {
	feud.BaseCounty
	安德鲁萨Androusa 	feud.Barony
	科龙Coron 	feud.Barony
	格里采纳Gritzena 	feud.Barony
	卡拉马塔Kalamata 	feud.Barony
	卡里泰纳Karytaina 	feud.Barony
	基帕里西亚Kiparissia 	feud.Barony
	莫顿Modon 	feud.Barony
	皮洛斯Pilos 	feud.Barony
}

func (c *墨托涅MethoneCounty) BAndrousa安德鲁萨() feud.Barony {
	return c.安德鲁萨Androusa
}
    
func (c *墨托涅MethoneCounty) BCoron科龙() feud.Barony {
	return c.科龙Coron
}
    
func (c *墨托涅MethoneCounty) BGritzena格里采纳() feud.Barony {
	return c.格里采纳Gritzena
}
    
func (c *墨托涅MethoneCounty) BKalamata卡拉马塔() feud.Barony {
	return c.卡拉马塔Kalamata
}
    
func (c *墨托涅MethoneCounty) BKarytaina卡里泰纳() feud.Barony {
	return c.卡里泰纳Karytaina
}
    
func (c *墨托涅MethoneCounty) BKiparissia基帕里西亚() feud.Barony {
	return c.基帕里西亚Kiparissia
}
    
func (c *墨托涅MethoneCounty) BModon莫顿() feud.Barony {
	return c.莫顿Modon
}
    
func (c *墨托涅MethoneCounty) BPilos皮洛斯() feud.Barony {
	return c.皮洛斯Pilos
}
    
var CMethone墨托涅 MethoneCounty = &墨托涅MethoneCounty{}

func init() {
	f := CMethone墨托涅.(*墨托涅MethoneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "477",
		Title:     "methone",
		TitleName: "墨托涅",
		TitleCode: "c_methone",
		Baronies:  map[string]feud.Barony{},
	}

	f.安德鲁萨Androusa = BAndrousa安德鲁萨
	f.安德鲁萨Androusa.SetParent(f)
	
	f.科龙Coron = BCoron科龙
	f.科龙Coron.SetParent(f)
	
	f.格里采纳Gritzena = BGritzena格里采纳
	f.格里采纳Gritzena.SetParent(f)
	
	f.卡拉马塔Kalamata = BKalamata卡拉马塔
	f.卡拉马塔Kalamata.SetParent(f)
	
	f.卡里泰纳Karytaina = BKarytaina卡里泰纳
	f.卡里泰纳Karytaina.SetParent(f)
	
	f.基帕里西亚Kiparissia = BKiparissia基帕里西亚
	f.基帕里西亚Kiparissia.SetParent(f)
	
	f.莫顿Modon = BModon莫顿
	f.莫顿Modon.SetParent(f)
	
	f.皮洛斯Pilos = BPilos皮洛斯
	f.皮洛斯Pilos.SetParent(f)
	
}
