package ili

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IliCounty interface {
    feud.County
    BIlibaliq亦力把里() 	feud.Barony
    BKax哈什() 	feud.Barony
    BKoktal科克塔尔() 	feud.Barony
    BKulja固勒扎() 	feud.Barony
    BKunes巩乃斯() 	feud.Barony
    BSharyn沙伦() 	feud.Barony
    BTekes特克斯() 	feud.Barony
}

type 伊丽IliCounty struct {
	feud.BaseCounty
	亦力把里Ilibaliq 	feud.Barony
	哈什Kax 	feud.Barony
	科克塔尔Koktal 	feud.Barony
	固勒扎Kulja 	feud.Barony
	巩乃斯Kunes 	feud.Barony
	沙伦Sharyn 	feud.Barony
	特克斯Tekes 	feud.Barony
}

func (c *伊丽IliCounty) BIlibaliq亦力把里() feud.Barony {
	return c.亦力把里Ilibaliq
}
    
func (c *伊丽IliCounty) BKax哈什() feud.Barony {
	return c.哈什Kax
}
    
func (c *伊丽IliCounty) BKoktal科克塔尔() feud.Barony {
	return c.科克塔尔Koktal
}
    
func (c *伊丽IliCounty) BKulja固勒扎() feud.Barony {
	return c.固勒扎Kulja
}
    
func (c *伊丽IliCounty) BKunes巩乃斯() feud.Barony {
	return c.巩乃斯Kunes
}
    
func (c *伊丽IliCounty) BSharyn沙伦() feud.Barony {
	return c.沙伦Sharyn
}
    
func (c *伊丽IliCounty) BTekes特克斯() feud.Barony {
	return c.特克斯Tekes
}
    
var CIli伊丽 IliCounty = &伊丽IliCounty{}

func init() {
	f := CIli伊丽.(*伊丽IliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1425",
		Title:     "ili",
		TitleName: "伊丽",
		TitleCode: "c_ili",
		Baronies:  map[string]feud.Barony{},
	}

	f.亦力把里Ilibaliq = BIlibaliq亦力把里
	f.亦力把里Ilibaliq.SetParent(f)
	
	f.哈什Kax = BKax哈什
	f.哈什Kax.SetParent(f)
	
	f.科克塔尔Koktal = BKoktal科克塔尔
	f.科克塔尔Koktal.SetParent(f)
	
	f.固勒扎Kulja = BKulja固勒扎
	f.固勒扎Kulja.SetParent(f)
	
	f.巩乃斯Kunes = BKunes巩乃斯
	f.巩乃斯Kunes.SetParent(f)
	
	f.沙伦Sharyn = BSharyn沙伦
	f.沙伦Sharyn.SetParent(f)
	
	f.特克斯Tekes = BTekes特克斯
	f.特克斯Tekes.SetParent(f)
	
}
