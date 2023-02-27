package kol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolCounty interface {
    feud.County
    BBarauli婆罗利() 	feud.Barony
    BDalmau达尔毛() 	feud.Barony
    BDor兜尔() 	feud.Barony
    BEtah埃达() 	feud.Barony
    BKol拘罗() 	feud.Barony
    BNilavati尼罗伐底() 	feud.Barony
    BSoron索龙() 	feud.Barony
}

type 拘罗KolCounty struct {
	feud.BaseCounty
	婆罗利Barauli 	feud.Barony
	达尔毛Dalmau 	feud.Barony
	兜尔Dor 	feud.Barony
	埃达Etah 	feud.Barony
	拘罗Kol 	feud.Barony
	尼罗伐底Nilavati 	feud.Barony
	索龙Soron 	feud.Barony
}

func (c *拘罗KolCounty) BBarauli婆罗利() feud.Barony {
	return c.婆罗利Barauli
}
    
func (c *拘罗KolCounty) BDalmau达尔毛() feud.Barony {
	return c.达尔毛Dalmau
}
    
func (c *拘罗KolCounty) BDor兜尔() feud.Barony {
	return c.兜尔Dor
}
    
func (c *拘罗KolCounty) BEtah埃达() feud.Barony {
	return c.埃达Etah
}
    
func (c *拘罗KolCounty) BKol拘罗() feud.Barony {
	return c.拘罗Kol
}
    
func (c *拘罗KolCounty) BNilavati尼罗伐底() feud.Barony {
	return c.尼罗伐底Nilavati
}
    
func (c *拘罗KolCounty) BSoron索龙() feud.Barony {
	return c.索龙Soron
}
    
var CKol拘罗 KolCounty = &拘罗KolCounty{}

func init() {
	f := CKol拘罗.(*拘罗KolCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1141",
		Title:     "kol",
		TitleName: "拘罗",
		TitleCode: "c_kol",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆罗利Barauli = BBarauli婆罗利
	f.婆罗利Barauli.SetParent(f)
	
	f.达尔毛Dalmau = BDalmau达尔毛
	f.达尔毛Dalmau.SetParent(f)
	
	f.兜尔Dor = BDor兜尔
	f.兜尔Dor.SetParent(f)
	
	f.埃达Etah = BEtah埃达
	f.埃达Etah.SetParent(f)
	
	f.拘罗Kol = BKol拘罗
	f.拘罗Kol.SetParent(f)
	
	f.尼罗伐底Nilavati = BNilavati尼罗伐底
	f.尼罗伐底Nilavati.SetParent(f)
	
	f.索龙Soron = BSoron索龙
	f.索龙Soron.SetParent(f)
	
}
