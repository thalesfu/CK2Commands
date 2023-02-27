package nalut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NalutCounty interface {
    feud.County
    BDar_el_mzali达尔姆扎利() 	feud.Barony
    BDouar_bessouf杜瓦尔贝苏夫() 	feud.Barony
    BEl_amarmria阿马姆里亚() 	feud.Barony
    BEl_mstiri姆斯蒂里() 	feud.Barony
    BFursata富尔塞塔() 	feud.Barony
    BNalut纳卢特() 	feud.Barony
    BTamaghzah塔马格扎() 	feud.Barony
}

type 纳卢特NalutCounty struct {
	feud.BaseCounty
	达尔姆扎利Dar_el_mzali 	feud.Barony
	杜瓦尔贝苏夫Douar_bessouf 	feud.Barony
	阿马姆里亚El_amarmria 	feud.Barony
	姆斯蒂里El_mstiri 	feud.Barony
	富尔塞塔Fursata 	feud.Barony
	纳卢特Nalut 	feud.Barony
	塔马格扎Tamaghzah 	feud.Barony
}

func (c *纳卢特NalutCounty) BDar_el_mzali达尔姆扎利() feud.Barony {
	return c.达尔姆扎利Dar_el_mzali
}
    
func (c *纳卢特NalutCounty) BDouar_bessouf杜瓦尔贝苏夫() feud.Barony {
	return c.杜瓦尔贝苏夫Douar_bessouf
}
    
func (c *纳卢特NalutCounty) BEl_amarmria阿马姆里亚() feud.Barony {
	return c.阿马姆里亚El_amarmria
}
    
func (c *纳卢特NalutCounty) BEl_mstiri姆斯蒂里() feud.Barony {
	return c.姆斯蒂里El_mstiri
}
    
func (c *纳卢特NalutCounty) BFursata富尔塞塔() feud.Barony {
	return c.富尔塞塔Fursata
}
    
func (c *纳卢特NalutCounty) BNalut纳卢特() feud.Barony {
	return c.纳卢特Nalut
}
    
func (c *纳卢特NalutCounty) BTamaghzah塔马格扎() feud.Barony {
	return c.塔马格扎Tamaghzah
}
    
var CNalut纳卢特 NalutCounty = &纳卢特NalutCounty{}

func init() {
	f := CNalut纳卢特.(*纳卢特NalutCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1724",
		Title:     "nalut",
		TitleName: "纳卢特",
		TitleCode: "c_nalut",
		Baronies:  map[string]feud.Barony{},
	}

	f.达尔姆扎利Dar_el_mzali = BDar_el_mzali达尔姆扎利
	f.达尔姆扎利Dar_el_mzali.SetParent(f)
	
	f.杜瓦尔贝苏夫Douar_bessouf = BDouar_bessouf杜瓦尔贝苏夫
	f.杜瓦尔贝苏夫Douar_bessouf.SetParent(f)
	
	f.阿马姆里亚El_amarmria = BEl_amarmria阿马姆里亚
	f.阿马姆里亚El_amarmria.SetParent(f)
	
	f.姆斯蒂里El_mstiri = BEl_mstiri姆斯蒂里
	f.姆斯蒂里El_mstiri.SetParent(f)
	
	f.富尔塞塔Fursata = BFursata富尔塞塔
	f.富尔塞塔Fursata.SetParent(f)
	
	f.纳卢特Nalut = BNalut纳卢特
	f.纳卢特Nalut.SetParent(f)
	
	f.塔马格扎Tamaghzah = BTamaghzah塔马格扎
	f.塔马格扎Tamaghzah.SetParent(f)
	
}
