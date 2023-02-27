package medelpad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedelpadCounty interface {
    feud.County
    BAlno阿尔讷() 	feud.Barony
    BLiden利登() 	feud.Barony
    BNjurunda纽伦达() 	feud.Barony
    BOtman奥特曼() 	feud.Barony
    BSelanger谢勒恩耶() 	feud.Barony
    BSkon申() 	feud.Barony
    BSundsvall松兹瓦尔() 	feud.Barony
    BTorp托普() 	feud.Barony
}

type 梅代尔帕德MedelpadCounty struct {
	feud.BaseCounty
	阿尔讷Alno 	feud.Barony
	利登Liden 	feud.Barony
	纽伦达Njurunda 	feud.Barony
	奥特曼Otman 	feud.Barony
	谢勒恩耶Selanger 	feud.Barony
	申Skon 	feud.Barony
	松兹瓦尔Sundsvall 	feud.Barony
	托普Torp 	feud.Barony
}

func (c *梅代尔帕德MedelpadCounty) BAlno阿尔讷() feud.Barony {
	return c.阿尔讷Alno
}
    
func (c *梅代尔帕德MedelpadCounty) BLiden利登() feud.Barony {
	return c.利登Liden
}
    
func (c *梅代尔帕德MedelpadCounty) BNjurunda纽伦达() feud.Barony {
	return c.纽伦达Njurunda
}
    
func (c *梅代尔帕德MedelpadCounty) BOtman奥特曼() feud.Barony {
	return c.奥特曼Otman
}
    
func (c *梅代尔帕德MedelpadCounty) BSelanger谢勒恩耶() feud.Barony {
	return c.谢勒恩耶Selanger
}
    
func (c *梅代尔帕德MedelpadCounty) BSkon申() feud.Barony {
	return c.申Skon
}
    
func (c *梅代尔帕德MedelpadCounty) BSundsvall松兹瓦尔() feud.Barony {
	return c.松兹瓦尔Sundsvall
}
    
func (c *梅代尔帕德MedelpadCounty) BTorp托普() feud.Barony {
	return c.托普Torp
}
    
var CMedelpad梅代尔帕德 MedelpadCounty = &梅代尔帕德MedelpadCounty{}

func init() {
	f := CMedelpad梅代尔帕德.(*梅代尔帕德MedelpadCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "283",
		Title:     "medelpad",
		TitleName: "梅代尔帕德",
		TitleCode: "c_medelpad",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔讷Alno = BAlno阿尔讷
	f.阿尔讷Alno.SetParent(f)
	
	f.利登Liden = BLiden利登
	f.利登Liden.SetParent(f)
	
	f.纽伦达Njurunda = BNjurunda纽伦达
	f.纽伦达Njurunda.SetParent(f)
	
	f.奥特曼Otman = BOtman奥特曼
	f.奥特曼Otman.SetParent(f)
	
	f.谢勒恩耶Selanger = BSelanger谢勒恩耶
	f.谢勒恩耶Selanger.SetParent(f)
	
	f.申Skon = BSkon申
	f.申Skon.SetParent(f)
	
	f.松兹瓦尔Sundsvall = BSundsvall松兹瓦尔
	f.松兹瓦尔Sundsvall.SetParent(f)
	
	f.托普Torp = BTorp托普
	f.托普Torp.SetParent(f)
	
}
