package gudbrandsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GudbrandsdalCounty interface {
    feud.County
    BDovre多夫勒() 	feud.Barony
    BGarmo加摩() 	feud.Barony
    BGausdal盖于斯达尔() 	feud.Barony
    BHeidal黑达尔() 	feud.Barony
    BLesja莱沙() 	feud.Barony
    BLom洛姆() 	feud.Barony
    BVaga沃戈() 	feud.Barony
}

type 居德布兰河谷GudbrandsdalCounty struct {
	feud.BaseCounty
	多夫勒Dovre 	feud.Barony
	加摩Garmo 	feud.Barony
	盖于斯达尔Gausdal 	feud.Barony
	黑达尔Heidal 	feud.Barony
	莱沙Lesja 	feud.Barony
	洛姆Lom 	feud.Barony
	沃戈Vaga 	feud.Barony
}

func (c *居德布兰河谷GudbrandsdalCounty) BDovre多夫勒() feud.Barony {
	return c.多夫勒Dovre
}
    
func (c *居德布兰河谷GudbrandsdalCounty) BGarmo加摩() feud.Barony {
	return c.加摩Garmo
}
    
func (c *居德布兰河谷GudbrandsdalCounty) BGausdal盖于斯达尔() feud.Barony {
	return c.盖于斯达尔Gausdal
}
    
func (c *居德布兰河谷GudbrandsdalCounty) BHeidal黑达尔() feud.Barony {
	return c.黑达尔Heidal
}
    
func (c *居德布兰河谷GudbrandsdalCounty) BLesja莱沙() feud.Barony {
	return c.莱沙Lesja
}
    
func (c *居德布兰河谷GudbrandsdalCounty) BLom洛姆() feud.Barony {
	return c.洛姆Lom
}
    
func (c *居德布兰河谷GudbrandsdalCounty) BVaga沃戈() feud.Barony {
	return c.沃戈Vaga
}
    
var CGudbrandsdal居德布兰河谷 GudbrandsdalCounty = &居德布兰河谷GudbrandsdalCounty{}

func init() {
	f := CGudbrandsdal居德布兰河谷.(*居德布兰河谷GudbrandsdalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1621",
		Title:     "gudbrandsdal",
		TitleName: "居德布兰河谷",
		TitleCode: "c_gudbrandsdal",
		Baronies:  map[string]feud.Barony{},
	}

	f.多夫勒Dovre = BDovre多夫勒
	f.多夫勒Dovre.SetParent(f)
	
	f.加摩Garmo = BGarmo加摩
	f.加摩Garmo.SetParent(f)
	
	f.盖于斯达尔Gausdal = BGausdal盖于斯达尔
	f.盖于斯达尔Gausdal.SetParent(f)
	
	f.黑达尔Heidal = BHeidal黑达尔
	f.黑达尔Heidal.SetParent(f)
	
	f.莱沙Lesja = BLesja莱沙
	f.莱沙Lesja.SetParent(f)
	
	f.洛姆Lom = BLom洛姆
	f.洛姆Lom.SetParent(f)
	
	f.沃戈Vaga = BVaga沃戈
	f.沃戈Vaga.SetParent(f)
	
}
