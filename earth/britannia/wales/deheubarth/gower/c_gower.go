package gower

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GowerCounty interface {
    feud.County
    BCaerphilly_gower卡菲利() 	feud.Barony
    BCarmarthen卡马森() 	feud.Barony
    BDinefwr迪内弗尔() 	feud.Barony
    BKidwelly基德韦利() 	feud.Barony
    BLlandeilo兰代洛() 	feud.Barony
    BLoughor拉赫尔() 	feud.Barony
    BSwansea斯旺西() 	feud.Barony
}

type 高尔GowerCounty struct {
	feud.BaseCounty
	卡菲利Caerphilly_gower 	feud.Barony
	卡马森Carmarthen 	feud.Barony
	迪内弗尔Dinefwr 	feud.Barony
	基德韦利Kidwelly 	feud.Barony
	兰代洛Llandeilo 	feud.Barony
	拉赫尔Loughor 	feud.Barony
	斯旺西Swansea 	feud.Barony
}

func (c *高尔GowerCounty) BCaerphilly_gower卡菲利() feud.Barony {
	return c.卡菲利Caerphilly_gower
}
    
func (c *高尔GowerCounty) BCarmarthen卡马森() feud.Barony {
	return c.卡马森Carmarthen
}
    
func (c *高尔GowerCounty) BDinefwr迪内弗尔() feud.Barony {
	return c.迪内弗尔Dinefwr
}
    
func (c *高尔GowerCounty) BKidwelly基德韦利() feud.Barony {
	return c.基德韦利Kidwelly
}
    
func (c *高尔GowerCounty) BLlandeilo兰代洛() feud.Barony {
	return c.兰代洛Llandeilo
}
    
func (c *高尔GowerCounty) BLoughor拉赫尔() feud.Barony {
	return c.拉赫尔Loughor
}
    
func (c *高尔GowerCounty) BSwansea斯旺西() feud.Barony {
	return c.斯旺西Swansea
}
    
var CGower高尔 GowerCounty = &高尔GowerCounty{}

func init() {
	f := CGower高尔.(*高尔GowerCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1946",
		Title:     "gower",
		TitleName: "高尔",
		TitleCode: "c_gower",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡菲利Caerphilly_gower = BCaerphilly_gower卡菲利
	f.卡菲利Caerphilly_gower.SetParent(f)
	
	f.卡马森Carmarthen = BCarmarthen卡马森
	f.卡马森Carmarthen.SetParent(f)
	
	f.迪内弗尔Dinefwr = BDinefwr迪内弗尔
	f.迪内弗尔Dinefwr.SetParent(f)
	
	f.基德韦利Kidwelly = BKidwelly基德韦利
	f.基德韦利Kidwelly.SetParent(f)
	
	f.兰代洛Llandeilo = BLlandeilo兰代洛
	f.兰代洛Llandeilo.SetParent(f)
	
	f.拉赫尔Loughor = BLoughor拉赫尔
	f.拉赫尔Loughor.SetParent(f)
	
	f.斯旺西Swansea = BSwansea斯旺西
	f.斯旺西Swansea.SetParent(f)
	
}
