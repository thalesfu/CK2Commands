package hlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HlynovCounty interface {
    feud.County
    BEgra叶格拉() 	feud.Barony
    BGlazkar格拉兹卡尔() 	feud.Barony
    BHlynov赫雷诺夫() 	feud.Barony
    BIzkar伊日卡尔() 	feud.Barony
    BKambarka坎巴尔卡() 	feud.Barony
    BMozjga莫日加() 	feud.Barony
    BWotka沃特卡() 	feud.Barony
}

type 赫雷诺夫HlynovCounty struct {
	feud.BaseCounty
	叶格拉Egra 	feud.Barony
	格拉兹卡尔Glazkar 	feud.Barony
	赫雷诺夫Hlynov 	feud.Barony
	伊日卡尔Izkar 	feud.Barony
	坎巴尔卡Kambarka 	feud.Barony
	莫日加Mozjga 	feud.Barony
	沃特卡Wotka 	feud.Barony
}

func (c *赫雷诺夫HlynovCounty) BEgra叶格拉() feud.Barony {
	return c.叶格拉Egra
}
    
func (c *赫雷诺夫HlynovCounty) BGlazkar格拉兹卡尔() feud.Barony {
	return c.格拉兹卡尔Glazkar
}
    
func (c *赫雷诺夫HlynovCounty) BHlynov赫雷诺夫() feud.Barony {
	return c.赫雷诺夫Hlynov
}
    
func (c *赫雷诺夫HlynovCounty) BIzkar伊日卡尔() feud.Barony {
	return c.伊日卡尔Izkar
}
    
func (c *赫雷诺夫HlynovCounty) BKambarka坎巴尔卡() feud.Barony {
	return c.坎巴尔卡Kambarka
}
    
func (c *赫雷诺夫HlynovCounty) BMozjga莫日加() feud.Barony {
	return c.莫日加Mozjga
}
    
func (c *赫雷诺夫HlynovCounty) BWotka沃特卡() feud.Barony {
	return c.沃特卡Wotka
}
    
var CHlynov赫雷诺夫 HlynovCounty = &赫雷诺夫HlynovCounty{}

func init() {
	f := CHlynov赫雷诺夫.(*赫雷诺夫HlynovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "401",
		Title:     "hlynov",
		TitleName: "赫雷诺夫",
		TitleCode: "c_hlynov",
		Baronies:  map[string]feud.Barony{},
	}

	f.叶格拉Egra = BEgra叶格拉
	f.叶格拉Egra.SetParent(f)
	
	f.格拉兹卡尔Glazkar = BGlazkar格拉兹卡尔
	f.格拉兹卡尔Glazkar.SetParent(f)
	
	f.赫雷诺夫Hlynov = BHlynov赫雷诺夫
	f.赫雷诺夫Hlynov.SetParent(f)
	
	f.伊日卡尔Izkar = BIzkar伊日卡尔
	f.伊日卡尔Izkar.SetParent(f)
	
	f.坎巴尔卡Kambarka = BKambarka坎巴尔卡
	f.坎巴尔卡Kambarka.SetParent(f)
	
	f.莫日加Mozjga = BMozjga莫日加
	f.莫日加Mozjga.SetParent(f)
	
	f.沃特卡Wotka = BWotka沃特卡
	f.沃特卡Wotka.SetParent(f)
	
}
