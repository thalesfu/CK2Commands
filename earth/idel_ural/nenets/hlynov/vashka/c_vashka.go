package vashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VashkaCounty interface {
    feud.County
    BKeba克巴() 	feud.Barony
    BOlema奥列马() 	feud.Barony
    BRezya列贾() 	feud.Barony
    BSel_za谢利扎() 	feud.Barony
    BTsilenga齐连加() 	feud.Barony
    BVashka瓦什卡() 	feud.Barony
    BVazhgort瓦日戈尔特() 	feud.Barony
}

type 瓦什卡VashkaCounty struct {
	feud.BaseCounty
	克巴Keba 	feud.Barony
	奥列马Olema 	feud.Barony
	列贾Rezya 	feud.Barony
	谢利扎Sel_za 	feud.Barony
	齐连加Tsilenga 	feud.Barony
	瓦什卡Vashka 	feud.Barony
	瓦日戈尔特Vazhgort 	feud.Barony
}

func (c *瓦什卡VashkaCounty) BKeba克巴() feud.Barony {
	return c.克巴Keba
}
    
func (c *瓦什卡VashkaCounty) BOlema奥列马() feud.Barony {
	return c.奥列马Olema
}
    
func (c *瓦什卡VashkaCounty) BRezya列贾() feud.Barony {
	return c.列贾Rezya
}
    
func (c *瓦什卡VashkaCounty) BSel_za谢利扎() feud.Barony {
	return c.谢利扎Sel_za
}
    
func (c *瓦什卡VashkaCounty) BTsilenga齐连加() feud.Barony {
	return c.齐连加Tsilenga
}
    
func (c *瓦什卡VashkaCounty) BVashka瓦什卡() feud.Barony {
	return c.瓦什卡Vashka
}
    
func (c *瓦什卡VashkaCounty) BVazhgort瓦日戈尔特() feud.Barony {
	return c.瓦日戈尔特Vazhgort
}
    
var CVashka瓦什卡 VashkaCounty = &瓦什卡VashkaCounty{}

func init() {
	f := CVashka瓦什卡.(*瓦什卡VashkaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1825",
		Title:     "vashka",
		TitleName: "瓦什卡",
		TitleCode: "c_vashka",
		Baronies:  map[string]feud.Barony{},
	}

	f.克巴Keba = BKeba克巴
	f.克巴Keba.SetParent(f)
	
	f.奥列马Olema = BOlema奥列马
	f.奥列马Olema.SetParent(f)
	
	f.列贾Rezya = BRezya列贾
	f.列贾Rezya.SetParent(f)
	
	f.谢利扎Sel_za = BSel_za谢利扎
	f.谢利扎Sel_za.SetParent(f)
	
	f.齐连加Tsilenga = BTsilenga齐连加
	f.齐连加Tsilenga.SetParent(f)
	
	f.瓦什卡Vashka = BVashka瓦什卡
	f.瓦什卡Vashka.SetParent(f)
	
	f.瓦日戈尔特Vazhgort = BVazhgort瓦日戈尔特
	f.瓦日戈尔特Vazhgort.SetParent(f)
	
}
