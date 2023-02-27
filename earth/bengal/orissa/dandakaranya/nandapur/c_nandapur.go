package nandapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NandapurCounty interface {
    feud.County
    BDharamgarh达摩姞利呬() 	feud.Barony
    BJeypore祇耶城() 	feud.Barony
    BJunagarh周那姞利呬() 	feud.Barony
    BKilpadappai吉波陀吠() 	feud.Barony
    BKoraput拘罗弗多() 	feud.Barony
    BMarhia摩罗醯耶() 	feud.Barony
    BNandapur难陀城() 	feud.Barony
}

type 难陀城NandapurCounty struct {
	feud.BaseCounty
	达摩姞利呬Dharamgarh 	feud.Barony
	祇耶城Jeypore 	feud.Barony
	周那姞利呬Junagarh 	feud.Barony
	吉波陀吠Kilpadappai 	feud.Barony
	拘罗弗多Koraput 	feud.Barony
	摩罗醯耶Marhia 	feud.Barony
	难陀城Nandapur 	feud.Barony
}

func (c *难陀城NandapurCounty) BDharamgarh达摩姞利呬() feud.Barony {
	return c.达摩姞利呬Dharamgarh
}
    
func (c *难陀城NandapurCounty) BJeypore祇耶城() feud.Barony {
	return c.祇耶城Jeypore
}
    
func (c *难陀城NandapurCounty) BJunagarh周那姞利呬() feud.Barony {
	return c.周那姞利呬Junagarh
}
    
func (c *难陀城NandapurCounty) BKilpadappai吉波陀吠() feud.Barony {
	return c.吉波陀吠Kilpadappai
}
    
func (c *难陀城NandapurCounty) BKoraput拘罗弗多() feud.Barony {
	return c.拘罗弗多Koraput
}
    
func (c *难陀城NandapurCounty) BMarhia摩罗醯耶() feud.Barony {
	return c.摩罗醯耶Marhia
}
    
func (c *难陀城NandapurCounty) BNandapur难陀城() feud.Barony {
	return c.难陀城Nandapur
}
    
var CNandapur难陀城 NandapurCounty = &难陀城NandapurCounty{}

func init() {
	f := CNandapur难陀城.(*难陀城NandapurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1226",
		Title:     "nandapur",
		TitleName: "难陀城",
		TitleCode: "c_nandapur",
		Baronies:  map[string]feud.Barony{},
	}

	f.达摩姞利呬Dharamgarh = BDharamgarh达摩姞利呬
	f.达摩姞利呬Dharamgarh.SetParent(f)
	
	f.祇耶城Jeypore = BJeypore祇耶城
	f.祇耶城Jeypore.SetParent(f)
	
	f.周那姞利呬Junagarh = BJunagarh周那姞利呬
	f.周那姞利呬Junagarh.SetParent(f)
	
	f.吉波陀吠Kilpadappai = BKilpadappai吉波陀吠
	f.吉波陀吠Kilpadappai.SetParent(f)
	
	f.拘罗弗多Koraput = BKoraput拘罗弗多
	f.拘罗弗多Koraput.SetParent(f)
	
	f.摩罗醯耶Marhia = BMarhia摩罗醯耶
	f.摩罗醯耶Marhia.SetParent(f)
	
	f.难陀城Nandapur = BNandapur难陀城
	f.难陀城Nandapur.SetParent(f)
	
}
