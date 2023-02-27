package fuqi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FuqiCounty interface {
    feud.County
    BCaierqiong才尔琼() 	feud.Barony
    BFuqi伏俟() 	feud.Barony
    BNaerji纳尔几() 	feud.Barony
    BQieji切吉() 	feud.Barony
    BShatuo沙陀() 	feud.Barony
    BShinaihai石乃亥() 	feud.Barony
    BTiebujia铁卜加() 	feud.Barony
}

type 伏俟FuqiCounty struct {
	feud.BaseCounty
	才尔琼Caierqiong 	feud.Barony
	伏俟Fuqi 	feud.Barony
	纳尔几Naerji 	feud.Barony
	切吉Qieji 	feud.Barony
	沙陀Shatuo 	feud.Barony
	石乃亥Shinaihai 	feud.Barony
	铁卜加Tiebujia 	feud.Barony
}

func (c *伏俟FuqiCounty) BCaierqiong才尔琼() feud.Barony {
	return c.才尔琼Caierqiong
}
    
func (c *伏俟FuqiCounty) BFuqi伏俟() feud.Barony {
	return c.伏俟Fuqi
}
    
func (c *伏俟FuqiCounty) BNaerji纳尔几() feud.Barony {
	return c.纳尔几Naerji
}
    
func (c *伏俟FuqiCounty) BQieji切吉() feud.Barony {
	return c.切吉Qieji
}
    
func (c *伏俟FuqiCounty) BShatuo沙陀() feud.Barony {
	return c.沙陀Shatuo
}
    
func (c *伏俟FuqiCounty) BShinaihai石乃亥() feud.Barony {
	return c.石乃亥Shinaihai
}
    
func (c *伏俟FuqiCounty) BTiebujia铁卜加() feud.Barony {
	return c.铁卜加Tiebujia
}
    
var CFuqi伏俟 FuqiCounty = &伏俟FuqiCounty{}

func init() {
	f := CFuqi伏俟.(*伏俟FuqiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1502",
		Title:     "fuqi",
		TitleName: "伏俟",
		TitleCode: "c_fuqi",
		Baronies:  map[string]feud.Barony{},
	}

	f.才尔琼Caierqiong = BCaierqiong才尔琼
	f.才尔琼Caierqiong.SetParent(f)
	
	f.伏俟Fuqi = BFuqi伏俟
	f.伏俟Fuqi.SetParent(f)
	
	f.纳尔几Naerji = BNaerji纳尔几
	f.纳尔几Naerji.SetParent(f)
	
	f.切吉Qieji = BQieji切吉
	f.切吉Qieji.SetParent(f)
	
	f.沙陀Shatuo = BShatuo沙陀
	f.沙陀Shatuo.SetParent(f)
	
	f.石乃亥Shinaihai = BShinaihai石乃亥
	f.石乃亥Shinaihai.SetParent(f)
	
	f.铁卜加Tiebujia = BTiebujia铁卜加
	f.铁卜加Tiebujia.SetParent(f)
	
}
