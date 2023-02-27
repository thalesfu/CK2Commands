package matamma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MatammaCounty interface {
    feud.County
    BBonga邦加() 	feud.Barony
    BChiri奇利基() 	feud.Barony
    BGaro伽罗() 	feud.Barony
    BGinbo津博() 	feud.Barony
    BGishabay吉什阿拜() 	feud.Barony
    BKaffa咖法() 	feud.Barony
    BWacha瓦豪() 	feud.Barony
    BYadeta雅德塔() 	feud.Barony
}

type 米提玛MatammaCounty struct {
	feud.BaseCounty
	邦加Bonga 	feud.Barony
	奇利基Chiri 	feud.Barony
	伽罗Garo 	feud.Barony
	津博Ginbo 	feud.Barony
	吉什阿拜Gishabay 	feud.Barony
	咖法Kaffa 	feud.Barony
	瓦豪Wacha 	feud.Barony
	雅德塔Yadeta 	feud.Barony
}

func (c *米提玛MatammaCounty) BBonga邦加() feud.Barony {
	return c.邦加Bonga
}
    
func (c *米提玛MatammaCounty) BChiri奇利基() feud.Barony {
	return c.奇利基Chiri
}
    
func (c *米提玛MatammaCounty) BGaro伽罗() feud.Barony {
	return c.伽罗Garo
}
    
func (c *米提玛MatammaCounty) BGinbo津博() feud.Barony {
	return c.津博Ginbo
}
    
func (c *米提玛MatammaCounty) BGishabay吉什阿拜() feud.Barony {
	return c.吉什阿拜Gishabay
}
    
func (c *米提玛MatammaCounty) BKaffa咖法() feud.Barony {
	return c.咖法Kaffa
}
    
func (c *米提玛MatammaCounty) BWacha瓦豪() feud.Barony {
	return c.瓦豪Wacha
}
    
func (c *米提玛MatammaCounty) BYadeta雅德塔() feud.Barony {
	return c.雅德塔Yadeta
}
    
var CMatamma米提玛 MatammaCounty = &米提玛MatammaCounty{}

func init() {
	f := CMatamma米提玛.(*米提玛MatammaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "885",
		Title:     "matamma",
		TitleName: "米提玛",
		TitleCode: "c_matamma",
		Baronies:  map[string]feud.Barony{},
	}

	f.邦加Bonga = BBonga邦加
	f.邦加Bonga.SetParent(f)
	
	f.奇利基Chiri = BChiri奇利基
	f.奇利基Chiri.SetParent(f)
	
	f.伽罗Garo = BGaro伽罗
	f.伽罗Garo.SetParent(f)
	
	f.津博Ginbo = BGinbo津博
	f.津博Ginbo.SetParent(f)
	
	f.吉什阿拜Gishabay = BGishabay吉什阿拜
	f.吉什阿拜Gishabay.SetParent(f)
	
	f.咖法Kaffa = BKaffa咖法
	f.咖法Kaffa.SetParent(f)
	
	f.瓦豪Wacha = BWacha瓦豪
	f.瓦豪Wacha.SetParent(f)
	
	f.雅德塔Yadeta = BYadeta雅德塔
	f.雅德塔Yadeta.SetParent(f)
	
}
