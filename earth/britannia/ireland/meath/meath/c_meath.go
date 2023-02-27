package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MeathCounty interface {
    feud.County
    BCiannachta奇纳赫塔() 	feud.Barony
    BForradh福拉() 	feud.Barony
    BKnowth瑙斯() 	feud.Barony
    BRaith_laoghaire拉莱里() 	feud.Barony
    BRaith_na_riogh拉纳里() 	feud.Barony
    BRaith_na_seanadh拉纳肖奈() 	feud.Barony
    BSaint_patricks圣帕特里克() 	feud.Barony
    BTara_meath塔拉() 	feud.Barony
}

type 米斯MeathCounty struct {
	feud.BaseCounty
	奇纳赫塔Ciannachta 	feud.Barony
	福拉Forradh 	feud.Barony
	瑙斯Knowth 	feud.Barony
	拉莱里Raith_laoghaire 	feud.Barony
	拉纳里Raith_na_riogh 	feud.Barony
	拉纳肖奈Raith_na_seanadh 	feud.Barony
	圣帕特里克Saint_patricks 	feud.Barony
	塔拉Tara_meath 	feud.Barony
}

func (c *米斯MeathCounty) BCiannachta奇纳赫塔() feud.Barony {
	return c.奇纳赫塔Ciannachta
}
    
func (c *米斯MeathCounty) BForradh福拉() feud.Barony {
	return c.福拉Forradh
}
    
func (c *米斯MeathCounty) BKnowth瑙斯() feud.Barony {
	return c.瑙斯Knowth
}
    
func (c *米斯MeathCounty) BRaith_laoghaire拉莱里() feud.Barony {
	return c.拉莱里Raith_laoghaire
}
    
func (c *米斯MeathCounty) BRaith_na_riogh拉纳里() feud.Barony {
	return c.拉纳里Raith_na_riogh
}
    
func (c *米斯MeathCounty) BRaith_na_seanadh拉纳肖奈() feud.Barony {
	return c.拉纳肖奈Raith_na_seanadh
}
    
func (c *米斯MeathCounty) BSaint_patricks圣帕特里克() feud.Barony {
	return c.圣帕特里克Saint_patricks
}
    
func (c *米斯MeathCounty) BTara_meath塔拉() feud.Barony {
	return c.塔拉Tara_meath
}
    
var CMeath米斯 MeathCounty = &米斯MeathCounty{}

func init() {
	f := CMeath米斯.(*米斯MeathCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1951",
		Title:     "meath",
		TitleName: "米斯",
		TitleCode: "c_meath",
		Baronies:  map[string]feud.Barony{},
	}

	f.奇纳赫塔Ciannachta = BCiannachta奇纳赫塔
	f.奇纳赫塔Ciannachta.SetParent(f)
	
	f.福拉Forradh = BForradh福拉
	f.福拉Forradh.SetParent(f)
	
	f.瑙斯Knowth = BKnowth瑙斯
	f.瑙斯Knowth.SetParent(f)
	
	f.拉莱里Raith_laoghaire = BRaith_laoghaire拉莱里
	f.拉莱里Raith_laoghaire.SetParent(f)
	
	f.拉纳里Raith_na_riogh = BRaith_na_riogh拉纳里
	f.拉纳里Raith_na_riogh.SetParent(f)
	
	f.拉纳肖奈Raith_na_seanadh = BRaith_na_seanadh拉纳肖奈
	f.拉纳肖奈Raith_na_seanadh.SetParent(f)
	
	f.圣帕特里克Saint_patricks = BSaint_patricks圣帕特里克
	f.圣帕特里克Saint_patricks.SetParent(f)
	
	f.塔拉Tara_meath = BTara_meath塔拉
	f.塔拉Tara_meath.SetParent(f)
	
}
