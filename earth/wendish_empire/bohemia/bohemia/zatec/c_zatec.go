package zatec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZatecCounty interface {
    feud.County
    BBilina比利纳() 	feud.Barony
    BChomutov霍穆托夫() 	feud.Barony
    BKadan卡丹() 	feud.Barony
    BLouny洛乌尼() 	feud.Barony
    BOsek奥塞克() 	feud.Barony
    BTeplice特普利采() 	feud.Barony
    BZatec扎泰茨() 	feud.Barony
}

type 扎泰茨ZatecCounty struct {
	feud.BaseCounty
	比利纳Bilina 	feud.Barony
	霍穆托夫Chomutov 	feud.Barony
	卡丹Kadan 	feud.Barony
	洛乌尼Louny 	feud.Barony
	奥塞克Osek 	feud.Barony
	特普利采Teplice 	feud.Barony
	扎泰茨Zatec 	feud.Barony
}

func (c *扎泰茨ZatecCounty) BBilina比利纳() feud.Barony {
	return c.比利纳Bilina
}
    
func (c *扎泰茨ZatecCounty) BChomutov霍穆托夫() feud.Barony {
	return c.霍穆托夫Chomutov
}
    
func (c *扎泰茨ZatecCounty) BKadan卡丹() feud.Barony {
	return c.卡丹Kadan
}
    
func (c *扎泰茨ZatecCounty) BLouny洛乌尼() feud.Barony {
	return c.洛乌尼Louny
}
    
func (c *扎泰茨ZatecCounty) BOsek奥塞克() feud.Barony {
	return c.奥塞克Osek
}
    
func (c *扎泰茨ZatecCounty) BTeplice特普利采() feud.Barony {
	return c.特普利采Teplice
}
    
func (c *扎泰茨ZatecCounty) BZatec扎泰茨() feud.Barony {
	return c.扎泰茨Zatec
}
    
var CZatec扎泰茨 ZatecCounty = &扎泰茨ZatecCounty{}

func init() {
	f := CZatec扎泰茨.(*扎泰茨ZatecCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1926",
		Title:     "zatec",
		TitleName: "扎泰茨",
		TitleCode: "c_zatec",
		Baronies:  map[string]feud.Barony{},
	}

	f.比利纳Bilina = BBilina比利纳
	f.比利纳Bilina.SetParent(f)
	
	f.霍穆托夫Chomutov = BChomutov霍穆托夫
	f.霍穆托夫Chomutov.SetParent(f)
	
	f.卡丹Kadan = BKadan卡丹
	f.卡丹Kadan.SetParent(f)
	
	f.洛乌尼Louny = BLouny洛乌尼
	f.洛乌尼Louny.SetParent(f)
	
	f.奥塞克Osek = BOsek奥塞克
	f.奥塞克Osek.SetParent(f)
	
	f.特普利采Teplice = BTeplice特普利采
	f.特普利采Teplice.SetParent(f)
	
	f.扎泰茨Zatec = BZatec扎泰茨
	f.扎泰茨Zatec.SetParent(f)
	
}
