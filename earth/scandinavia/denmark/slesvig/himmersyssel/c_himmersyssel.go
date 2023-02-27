package himmersyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HimmersysselCounty interface {
    feud.County
    BAalborg奥尔堡() 	feud.Barony
    BAar奥尔() 	feud.Barony
    BGislum吉斯卢姆() 	feud.Barony
    BHadesundt海松() 	feud.Barony
    BHornum霍努姆() 	feud.Barony
    BRins林斯() 	feud.Barony
    BViborg维堡() 	feud.Barony
}

type 奥尔堡HimmersysselCounty struct {
	feud.BaseCounty
	奥尔堡Aalborg 	feud.Barony
	奥尔Aar 	feud.Barony
	吉斯卢姆Gislum 	feud.Barony
	海松Hadesundt 	feud.Barony
	霍努姆Hornum 	feud.Barony
	林斯Rins 	feud.Barony
	维堡Viborg 	feud.Barony
}

func (c *奥尔堡HimmersysselCounty) BAalborg奥尔堡() feud.Barony {
	return c.奥尔堡Aalborg
}
    
func (c *奥尔堡HimmersysselCounty) BAar奥尔() feud.Barony {
	return c.奥尔Aar
}
    
func (c *奥尔堡HimmersysselCounty) BGislum吉斯卢姆() feud.Barony {
	return c.吉斯卢姆Gislum
}
    
func (c *奥尔堡HimmersysselCounty) BHadesundt海松() feud.Barony {
	return c.海松Hadesundt
}
    
func (c *奥尔堡HimmersysselCounty) BHornum霍努姆() feud.Barony {
	return c.霍努姆Hornum
}
    
func (c *奥尔堡HimmersysselCounty) BRins林斯() feud.Barony {
	return c.林斯Rins
}
    
func (c *奥尔堡HimmersysselCounty) BViborg维堡() feud.Barony {
	return c.维堡Viborg
}
    
var CHimmersyssel奥尔堡 HimmersysselCounty = &奥尔堡HimmersysselCounty{}

func init() {
	f := CHimmersyssel奥尔堡.(*奥尔堡HimmersysselCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1685",
		Title:     "himmersyssel",
		TitleName: "奥尔堡",
		TitleCode: "c_himmersyssel",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥尔堡Aalborg = BAalborg奥尔堡
	f.奥尔堡Aalborg.SetParent(f)
	
	f.奥尔Aar = BAar奥尔
	f.奥尔Aar.SetParent(f)
	
	f.吉斯卢姆Gislum = BGislum吉斯卢姆
	f.吉斯卢姆Gislum.SetParent(f)
	
	f.海松Hadesundt = BHadesundt海松
	f.海松Hadesundt.SetParent(f)
	
	f.霍努姆Hornum = BHornum霍努姆
	f.霍努姆Hornum.SetParent(f)
	
	f.林斯Rins = BRins林斯
	f.林斯Rins.SetParent(f)
	
	f.维堡Viborg = BViborg维堡
	f.维堡Viborg.SetParent(f)
	
}
