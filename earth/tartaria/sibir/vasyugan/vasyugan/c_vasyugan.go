package vasyugan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VasyuganCounty interface {
    feud.County
    BBylino贝利诺() 	feud.Barony
    BMegion梅吉翁() 	feud.Barony
    BNazino纳济诺() 	feud.Barony
    BSredny斯列德尼() 	feud.Barony
    BVampugol万普戈尔() 	feud.Barony
    BVasyugan瓦休甘() 	feud.Barony
    BZaytseva宰采瓦() 	feud.Barony
}

type 瓦休甘VasyuganCounty struct {
	feud.BaseCounty
	贝利诺Bylino 	feud.Barony
	梅吉翁Megion 	feud.Barony
	纳济诺Nazino 	feud.Barony
	斯列德尼Sredny 	feud.Barony
	万普戈尔Vampugol 	feud.Barony
	瓦休甘Vasyugan 	feud.Barony
	宰采瓦Zaytseva 	feud.Barony
}

func (c *瓦休甘VasyuganCounty) BBylino贝利诺() feud.Barony {
	return c.贝利诺Bylino
}
    
func (c *瓦休甘VasyuganCounty) BMegion梅吉翁() feud.Barony {
	return c.梅吉翁Megion
}
    
func (c *瓦休甘VasyuganCounty) BNazino纳济诺() feud.Barony {
	return c.纳济诺Nazino
}
    
func (c *瓦休甘VasyuganCounty) BSredny斯列德尼() feud.Barony {
	return c.斯列德尼Sredny
}
    
func (c *瓦休甘VasyuganCounty) BVampugol万普戈尔() feud.Barony {
	return c.万普戈尔Vampugol
}
    
func (c *瓦休甘VasyuganCounty) BVasyugan瓦休甘() feud.Barony {
	return c.瓦休甘Vasyugan
}
    
func (c *瓦休甘VasyuganCounty) BZaytseva宰采瓦() feud.Barony {
	return c.宰采瓦Zaytseva
}
    
var CVasyugan瓦休甘 VasyuganCounty = &瓦休甘VasyuganCounty{}

func init() {
	f := CVasyugan瓦休甘.(*瓦休甘VasyuganCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1860",
		Title:     "vasyugan",
		TitleName: "瓦休甘",
		TitleCode: "c_vasyugan",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝利诺Bylino = BBylino贝利诺
	f.贝利诺Bylino.SetParent(f)
	
	f.梅吉翁Megion = BMegion梅吉翁
	f.梅吉翁Megion.SetParent(f)
	
	f.纳济诺Nazino = BNazino纳济诺
	f.纳济诺Nazino.SetParent(f)
	
	f.斯列德尼Sredny = BSredny斯列德尼
	f.斯列德尼Sredny.SetParent(f)
	
	f.万普戈尔Vampugol = BVampugol万普戈尔
	f.万普戈尔Vampugol.SetParent(f)
	
	f.瓦休甘Vasyugan = BVasyugan瓦休甘
	f.瓦休甘Vasyugan.SetParent(f)
	
	f.宰采瓦Zaytseva = BZaytseva宰采瓦
	f.宰采瓦Zaytseva.SetParent(f)
	
}
