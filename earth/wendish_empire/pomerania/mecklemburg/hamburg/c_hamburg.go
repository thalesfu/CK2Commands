package hamburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HamburgCounty interface {
    feud.County
    BAltona阿尔托纳() 	feud.Barony
    BBrunsbuttel布伦斯比特尔() 	feud.Barony
    BBuxtehude布克斯特胡德() 	feud.Barony
    BHamburg汉堡() 	feud.Barony
    BHarburg哈尔堡() 	feud.Barony
    BLokstedt洛克施泰特() 	feud.Barony
    BNiendorf宁多夫() 	feud.Barony
}

type 汉堡HamburgCounty struct {
	feud.BaseCounty
	阿尔托纳Altona 	feud.Barony
	布伦斯比特尔Brunsbuttel 	feud.Barony
	布克斯特胡德Buxtehude 	feud.Barony
	汉堡Hamburg 	feud.Barony
	哈尔堡Harburg 	feud.Barony
	洛克施泰特Lokstedt 	feud.Barony
	宁多夫Niendorf 	feud.Barony
}

func (c *汉堡HamburgCounty) BAltona阿尔托纳() feud.Barony {
	return c.阿尔托纳Altona
}
    
func (c *汉堡HamburgCounty) BBrunsbuttel布伦斯比特尔() feud.Barony {
	return c.布伦斯比特尔Brunsbuttel
}
    
func (c *汉堡HamburgCounty) BBuxtehude布克斯特胡德() feud.Barony {
	return c.布克斯特胡德Buxtehude
}
    
func (c *汉堡HamburgCounty) BHamburg汉堡() feud.Barony {
	return c.汉堡Hamburg
}
    
func (c *汉堡HamburgCounty) BHarburg哈尔堡() feud.Barony {
	return c.哈尔堡Harburg
}
    
func (c *汉堡HamburgCounty) BLokstedt洛克施泰特() feud.Barony {
	return c.洛克施泰特Lokstedt
}
    
func (c *汉堡HamburgCounty) BNiendorf宁多夫() feud.Barony {
	return c.宁多夫Niendorf
}
    
var CHamburg汉堡 HamburgCounty = &汉堡HamburgCounty{}

func init() {
	f := CHamburg汉堡.(*汉堡HamburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "261",
		Title:     "hamburg",
		TitleName: "汉堡",
		TitleCode: "c_hamburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔托纳Altona = BAltona阿尔托纳
	f.阿尔托纳Altona.SetParent(f)
	
	f.布伦斯比特尔Brunsbuttel = BBrunsbuttel布伦斯比特尔
	f.布伦斯比特尔Brunsbuttel.SetParent(f)
	
	f.布克斯特胡德Buxtehude = BBuxtehude布克斯特胡德
	f.布克斯特胡德Buxtehude.SetParent(f)
	
	f.汉堡Hamburg = BHamburg汉堡
	f.汉堡Hamburg.SetParent(f)
	
	f.哈尔堡Harburg = BHarburg哈尔堡
	f.哈尔堡Harburg.SetParent(f)
	
	f.洛克施泰特Lokstedt = BLokstedt洛克施泰特
	f.洛克施泰特Lokstedt.SetParent(f)
	
	f.宁多夫Niendorf = BNiendorf宁多夫
	f.宁多夫Niendorf.SetParent(f)
	
}
