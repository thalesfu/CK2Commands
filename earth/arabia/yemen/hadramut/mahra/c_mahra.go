package mahra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MahraCounty interface {
    feud.County
    BAlhalya哈勒亚() 	feud.Barony
    BAlkurah库拉() 	feud.Barony
    BDamqwat达玛瓦特() 	feud.Barony
    BGhaydah盖达() 	feud.Barony
    BHaswayn哈斯韦因() 	feud.Barony
    BItab伊塔卜() 	feud.Barony
    BNishtun尼什通() 	feud.Barony
    BQishn基什恩() 	feud.Barony
}

type 马哈拉MahraCounty struct {
	feud.BaseCounty
	哈勒亚Alhalya 	feud.Barony
	库拉Alkurah 	feud.Barony
	达玛瓦特Damqwat 	feud.Barony
	盖达Ghaydah 	feud.Barony
	哈斯韦因Haswayn 	feud.Barony
	伊塔卜Itab 	feud.Barony
	尼什通Nishtun 	feud.Barony
	基什恩Qishn 	feud.Barony
}

func (c *马哈拉MahraCounty) BAlhalya哈勒亚() feud.Barony {
	return c.哈勒亚Alhalya
}
    
func (c *马哈拉MahraCounty) BAlkurah库拉() feud.Barony {
	return c.库拉Alkurah
}
    
func (c *马哈拉MahraCounty) BDamqwat达玛瓦特() feud.Barony {
	return c.达玛瓦特Damqwat
}
    
func (c *马哈拉MahraCounty) BGhaydah盖达() feud.Barony {
	return c.盖达Ghaydah
}
    
func (c *马哈拉MahraCounty) BHaswayn哈斯韦因() feud.Barony {
	return c.哈斯韦因Haswayn
}
    
func (c *马哈拉MahraCounty) BItab伊塔卜() feud.Barony {
	return c.伊塔卜Itab
}
    
func (c *马哈拉MahraCounty) BNishtun尼什通() feud.Barony {
	return c.尼什通Nishtun
}
    
func (c *马哈拉MahraCounty) BQishn基什恩() feud.Barony {
	return c.基什恩Qishn
}
    
var CMahra马哈拉 MahraCounty = &马哈拉MahraCounty{}

func init() {
	f := CMahra马哈拉.(*马哈拉MahraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "855",
		Title:     "mahra",
		TitleName: "马哈拉",
		TitleCode: "c_mahra",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈勒亚Alhalya = BAlhalya哈勒亚
	f.哈勒亚Alhalya.SetParent(f)
	
	f.库拉Alkurah = BAlkurah库拉
	f.库拉Alkurah.SetParent(f)
	
	f.达玛瓦特Damqwat = BDamqwat达玛瓦特
	f.达玛瓦特Damqwat.SetParent(f)
	
	f.盖达Ghaydah = BGhaydah盖达
	f.盖达Ghaydah.SetParent(f)
	
	f.哈斯韦因Haswayn = BHaswayn哈斯韦因
	f.哈斯韦因Haswayn.SetParent(f)
	
	f.伊塔卜Itab = BItab伊塔卜
	f.伊塔卜Itab.SetParent(f)
	
	f.尼什通Nishtun = BNishtun尼什通
	f.尼什通Nishtun.SetParent(f)
	
	f.基什恩Qishn = BQishn基什恩
	f.基什恩Qishn.SetParent(f)
	
}
