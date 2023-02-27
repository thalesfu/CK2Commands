package aswan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AswanCounty interface {
    feud.County
    BAswan阿斯旺() 	feud.Barony
    BBigeh比格() 	feud.Barony
    BElefantina象岛() 	feud.Barony
    BKalabsha卡拉布萨() 	feud.Barony
    BPhilae菲莱() 	feud.Barony
    BShellal谢拉尔() 	feud.Barony
    BVeset维瑟特() 	feud.Barony
}

type 阿斯旺AswanCounty struct {
	feud.BaseCounty
	阿斯旺Aswan 	feud.Barony
	比格Bigeh 	feud.Barony
	象岛Elefantina 	feud.Barony
	卡拉布萨Kalabsha 	feud.Barony
	菲莱Philae 	feud.Barony
	谢拉尔Shellal 	feud.Barony
	维瑟特Veset 	feud.Barony
}

func (c *阿斯旺AswanCounty) BAswan阿斯旺() feud.Barony {
	return c.阿斯旺Aswan
}
    
func (c *阿斯旺AswanCounty) BBigeh比格() feud.Barony {
	return c.比格Bigeh
}
    
func (c *阿斯旺AswanCounty) BElefantina象岛() feud.Barony {
	return c.象岛Elefantina
}
    
func (c *阿斯旺AswanCounty) BKalabsha卡拉布萨() feud.Barony {
	return c.卡拉布萨Kalabsha
}
    
func (c *阿斯旺AswanCounty) BPhilae菲莱() feud.Barony {
	return c.菲莱Philae
}
    
func (c *阿斯旺AswanCounty) BShellal谢拉尔() feud.Barony {
	return c.谢拉尔Shellal
}
    
func (c *阿斯旺AswanCounty) BVeset维瑟特() feud.Barony {
	return c.维瑟特Veset
}
    
var CAswan阿斯旺 AswanCounty = &阿斯旺AswanCounty{}

func init() {
	f := CAswan阿斯旺.(*阿斯旺AswanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "794",
		Title:     "aswan",
		TitleName: "阿斯旺",
		TitleCode: "c_aswan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯旺Aswan = BAswan阿斯旺
	f.阿斯旺Aswan.SetParent(f)
	
	f.比格Bigeh = BBigeh比格
	f.比格Bigeh.SetParent(f)
	
	f.象岛Elefantina = BElefantina象岛
	f.象岛Elefantina.SetParent(f)
	
	f.卡拉布萨Kalabsha = BKalabsha卡拉布萨
	f.卡拉布萨Kalabsha.SetParent(f)
	
	f.菲莱Philae = BPhilae菲莱
	f.菲莱Philae.SetParent(f)
	
	f.谢拉尔Shellal = BShellal谢拉尔
	f.谢拉尔Shellal.SetParent(f)
	
	f.维瑟特Veset = BVeset维瑟特
	f.维瑟特Veset.SetParent(f)
	
}
