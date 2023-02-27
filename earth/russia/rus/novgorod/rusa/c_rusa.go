package rusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RusaCounty interface {
    feud.County
    BBelebelka别列别尔卡() 	feud.Barony
    BKorchevka科尔乔夫卡() 	feud.Barony
    BLyublino柳布利诺() 	feud.Barony
    BPeregino佩列吉诺() 	feud.Barony
    BPoddorye波多里耶() 	feud.Barony
    BRamushevo拉穆舍沃() 	feud.Barony
    BRusa鲁萨() 	feud.Barony
}

type 鲁萨RusaCounty struct {
	feud.BaseCounty
	别列别尔卡Belebelka 	feud.Barony
	科尔乔夫卡Korchevka 	feud.Barony
	柳布利诺Lyublino 	feud.Barony
	佩列吉诺Peregino 	feud.Barony
	波多里耶Poddorye 	feud.Barony
	拉穆舍沃Ramushevo 	feud.Barony
	鲁萨Rusa 	feud.Barony
}

func (c *鲁萨RusaCounty) BBelebelka别列别尔卡() feud.Barony {
	return c.别列别尔卡Belebelka
}
    
func (c *鲁萨RusaCounty) BKorchevka科尔乔夫卡() feud.Barony {
	return c.科尔乔夫卡Korchevka
}
    
func (c *鲁萨RusaCounty) BLyublino柳布利诺() feud.Barony {
	return c.柳布利诺Lyublino
}
    
func (c *鲁萨RusaCounty) BPeregino佩列吉诺() feud.Barony {
	return c.佩列吉诺Peregino
}
    
func (c *鲁萨RusaCounty) BPoddorye波多里耶() feud.Barony {
	return c.波多里耶Poddorye
}
    
func (c *鲁萨RusaCounty) BRamushevo拉穆舍沃() feud.Barony {
	return c.拉穆舍沃Ramushevo
}
    
func (c *鲁萨RusaCounty) BRusa鲁萨() feud.Barony {
	return c.鲁萨Rusa
}
    
var CRusa鲁萨 RusaCounty = &鲁萨RusaCounty{}

func init() {
	f := CRusa鲁萨.(*鲁萨RusaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1662",
		Title:     "rusa",
		TitleName: "鲁萨",
		TitleCode: "c_rusa",
		Baronies:  map[string]feud.Barony{},
	}

	f.别列别尔卡Belebelka = BBelebelka别列别尔卡
	f.别列别尔卡Belebelka.SetParent(f)
	
	f.科尔乔夫卡Korchevka = BKorchevka科尔乔夫卡
	f.科尔乔夫卡Korchevka.SetParent(f)
	
	f.柳布利诺Lyublino = BLyublino柳布利诺
	f.柳布利诺Lyublino.SetParent(f)
	
	f.佩列吉诺Peregino = BPeregino佩列吉诺
	f.佩列吉诺Peregino.SetParent(f)
	
	f.波多里耶Poddorye = BPoddorye波多里耶
	f.波多里耶Poddorye.SetParent(f)
	
	f.拉穆舍沃Ramushevo = BRamushevo拉穆舍沃
	f.拉穆舍沃Ramushevo.SetParent(f)
	
	f.鲁萨Rusa = BRusa鲁萨
	f.鲁萨Rusa.SetParent(f)
	
}
