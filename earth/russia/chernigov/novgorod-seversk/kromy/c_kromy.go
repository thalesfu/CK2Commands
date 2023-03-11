package kromy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KromyCounty interface {
    feud.County
    BDmitrievka季米特里耶夫卡() 	feud.Barony
    BKromy克罗梅() 	feud.Barony
    BLokot洛科季() 	feud.Barony
    BNavlya纳夫利亚() 	feud.Barony
    BNechayeva涅恰耶瓦() 	feud.Barony
    BSakovnina萨科夫尼纳() 	feud.Barony
    BSevsk谢夫斯克() 	feud.Barony
}

type 克罗梅KromyCounty struct {
	feud.BaseCounty
	季米特里耶夫卡Dmitrievka 	feud.Barony
	克罗梅Kromy 	feud.Barony
	洛科季Lokot 	feud.Barony
	纳夫利亚Navlya 	feud.Barony
	涅恰耶瓦Nechayeva 	feud.Barony
	萨科夫尼纳Sakovnina 	feud.Barony
	谢夫斯克Sevsk 	feud.Barony
}

func (c *克罗梅KromyCounty) BDmitrievka季米特里耶夫卡() feud.Barony {
	return c.季米特里耶夫卡Dmitrievka
}
    
func (c *克罗梅KromyCounty) BKromy克罗梅() feud.Barony {
	return c.克罗梅Kromy
}
    
func (c *克罗梅KromyCounty) BLokot洛科季() feud.Barony {
	return c.洛科季Lokot
}
    
func (c *克罗梅KromyCounty) BNavlya纳夫利亚() feud.Barony {
	return c.纳夫利亚Navlya
}
    
func (c *克罗梅KromyCounty) BNechayeva涅恰耶瓦() feud.Barony {
	return c.涅恰耶瓦Nechayeva
}
    
func (c *克罗梅KromyCounty) BSakovnina萨科夫尼纳() feud.Barony {
	return c.萨科夫尼纳Sakovnina
}
    
func (c *克罗梅KromyCounty) BSevsk谢夫斯克() feud.Barony {
	return c.谢夫斯克Sevsk
}
    
var CKromy克罗梅 KromyCounty = &克罗梅KromyCounty{}

func init() {
	f := CKromy克罗梅.(*克罗梅KromyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1676",
		Title:     "kromy",
		TitleName: "克罗梅",
		TitleCode: "c_kromy",
		Baronies:  map[string]feud.Barony{},
	}

	f.季米特里耶夫卡Dmitrievka = BDmitrievka季米特里耶夫卡
	f.季米特里耶夫卡Dmitrievka.SetParent(f)
	
	f.克罗梅Kromy = BKromy克罗梅
	f.克罗梅Kromy.SetParent(f)
	
	f.洛科季Lokot = BLokot洛科季
	f.洛科季Lokot.SetParent(f)
	
	f.纳夫利亚Navlya = BNavlya纳夫利亚
	f.纳夫利亚Navlya.SetParent(f)
	
	f.涅恰耶瓦Nechayeva = BNechayeva涅恰耶瓦
	f.涅恰耶瓦Nechayeva.SetParent(f)
	
	f.萨科夫尼纳Sakovnina = BSakovnina萨科夫尼纳
	f.萨科夫尼纳Sakovnina.SetParent(f)
	
	f.谢夫斯克Sevsk = BSevsk谢夫斯克
	f.谢夫斯克Sevsk.SetParent(f)
	
}
