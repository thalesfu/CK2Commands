package diafunu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DiafunuCounty interface {
    feud.County
    BDiafunu迪亚富努() 	feud.Barony
    BGalam加拉姆() 	feud.Barony
    BGuidimaka吉迪马卡() 	feud.Barony
    BKokoye科科耶() 	feud.Barony
    BMissiribougou米西里布古() 	feud.Barony
    BSiela西埃拉() 	feud.Barony
    BTinngoba廷戈巴() 	feud.Barony
}

type 迪亚富努DiafunuCounty struct {
	feud.BaseCounty
	迪亚富努Diafunu 	feud.Barony
	加拉姆Galam 	feud.Barony
	吉迪马卡Guidimaka 	feud.Barony
	科科耶Kokoye 	feud.Barony
	米西里布古Missiribougou 	feud.Barony
	西埃拉Siela 	feud.Barony
	廷戈巴Tinngoba 	feud.Barony
}

func (c *迪亚富努DiafunuCounty) BDiafunu迪亚富努() feud.Barony {
	return c.迪亚富努Diafunu
}
    
func (c *迪亚富努DiafunuCounty) BGalam加拉姆() feud.Barony {
	return c.加拉姆Galam
}
    
func (c *迪亚富努DiafunuCounty) BGuidimaka吉迪马卡() feud.Barony {
	return c.吉迪马卡Guidimaka
}
    
func (c *迪亚富努DiafunuCounty) BKokoye科科耶() feud.Barony {
	return c.科科耶Kokoye
}
    
func (c *迪亚富努DiafunuCounty) BMissiribougou米西里布古() feud.Barony {
	return c.米西里布古Missiribougou
}
    
func (c *迪亚富努DiafunuCounty) BSiela西埃拉() feud.Barony {
	return c.西埃拉Siela
}
    
func (c *迪亚富努DiafunuCounty) BTinngoba廷戈巴() feud.Barony {
	return c.廷戈巴Tinngoba
}
    
var CDiafunu迪亚富努 DiafunuCounty = &迪亚富努DiafunuCounty{}

func init() {
	f := CDiafunu迪亚富努.(*迪亚富努DiafunuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1764",
		Title:     "diafunu",
		TitleName: "迪亚富努",
		TitleCode: "c_diafunu",
		Baronies:  map[string]feud.Barony{},
	}

	f.迪亚富努Diafunu = BDiafunu迪亚富努
	f.迪亚富努Diafunu.SetParent(f)
	
	f.加拉姆Galam = BGalam加拉姆
	f.加拉姆Galam.SetParent(f)
	
	f.吉迪马卡Guidimaka = BGuidimaka吉迪马卡
	f.吉迪马卡Guidimaka.SetParent(f)
	
	f.科科耶Kokoye = BKokoye科科耶
	f.科科耶Kokoye.SetParent(f)
	
	f.米西里布古Missiribougou = BMissiribougou米西里布古
	f.米西里布古Missiribougou.SetParent(f)
	
	f.西埃拉Siela = BSiela西埃拉
	f.西埃拉Siela.SetParent(f)
	
	f.廷戈巴Tinngoba = BTinngoba廷戈巴
	f.廷戈巴Tinngoba.SetParent(f)
	
}
