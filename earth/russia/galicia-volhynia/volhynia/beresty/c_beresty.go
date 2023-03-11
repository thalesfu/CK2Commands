package beresty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BerestyCounty interface {
    feud.County
    BBeresty别列斯捷() 	feud.Barony
    BBielsk别尔斯克() 	feud.Barony
    BKamyanyets卡梅涅茨() 	feud.Barony
    BKobryn科布林() 	feud.Barony
    BKoden科登() 	feud.Barony
    BMielnik梅尔尼克() 	feud.Barony
    BParwa帕尔瓦() 	feud.Barony
    BWlodawa弗沃达瓦() 	feud.Barony
}

type 别列斯捷BerestyCounty struct {
	feud.BaseCounty
	别列斯捷Beresty 	feud.Barony
	别尔斯克Bielsk 	feud.Barony
	卡梅涅茨Kamyanyets 	feud.Barony
	科布林Kobryn 	feud.Barony
	科登Koden 	feud.Barony
	梅尔尼克Mielnik 	feud.Barony
	帕尔瓦Parwa 	feud.Barony
	弗沃达瓦Wlodawa 	feud.Barony
}

func (c *别列斯捷BerestyCounty) BBeresty别列斯捷() feud.Barony {
	return c.别列斯捷Beresty
}
    
func (c *别列斯捷BerestyCounty) BBielsk别尔斯克() feud.Barony {
	return c.别尔斯克Bielsk
}
    
func (c *别列斯捷BerestyCounty) BKamyanyets卡梅涅茨() feud.Barony {
	return c.卡梅涅茨Kamyanyets
}
    
func (c *别列斯捷BerestyCounty) BKobryn科布林() feud.Barony {
	return c.科布林Kobryn
}
    
func (c *别列斯捷BerestyCounty) BKoden科登() feud.Barony {
	return c.科登Koden
}
    
func (c *别列斯捷BerestyCounty) BMielnik梅尔尼克() feud.Barony {
	return c.梅尔尼克Mielnik
}
    
func (c *别列斯捷BerestyCounty) BParwa帕尔瓦() feud.Barony {
	return c.帕尔瓦Parwa
}
    
func (c *别列斯捷BerestyCounty) BWlodawa弗沃达瓦() feud.Barony {
	return c.弗沃达瓦Wlodawa
}
    
var CBeresty别列斯捷 BerestyCounty = &别列斯捷BerestyCounty{}

func init() {
	f := CBeresty别列斯捷.(*别列斯捷BerestyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "549",
		Title:     "beresty",
		TitleName: "别列斯捷",
		TitleCode: "c_beresty",
		Baronies:  map[string]feud.Barony{},
	}

	f.别列斯捷Beresty = BBeresty别列斯捷
	f.别列斯捷Beresty.SetParent(f)
	
	f.别尔斯克Bielsk = BBielsk别尔斯克
	f.别尔斯克Bielsk.SetParent(f)
	
	f.卡梅涅茨Kamyanyets = BKamyanyets卡梅涅茨
	f.卡梅涅茨Kamyanyets.SetParent(f)
	
	f.科布林Kobryn = BKobryn科布林
	f.科布林Kobryn.SetParent(f)
	
	f.科登Koden = BKoden科登
	f.科登Koden.SetParent(f)
	
	f.梅尔尼克Mielnik = BMielnik梅尔尼克
	f.梅尔尼克Mielnik.SetParent(f)
	
	f.帕尔瓦Parwa = BParwa帕尔瓦
	f.帕尔瓦Parwa.SetParent(f)
	
	f.弗沃达瓦Wlodawa = BWlodawa弗沃达瓦
	f.弗沃达瓦Wlodawa.SetParent(f)
	
}
