package bergenshus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BergenshusCounty interface {
    feud.County
    BBergen卑尔根() 	feud.Barony
    BBergenhus卑尔根胡斯() 	feud.Barony
    BFedje费迪厄() 	feud.Barony
    BHove霍沃() 	feud.Barony
    BKinsarvik欣萨维克() 	feud.Barony
    BLyse吕瑟() 	feud.Barony
    BStrandvik斯特兰维克() 	feud.Barony
}

type 卑尔根胡斯BergenshusCounty struct {
	feud.BaseCounty
	卑尔根Bergen 	feud.Barony
	卑尔根胡斯Bergenhus 	feud.Barony
	费迪厄Fedje 	feud.Barony
	霍沃Hove 	feud.Barony
	欣萨维克Kinsarvik 	feud.Barony
	吕瑟Lyse 	feud.Barony
	斯特兰维克Strandvik 	feud.Barony
}

func (c *卑尔根胡斯BergenshusCounty) BBergen卑尔根() feud.Barony {
	return c.卑尔根Bergen
}
    
func (c *卑尔根胡斯BergenshusCounty) BBergenhus卑尔根胡斯() feud.Barony {
	return c.卑尔根胡斯Bergenhus
}
    
func (c *卑尔根胡斯BergenshusCounty) BFedje费迪厄() feud.Barony {
	return c.费迪厄Fedje
}
    
func (c *卑尔根胡斯BergenshusCounty) BHove霍沃() feud.Barony {
	return c.霍沃Hove
}
    
func (c *卑尔根胡斯BergenshusCounty) BKinsarvik欣萨维克() feud.Barony {
	return c.欣萨维克Kinsarvik
}
    
func (c *卑尔根胡斯BergenshusCounty) BLyse吕瑟() feud.Barony {
	return c.吕瑟Lyse
}
    
func (c *卑尔根胡斯BergenshusCounty) BStrandvik斯特兰维克() feud.Barony {
	return c.斯特兰维克Strandvik
}
    
var CBergenshus卑尔根胡斯 BergenshusCounty = &卑尔根胡斯BergenshusCounty{}

func init() {
	f := CBergenshus卑尔根胡斯.(*卑尔根胡斯BergenshusCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "274",
		Title:     "bergenshus",
		TitleName: "卑尔根胡斯",
		TitleCode: "c_bergenshus",
		Baronies:  map[string]feud.Barony{},
	}

	f.卑尔根Bergen = BBergen卑尔根
	f.卑尔根Bergen.SetParent(f)
	
	f.卑尔根胡斯Bergenhus = BBergenhus卑尔根胡斯
	f.卑尔根胡斯Bergenhus.SetParent(f)
	
	f.费迪厄Fedje = BFedje费迪厄
	f.费迪厄Fedje.SetParent(f)
	
	f.霍沃Hove = BHove霍沃
	f.霍沃Hove.SetParent(f)
	
	f.欣萨维克Kinsarvik = BKinsarvik欣萨维克
	f.欣萨维克Kinsarvik.SetParent(f)
	
	f.吕瑟Lyse = BLyse吕瑟
	f.吕瑟Lyse.SetParent(f)
	
	f.斯特兰维克Strandvik = BStrandvik斯特兰维克
	f.斯特兰维克Strandvik.SetParent(f)
	
}
