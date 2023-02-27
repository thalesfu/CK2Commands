package durdzukia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DurdzukiaCounty interface {
    feud.County
    BAlagir阿拉吉尔() 	feud.Barony
    BArdon阿尔东() 	feud.Barony
    BBeslan别斯兰() 	feud.Barony
    BDargavs达尔加夫斯() 	feud.Barony
    BDurdzukia杜尔祖基亚() 	feud.Barony
    BMizur米祖尔() 	feud.Barony
    BNazran纳兹兰() 	feud.Barony
}

type 杜尔祖基亚DurdzukiaCounty struct {
	feud.BaseCounty
	阿拉吉尔Alagir 	feud.Barony
	阿尔东Ardon 	feud.Barony
	别斯兰Beslan 	feud.Barony
	达尔加夫斯Dargavs 	feud.Barony
	杜尔祖基亚Durdzukia 	feud.Barony
	米祖尔Mizur 	feud.Barony
	纳兹兰Nazran 	feud.Barony
}

func (c *杜尔祖基亚DurdzukiaCounty) BAlagir阿拉吉尔() feud.Barony {
	return c.阿拉吉尔Alagir
}
    
func (c *杜尔祖基亚DurdzukiaCounty) BArdon阿尔东() feud.Barony {
	return c.阿尔东Ardon
}
    
func (c *杜尔祖基亚DurdzukiaCounty) BBeslan别斯兰() feud.Barony {
	return c.别斯兰Beslan
}
    
func (c *杜尔祖基亚DurdzukiaCounty) BDargavs达尔加夫斯() feud.Barony {
	return c.达尔加夫斯Dargavs
}
    
func (c *杜尔祖基亚DurdzukiaCounty) BDurdzukia杜尔祖基亚() feud.Barony {
	return c.杜尔祖基亚Durdzukia
}
    
func (c *杜尔祖基亚DurdzukiaCounty) BMizur米祖尔() feud.Barony {
	return c.米祖尔Mizur
}
    
func (c *杜尔祖基亚DurdzukiaCounty) BNazran纳兹兰() feud.Barony {
	return c.纳兹兰Nazran
}
    
var CDurdzukia杜尔祖基亚 DurdzukiaCounty = &杜尔祖基亚DurdzukiaCounty{}

func init() {
	f := CDurdzukia杜尔祖基亚.(*杜尔祖基亚DurdzukiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1809",
		Title:     "durdzukia",
		TitleName: "杜尔祖基亚",
		TitleCode: "c_durdzukia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉吉尔Alagir = BAlagir阿拉吉尔
	f.阿拉吉尔Alagir.SetParent(f)
	
	f.阿尔东Ardon = BArdon阿尔东
	f.阿尔东Ardon.SetParent(f)
	
	f.别斯兰Beslan = BBeslan别斯兰
	f.别斯兰Beslan.SetParent(f)
	
	f.达尔加夫斯Dargavs = BDargavs达尔加夫斯
	f.达尔加夫斯Dargavs.SetParent(f)
	
	f.杜尔祖基亚Durdzukia = BDurdzukia杜尔祖基亚
	f.杜尔祖基亚Durdzukia.SetParent(f)
	
	f.米祖尔Mizur = BMizur米祖尔
	f.米祖尔Mizur.SetParent(f)
	
	f.纳兹兰Nazran = BNazran纳兹兰
	f.纳兹兰Nazran.SetParent(f)
	
}
