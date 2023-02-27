package suffolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SuffolkCounty interface {
    feud.County
    BBungay邦吉() 	feud.Barony
    BBury贝里() 	feud.Barony
    BDunwich邓尼奇() 	feud.Barony
    BEly伊利() 	feud.Barony
    BFramlingham弗瑞林姆() 	feud.Barony
    BIpswich伊普斯威奇() 	feud.Barony
    BLowestoft洛斯托夫特() 	feud.Barony
    BStow斯托() 	feud.Barony
}

type 萨福克SuffolkCounty struct {
	feud.BaseCounty
	邦吉Bungay 	feud.Barony
	贝里Bury 	feud.Barony
	邓尼奇Dunwich 	feud.Barony
	伊利Ely 	feud.Barony
	弗瑞林姆Framlingham 	feud.Barony
	伊普斯威奇Ipswich 	feud.Barony
	洛斯托夫特Lowestoft 	feud.Barony
	斯托Stow 	feud.Barony
}

func (c *萨福克SuffolkCounty) BBungay邦吉() feud.Barony {
	return c.邦吉Bungay
}
    
func (c *萨福克SuffolkCounty) BBury贝里() feud.Barony {
	return c.贝里Bury
}
    
func (c *萨福克SuffolkCounty) BDunwich邓尼奇() feud.Barony {
	return c.邓尼奇Dunwich
}
    
func (c *萨福克SuffolkCounty) BEly伊利() feud.Barony {
	return c.伊利Ely
}
    
func (c *萨福克SuffolkCounty) BFramlingham弗瑞林姆() feud.Barony {
	return c.弗瑞林姆Framlingham
}
    
func (c *萨福克SuffolkCounty) BIpswich伊普斯威奇() feud.Barony {
	return c.伊普斯威奇Ipswich
}
    
func (c *萨福克SuffolkCounty) BLowestoft洛斯托夫特() feud.Barony {
	return c.洛斯托夫特Lowestoft
}
    
func (c *萨福克SuffolkCounty) BStow斯托() feud.Barony {
	return c.斯托Stow
}
    
var CSuffolk萨福克 SuffolkCounty = &萨福克SuffolkCounty{}

func init() {
	f := CSuffolk萨福克.(*萨福克SuffolkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "71",
		Title:     "suffolk",
		TitleName: "萨福克",
		TitleCode: "c_suffolk",
		Baronies:  map[string]feud.Barony{},
	}

	f.邦吉Bungay = BBungay邦吉
	f.邦吉Bungay.SetParent(f)
	
	f.贝里Bury = BBury贝里
	f.贝里Bury.SetParent(f)
	
	f.邓尼奇Dunwich = BDunwich邓尼奇
	f.邓尼奇Dunwich.SetParent(f)
	
	f.伊利Ely = BEly伊利
	f.伊利Ely.SetParent(f)
	
	f.弗瑞林姆Framlingham = BFramlingham弗瑞林姆
	f.弗瑞林姆Framlingham.SetParent(f)
	
	f.伊普斯威奇Ipswich = BIpswich伊普斯威奇
	f.伊普斯威奇Ipswich.SetParent(f)
	
	f.洛斯托夫特Lowestoft = BLowestoft洛斯托夫特
	f.洛斯托夫特Lowestoft.SetParent(f)
	
	f.斯托Stow = BStow斯托
	f.斯托Stow.SetParent(f)
	
}
