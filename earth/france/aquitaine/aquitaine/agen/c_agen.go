package agen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AgenCounty interface {
    feud.County
    BAgen阿让() 	feud.Barony
    BBlanquefort布朗克福() 	feud.Barony
    BCahors卡奥尔() 	feud.Barony
    BFigeac菲雅克() 	feud.Barony
    BLuzech吕泽克() 	feud.Barony
    BMoissac穆瓦萨克() 	feud.Barony
    BPenne佩纳() 	feud.Barony
    BRocamadour罗卡马杜尔() 	feud.Barony
}

type 阿让AgenCounty struct {
	feud.BaseCounty
	阿让Agen 	feud.Barony
	布朗克福Blanquefort 	feud.Barony
	卡奥尔Cahors 	feud.Barony
	菲雅克Figeac 	feud.Barony
	吕泽克Luzech 	feud.Barony
	穆瓦萨克Moissac 	feud.Barony
	佩纳Penne 	feud.Barony
	罗卡马杜尔Rocamadour 	feud.Barony
}

func (c *阿让AgenCounty) BAgen阿让() feud.Barony {
	return c.阿让Agen
}
    
func (c *阿让AgenCounty) BBlanquefort布朗克福() feud.Barony {
	return c.布朗克福Blanquefort
}
    
func (c *阿让AgenCounty) BCahors卡奥尔() feud.Barony {
	return c.卡奥尔Cahors
}
    
func (c *阿让AgenCounty) BFigeac菲雅克() feud.Barony {
	return c.菲雅克Figeac
}
    
func (c *阿让AgenCounty) BLuzech吕泽克() feud.Barony {
	return c.吕泽克Luzech
}
    
func (c *阿让AgenCounty) BMoissac穆瓦萨克() feud.Barony {
	return c.穆瓦萨克Moissac
}
    
func (c *阿让AgenCounty) BPenne佩纳() feud.Barony {
	return c.佩纳Penne
}
    
func (c *阿让AgenCounty) BRocamadour罗卡马杜尔() feud.Barony {
	return c.罗卡马杜尔Rocamadour
}
    
var CAgen阿让 AgenCounty = &阿让AgenCounty{}

func init() {
	f := CAgen阿让.(*阿让AgenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "215",
		Title:     "agen",
		TitleName: "阿让",
		TitleCode: "c_agen",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿让Agen = BAgen阿让
	f.阿让Agen.SetParent(f)
	
	f.布朗克福Blanquefort = BBlanquefort布朗克福
	f.布朗克福Blanquefort.SetParent(f)
	
	f.卡奥尔Cahors = BCahors卡奥尔
	f.卡奥尔Cahors.SetParent(f)
	
	f.菲雅克Figeac = BFigeac菲雅克
	f.菲雅克Figeac.SetParent(f)
	
	f.吕泽克Luzech = BLuzech吕泽克
	f.吕泽克Luzech.SetParent(f)
	
	f.穆瓦萨克Moissac = BMoissac穆瓦萨克
	f.穆瓦萨克Moissac.SetParent(f)
	
	f.佩纳Penne = BPenne佩纳
	f.佩纳Penne.SetParent(f)
	
	f.罗卡马杜尔Rocamadour = BRocamadour罗卡马杜尔
	f.罗卡马杜尔Rocamadour.SetParent(f)
	
}
