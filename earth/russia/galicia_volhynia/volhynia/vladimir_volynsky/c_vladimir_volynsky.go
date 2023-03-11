package vladimir_volynsky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Vladimir_volynskyCounty interface {
    feud.County
    BKovel科韦利() 	feud.Barony
    BLityn利京() 	feud.Barony
    BLuboml柳博姆利() 	feud.Barony
    BLukiv卢基夫() 	feud.Barony
    BMochalky莫恰尔基() 	feud.Barony
    BOkhotnyky奥霍特尼基() 	feud.Barony
    BVladimirvolynsky弗拉基米尔沃伦斯基() 	feud.Barony
}

type 弗拉基米尔沃伦斯基Vladimir_volynskyCounty struct {
	feud.BaseCounty
	科韦利Kovel 	feud.Barony
	利京Lityn 	feud.Barony
	柳博姆利Luboml 	feud.Barony
	卢基夫Lukiv 	feud.Barony
	莫恰尔基Mochalky 	feud.Barony
	奥霍特尼基Okhotnyky 	feud.Barony
	弗拉基米尔沃伦斯基Vladimirvolynsky 	feud.Barony
}

func (c *弗拉基米尔沃伦斯基Vladimir_volynskyCounty) BKovel科韦利() feud.Barony {
	return c.科韦利Kovel
}
    
func (c *弗拉基米尔沃伦斯基Vladimir_volynskyCounty) BLityn利京() feud.Barony {
	return c.利京Lityn
}
    
func (c *弗拉基米尔沃伦斯基Vladimir_volynskyCounty) BLuboml柳博姆利() feud.Barony {
	return c.柳博姆利Luboml
}
    
func (c *弗拉基米尔沃伦斯基Vladimir_volynskyCounty) BLukiv卢基夫() feud.Barony {
	return c.卢基夫Lukiv
}
    
func (c *弗拉基米尔沃伦斯基Vladimir_volynskyCounty) BMochalky莫恰尔基() feud.Barony {
	return c.莫恰尔基Mochalky
}
    
func (c *弗拉基米尔沃伦斯基Vladimir_volynskyCounty) BOkhotnyky奥霍特尼基() feud.Barony {
	return c.奥霍特尼基Okhotnyky
}
    
func (c *弗拉基米尔沃伦斯基Vladimir_volynskyCounty) BVladimirvolynsky弗拉基米尔沃伦斯基() feud.Barony {
	return c.弗拉基米尔沃伦斯基Vladimirvolynsky
}
    
var CVladimir_volynsky弗拉基米尔沃伦斯基 Vladimir_volynskyCounty = &弗拉基米尔沃伦斯基Vladimir_volynskyCounty{}

func init() {
	f := CVladimir_volynsky弗拉基米尔沃伦斯基.(*弗拉基米尔沃伦斯基Vladimir_volynskyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "535",
		Title:     "vladimir_volynsky",
		TitleName: "弗拉基米尔沃伦斯基",
		TitleCode: "c_vladimir_volynsky",
		Baronies:  map[string]feud.Barony{},
	}

	f.科韦利Kovel = BKovel科韦利
	f.科韦利Kovel.SetParent(f)
	
	f.利京Lityn = BLityn利京
	f.利京Lityn.SetParent(f)
	
	f.柳博姆利Luboml = BLuboml柳博姆利
	f.柳博姆利Luboml.SetParent(f)
	
	f.卢基夫Lukiv = BLukiv卢基夫
	f.卢基夫Lukiv.SetParent(f)
	
	f.莫恰尔基Mochalky = BMochalky莫恰尔基
	f.莫恰尔基Mochalky.SetParent(f)
	
	f.奥霍特尼基Okhotnyky = BOkhotnyky奥霍特尼基
	f.奥霍特尼基Okhotnyky.SetParent(f)
	
	f.弗拉基米尔沃伦斯基Vladimirvolynsky = BVladimirvolynsky弗拉基米尔沃伦斯基
	f.弗拉基米尔沃伦斯基Vladimirvolynsky.SetParent(f)
	
}
