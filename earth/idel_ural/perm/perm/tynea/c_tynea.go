package tynea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TyneaCounty interface {
    feud.County
    BAkbash阿克巴什() 	feud.Barony
    BFedorki费多尔基() 	feud.Barony
    BOsa_tynea特涅阿() 	feud.Barony
    BPosha波沙() 	feud.Barony
    BTynea特涅阿() 	feud.Barony
    BVoyady沃亚德() 	feud.Barony
    BYelovo叶洛沃() 	feud.Barony
}

type 特涅阿TyneaCounty struct {
	feud.BaseCounty
	阿克巴什Akbash 	feud.Barony
	费多尔基Fedorki 	feud.Barony
	特涅阿Osa_tynea 	feud.Barony
	波沙Posha 	feud.Barony
	特涅阿Tynea 	feud.Barony
	沃亚德Voyady 	feud.Barony
	叶洛沃Yelovo 	feud.Barony
}

func (c *特涅阿TyneaCounty) BAkbash阿克巴什() feud.Barony {
	return c.阿克巴什Akbash
}
    
func (c *特涅阿TyneaCounty) BFedorki费多尔基() feud.Barony {
	return c.费多尔基Fedorki
}
    
func (c *特涅阿TyneaCounty) BOsa_tynea特涅阿() feud.Barony {
	return c.特涅阿Osa_tynea
}
    
func (c *特涅阿TyneaCounty) BPosha波沙() feud.Barony {
	return c.波沙Posha
}
    
func (c *特涅阿TyneaCounty) BTynea特涅阿() feud.Barony {
	return c.特涅阿Tynea
}
    
func (c *特涅阿TyneaCounty) BVoyady沃亚德() feud.Barony {
	return c.沃亚德Voyady
}
    
func (c *特涅阿TyneaCounty) BYelovo叶洛沃() feud.Barony {
	return c.叶洛沃Yelovo
}
    
var CTynea特涅阿 TyneaCounty = &特涅阿TyneaCounty{}

func init() {
	f := CTynea特涅阿.(*特涅阿TyneaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1845",
		Title:     "tynea",
		TitleName: "特涅阿",
		TitleCode: "c_tynea",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克巴什Akbash = BAkbash阿克巴什
	f.阿克巴什Akbash.SetParent(f)
	
	f.费多尔基Fedorki = BFedorki费多尔基
	f.费多尔基Fedorki.SetParent(f)
	
	f.特涅阿Osa_tynea = BOsa_tynea特涅阿
	f.特涅阿Osa_tynea.SetParent(f)
	
	f.波沙Posha = BPosha波沙
	f.波沙Posha.SetParent(f)
	
	f.特涅阿Tynea = BTynea特涅阿
	f.特涅阿Tynea.SetParent(f)
	
	f.沃亚德Voyady = BVoyady沃亚德
	f.沃亚德Voyady.SetParent(f)
	
	f.叶洛沃Yelovo = BYelovo叶洛沃
	f.叶洛沃Yelovo.SetParent(f)
	
}
