package kursk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KurskCounty interface {
    feud.County
    BFatej法捷日() 	feud.Barony
    BKourchatov库尔恰托夫() 	feud.Barony
    BKursk库尔斯克() 	feud.Barony
    BOlgov奥尔戈夫() 	feud.Barony
    BPryamitsyno普里亚米齐诺() 	feud.Barony
    BSukovo苏科沃() 	feud.Barony
    BZolotukhino佐洛图希诺() 	feud.Barony
}

type 库尔斯克KurskCounty struct {
	feud.BaseCounty
	法捷日Fatej 	feud.Barony
	库尔恰托夫Kourchatov 	feud.Barony
	库尔斯克Kursk 	feud.Barony
	奥尔戈夫Olgov 	feud.Barony
	普里亚米齐诺Pryamitsyno 	feud.Barony
	苏科沃Sukovo 	feud.Barony
	佐洛图希诺Zolotukhino 	feud.Barony
}

func (c *库尔斯克KurskCounty) BFatej法捷日() feud.Barony {
	return c.法捷日Fatej
}
    
func (c *库尔斯克KurskCounty) BKourchatov库尔恰托夫() feud.Barony {
	return c.库尔恰托夫Kourchatov
}
    
func (c *库尔斯克KurskCounty) BKursk库尔斯克() feud.Barony {
	return c.库尔斯克Kursk
}
    
func (c *库尔斯克KurskCounty) BOlgov奥尔戈夫() feud.Barony {
	return c.奥尔戈夫Olgov
}
    
func (c *库尔斯克KurskCounty) BPryamitsyno普里亚米齐诺() feud.Barony {
	return c.普里亚米齐诺Pryamitsyno
}
    
func (c *库尔斯克KurskCounty) BSukovo苏科沃() feud.Barony {
	return c.苏科沃Sukovo
}
    
func (c *库尔斯克KurskCounty) BZolotukhino佐洛图希诺() feud.Barony {
	return c.佐洛图希诺Zolotukhino
}
    
var CKursk库尔斯克 KurskCounty = &库尔斯克KurskCounty{}

func init() {
	f := CKursk库尔斯克.(*库尔斯克KurskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1681",
		Title:     "kursk",
		TitleName: "库尔斯克",
		TitleCode: "c_kursk",
		Baronies:  map[string]feud.Barony{},
	}

	f.法捷日Fatej = BFatej法捷日
	f.法捷日Fatej.SetParent(f)
	
	f.库尔恰托夫Kourchatov = BKourchatov库尔恰托夫
	f.库尔恰托夫Kourchatov.SetParent(f)
	
	f.库尔斯克Kursk = BKursk库尔斯克
	f.库尔斯克Kursk.SetParent(f)
	
	f.奥尔戈夫Olgov = BOlgov奥尔戈夫
	f.奥尔戈夫Olgov.SetParent(f)
	
	f.普里亚米齐诺Pryamitsyno = BPryamitsyno普里亚米齐诺
	f.普里亚米齐诺Pryamitsyno.SetParent(f)
	
	f.苏科沃Sukovo = BSukovo苏科沃
	f.苏科沃Sukovo.SetParent(f)
	
	f.佐洛图希诺Zolotukhino = BZolotukhino佐洛图希诺
	f.佐洛图希诺Zolotukhino.SetParent(f)
	
}
