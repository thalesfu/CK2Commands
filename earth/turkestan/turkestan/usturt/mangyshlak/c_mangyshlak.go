package mangyshlak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MangyshlakCounty interface {
    feud.County
    BAmankyzylit阿曼克孜利特() 	feud.Barony
    BAqtaw阿克套() 	feud.Barony
    BAraldy阿拉尔德() 	feud.Barony
    BAshchimuryn阿希穆伦() 	feud.Barony
    BKzyluzen克孜勒乌赞() 	feud.Barony
    BSayutes赛_厄捷斯() 	feud.Barony
    BTigen季根() 	feud.Barony
    BUzen乌津() 	feud.Barony
}

type 曼格什拉克MangyshlakCounty struct {
	feud.BaseCounty
	阿曼克孜利特Amankyzylit 	feud.Barony
	阿克套Aqtaw 	feud.Barony
	阿拉尔德Araldy 	feud.Barony
	阿希穆伦Ashchimuryn 	feud.Barony
	克孜勒乌赞Kzyluzen 	feud.Barony
	赛_厄捷斯Sayutes 	feud.Barony
	季根Tigen 	feud.Barony
	乌津Uzen 	feud.Barony
}

func (c *曼格什拉克MangyshlakCounty) BAmankyzylit阿曼克孜利特() feud.Barony {
	return c.阿曼克孜利特Amankyzylit
}
    
func (c *曼格什拉克MangyshlakCounty) BAqtaw阿克套() feud.Barony {
	return c.阿克套Aqtaw
}
    
func (c *曼格什拉克MangyshlakCounty) BAraldy阿拉尔德() feud.Barony {
	return c.阿拉尔德Araldy
}
    
func (c *曼格什拉克MangyshlakCounty) BAshchimuryn阿希穆伦() feud.Barony {
	return c.阿希穆伦Ashchimuryn
}
    
func (c *曼格什拉克MangyshlakCounty) BKzyluzen克孜勒乌赞() feud.Barony {
	return c.克孜勒乌赞Kzyluzen
}
    
func (c *曼格什拉克MangyshlakCounty) BSayutes赛_厄捷斯() feud.Barony {
	return c.赛_厄捷斯Sayutes
}
    
func (c *曼格什拉克MangyshlakCounty) BTigen季根() feud.Barony {
	return c.季根Tigen
}
    
func (c *曼格什拉克MangyshlakCounty) BUzen乌津() feud.Barony {
	return c.乌津Uzen
}
    
var CMangyshlak曼格什拉克 MangyshlakCounty = &曼格什拉克MangyshlakCounty{}

func init() {
	f := CMangyshlak曼格什拉克.(*曼格什拉克MangyshlakCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "624",
		Title:     "mangyshlak",
		TitleName: "曼格什拉克",
		TitleCode: "c_mangyshlak",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿曼克孜利特Amankyzylit = BAmankyzylit阿曼克孜利特
	f.阿曼克孜利特Amankyzylit.SetParent(f)
	
	f.阿克套Aqtaw = BAqtaw阿克套
	f.阿克套Aqtaw.SetParent(f)
	
	f.阿拉尔德Araldy = BAraldy阿拉尔德
	f.阿拉尔德Araldy.SetParent(f)
	
	f.阿希穆伦Ashchimuryn = BAshchimuryn阿希穆伦
	f.阿希穆伦Ashchimuryn.SetParent(f)
	
	f.克孜勒乌赞Kzyluzen = BKzyluzen克孜勒乌赞
	f.克孜勒乌赞Kzyluzen.SetParent(f)
	
	f.赛_厄捷斯Sayutes = BSayutes赛_厄捷斯
	f.赛_厄捷斯Sayutes.SetParent(f)
	
	f.季根Tigen = BTigen季根
	f.季根Tigen.SetParent(f)
	
	f.乌津Uzen = BUzen乌津
	f.乌津Uzen.SetParent(f)
	
}
