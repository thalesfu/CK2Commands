package belaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BelayaCounty interface {
    feud.County
    BAkberdino阿克别尔季诺() 	feud.Barony
    BAsy阿瑟() 	feud.Barony
    BBalazhi巴拉日() 	feud.Barony
    BBelaya别拉亚() 	feud.Barony
    BShipovo希波沃() 	feud.Barony
    BSlutka斯卢特卡() 	feud.Barony
    BUrunda乌伦达() 	feud.Barony
}

type 别拉亚BelayaCounty struct {
	feud.BaseCounty
	阿克别尔季诺Akberdino 	feud.Barony
	阿瑟Asy 	feud.Barony
	巴拉日Balazhi 	feud.Barony
	别拉亚Belaya 	feud.Barony
	希波沃Shipovo 	feud.Barony
	斯卢特卡Slutka 	feud.Barony
	乌伦达Urunda 	feud.Barony
}

func (c *别拉亚BelayaCounty) BAkberdino阿克别尔季诺() feud.Barony {
	return c.阿克别尔季诺Akberdino
}
    
func (c *别拉亚BelayaCounty) BAsy阿瑟() feud.Barony {
	return c.阿瑟Asy
}
    
func (c *别拉亚BelayaCounty) BBalazhi巴拉日() feud.Barony {
	return c.巴拉日Balazhi
}
    
func (c *别拉亚BelayaCounty) BBelaya别拉亚() feud.Barony {
	return c.别拉亚Belaya
}
    
func (c *别拉亚BelayaCounty) BShipovo希波沃() feud.Barony {
	return c.希波沃Shipovo
}
    
func (c *别拉亚BelayaCounty) BSlutka斯卢特卡() feud.Barony {
	return c.斯卢特卡Slutka
}
    
func (c *别拉亚BelayaCounty) BUrunda乌伦达() feud.Barony {
	return c.乌伦达Urunda
}
    
var CBelaya别拉亚 BelayaCounty = &别拉亚BelayaCounty{}

func init() {
	f := CBelaya别拉亚.(*别拉亚BelayaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1816",
		Title:     "belaya",
		TitleName: "别拉亚",
		TitleCode: "c_belaya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克别尔季诺Akberdino = BAkberdino阿克别尔季诺
	f.阿克别尔季诺Akberdino.SetParent(f)
	
	f.阿瑟Asy = BAsy阿瑟
	f.阿瑟Asy.SetParent(f)
	
	f.巴拉日Balazhi = BBalazhi巴拉日
	f.巴拉日Balazhi.SetParent(f)
	
	f.别拉亚Belaya = BBelaya别拉亚
	f.别拉亚Belaya.SetParent(f)
	
	f.希波沃Shipovo = BShipovo希波沃
	f.希波沃Shipovo.SetParent(f)
	
	f.斯卢特卡Slutka = BSlutka斯卢特卡
	f.斯卢特卡Slutka.SetParent(f)
	
	f.乌伦达Urunda = BUrunda乌伦达
	f.乌伦达Urunda.SetParent(f)
	
}
