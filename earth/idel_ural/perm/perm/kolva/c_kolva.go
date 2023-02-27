package kolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolvaCounty interface {
    feud.County
    BAyya艾亚() 	feud.Barony
    BKolva科尔瓦() 	feud.Barony
    BKumay库马伊() 	feud.Barony
    BMudyl穆德利() 	feud.Barony
    BNizva尼济瓦() 	feud.Barony
    BSeleya谢列亚() 	feud.Barony
    BTulpan图尔潘() 	feud.Barony
}

type 科尔瓦KolvaCounty struct {
	feud.BaseCounty
	艾亚Ayya 	feud.Barony
	科尔瓦Kolva 	feud.Barony
	库马伊Kumay 	feud.Barony
	穆德利Mudyl 	feud.Barony
	尼济瓦Nizva 	feud.Barony
	谢列亚Seleya 	feud.Barony
	图尔潘Tulpan 	feud.Barony
}

func (c *科尔瓦KolvaCounty) BAyya艾亚() feud.Barony {
	return c.艾亚Ayya
}
    
func (c *科尔瓦KolvaCounty) BKolva科尔瓦() feud.Barony {
	return c.科尔瓦Kolva
}
    
func (c *科尔瓦KolvaCounty) BKumay库马伊() feud.Barony {
	return c.库马伊Kumay
}
    
func (c *科尔瓦KolvaCounty) BMudyl穆德利() feud.Barony {
	return c.穆德利Mudyl
}
    
func (c *科尔瓦KolvaCounty) BNizva尼济瓦() feud.Barony {
	return c.尼济瓦Nizva
}
    
func (c *科尔瓦KolvaCounty) BSeleya谢列亚() feud.Barony {
	return c.谢列亚Seleya
}
    
func (c *科尔瓦KolvaCounty) BTulpan图尔潘() feud.Barony {
	return c.图尔潘Tulpan
}
    
var CKolva科尔瓦 KolvaCounty = &科尔瓦KolvaCounty{}

func init() {
	f := CKolva科尔瓦.(*科尔瓦KolvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1841",
		Title:     "kolva",
		TitleName: "科尔瓦",
		TitleCode: "c_kolva",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾亚Ayya = BAyya艾亚
	f.艾亚Ayya.SetParent(f)
	
	f.科尔瓦Kolva = BKolva科尔瓦
	f.科尔瓦Kolva.SetParent(f)
	
	f.库马伊Kumay = BKumay库马伊
	f.库马伊Kumay.SetParent(f)
	
	f.穆德利Mudyl = BMudyl穆德利
	f.穆德利Mudyl.SetParent(f)
	
	f.尼济瓦Nizva = BNizva尼济瓦
	f.尼济瓦Nizva.SetParent(f)
	
	f.谢列亚Seleya = BSeleya谢列亚
	f.谢列亚Seleya.SetParent(f)
	
	f.图尔潘Tulpan = BTulpan图尔潘
	f.图尔潘Tulpan.SetParent(f)
	
}
