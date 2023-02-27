package sacz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaczCounty interface {
    feud.County
    BBiecz别奇() 	feud.Barony
    BGorlice戈尔利采() 	feud.Barony
    BMielec梅莱茨() 	feud.Barony
    BNowysacz新松奇() 	feud.Barony
    BPilzno皮尔兹诺() 	feud.Barony
    BSzczyrzycz什哲日茨() 	feud.Barony
    BTarnow塔尔努夫() 	feud.Barony
}

type 松奇SaczCounty struct {
	feud.BaseCounty
	别奇Biecz 	feud.Barony
	戈尔利采Gorlice 	feud.Barony
	梅莱茨Mielec 	feud.Barony
	新松奇Nowysacz 	feud.Barony
	皮尔兹诺Pilzno 	feud.Barony
	什哲日茨Szczyrzycz 	feud.Barony
	塔尔努夫Tarnow 	feud.Barony
}

func (c *松奇SaczCounty) BBiecz别奇() feud.Barony {
	return c.别奇Biecz
}
    
func (c *松奇SaczCounty) BGorlice戈尔利采() feud.Barony {
	return c.戈尔利采Gorlice
}
    
func (c *松奇SaczCounty) BMielec梅莱茨() feud.Barony {
	return c.梅莱茨Mielec
}
    
func (c *松奇SaczCounty) BNowysacz新松奇() feud.Barony {
	return c.新松奇Nowysacz
}
    
func (c *松奇SaczCounty) BPilzno皮尔兹诺() feud.Barony {
	return c.皮尔兹诺Pilzno
}
    
func (c *松奇SaczCounty) BSzczyrzycz什哲日茨() feud.Barony {
	return c.什哲日茨Szczyrzycz
}
    
func (c *松奇SaczCounty) BTarnow塔尔努夫() feud.Barony {
	return c.塔尔努夫Tarnow
}
    
var CSacz松奇 SaczCounty = &松奇SaczCounty{}

func init() {
	f := CSacz松奇.(*松奇SaczCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "532",
		Title:     "sacz",
		TitleName: "松奇",
		TitleCode: "c_sacz",
		Baronies:  map[string]feud.Barony{},
	}

	f.别奇Biecz = BBiecz别奇
	f.别奇Biecz.SetParent(f)
	
	f.戈尔利采Gorlice = BGorlice戈尔利采
	f.戈尔利采Gorlice.SetParent(f)
	
	f.梅莱茨Mielec = BMielec梅莱茨
	f.梅莱茨Mielec.SetParent(f)
	
	f.新松奇Nowysacz = BNowysacz新松奇
	f.新松奇Nowysacz.SetParent(f)
	
	f.皮尔兹诺Pilzno = BPilzno皮尔兹诺
	f.皮尔兹诺Pilzno.SetParent(f)
	
	f.什哲日茨Szczyrzycz = BSzczyrzycz什哲日茨
	f.什哲日茨Szczyrzycz.SetParent(f)
	
	f.塔尔努夫Tarnow = BTarnow塔尔努夫
	f.塔尔努夫Tarnow.SetParent(f)
	
}
