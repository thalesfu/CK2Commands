package katun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KatunCounty interface {
    feud.County
    BAya阿亚() 	feud.Barony
    BKarlushka卡尔卢什卡() 	feud.Barony
    BKatun卡通() 	feud.Barony
    BManzherok曼热罗克() 	feud.Barony
    BMayma迈马() 	feud.Barony
    BSouzga索乌兹加() 	feud.Barony
    BTurbaza图尔巴扎() 	feud.Barony
}

type 卡通KatunCounty struct {
	feud.BaseCounty
	阿亚Aya 	feud.Barony
	卡尔卢什卡Karlushka 	feud.Barony
	卡通Katun 	feud.Barony
	曼热罗克Manzherok 	feud.Barony
	迈马Mayma 	feud.Barony
	索乌兹加Souzga 	feud.Barony
	图尔巴扎Turbaza 	feud.Barony
}

func (c *卡通KatunCounty) BAya阿亚() feud.Barony {
	return c.阿亚Aya
}
    
func (c *卡通KatunCounty) BKarlushka卡尔卢什卡() feud.Barony {
	return c.卡尔卢什卡Karlushka
}
    
func (c *卡通KatunCounty) BKatun卡通() feud.Barony {
	return c.卡通Katun
}
    
func (c *卡通KatunCounty) BManzherok曼热罗克() feud.Barony {
	return c.曼热罗克Manzherok
}
    
func (c *卡通KatunCounty) BMayma迈马() feud.Barony {
	return c.迈马Mayma
}
    
func (c *卡通KatunCounty) BSouzga索乌兹加() feud.Barony {
	return c.索乌兹加Souzga
}
    
func (c *卡通KatunCounty) BTurbaza图尔巴扎() feud.Barony {
	return c.图尔巴扎Turbaza
}
    
var CKatun卡通 KatunCounty = &卡通KatunCounty{}

func init() {
	f := CKatun卡通.(*卡通KatunCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1871",
		Title:     "katun",
		TitleName: "卡通",
		TitleCode: "c_katun",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿亚Aya = BAya阿亚
	f.阿亚Aya.SetParent(f)
	
	f.卡尔卢什卡Karlushka = BKarlushka卡尔卢什卡
	f.卡尔卢什卡Karlushka.SetParent(f)
	
	f.卡通Katun = BKatun卡通
	f.卡通Katun.SetParent(f)
	
	f.曼热罗克Manzherok = BManzherok曼热罗克
	f.曼热罗克Manzherok.SetParent(f)
	
	f.迈马Mayma = BMayma迈马
	f.迈马Mayma.SetParent(f)
	
	f.索乌兹加Souzga = BSouzga索乌兹加
	f.索乌兹加Souzga.SetParent(f)
	
	f.图尔巴扎Turbaza = BTurbaza图尔巴扎
	f.图尔巴扎Turbaza.SetParent(f)
	
}
