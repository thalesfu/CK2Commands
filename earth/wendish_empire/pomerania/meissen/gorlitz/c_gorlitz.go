package gorlitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GorlitzCounty interface {
    feud.County
    BBautzen包岑() 	feud.Barony
    BBischolow比绍洛夫() 	feud.Barony
    BGorlitz格尔利茨() 	feud.Barony
    BHerrhut黑尔恩胡特() 	feud.Barony
    BPriebus普里布斯() 	feud.Barony
    BSpremberg施普伦贝格() 	feud.Barony
    BYlustiaberg伊卢施蒂亚贝格() 	feud.Barony
}

type 格尔利茨GorlitzCounty struct {
	feud.BaseCounty
	包岑Bautzen 	feud.Barony
	比绍洛夫Bischolow 	feud.Barony
	格尔利茨Gorlitz 	feud.Barony
	黑尔恩胡特Herrhut 	feud.Barony
	普里布斯Priebus 	feud.Barony
	施普伦贝格Spremberg 	feud.Barony
	伊卢施蒂亚贝格Ylustiaberg 	feud.Barony
}

func (c *格尔利茨GorlitzCounty) BBautzen包岑() feud.Barony {
	return c.包岑Bautzen
}
    
func (c *格尔利茨GorlitzCounty) BBischolow比绍洛夫() feud.Barony {
	return c.比绍洛夫Bischolow
}
    
func (c *格尔利茨GorlitzCounty) BGorlitz格尔利茨() feud.Barony {
	return c.格尔利茨Gorlitz
}
    
func (c *格尔利茨GorlitzCounty) BHerrhut黑尔恩胡特() feud.Barony {
	return c.黑尔恩胡特Herrhut
}
    
func (c *格尔利茨GorlitzCounty) BPriebus普里布斯() feud.Barony {
	return c.普里布斯Priebus
}
    
func (c *格尔利茨GorlitzCounty) BSpremberg施普伦贝格() feud.Barony {
	return c.施普伦贝格Spremberg
}
    
func (c *格尔利茨GorlitzCounty) BYlustiaberg伊卢施蒂亚贝格() feud.Barony {
	return c.伊卢施蒂亚贝格Ylustiaberg
}
    
var CGorlitz格尔利茨 GorlitzCounty = &格尔利茨GorlitzCounty{}

func init() {
	f := CGorlitz格尔利茨.(*格尔利茨GorlitzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1589",
		Title:     "gorlitz",
		TitleName: "格尔利茨",
		TitleCode: "c_gorlitz",
		Baronies:  map[string]feud.Barony{},
	}

	f.包岑Bautzen = BBautzen包岑
	f.包岑Bautzen.SetParent(f)
	
	f.比绍洛夫Bischolow = BBischolow比绍洛夫
	f.比绍洛夫Bischolow.SetParent(f)
	
	f.格尔利茨Gorlitz = BGorlitz格尔利茨
	f.格尔利茨Gorlitz.SetParent(f)
	
	f.黑尔恩胡特Herrhut = BHerrhut黑尔恩胡特
	f.黑尔恩胡特Herrhut.SetParent(f)
	
	f.普里布斯Priebus = BPriebus普里布斯
	f.普里布斯Priebus.SetParent(f)
	
	f.施普伦贝格Spremberg = BSpremberg施普伦贝格
	f.施普伦贝格Spremberg.SetParent(f)
	
	f.伊卢施蒂亚贝格Ylustiaberg = BYlustiaberg伊卢施蒂亚贝格
	f.伊卢施蒂亚贝格Ylustiaberg.SetParent(f)
	
}
