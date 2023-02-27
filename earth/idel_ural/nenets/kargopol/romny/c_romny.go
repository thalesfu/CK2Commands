package romny

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RomnyCounty interface {
    feud.County
    BBykovskaya贝科夫斯卡亚() 	feud.Barony
    BKholuy霍卢伊() 	feud.Barony
    BMarinskaya马林斯卡亚() 	feud.Barony
    BUgol乌戈尔() 	feud.Barony
    BVozhega沃热加() 	feud.Barony
    BVysokaya维索卡亚() 	feud.Barony
    BYuchka尤奇卡() 	feud.Barony
}

type 罗姆内RomnyCounty struct {
	feud.BaseCounty
	贝科夫斯卡亚Bykovskaya 	feud.Barony
	霍卢伊Kholuy 	feud.Barony
	马林斯卡亚Marinskaya 	feud.Barony
	乌戈尔Ugol 	feud.Barony
	沃热加Vozhega 	feud.Barony
	维索卡亚Vysokaya 	feud.Barony
	尤奇卡Yuchka 	feud.Barony
}

func (c *罗姆内RomnyCounty) BBykovskaya贝科夫斯卡亚() feud.Barony {
	return c.贝科夫斯卡亚Bykovskaya
}
    
func (c *罗姆内RomnyCounty) BKholuy霍卢伊() feud.Barony {
	return c.霍卢伊Kholuy
}
    
func (c *罗姆内RomnyCounty) BMarinskaya马林斯卡亚() feud.Barony {
	return c.马林斯卡亚Marinskaya
}
    
func (c *罗姆内RomnyCounty) BUgol乌戈尔() feud.Barony {
	return c.乌戈尔Ugol
}
    
func (c *罗姆内RomnyCounty) BVozhega沃热加() feud.Barony {
	return c.沃热加Vozhega
}
    
func (c *罗姆内RomnyCounty) BVysokaya维索卡亚() feud.Barony {
	return c.维索卡亚Vysokaya
}
    
func (c *罗姆内RomnyCounty) BYuchka尤奇卡() feud.Barony {
	return c.尤奇卡Yuchka
}
    
var CRomny罗姆内 RomnyCounty = &罗姆内RomnyCounty{}

func init() {
	f := CRomny罗姆内.(*罗姆内RomnyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "403",
		Title:     "romny",
		TitleName: "罗姆内",
		TitleCode: "c_romny",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝科夫斯卡亚Bykovskaya = BBykovskaya贝科夫斯卡亚
	f.贝科夫斯卡亚Bykovskaya.SetParent(f)
	
	f.霍卢伊Kholuy = BKholuy霍卢伊
	f.霍卢伊Kholuy.SetParent(f)
	
	f.马林斯卡亚Marinskaya = BMarinskaya马林斯卡亚
	f.马林斯卡亚Marinskaya.SetParent(f)
	
	f.乌戈尔Ugol = BUgol乌戈尔
	f.乌戈尔Ugol.SetParent(f)
	
	f.沃热加Vozhega = BVozhega沃热加
	f.沃热加Vozhega.SetParent(f)
	
	f.维索卡亚Vysokaya = BVysokaya维索卡亚
	f.维索卡亚Vysokaya.SetParent(f)
	
	f.尤奇卡Yuchka = BYuchka尤奇卡
	f.尤奇卡Yuchka.SetParent(f)
	
}
