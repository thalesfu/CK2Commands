package taza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TazaCounty interface {
    feud.County
    BGeurcif盖尔西夫() 	feud.Barony
    BMania马尼亚() 	feud.Barony
    BSafsaf塞夫萨夫() 	feud.Barony
    BSaka萨卡() 	feud.Barony
    BTamdrost塔姆德鲁斯特() 	feud.Barony
    BTaza塔扎() 	feud.Barony
    BTighdouine蒂格杜因() 	feud.Barony
}

type 塔扎TazaCounty struct {
	feud.BaseCounty
	盖尔西夫Geurcif 	feud.Barony
	马尼亚Mania 	feud.Barony
	塞夫萨夫Safsaf 	feud.Barony
	萨卡Saka 	feud.Barony
	塔姆德鲁斯特Tamdrost 	feud.Barony
	塔扎Taza 	feud.Barony
	蒂格杜因Tighdouine 	feud.Barony
}

func (c *塔扎TazaCounty) BGeurcif盖尔西夫() feud.Barony {
	return c.盖尔西夫Geurcif
}
    
func (c *塔扎TazaCounty) BMania马尼亚() feud.Barony {
	return c.马尼亚Mania
}
    
func (c *塔扎TazaCounty) BSafsaf塞夫萨夫() feud.Barony {
	return c.塞夫萨夫Safsaf
}
    
func (c *塔扎TazaCounty) BSaka萨卡() feud.Barony {
	return c.萨卡Saka
}
    
func (c *塔扎TazaCounty) BTamdrost塔姆德鲁斯特() feud.Barony {
	return c.塔姆德鲁斯特Tamdrost
}
    
func (c *塔扎TazaCounty) BTaza塔扎() feud.Barony {
	return c.塔扎Taza
}
    
func (c *塔扎TazaCounty) BTighdouine蒂格杜因() feud.Barony {
	return c.蒂格杜因Tighdouine
}
    
var CTaza塔扎 TazaCounty = &塔扎TazaCounty{}

func init() {
	f := CTaza塔扎.(*塔扎TazaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1773",
		Title:     "taza",
		TitleName: "塔扎",
		TitleCode: "c_taza",
		Baronies:  map[string]feud.Barony{},
	}

	f.盖尔西夫Geurcif = BGeurcif盖尔西夫
	f.盖尔西夫Geurcif.SetParent(f)
	
	f.马尼亚Mania = BMania马尼亚
	f.马尼亚Mania.SetParent(f)
	
	f.塞夫萨夫Safsaf = BSafsaf塞夫萨夫
	f.塞夫萨夫Safsaf.SetParent(f)
	
	f.萨卡Saka = BSaka萨卡
	f.萨卡Saka.SetParent(f)
	
	f.塔姆德鲁斯特Tamdrost = BTamdrost塔姆德鲁斯特
	f.塔姆德鲁斯特Tamdrost.SetParent(f)
	
	f.塔扎Taza = BTaza塔扎
	f.塔扎Taza.SetParent(f)
	
	f.蒂格杜因Tighdouine = BTighdouine蒂格杜因
	f.蒂格杜因Tighdouine.SetParent(f)
	
}
