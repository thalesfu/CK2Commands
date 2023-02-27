package lhasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LhasaCounty interface {
    feud.County
    BBarkhor八廓() 	feud.Barony
    BDonggar东嘎() 	feud.Barony
    BJokhang觉康() 	feud.Barony
    BLhasa逻些() 	feud.Barony
    BMarpori玛布日() 	feud.Barony
    BNaiqung乃琼() 	feud.Barony
    BRamoche绕木切() 	feud.Barony
}

type 逻些LhasaCounty struct {
	feud.BaseCounty
	八廓Barkhor 	feud.Barony
	东嘎Donggar 	feud.Barony
	觉康Jokhang 	feud.Barony
	逻些Lhasa 	feud.Barony
	玛布日Marpori 	feud.Barony
	乃琼Naiqung 	feud.Barony
	绕木切Ramoche 	feud.Barony
}

func (c *逻些LhasaCounty) BBarkhor八廓() feud.Barony {
	return c.八廓Barkhor
}
    
func (c *逻些LhasaCounty) BDonggar东嘎() feud.Barony {
	return c.东嘎Donggar
}
    
func (c *逻些LhasaCounty) BJokhang觉康() feud.Barony {
	return c.觉康Jokhang
}
    
func (c *逻些LhasaCounty) BLhasa逻些() feud.Barony {
	return c.逻些Lhasa
}
    
func (c *逻些LhasaCounty) BMarpori玛布日() feud.Barony {
	return c.玛布日Marpori
}
    
func (c *逻些LhasaCounty) BNaiqung乃琼() feud.Barony {
	return c.乃琼Naiqung
}
    
func (c *逻些LhasaCounty) BRamoche绕木切() feud.Barony {
	return c.绕木切Ramoche
}
    
var CLhasa逻些 LhasaCounty = &逻些LhasaCounty{}

func init() {
	f := CLhasa逻些.(*逻些LhasaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1498",
		Title:     "lhasa",
		TitleName: "逻些",
		TitleCode: "c_lhasa",
		Baronies:  map[string]feud.Barony{},
	}

	f.八廓Barkhor = BBarkhor八廓
	f.八廓Barkhor.SetParent(f)
	
	f.东嘎Donggar = BDonggar东嘎
	f.东嘎Donggar.SetParent(f)
	
	f.觉康Jokhang = BJokhang觉康
	f.觉康Jokhang.SetParent(f)
	
	f.逻些Lhasa = BLhasa逻些
	f.逻些Lhasa.SetParent(f)
	
	f.玛布日Marpori = BMarpori玛布日
	f.玛布日Marpori.SetParent(f)
	
	f.乃琼Naiqung = BNaiqung乃琼
	f.乃琼Naiqung.SetParent(f)
	
	f.绕木切Ramoche = BRamoche绕木切
	f.绕木切Ramoche.SetParent(f)
	
}
