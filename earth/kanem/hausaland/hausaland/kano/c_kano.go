package kano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KanoCounty interface {
    feud.County
    BGamaja加马贾() 	feud.Barony
    BHadejia哈代贾() 	feud.Barony
    BKano卡诺() 	feud.Barony
    BLabozhi拉博日() 	feud.Barony
    BNok诺克() 	feud.Barony
    BTigaro蒂加罗() 	feud.Barony
    BTurunku图伦库() 	feud.Barony
}

type 卡诺KanoCounty struct {
	feud.BaseCounty
	加马贾Gamaja 	feud.Barony
	哈代贾Hadejia 	feud.Barony
	卡诺Kano 	feud.Barony
	拉博日Labozhi 	feud.Barony
	诺克Nok 	feud.Barony
	蒂加罗Tigaro 	feud.Barony
	图伦库Turunku 	feud.Barony
}

func (c *卡诺KanoCounty) BGamaja加马贾() feud.Barony {
	return c.加马贾Gamaja
}
    
func (c *卡诺KanoCounty) BHadejia哈代贾() feud.Barony {
	return c.哈代贾Hadejia
}
    
func (c *卡诺KanoCounty) BKano卡诺() feud.Barony {
	return c.卡诺Kano
}
    
func (c *卡诺KanoCounty) BLabozhi拉博日() feud.Barony {
	return c.拉博日Labozhi
}
    
func (c *卡诺KanoCounty) BNok诺克() feud.Barony {
	return c.诺克Nok
}
    
func (c *卡诺KanoCounty) BTigaro蒂加罗() feud.Barony {
	return c.蒂加罗Tigaro
}
    
func (c *卡诺KanoCounty) BTurunku图伦库() feud.Barony {
	return c.图伦库Turunku
}
    
var CKano卡诺 KanoCounty = &卡诺KanoCounty{}

func init() {
	f := CKano卡诺.(*卡诺KanoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1746",
		Title:     "kano",
		TitleName: "卡诺",
		TitleCode: "c_kano",
		Baronies:  map[string]feud.Barony{},
	}

	f.加马贾Gamaja = BGamaja加马贾
	f.加马贾Gamaja.SetParent(f)
	
	f.哈代贾Hadejia = BHadejia哈代贾
	f.哈代贾Hadejia.SetParent(f)
	
	f.卡诺Kano = BKano卡诺
	f.卡诺Kano.SetParent(f)
	
	f.拉博日Labozhi = BLabozhi拉博日
	f.拉博日Labozhi.SetParent(f)
	
	f.诺克Nok = BNok诺克
	f.诺克Nok.SetParent(f)
	
	f.蒂加罗Tigaro = BTigaro蒂加罗
	f.蒂加罗Tigaro.SetParent(f)
	
	f.图伦库Turunku = BTurunku图伦库
	f.图伦库Turunku.SetParent(f)
	
}
