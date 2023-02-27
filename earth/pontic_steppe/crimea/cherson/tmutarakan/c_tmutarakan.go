package tmutarakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TmutarakanCounty interface {
    feud.County
    BBata巴塔() 	feud.Barony
    BJevlisia耶夫利西亚() 	feud.Barony
    BMapa马帕() 	feud.Barony
    BSujukqale索乌贾克() 	feud.Barony
    BTaman塔曼() 	feud.Barony
    BTmutarakan特穆塔拉坎() 	feud.Barony
    BTsemes采梅斯() 	feud.Barony
    BTumnev图姆涅夫() 	feud.Barony
}

type 特穆塔拉坎TmutarakanCounty struct {
	feud.BaseCounty
	巴塔Bata 	feud.Barony
	耶夫利西亚Jevlisia 	feud.Barony
	马帕Mapa 	feud.Barony
	索乌贾克Sujukqale 	feud.Barony
	塔曼Taman 	feud.Barony
	特穆塔拉坎Tmutarakan 	feud.Barony
	采梅斯Tsemes 	feud.Barony
	图姆涅夫Tumnev 	feud.Barony
}

func (c *特穆塔拉坎TmutarakanCounty) BBata巴塔() feud.Barony {
	return c.巴塔Bata
}
    
func (c *特穆塔拉坎TmutarakanCounty) BJevlisia耶夫利西亚() feud.Barony {
	return c.耶夫利西亚Jevlisia
}
    
func (c *特穆塔拉坎TmutarakanCounty) BMapa马帕() feud.Barony {
	return c.马帕Mapa
}
    
func (c *特穆塔拉坎TmutarakanCounty) BSujukqale索乌贾克() feud.Barony {
	return c.索乌贾克Sujukqale
}
    
func (c *特穆塔拉坎TmutarakanCounty) BTaman塔曼() feud.Barony {
	return c.塔曼Taman
}
    
func (c *特穆塔拉坎TmutarakanCounty) BTmutarakan特穆塔拉坎() feud.Barony {
	return c.特穆塔拉坎Tmutarakan
}
    
func (c *特穆塔拉坎TmutarakanCounty) BTsemes采梅斯() feud.Barony {
	return c.采梅斯Tsemes
}
    
func (c *特穆塔拉坎TmutarakanCounty) BTumnev图姆涅夫() feud.Barony {
	return c.图姆涅夫Tumnev
}
    
var CTmutarakan特穆塔拉坎 TmutarakanCounty = &特穆塔拉坎TmutarakanCounty{}

func init() {
	f := CTmutarakan特穆塔拉坎.(*特穆塔拉坎TmutarakanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "598",
		Title:     "tmutarakan",
		TitleName: "特穆塔拉坎",
		TitleCode: "c_tmutarakan",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴塔Bata = BBata巴塔
	f.巴塔Bata.SetParent(f)
	
	f.耶夫利西亚Jevlisia = BJevlisia耶夫利西亚
	f.耶夫利西亚Jevlisia.SetParent(f)
	
	f.马帕Mapa = BMapa马帕
	f.马帕Mapa.SetParent(f)
	
	f.索乌贾克Sujukqale = BSujukqale索乌贾克
	f.索乌贾克Sujukqale.SetParent(f)
	
	f.塔曼Taman = BTaman塔曼
	f.塔曼Taman.SetParent(f)
	
	f.特穆塔拉坎Tmutarakan = BTmutarakan特穆塔拉坎
	f.特穆塔拉坎Tmutarakan.SetParent(f)
	
	f.采梅斯Tsemes = BTsemes采梅斯
	f.采梅斯Tsemes.SetParent(f)
	
	f.图姆涅夫Tumnev = BTumnev图姆涅夫
	f.图姆涅夫Tumnev.SetParent(f)
	
}
