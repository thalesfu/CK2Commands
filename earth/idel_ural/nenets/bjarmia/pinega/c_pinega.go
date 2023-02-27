package pinega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PinegaCounty interface {
    feud.County
    BKuloy库洛伊() 	feud.Barony
    BKyolda基奥尔达() 	feud.Barony
    BNemnyuga涅姆纽加() 	feud.Barony
    BNizhneye尼日涅耶() 	feud.Barony
    BSotka索特卡() 	feud.Barony
    BSoyana索亚纳() 	feud.Barony
    BVilan维兰() 	feud.Barony
}

type 库洛伊PinegaCounty struct {
	feud.BaseCounty
	库洛伊Kuloy 	feud.Barony
	基奥尔达Kyolda 	feud.Barony
	涅姆纽加Nemnyuga 	feud.Barony
	尼日涅耶Nizhneye 	feud.Barony
	索特卡Sotka 	feud.Barony
	索亚纳Soyana 	feud.Barony
	维兰Vilan 	feud.Barony
}

func (c *库洛伊PinegaCounty) BKuloy库洛伊() feud.Barony {
	return c.库洛伊Kuloy
}
    
func (c *库洛伊PinegaCounty) BKyolda基奥尔达() feud.Barony {
	return c.基奥尔达Kyolda
}
    
func (c *库洛伊PinegaCounty) BNemnyuga涅姆纽加() feud.Barony {
	return c.涅姆纽加Nemnyuga
}
    
func (c *库洛伊PinegaCounty) BNizhneye尼日涅耶() feud.Barony {
	return c.尼日涅耶Nizhneye
}
    
func (c *库洛伊PinegaCounty) BSotka索特卡() feud.Barony {
	return c.索特卡Sotka
}
    
func (c *库洛伊PinegaCounty) BSoyana索亚纳() feud.Barony {
	return c.索亚纳Soyana
}
    
func (c *库洛伊PinegaCounty) BVilan维兰() feud.Barony {
	return c.维兰Vilan
}
    
var CPinega库洛伊 PinegaCounty = &库洛伊PinegaCounty{}

func init() {
	f := CPinega库洛伊.(*库洛伊PinegaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1822",
		Title:     "pinega",
		TitleName: "库洛伊",
		TitleCode: "c_pinega",
		Baronies:  map[string]feud.Barony{},
	}

	f.库洛伊Kuloy = BKuloy库洛伊
	f.库洛伊Kuloy.SetParent(f)
	
	f.基奥尔达Kyolda = BKyolda基奥尔达
	f.基奥尔达Kyolda.SetParent(f)
	
	f.涅姆纽加Nemnyuga = BNemnyuga涅姆纽加
	f.涅姆纽加Nemnyuga.SetParent(f)
	
	f.尼日涅耶Nizhneye = BNizhneye尼日涅耶
	f.尼日涅耶Nizhneye.SetParent(f)
	
	f.索特卡Sotka = BSotka索特卡
	f.索特卡Sotka.SetParent(f)
	
	f.索亚纳Soyana = BSoyana索亚纳
	f.索亚纳Soyana.SetParent(f)
	
	f.维兰Vilan = BVilan维兰
	f.维兰Vilan.SetParent(f)
	
}
